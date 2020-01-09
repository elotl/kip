package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/virtual-kubelet/virtual-kubelet/log"
	"github.com/virtual-kubelet/virtual-kubelet/node/api"
	"github.com/virtual-kubelet/virtual-kubelet/trace"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	stats "k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1"
)

const (
	// Provider configuration defaults.
	defaultCPUCapacity    = "20"
	defaultMemoryCapacity = "100Gi"
	defaultPodCapacity    = "20"

	// Values used in tracing as attribute keys.
	namespaceKey     = "namespace"
	nameKey          = "name"
	containerNameKey = "containerName"
)

type InstanceProvider struct {
	nodeName           string
	operatingSystem    string
	internalIP         string
	daemonEndpointPort int32
	config             Config
	startTime          time.Time
	notifier           func(*v1.Pod)
}

type Config struct {
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
	Pods   string `json:"pods,omitempty"`
}

func NewProvider(providerConfig, nodeName, operatingSystem string, internalIP string, daemonEndpointPort int32) (*InstanceProvider, error) {
	config, err := loadConfig(providerConfig, nodeName)
	if err != nil {
		return nil, err
	}
	provider := InstanceProvider{
		nodeName:           nodeName,
		operatingSystem:    operatingSystem,
		internalIP:         internalIP,
		daemonEndpointPort: daemonEndpointPort,
		//pods:               make(map[string]*v1.Pod),
		config:    config,
		startTime: time.Now(),
	}
	return &provider, nil
}

func loadConfig(providerConfig, nodeName string) (config Config, err error) {
	data, err := ioutil.ReadFile(providerConfig)
	if err != nil {
		return config, fmt.Errorf("reading config %q: %v", providerConfig, err)
	}
	configMap := map[string]Config{}
	err = json.Unmarshal(data, &configMap)
	if err != nil {
		log.G(context.TODO()).Errorf(
			"parsing config %q: %v", providerConfig, err)
		return config, err
	}
	if _, exist := configMap[nodeName]; exist {
		log.G(context.TODO()).Infof("found config for node %q", nodeName)
		config = configMap[nodeName]
	}
	if config.CPU == "" {
		config.CPU = defaultCPUCapacity
	}
	if config.Memory == "" {
		config.Memory = defaultMemoryCapacity
	}
	if config.Pods == "" {
		config.Pods = defaultPodCapacity
	}
	if _, err = resource.ParseQuantity(config.CPU); err != nil {
		return config, fmt.Errorf("Invalid CPU value %v", config.CPU)
	}
	if _, err = resource.ParseQuantity(config.Memory); err != nil {
		return config, fmt.Errorf("Invalid memory value %v", config.Memory)
	}
	if _, err = resource.ParseQuantity(config.Pods); err != nil {
		return config, fmt.Errorf("Invalid pods value %v", config.Pods)
	}
	return config, nil
}

func (p *InstanceProvider) CreatePod(ctx context.Context, pod *v1.Pod) error {
	ctx, span := trace.StartSpan(ctx, "CreatePod")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, pod.Namespace, nameKey, pod.Name)
	log.G(ctx).Infof("CreatePod %q", pod.Name)
	//p.notifier(pod)
	return fmt.Errorf("not implemented")
}

func (p *InstanceProvider) UpdatePod(ctx context.Context, pod *v1.Pod) error {
	ctx, span := trace.StartSpan(ctx, "UpdatePod")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, pod.Namespace, nameKey, pod.Name)
	log.G(ctx).Infof("UpdatePod %q", pod.Name)
	//p.notifier(pod)
	return fmt.Errorf("not implemented")
}

// DeletePod deletes the specified pod out of memory.
func (p *InstanceProvider) DeletePod(ctx context.Context, pod *v1.Pod) (err error) {
	ctx, span := trace.StartSpan(ctx, "DeletePod")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, pod.Namespace, nameKey, pod.Name)
	log.G(ctx).Infof("DeletePod %q", pod.Name)
	//p.notifier(pod)
	return fmt.Errorf("not implemented")
}

func (p *InstanceProvider) GetPod(ctx context.Context, namespace, name string) (pod *v1.Pod, err error) {
	ctx, span := trace.StartSpan(ctx, "GetPod")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, namespace, nameKey, name)
	log.G(ctx).Infof("GetPod %q", name)
	//p.notifier(pod)
	return nil, fmt.Errorf("not implemented")
}

func (p *InstanceProvider) GetContainerLogs(ctx context.Context, namespace, podName, containerName string, opts api.ContainerLogOpts) (io.ReadCloser, error) {
	ctx, span := trace.StartSpan(ctx, "GetContainerLogs")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, namespace, nameKey, podName, containerNameKey, containerName)
	log.G(ctx).Infof("GetContainerLogs %q", podName)
	//p.notifier(pod)
	return nil, fmt.Errorf("not implemented")
}

func (p *InstanceProvider) RunInContainer(ctx context.Context, namespace, podName, containerName string, cmd []string, attach api.AttachIO) error {
	ctx, span := trace.StartSpan(ctx, "RunInContainer")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, namespace, nameKey, podName, containerNameKey, containerName)
	log.G(ctx).Infof("RunInContainer %q %v", podName, cmd)
	//p.notifier(pod)
	return fmt.Errorf("not implemented")
}

func (p *InstanceProvider) GetPodStatus(ctx context.Context, namespace, name string) (*v1.PodStatus, error) {
	ctx, span := trace.StartSpan(ctx, "GetPodStatus")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, namespace, nameKey, name)
	log.G(ctx).Infof("GetPodStatus %q", name)
	//p.notifier(pod)
	return nil, fmt.Errorf("not implemented")
}

// GetPods returns a list of all pods known to be "running".
func (p *InstanceProvider) GetPods(ctx context.Context) ([]*v1.Pod, error) {
	ctx, span := trace.StartSpan(ctx, "GetPods")
	defer span.End()
	log.G(ctx).Infof("GetPods")
	//p.notifier(pod)
	return nil, fmt.Errorf("not implemented")
}

func (p *InstanceProvider) ConfigureNode(ctx context.Context, n *v1.Node) {
	ctx, span := trace.StartSpan(ctx, "ConfigureNode")
	defer span.End()
	log.G(ctx).Infof("ConfigureNode")
	n.Status.Capacity = p.capacity()
	n.Status.Allocatable = p.capacity()
	n.Status.Conditions = p.nodeConditions()
	n.Status.Addresses = p.nodeAddresses()
	n.Status.DaemonEndpoints = p.nodeDaemonEndpoints()
	os := p.operatingSystem
	if os == "" {
		os = "Linux"
	}
	n.Status.NodeInfo.OperatingSystem = os
	n.Status.NodeInfo.Architecture = "amd64"
}

func (p *InstanceProvider) capacity() v1.ResourceList {
	return v1.ResourceList{
		"cpu":    resource.MustParse(p.config.CPU),
		"memory": resource.MustParse(p.config.Memory),
		"pods":   resource.MustParse(p.config.Pods),
	}
}

func (p *InstanceProvider) nodeConditions() []v1.NodeCondition {
	return []v1.NodeCondition{
		{
			Type:               "Ready",
			Status:             v1.ConditionTrue,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "KubeletReady",
			Message:            "kubelet is ready",
		},
		{
			Type:               "OutOfDisk",
			Status:             v1.ConditionFalse,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "KubeletHasSufficientDisk",
			Message:            "kubelet has sufficient disk space available",
		},
		{
			Type:               "MemoryPressure",
			Status:             v1.ConditionFalse,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "KubeletHasSufficientMemory",
			Message:            "kubelet has sufficient memory available",
		},
		{
			Type:               "DiskPressure",
			Status:             v1.ConditionFalse,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "KubeletHasNoDiskPressure",
			Message:            "kubelet has no disk pressure",
		},
		{
			Type:               "NetworkUnavailable",
			Status:             v1.ConditionFalse,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "RouteCreated",
			Message:            "RouteController created a route",
		},
	}

}

func (p *InstanceProvider) nodeAddresses() []v1.NodeAddress {
	return []v1.NodeAddress{
		{
			Type:    "InternalIP",
			Address: p.internalIP,
		},
	}
}

func (p *InstanceProvider) nodeDaemonEndpoints() v1.NodeDaemonEndpoints {
	return v1.NodeDaemonEndpoints{
		KubeletEndpoint: v1.DaemonEndpoint{
			Port: p.daemonEndpointPort,
		},
	}
}

func (p *InstanceProvider) GetStatsSummary(ctx context.Context) (*stats.Summary, error) {
	var span trace.Span
	ctx, span = trace.StartSpan(ctx, "GetStatsSummary")
	defer span.End()
	res := &stats.Summary{}
	res.Node = stats.NodeStats{
		NodeName:  p.nodeName,
		StartTime: metav1.NewTime(p.startTime),
	}
	//	time := metav1.NewTime(time.Now())
	//	for _, pod := range p.pods {
	//		var (
	//			totalUsageNanoCores uint64
	//			totalUsageBytes uint64
	//		)
	//		pss := stats.PodStats{
	//			PodRef: stats.PodReference{
	//				Name:      pod.Name,
	//				Namespace: pod.Namespace,
	//				UID:       string(pod.UID),
	//			},
	//			StartTime: pod.CreationTimestamp,
	//		}
	//		for _, container := range pod.Spec.Containers {
	//			dummyUsageNanoCores := uint64(rand.Uint32())
	//			totalUsageNanoCores += dummyUsageNanoCores
	//			dummyUsageBytes := uint64(rand.Uint32())
	//			totalUsageBytes += dummyUsageBytes
	//			pss.Containers = append(pss.Containers, stats.ContainerStats{
	//				Name:      container.Name,
	//				StartTime: pod.CreationTimestamp,
	//				CPU: &stats.CPUStats{
	//					Time:           time,
	//					UsageNanoCores: &dummyUsageNanoCores,
	//				},
	//				Memory: &stats.MemoryStats{
	//					Time:       time,
	//					UsageBytes: &dummyUsageBytes,
	//				},
	//			})
	//		}
	//		pss.CPU = &stats.CPUStats{
	//			Time:           time,
	//			UsageNanoCores: &totalUsageNanoCores,
	//		}
	//		pss.Memory = &stats.MemoryStats{
	//			Time:       time,
	//			UsageBytes: &totalUsageBytes,
	//		}
	//		res.Pods = append(res.Pods, pss)
	//	}
	return res, nil
}

// NotifyPods is called to set a pod notifier callback function. This should be
// called before any operations are done within the provider.
func (p *InstanceProvider) NotifyPods(ctx context.Context, notifier func(*v1.Pod)) {
	p.notifier = notifier
}

func addAttributes(ctx context.Context, span trace.Span, attrs ...string) context.Context {
	if len(attrs)%2 == 1 {
		return ctx
	}
	for i := 0; i < len(attrs); i += 2 {
		ctx = span.WithField(ctx, attrs[i], attrs[i+1])
	}
	return ctx
}
