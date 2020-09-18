package server

import (
	"context"
	"time"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/util"
	"github.com/virtual-kubelet/virtual-kubelet/podutils"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/record"
	podshelper "k8s.io/kubernetes/pkg/apis/core/pods"
	"k8s.io/kubernetes/pkg/fieldpath"
)

const (
	DefaultInformerResyncPeriod = 1 * time.Minute
)

type EnvironmentResolver interface {
	Resolve(*api.Pod) error
}

type KipPodEnvironmentResolver struct {
	nodeName        string
	hostIP          string
	recorder        record.EventRecorder
	secretLister    corev1listers.SecretLister
	configMapLister corev1listers.ConfigMapLister
	serviceLister   corev1listers.ServiceLister
}

func NewKipPodEnvironmentResolver(client kubernetes.Interface, recorder record.EventRecorder, nodeName, hostIP string) *KipPodEnvironmentResolver {
	scmInformerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(client, DefaultInformerResyncPeriod)
	secretInformer := scmInformerFactory.Core().V1().Secrets()
	configMapInformer := scmInformerFactory.Core().V1().ConfigMaps()
	serviceInformer := scmInformerFactory.Core().V1().Services()
	return &KipPodEnvironmentResolver{
		nodeName:        nodeName,
		hostIP:          hostIP,
		recorder:        recorder,
		secretLister:    secretInformer.Lister(),
		configMapLister: configMapInformer.Lister(),
		serviceLister:   serviceInformer.Lister(),
	}
}

func (e *KipPodEnvironmentResolver) Resolve(pod *api.Pod) error {
	namespace, name := util.SplitNamespaceAndName(pod.Name)
	kpod := &v1.Pod{}
	kpod.Name = name
	kpod.Namespace = namespace
	kpod.Annotations = pod.Annotations
	kpod.Labels = pod.Labels
	kpod.UID = types.UID(pod.UID)
	kpod.Spec.RestartPolicy = v1.RestartPolicy(pod.Spec.RestartPolicy)
	kpod.Spec.NodeName = e.nodeName
	kpod.Status.HostIP = e.hostIP

	// Fixme!
	kpod.Spec.ServiceAccountName = "" // todo
	kpod.Spec.SchedulerName = ""      // todo

	kpod.Status.Phase = milpaToK8sPhase(pod, false)
	kpod.Status.PodIP = api.GetPodIP(pod.Status.Addresses)

	kpod.Spec.InitContainers = make([]v1.Container, len(pod.Spec.InitUnits))
	for i := range pod.Spec.InitUnits {
		kpod.Spec.InitContainers[i].Env = milpaToK8sEnv(pod.Spec.InitUnits[i].Env)
		kpod.Spec.InitContainers[i].EnvFrom = milpaToK8sEnvFrom(pod.Spec.InitUnits[i].EnvFrom)
	}
	kpod.Spec.Containers = make([]v1.Container, len(pod.Spec.Units))
	for i := range pod.Spec.Units {
		kpod.Spec.Containers[i].Env = milpaToK8sEnv(pod.Spec.Units[i].Env)
		kpod.Spec.Containers[i].EnvFrom = milpaToK8sEnvFrom(pod.Spec.Units[i].EnvFrom)
	}
	err := podutils.ResolveConfigMapRefs(context.Background(), kpod, e.configMapLister, e.recorder)
	if err != nil {
		return err
	}
	err = podutils.ResolveSecretRefs(context.Background(), kpod, e.secretLister, e.recorder)
	if err != nil {
		return err
	}
	podutils.RemoveEnvFromVars(kpod)
	err = podutils.InsertServiceEnvVars(context.Background(), kpod, e.serviceLister)
	if err != nil {
		return err
	}
	err = podutils.ResolveFieldRefs(kpod)
	if err != nil {
		return err
	}
	// Remove unresolved vars before resolving expansions or else
	// expansions that reference an unresolved variable will expand to
	// empty string ("") instead of remaining unexpanded.
	podutils.RemoveUnresolvedEnvVars(kpod)
	podutils.ResolveEnvVarExpansions(kpod)
	podutils.UniqifyEnvVars(kpod)
	for i := range kpod.Spec.InitContainers {
		pod.Spec.InitUnits[i].Env = k8sToMilpaEnv(kpod.Spec.InitContainers[i].Env)
	}
	for i := range pod.Spec.Units {
		pod.Spec.Units[i].Env = k8sToMilpaEnv(kpod.Spec.Containers[i].Env)
	}

	return nil
}

// Run through all containers in a pod and apply f each container
func forEachContainer(pod *v1.Pod, f func(container *v1.Container)) {
	for i := range pod.Spec.InitContainers {
		f(&pod.Spec.InitContainers[i])
	}
	for i := range pod.Spec.Containers {
		f(&pod.Spec.Containers[i])
	}
}

// Run through all containers in a pod and apply f to each container,
// stop iterating if f returns an error
func forEachContainerWithError(pod *v1.Pod, f func(container *v1.Container) error) error {
	for i := range pod.Spec.InitContainers {
		err := f(&pod.Spec.InitContainers[i])
		if err != nil {
			return err
		}
	}
	for i := range pod.Spec.Containers {
		err := f(&pod.Spec.Containers[i])
		if err != nil {
			return err
		}
	}
	return nil
}

// Resolves the runtime value of all field selectors in a pod's envs.
func (e *KipPodEnvironmentResolver) ResolveFieldRefs(pod *v1.Pod) error {
	err := forEachContainerWithError(pod, func(container *v1.Container) error {
		return e.resolveContainerFieldRefs(pod, container)
	})
	return err
}

// Resolves the runtime value of all field selectors in a pod's envs.
func (e *KipPodEnvironmentResolver) resolveContainerFieldRefs(pod *v1.Pod, container *v1.Container) error {
	for i := range container.Env {
		if container.Env[i].ValueFrom == nil ||
			container.Env[i].ValueFrom.FieldRef == nil {
			continue
		}
		val, err := e.podFieldSelectorRuntimeValue(container.Env[i].ValueFrom.FieldRef, pod)
		if err != nil {
			return err
		}
		container.Env[i].Value = val
		container.Env[i].ValueFrom = nil
	}
	return nil
}

// podFieldSelectorRuntimeValue returns the runtime value of the given
// selector for a pod.  Will throw an error if asked to resolve
// status.hostIP, status.podIP, status.podIPs
func (e *KipPodEnvironmentResolver) podFieldSelectorRuntimeValue(fs *v1.ObjectFieldSelector, pod *v1.Pod) (string, error) {
	internalFieldPath, _, err := podshelper.ConvertDownwardAPIFieldLabel(fs.APIVersion, fs.FieldPath, "")
	if err != nil {
		return "", err
	}
	switch internalFieldPath {
	case "spec.nodeName":
		return pod.Spec.NodeName, nil
	case "spec.serviceAccountName":
		return pod.Spec.ServiceAccountName, nil
	case "status.hostIP":
		return pod.Status.HostIP, nil
	case "status.podIP", "status.podIPs":
		return pod.Status.PodIP, nil
	}
	return fieldpath.ExtractFieldPathAsString(pod, internalFieldPath)
}
