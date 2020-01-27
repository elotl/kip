package server

import (
	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	k8sutil "github.com/elotl/cloud-instance-provider/pkg/util/k8s"
	"github.com/golang/glog"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func (p *InstanceProvider) getStatus(milpaPod *api.Pod, pod *v1.Pod) v1.PodStatus {
	phase := v1.PodUnknown
	switch milpaPod.Status.Phase {
	case api.PodWaiting:
		phase = v1.PodPending
	case api.PodDispatching:
		phase = v1.PodPending
	case api.PodRunning:
		phase = v1.PodRunning
	case api.PodSucceeded:
		phase = v1.PodSucceeded
	case api.PodFailed:
		phase = v1.PodFailed
	case api.PodTerminated:
		phase = v1.PodFailed
	}
	startTime := metav1.Time{}
	if milpaPod.Status.ReadyTime != nil {
		startTime = metav1.NewTime(milpaPod.Status.ReadyTime.Time)
	}
	privateIPv4Address := ""
	for _, address := range milpaPod.Status.Addresses {
		if address.Type == api.PrivateIP {
			privateIPv4Address = address.Address
		}
	}
	initContainerStatuses := make([]v1.ContainerStatus, len(milpaPod.Status.InitUnitStatuses))
	for i, st := range milpaPod.Status.InitUnitStatuses {
		initContainerStatuses[i] = unitToContainerStatus(st)
	}
	containerStatuses := make([]v1.ContainerStatus, len(milpaPod.Status.UnitStatuses))
	for i, st := range milpaPod.Status.UnitStatuses {
		containerStatuses[i] = unitToContainerStatus(st)
	}
	// We use the implementation from Kubernetes here to determine conditions.
	conditions := []v1.PodCondition{}
	conditions = append(conditions, k8sutil.GeneratePodInitializedCondition(&pod.Spec, initContainerStatuses, pod.Status.Phase))
	conditions = append(conditions, k8sutil.GeneratePodReadyCondition(&pod.Spec, conditions, containerStatuses, pod.Status.Phase))
	conditions = append(conditions, k8sutil.GenerateContainersReadyCondition(&pod.Spec, containerStatuses, pod.Status.Phase))
	// PodScheduled is always true when the pod gets to the kubelet.
	conditions = append(conditions, v1.PodCondition{
		Type:   v1.PodScheduled,
		Status: v1.ConditionTrue,
	})
	return v1.PodStatus{
		Phase:                 phase,
		Conditions:            conditions,
		Message:               "",
		Reason:                "",
		HostIP:                p.internalIP,
		PodIP:                 privateIPv4Address,
		StartTime:             &startTime,
		InitContainerStatuses: initContainerStatuses,
		ContainerStatuses:     containerStatuses,
		QOSClass:              v1.PodQOSBestEffort,
	}
}

func unitToContainerStatus(st api.UnitStatus) v1.ContainerStatus {
	cst := v1.ContainerStatus{
		Name:         st.Name,
		Image:        st.Image,
		ImageID:      st.Image,
		RestartCount: st.RestartCount,
	}
	if st.State.Waiting != nil {
		cst.Ready = false
		cst.State.Waiting = &v1.ContainerStateWaiting{
			Reason: st.State.Waiting.Reason,
		}
	}
	if st.State.Running != nil {
		cst.Ready = true // TODO: use readiness probe result.
		cst.State.Running = &v1.ContainerStateRunning{
			StartedAt: metav1.NewTime(st.State.Running.StartedAt.Time),
		}
	}
	if st.State.Terminated != nil {
		cst.Ready = false
		cst.State.Terminated = &v1.ContainerStateTerminated{
			ExitCode:   st.State.Terminated.ExitCode,
			FinishedAt: metav1.NewTime(st.State.Terminated.FinishedAt.Time),
			//ContainerID:
		}
	}
	return cst
}

func containerToUnit(container v1.Container) api.Unit {
	unit := api.Unit{
		Name:       container.Name,
		Image:      container.Image,
		Command:    container.Command,
		Args:       container.Args,
		WorkingDir: container.WorkingDir,
	}
	//Resources: v1.ResourceRequirements{
	//	Limits: v1.ResourceList{
	//		v1.ResourceCPU:    resource.MustParse(fmt.Sprintf("%d", *cntrDef.Cpu)),
	//		v1.ResourceMemory: resource.MustParse(fmt.Sprintf("%dMi", *cntrDef.Memory)),
	//	},
	//	Requests: v1.ResourceList{
	//		v1.ResourceCPU:    resource.MustParse(fmt.Sprintf("%d", *cntrDef.Cpu)),
	//		v1.ResourceMemory: resource.MustParse(fmt.Sprintf("%dMi", *cntrDef.MemoryReservation)),
	//	},
	//},
	for _, e := range container.Env {
		unit.Env = append(unit.Env, api.EnvVar{
			Name:  e.Name,
			Value: e.Value,
		})
	}
	if container.SecurityContext != nil {
		unit.SecurityContext = &api.SecurityContext{
			RunAsUser:  container.SecurityContext.RunAsUser,
			RunAsGroup: container.SecurityContext.RunAsGroup,
		}
		ccaps := container.SecurityContext.Capabilities
		if ccaps != nil {
			caps := &api.Capabilities{
				Add:  make([]string, len(ccaps.Add)),
				Drop: make([]string, len(ccaps.Drop)),
			}
			for i, a := range ccaps.Add {
				caps.Add[i] = string(a)
			}
			for i, d := range ccaps.Drop {
				caps.Drop[i] = string(d)
			}
			unit.SecurityContext.Capabilities = caps
		}
	}
	for _, port := range container.Ports {
		unit.Ports = append(unit.Ports,
			api.ServicePort{
				Port:     int(port.ContainerPort),
				NodePort: int(port.HostPort),
				Protocol: api.Protocol(string(port.Protocol)),
			})
	}
	for _, vm := range container.VolumeMounts {
		unit.VolumeMounts = append(unit.VolumeMounts, api.VolumeMount{
			Name:      vm.Name,
			MountPath: vm.MountPath,
		})
	}
	//container.EnvFrom,
	return unit
}

func unitToContainer(unit api.Unit, container *v1.Container) v1.Container {
	if container == nil {
		container = &v1.Container{}
	}
	container.Name = unit.Name
	container.Image = unit.Image
	container.Command = unit.Command
	container.Args = unit.Args
	container.WorkingDir = unit.WorkingDir
	container.Env = make([]v1.EnvVar, len(unit.Env))
	for i, e := range unit.Env {
		container.Env[i] = v1.EnvVar{
			Name:  e.Name,
			Value: e.Value,
		}
	}
	for _, port := range unit.Ports {
		container.Ports = append(container.Ports,
			v1.ContainerPort{
				ContainerPort: int32(port.Port),
				HostPort:      int32(port.NodePort),
				Protocol:      v1.Protocol(string(port.Protocol)),
			})
	}
	usc := unit.SecurityContext
	if usc != nil {
		if container.SecurityContext == nil {
			container.SecurityContext = &v1.SecurityContext{}
		}
		csc := container.SecurityContext
		csc.RunAsUser = usc.RunAsUser
		csc.RunAsGroup = usc.RunAsGroup
		ucaps := usc.Capabilities
		if ucaps != nil {
			caps := &v1.Capabilities{
				Add:  make([]v1.Capability, len(ucaps.Add)),
				Drop: make([]v1.Capability, len(ucaps.Drop)),
			}
			for i, a := range ucaps.Add {
				caps.Add[i] = v1.Capability(string(a))
			}
			for i, d := range ucaps.Drop {
				caps.Drop[i] = v1.Capability(string(d))
			}
			csc.Capabilities = caps
		}
	}
	for _, vm := range unit.VolumeMounts {
		container.VolumeMounts = append(container.VolumeMounts, v1.VolumeMount{
			Name:      vm.Name,
			MountPath: vm.MountPath,
		})
	}

	return *container
}

// For development, I'm making the assumption we can safely copy these
// values instead of doing deep copies?
// Todo: verify items coming out of our informers/listers are copies
func k8sToMilpaVolume(vol v1.Volume) *api.Volume {
	convertKeyToPath := func(k8s []v1.KeyToPath) []api.KeyToPath {
		milpa := make([]api.KeyToPath, len(k8s))
		for i := range k8s {
			milpa = append(milpa, api.KeyToPath{
				Key:  k8s[i].Key,
				Path: k8s[i].Path,
				Mode: k8s[i].Mode,
			})
		}
		return milpa
	}
	if vol.Secret != nil {
		return &api.Volume{
			Name: vol.Name,
			VolumeSource: api.VolumeSource{
				Secret: &api.SecretVolumeSource{
					SecretName:  vol.Secret.SecretName,
					Items:       convertKeyToPath(vol.Secret.Items),
					DefaultMode: vol.Secret.DefaultMode,
					Optional:    vol.Secret.Optional,
				},
			},
		}
	} else if vol.ConfigMap != nil {
		return &api.Volume{
			Name: vol.Name,
			VolumeSource: api.VolumeSource{
				ConfigMap: &api.ConfigMapVolumeSource{
					LocalObjectReference: api.LocalObjectReference{
						Name: vol.ConfigMap.Name,
					},
					Items:       convertKeyToPath(vol.ConfigMap.Items),
					DefaultMode: vol.ConfigMap.DefaultMode,
					Optional:    vol.ConfigMap.Optional,
				},
			},
		}
	} else if vol.EmptyDir != nil {
		var sizeLimit int64
		if vol.EmptyDir.SizeLimit != nil {
			sizeLimit, _ = vol.EmptyDir.SizeLimit.AsInt64()
		}
		return &api.Volume{
			Name: vol.Name,
			VolumeSource: api.VolumeSource{
				EmptyDir: &api.EmptyDir{
					Medium:    api.StorageMedium(string(vol.EmptyDir.Medium)),
					SizeLimit: sizeLimit,
				},
			},
		}
	} else {
		glog.Warningf("Unsupported volume type for volume: %s", vol.Name)
		return &api.Volume{
			Name: vol.Name,
			VolumeSource: api.VolumeSource{
				EmptyDir: &api.EmptyDir{
					Medium: api.StorageMediumMemory,
				},
			},
		}
	}
}

// For development, I'm making the assumption we can safely copy these
// values instead of doing deep copies?
// Todo: verify items coming out of our informers/listers are copies
func milpaToK8sVolume(vol api.Volume) *v1.Volume {
	convertKeyToPath := func(milpa []api.KeyToPath) []v1.KeyToPath {
		k8s := make([]v1.KeyToPath, len(milpa))
		for i := range milpa {
			k8s = append(k8s, v1.KeyToPath{
				Key:  milpa[i].Key,
				Path: milpa[i].Path,
				Mode: milpa[i].Mode,
			})
		}
		return k8s
	}
	if vol.Secret != nil {
		return &v1.Volume{
			Name: vol.Name,
			VolumeSource: v1.VolumeSource{
				Secret: &v1.SecretVolumeSource{
					SecretName:  vol.Secret.SecretName,
					Items:       convertKeyToPath(vol.Secret.Items),
					DefaultMode: vol.Secret.DefaultMode,
					Optional:    vol.Secret.Optional,
				},
			},
		}
	} else if vol.ConfigMap != nil {
		return &v1.Volume{
			Name: vol.Name,
			VolumeSource: v1.VolumeSource{
				ConfigMap: &v1.ConfigMapVolumeSource{
					LocalObjectReference: v1.LocalObjectReference{
						Name: vol.ConfigMap.Name,
					},
					Items:       convertKeyToPath(vol.ConfigMap.Items),
					DefaultMode: vol.ConfigMap.DefaultMode,
					Optional:    vol.ConfigMap.Optional,
				},
			},
		}
	} else if vol.EmptyDir != nil {
		var sizeLimit *resource.Quantity
		if vol.EmptyDir.SizeLimit != 0 {
			sizeLimit = resource.NewQuantity(vol.EmptyDir.SizeLimit, resource.BinarySI)
		}
		return &v1.Volume{
			Name: vol.Name,
			VolumeSource: v1.VolumeSource{
				EmptyDir: &v1.EmptyDirVolumeSource{
					Medium:    v1.StorageMedium(string(vol.EmptyDir.Medium)),
					SizeLimit: sizeLimit,
				},
			},
		}
	} else {
		glog.Warningf("Unspported volume type for volume: %s", vol.Name)
	}
	return nil
}

func (p *InstanceProvider) k8sToMilpaPod(pod *v1.Pod) (*api.Pod, error) {
	milpapod := api.NewPod()
	milpapod.Name = util.WithNamespace(pod.Namespace, pod.Name)
	milpapod.Namespace = pod.Namespace
	milpapod.UID = string(pod.UID)
	milpapod.CreationTimestamp = api.Time{Time: pod.CreationTimestamp.Time}
	milpapod.Labels = pod.Labels
	milpapod.Annotations = pod.Annotations
	milpapod.Spec.RestartPolicy = api.RestartPolicy(string(pod.Spec.RestartPolicy))
	podsc := pod.Spec.SecurityContext
	if podsc != nil {
		mpsc := &api.PodSecurityContext{
			RunAsUser:          podsc.RunAsUser,
			RunAsGroup:         podsc.RunAsGroup,
			SupplementalGroups: podsc.SupplementalGroups,
		}
		mpsc.NamespaceOptions = &api.NamespaceOption{
			Network: api.NamespaceModePod,
			Pid:     api.NamespaceModeContainer,
			Ipc:     api.NamespaceModeContainer,
		}
		if pod.Spec.HostNetwork {
			mpsc.NamespaceOptions.Network = api.NamespaceModeNode
		}
		if pod.Spec.HostPID {
			mpsc.NamespaceOptions.Pid = api.NamespaceModeNode
		}
		if pod.Spec.HostIPC {
			mpsc.NamespaceOptions.Ipc = api.NamespaceModeNode
		}
		if pod.Spec.ShareProcessNamespace != nil {
			// TODO: containers share pid namespace in the pod.
		}
		mpsc.Sysctls = make([]api.Sysctl, len(podsc.Sysctls))
		for i, sysctl := range podsc.Sysctls {
			mpsc.Sysctls[i] = api.Sysctl{
				Name:  sysctl.Name,
				Value: sysctl.Value,
			}
		}
		milpapod.Spec.SecurityContext = mpsc
	}
	for _, initContainer := range pod.Spec.InitContainers {
		initUnit := containerToUnit(initContainer)
		milpapod.Spec.InitUnits = append(milpapod.Spec.InitUnits, initUnit)
	}
	for _, container := range pod.Spec.Containers {
		unit := containerToUnit(container)
		milpapod.Spec.Units = append(milpapod.Spec.Units, unit)
	}
	for _, volume := range pod.Spec.Volumes {
		volume := k8sToMilpaVolume(volume)
		if volume != nil {
			milpapod.Spec.Volumes = append(milpapod.Spec.Volumes, *volume)
		}
	}
	return milpapod, nil
}

func (p *InstanceProvider) milpaToK8sPod(milpaPod *api.Pod) (*v1.Pod, error) {
	namespace, name := util.SplitNamespaceAndName(milpaPod.Name)
	pod := &v1.Pod{}
	pod.Kind = "Pod"
	pod.APIVersion = "v1"
	pod.Name = name
	pod.Namespace = namespace
	pod.UID = types.UID(milpaPod.UID)
	pod.CreationTimestamp = metav1.NewTime(milpaPod.CreationTimestamp.Time)
	pod.Labels = milpaPod.Labels
	pod.Annotations = milpaPod.Annotations
	pod.Spec.NodeName = p.nodeName
	pod.Spec.Volumes = []v1.Volume{}
	pod.Spec.RestartPolicy = v1.RestartPolicy(string(milpaPod.Spec.RestartPolicy))
	mpsc := milpaPod.Spec.SecurityContext
	if mpsc != nil {
		if pod.Spec.SecurityContext == nil {
			pod.Spec.SecurityContext = &v1.PodSecurityContext{}
		}
		psc := pod.Spec.SecurityContext
		psc.RunAsUser = mpsc.RunAsUser
		psc.RunAsGroup = mpsc.RunAsGroup
		psc.SupplementalGroups = mpsc.SupplementalGroups
		if mpsc.NamespaceOptions != nil {
			if mpsc.NamespaceOptions.Network == api.NamespaceModeNode {
				pod.Spec.HostNetwork = true
			}
			if mpsc.NamespaceOptions.Pid == api.NamespaceModeNode {
				pod.Spec.HostPID = true
			}
			if mpsc.NamespaceOptions.Ipc == api.NamespaceModeNode {
				pod.Spec.HostIPC = true
			}
		}
		psc.Sysctls = make([]v1.Sysctl, len(mpsc.Sysctls))
		for i, sysctl := range mpsc.Sysctls {
			psc.Sysctls[i] = v1.Sysctl{
				Name:  sysctl.Name,
				Value: sysctl.Value,
			}
		}
		pod.Spec.SecurityContext = psc
	}
	initContainerMap := make(map[string]v1.Container)
	for _, initContainer := range pod.Spec.InitContainers {
		initContainerMap[initContainer.Name] = initContainer
	}
	containerMap := make(map[string]v1.Container)
	for _, container := range pod.Spec.Containers {
		containerMap[container.Name] = container
	}
	for _, initUnit := range milpaPod.Spec.InitUnits {
		initContainer, exists := initContainerMap[initUnit.Name]
		ptr := &initContainer
		if !exists {
			ptr = nil
		}
		initContainer = unitToContainer(initUnit, ptr)
		pod.Spec.InitContainers = append(pod.Spec.InitContainers, initContainer)
	}
	for _, unit := range milpaPod.Spec.Units {
		container, exists := containerMap[unit.Name]
		ptr := &container
		if !exists {
			ptr = nil
		}
		container = unitToContainer(unit, ptr)
		pod.Spec.Containers = append(pod.Spec.Containers, container)
	}
	for _, volume := range milpaPod.Spec.Volumes {
		volume := milpaToK8sVolume(volume)
		if volume != nil {
			pod.Spec.Volumes = append(pod.Spec.Volumes, *volume)
		}
	}
	pod.Status = p.getStatus(milpaPod, pod)
	return pod, nil
}
