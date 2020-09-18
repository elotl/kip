package server

import (
	"time"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/util"
	"github.com/virtual-kubelet/virtual-kubelet/podutils"
	v1 "k8s.io/api/core/v1"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	corev1listers "k8s.io/client-go/listers/core/v1"
)

const (
	DefaultInformerResyncPeriod = 1 * time.Minute
)

type EnvironmentResolver interface {
	Resolve(*api.Pod) error
}

type KipPodEnvironmentResolver struct {
	secretLister    corev1listers.SecretLister
	configMapLister corev1listers.ConfigMapLister
	serviceLister   corev1listers.ServiceLister
}

func NewKipPodEnvironmentResolver(client kubernetes.Interface) *KipPodEnvironmentResolver {
	scmInformerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(client, DefaultInformerResyncPeriod)
	secretInformer := scmInformerFactory.Core().V1().Secrets()
	configMapInformer := scmInformerFactory.Core().V1().ConfigMaps()
	serviceInformer := scmInformerFactory.Core().V1().Services()
	return &KipPodEnvironmentResolver{
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
	kpod.Spec.InitContainers = make([]v1.Container, len(pod.Spec.InitUnits))
	for i := range pod.Spec.InitUnits {
		kpod.Spec.InitContainers[i].Env = milpaToK8sEnv(pod.Spec.InitUnits[i].Env)
	}
	kpod.Spec.Containers = make([]v1.Container, len(pod.Spec.Units))
	for i := range pod.Spec.Units {
		kpod.Spec.Containers[i].Env = milpaToK8sEnv(pod.Spec.Units[i].Env)
	}
	// err := podutils.ResolveConfigMapRefs(ctx, pod, configMapLister, recorder)
	// if err != nil {
	// 	return err
	// }
	// err = podutils.ResolveSecretRefs(ctx, pod, secretLister, recorder)
	// if err != nil {
	// 	return err
	// }
	// RemoveEnvFromVars(pod)
	// err = podutils.InsertServiceEnvVars(ctx, pod, serviceLister)
	// if err != nil {
	// 	return err
	// }
	// err = podutils.ResolveFieldRefs(pod)
	// if err != nil {
	// 	return err
	// }
	// Remove unresolved vars before resolving expansions or else
	// expansions that reference an unresolved variable will expand to
	// empty string ("") instead of remaining unexpanded.
	// podutils.RemoveUnresolvedEnvVars(pod)

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
