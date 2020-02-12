package server

import (
	"fmt"
	"net"
	"sort"
	"sync"
	"time"

	"github.com/docker/libkv/store"
	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/api/validation"
	"github.com/elotl/cloud-instance-provider/pkg/certs"
	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/cloud-instance-provider/pkg/etcd"
	"github.com/elotl/cloud-instance-provider/pkg/nodeclient"
	"github.com/elotl/cloud-instance-provider/pkg/portmanager"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud/azure"
	"github.com/elotl/cloud-instance-provider/pkg/server/events"
	"github.com/elotl/cloud-instance-provider/pkg/server/nodemanager"
	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/elotl/cloud-instance-provider/pkg/util/cloudinitfile"
	"github.com/elotl/cloud-instance-provider/pkg/util/conmap"
	"github.com/elotl/cloud-instance-provider/pkg/util/instanceselector"
	"github.com/elotl/cloud-instance-provider/pkg/util/timeoutmap"
	"github.com/elotl/cloud-instance-provider/pkg/util/validation/field"
	"github.com/virtual-kubelet/node-cli/manager"
	"github.com/virtual-kubelet/virtual-kubelet/errdefs"
	"github.com/virtual-kubelet/virtual-kubelet/trace"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
	utiliptables "k8s.io/kubernetes/pkg/util/iptables"
	utilexec "k8s.io/utils/exec"
)

const (
	// Values used in tracing as attribute keys.
	namespaceKey          = "namespace"
	nameKey               = "name"
	containerNameKey      = "containerName"
	etcdClusterRegionPath = "milpa/cluster/region"
	kubernetesPodKey      = "elotl.co/kubernetes-pod"
	defaultPort           = 54555
	defaultProtocol       = "tcp"
)

var (
	MaxEventListSize = 4000 // modified for testing
)

type Controller interface {
	Start(quit <-chan struct{}, wg *sync.WaitGroup)
	Dump() []byte
}

type InstanceProvider struct {
	Registries        map[string]registry.Registryer
	Encoder           api.MilpaCodec
	SystemQuit        <-chan struct{}
	SystemWaitGroup   *sync.WaitGroup
	Controllers       map[string]Controller
	ItzoClientFactory nodeclient.ItzoClientFactoryer
	cloudClient       cloud.CloudClient
	controllerManager *ControllerManager
	// Unsure how many of these are actually needed here...
	nodeName           string
	internalIP         string
	daemonEndpointPort int32
	kubeletConfig      KubeletConfig
	startTime          time.Time
	notifier           func(*v1.Pod)
	portManager        *portmanager.PortManager
}

func validateWriteToEtcd(client *etcd.SimpleEtcd) error {
	klog.V(2).Info("validating write access to etcd (will block until we can connect)")
	wo := &store.WriteOptions{
		IsDir: false,
		TTL:   2 * time.Second,
	}

	err := client.PutNoTimeout("/milpa/startup", []byte("OK"), wo)
	if err != nil {
		return err
	}
	klog.V(2).Info("write to etcd successful")
	return nil
}

func setupEtcd(configFile, dataDir string, quit <-chan struct{}, wg *sync.WaitGroup) (*etcd.SimpleEtcd, error) {
	// if we have client endpoints, don't start the server. This could
	// change in the future if we want the embedded server to join
	// existing etcd server, but, for now just don't start it.
	var client *etcd.SimpleEtcd
	klog.V(2).Infof("starting internal etcd")
	etcdServer := etcd.EtcdServer{
		ConfigFile: configFile,
		DataDir:    dataDir,
	}
	err := etcdServer.Start(quit, wg)
	if err != nil {
		return nil, util.WrapError(
			err, "error creating internal etcd storage backend")
	}
	client = etcdServer.Client
	err = validateWriteToEtcd(client)
	if err != nil {
		return nil, util.WrapError(err, "fatal error: Could not write to etcd")
	}
	return client, err
}

func ensureRegionUnchanged(etcdClient *etcd.SimpleEtcd, region string) error {
	klog.V(2).Infof("ensuring region has not changed")
	var savedRegion string
	pair, err := etcdClient.Get(etcdClusterRegionPath)
	if err != nil {
		if err != store.ErrKeyNotFound {
			return err
		}
		_, _, err = etcdClient.AtomicPut(etcdClusterRegionPath, []byte(region), nil, nil)
		return err
	}
	savedRegion = string(pair.Value)
	if region != savedRegion {
		return fmt.Errorf(
			"error: region has changed from %s to %s. "+
				"This is unsupported. "+
				"Please delete all cluster resources and rename your cluster",
			savedRegion, region)
	}
	return nil
}

// InstanceProvider should implement node.PodLifecycleHandler
func NewInstanceProvider(configFilePath, nodeName, internalIP, serverURL, networkAgentSecret string, daemonEndpointPort int32, debugServer bool, rm *manager.ResourceManager, systemQuit <-chan struct{}) (*InstanceProvider, error) {
	systemWG := &sync.WaitGroup{}

	execer := utilexec.New()
	ipt := utiliptables.New(execer, utiliptables.ProtocolIpv4)
	portManager := portmanager.NewPortManager(ipt)

	serverConfigFile, err := ParseConfig(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("loading Config file (%s) failed with error: %s",
			configFilePath, err.Error())
	}
	errs := validateServerConfigFile(serverConfigFile)
	if len(errs) > 0 {
		return nil, fmt.Errorf("invalid server config: %v", errs.ToAggregate())
	}

	etcdClient, err := setupEtcd(
		serverConfigFile.Etcd.Internal.ConfigFile,
		serverConfigFile.Etcd.Internal.DataDir,
		systemQuit,
		systemWG,
	)
	if err != nil {
		return nil, fmt.Errorf("etcd error: %s", err)
	}
	controllerID, err := getControllerID(etcdClient)
	if err != nil {
		return nil, fmt.Errorf("controller ID error: %s", err)
	}
	if serverConfigFile.Testing.ControllerID != "" {
		controllerID = serverConfigFile.Testing.ControllerID
	}
	nametag := serverConfigFile.Cells.Nametag
	if nametag == "" {
		nametag = controllerID
	}

	klog.V(2).Infof("ControllerID: %s", controllerID)

	certFactory, err := certs.New(etcdClient)
	if err != nil {
		return nil, fmt.Errorf("error setting up certificate factory: %v", err)
	}

	cloudClient, err := ConfigureCloud(serverConfigFile, controllerID, nametag)
	if err != nil {
		return nil, fmt.Errorf("error configuring cloud client: %v", err)
	}
	cloudRegion := cloudClient.GetAttributes().Region
	err = ensureRegionUnchanged(etcdClient, cloudRegion)
	if err != nil {
		return nil, fmt.Errorf("error ensuring Milpa region is unchanged: %v", err)
	}
	clientCert, err := certFactory.CreateClientCert()
	if err != nil {
		return nil, fmt.Errorf("error creating node client certificate: %v", err)
	}
	cloudStatus := cloudClient.CloudStatusKeeper()
	cloudStatus.Start()
	statefulValidator := validation.NewStatefulValidator(
		cloudStatus,
		cloudClient.GetAttributes().Provider,
		cloudClient.GetVPCCIDRs(),
	)
	err = instanceselector.Setup(
		cloudClient.GetAttributes().Provider,
		cloudRegion,
		serverConfigFile.Cells.DefaultInstanceType)
	if err != nil {
		return nil, fmt.Errorf("error setting up instance selector %s", err)
	}
	// Ugly: need to do validation of this field after we have setup
	// the instanceselector
	errs = validation.ValidateInstanceType(serverConfigFile.Cells.DefaultInstanceType, field.NewPath("nodes.defaultInstanceType"))
	if len(errs) > 0 {
		return nil, fmt.Errorf("error validating server.yml: %v", errs.ToAggregate())
	}

	klog.V(2).Infof("setting up events")
	eventSystem := events.NewEventSystem(systemQuit, systemWG)

	klog.V(2).Infof("setting up registry")
	podRegistry := registry.NewPodRegistry(
		etcdClient, api.VersioningCodec{}, eventSystem, statefulValidator)
	nodeRegistry := registry.NewNodeRegistry(
		etcdClient, api.VersioningCodec{}, eventSystem)
	eventRegistry := registry.NewEventRegistry(
		etcdClient, api.VersioningCodec{}, eventSystem)
	logRegistry := registry.NewLogRegistry(
		etcdClient, api.VersioningCodec{}, eventSystem)
	metricsRegistry := registry.NewMetricsRegistry(240)
	kv := map[string]registry.Registryer{
		"Pod":    podRegistry,
		"Node":   nodeRegistry,
		"Event":  eventRegistry,
		"Log":    logRegistry,
		"Metric": metricsRegistry,
	}

	usePublicIPs := !cloudClient.ControllerInsideVPC()
	itzoClientFactory := nodeclient.NewItzoFactory(
		&certFactory.Root, *clientCert, usePublicIPs)
	nodeDispenser := nodemanager.NewNodeDispenser()
	podController := &PodController{
		podRegistry:        podRegistry,
		logRegistry:        logRegistry,
		metricsRegistry:    metricsRegistry,
		nodeLister:         nodeRegistry,
		resourceManager:    rm,
		nodeDispenser:      nodeDispenser,
		nodeClientFactory:  itzoClientFactory,
		events:             eventSystem,
		cloudClient:        cloudClient,
		controllerID:       controllerID,
		nametag:            nametag,
		lastStatusReply:    conmap.NewStringTimeTime(),
		serverURL:          serverURL,
		networkAgentSecret: networkAgentSecret,
	}
	imageIdCache := timeoutmap.New(false, nil)
	cloudInitFile := cloudinitfile.New(serverConfigFile.Cells.CloudInitFile)
	fixedSizeVolume := cloudClient.GetAttributes().FixedSizeVolume
	nodeController := &nodemanager.NodeController{
		Config: nodemanager.NodeControllerConfig{
			PoolInterval:      7 * time.Second,
			HeartbeatInterval: 10 * time.Second,
			ReaperInterval:    10 * time.Second,
			ItzoVersion:       serverConfigFile.Cells.Itzo.Version,
			ItzoURL:           serverConfigFile.Cells.Itzo.URL,
		},
		NodeRegistry:  nodeRegistry,
		LogRegistry:   logRegistry,
		PodReader:     podRegistry,
		NodeDispenser: nodeDispenser,
		NodeScaler: nodemanager.NewBindingNodeScaler(
			nodeRegistry,
			serverConfigFile.Cells.StandbyCells,
			cloudStatus,
			serverConfigFile.Cells.DefaultVolumeSize,
			fixedSizeVolume,
		),
		CloudClient:        cloudClient,
		NodeClientFactory:  itzoClientFactory,
		Events:             eventSystem,
		ImageIdCache:       imageIdCache,
		CloudInitFile:      cloudInitFile,
		CertificateFactory: certFactory,
		CloudStatus:        cloudStatus,
		BootImageTags:      serverConfigFile.Cells.BootImageTags,
	}
	garbageController := &GarbageController{
		config: GarbageControllerConfig{
			CleanInstancesInterval:  60 * time.Second,
			CleanTerminatedInterval: 10 * time.Second,
		},
		podRegistry:  podRegistry,
		nodeRegistry: nodeRegistry,
		cloudClient:  cloudClient,
		controllerID: controllerID,
	}
	metricsController := &MetricsController{
		metricsRegistry: metricsRegistry,
		podLister:       podRegistry,
	}
	controllers := map[string]Controller{
		"PodController":     podController,
		"NodeController":    nodeController,
		"GarbageController": garbageController,
		"MetricsController": metricsController,
	}

	if azClient, ok := cloudClient.(*azure.AzureClient); ok {
		azureImageController := azure.NewImageController(
			controllerID, serverConfigFile.Cells.BootImageTags, azClient)
		controllers["ImageController"] = azureImageController
	}
	controllerManager := NewControllerManager(controllers)

	s := &InstanceProvider{
		Registries:        kv,
		Encoder:           api.VersioningCodec{},
		SystemQuit:        systemQuit,
		SystemWaitGroup:   systemWG,
		ItzoClientFactory: itzoClientFactory,
		cloudClient:       cloudClient,
		controllerManager: controllerManager,

		// Todo: cleanup these parameters after initial commit
		nodeName:           nodeName,
		internalIP:         internalIP,
		daemonEndpointPort: daemonEndpointPort,
		kubeletConfig:      serverConfigFile.Kubelet,
		startTime:          time.Now(),
		portManager:        portManager,
	}
	eventSystem.RegisterHandler(events.PodRunning, s)
	eventSystem.RegisterHandler(events.PodTerminated, s)
	eventSystem.RegisterHandler(events.PodUpdated, s)
	eventSystem.RegisterHandler(events.PodEjected, s)

	go controllerManager.Start()
	go controllerManager.WaitForShutdown(systemQuit, systemWG)

	controllerManager.StartControllers()

	if ctrl, ok := controllers["ImageController"]; ok {
		azureImageController := ctrl.(*azure.ImageController)
		klog.V(2).Infof("downloading Milpa node image to local Azure subscription (this could take a few minutes)")
		azureImageController.WaitForAvailable()
	}

	if debugServer {
		if err := s.setupDebugServer(); err != nil {
			return nil, err
		}
	}

	err = validateBootImageTags(
		serverConfigFile.Cells.BootImageTags, cloudClient)
	if err != nil {
		return nil, fmt.Errorf("failed to validate boot image tags")
	}

	return s, err
}

func (p *InstanceProvider) setupDebugServer() error {
	lis, err := net.Listen(defaultProtocol, fmt.Sprintf("127.0.0.1:%d", defaultPort))
	if err != nil {
		return fmt.Errorf("error setting up debug %s listener on port %d", defaultProtocol, defaultPort)
	}
	grpcServer := grpc.NewServer()
	clientapi.RegisterMilpaServer(grpcServer, p)
	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			klog.Errorln("Error returned from Serve:", err)
		}
	}()
	return nil
}

func getPortMappings(containers []v1.Container) []v1.ContainerPort {
	var portMappings []v1.ContainerPort
	for _, container := range containers {
		for _, pm := range container.Ports {
			if pm.ContainerPort > 0 && pm.HostPort > 0 {
				portMappings = append(portMappings, pm)
			}
		}
	}
	return portMappings
}

func (p *InstanceProvider) addOrRemovePodPortMappings(pod *v1.Pod, add bool) error {
	portMappings := getPortMappings(
		append(pod.Spec.InitContainers, pod.Spec.Containers...))
	if len(portMappings) == 0 {
		return nil
	}
	podIP := net.ParseIP(pod.Status.PodIP)
	if podIP.IsUnspecified() {
		return fmt.Errorf("empty pod IP for %q %+v", pod.Name, portMappings)
	}
	if add {
		klog.V(4).Infof("adding %q port mappings %+v", pod.Name, portMappings)
		return p.portManager.AddPodPortMappings(podIP.String(), portMappings)
	}
	klog.V(4).Infof("removing %q port mappings %+v", pod.Name, portMappings)
	p.portManager.RemovePodPortMappings(podIP.String())
	return nil
}

func (p *InstanceProvider) Handle(ev events.Event) error {
	milpaPod, ok := ev.Object.(*api.Pod)
	if !ok {
		klog.Errorf("event %v with unknown object", ev)
		return nil
	}
	pod, err := p.milpaToK8sPod(milpaPod)
	if err != nil {
		klog.Errorf("converting milpa pod %s: %v", milpaPod.Name, err)
		return nil
	}
	if ev.Status == events.PodUpdated {
		if milpaPod.Status.Phase == api.PodRunning &&
			pod.Status.PodIP != "" {
			// Pod is up and running, let's set up its hostport mappings.
			if err := p.addOrRemovePodPortMappings(pod, true); err != nil {
				klog.Warningf("adding hostports %q: %v", milpaPod.Name, err)
			}
		} else if api.IsTerminalPodPhase(milpaPod.Status.Phase) {
			// Remove port mappings if pod has been terminated or stopped.
			if err := p.addOrRemovePodPortMappings(pod, false); err != nil {
				klog.Warningf("removing hostports %q: %v", milpaPod.Name, err)
			}
		}
	}
	klog.V(4).Infof("milpa pod %q event %v", milpaPod.Name, ev)
	p.notifier(pod)
	return nil
}

func (p *InstanceProvider) Stop() {
	quitTimeout := time.Duration(10)
	waitGroupDone := make(chan struct{})
	go waitForWaitGroup(p.SystemWaitGroup, waitGroupDone)
	select {
	case <-waitGroupDone:
		return
	case <-time.After(time.Second * quitTimeout):
		klog.Errorf(
			"Loops were still running after %d seconds, forcing exit",
			quitTimeout)
		return
	}
}

func waitForWaitGroup(wg *sync.WaitGroup, waitGroupDone chan struct{}) {
	wg.Wait()
	klog.V(2).Info("all controllers have exited")
	waitGroupDone <- struct{}{}
}

func filterEventList(eventList *api.EventList) *api.EventList {
	if eventList != nil && len(eventList.Items) > MaxEventListSize {
		// Take the most recent MaxEventListSize items
		sort.Slice(eventList.Items, func(i, j int) bool {
			return eventList.Items[i].CreationTimestamp.Before(eventList.Items[j].CreationTimestamp)
		})
		size := len(eventList.Items)
		start := size - MaxEventListSize
		eventList.Items = eventList.Items[start:]
	}
	return eventList
}

func filterReplyObject(obj api.MilpaObject) api.MilpaObject {
	switch v := obj.(type) {
	case *api.EventList:
		return filterEventList(v)
	}
	return obj
}

func (p *InstanceProvider) getPodRegistry() *registry.PodRegistry {
	reg := p.Registries["Pod"]
	return reg.(*registry.PodRegistry)
}

func (p *InstanceProvider) getNodeRegistry() *registry.NodeRegistry {
	reg := p.Registries["Node"]
	return reg.(*registry.NodeRegistry)
}

func (p *InstanceProvider) getMetricsRegistry() *registry.MetricsRegistry {
	reg := p.Registries["Metric"]
	return reg.(*registry.MetricsRegistry)
}

func (p *InstanceProvider) CreatePod(ctx context.Context, pod *v1.Pod) error {
	ctx, span := trace.StartSpan(ctx, "CreatePod")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, pod.Namespace, nameKey, pod.Name)
	klog.V(5).Infof("CreatePod %q", pod.Name)
	milpaPod, err := p.k8sToMilpaPod(pod)
	if err != nil {
		klog.Errorf("CreatePod %q: %v", pod.Name, err)
		return err
	}
	podRegistry := p.getPodRegistry()
	_, err = podRegistry.CreatePod(milpaPod)
	if err != nil {
		klog.Errorf("CreatePod %q: %v", pod.Name, err)
		return err
	}
	p.notifier(pod)
	return nil
}

func (p *InstanceProvider) UpdatePod(ctx context.Context, pod *v1.Pod) error {
	ctx, span := trace.StartSpan(ctx, "UpdatePod")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, pod.Namespace, nameKey, pod.Name)
	klog.V(5).Infof("UpdatePod %q", pod.Name)
	milpaPod, err := p.k8sToMilpaPod(pod)
	if err != nil {
		klog.Errorf("UpdatePod %q: %v", pod.Name, err)
		return err
	}
	podRegistry := p.getPodRegistry()
	_, err = podRegistry.UpdatePodSpecAndLabels(milpaPod)
	if err != nil {
		klog.Errorf("UpdatePod %q: %v", pod.Name, err)
		return err
	}
	p.notifier(pod)
	return nil
}

func (p *InstanceProvider) DeletePod(ctx context.Context, pod *v1.Pod) (err error) {
	ctx, span := trace.StartSpan(ctx, "DeletePod")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, pod.Namespace, nameKey, pod.Name)
	klog.V(5).Infof("DeletePod %q", pod.Name)
	milpaPod, err := p.k8sToMilpaPod(pod)
	if err != nil {
		klog.Errorf("DeletePod %q: %v", pod.Name, err)
		return err
	}
	podRegistry := p.getPodRegistry()
	_, err = podRegistry.Delete(milpaPod.Name)
	if err != nil {
		klog.Errorf("DeletePod %q: %v", pod.Name, err)
		return err
	}
	p.notifier(pod)
	return nil
}

func (p *InstanceProvider) GetPod(ctx context.Context, namespace, name string) (*v1.Pod, error) {
	ctx, span := trace.StartSpan(ctx, "GetPod")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, namespace, nameKey, name)
	klog.V(5).Infof("GetPod %q", name)
	podRegistry := p.getPodRegistry()
	milpaPod, err := podRegistry.GetPod(util.WithNamespace(namespace, name))
	if err != nil {
		if err == store.ErrKeyNotFound {
			return nil, errdefs.NotFoundf("pod %s/%s is not found", namespace, name)
		}
		klog.Errorf("GetPod %q: %v", name, err)
		return nil, err
	}
	pod, err := p.milpaToK8sPod(milpaPod)
	if err != nil {
		klog.Errorf("GetPod %q: %v", name, err)
		return nil, err
	}
	return pod, nil
}

func (p *InstanceProvider) GetPodStatus(ctx context.Context, namespace, name string) (*v1.PodStatus, error) {
	ctx, span := trace.StartSpan(ctx, "GetPodStatus")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, namespace, nameKey, name)
	klog.V(5).Infof("GetPodStatus %q", name)
	podRegistry := p.getPodRegistry()
	milpaPod, err := podRegistry.GetPod(util.WithNamespace(namespace, name))
	if err != nil {
		klog.Errorf("GetPodStatus %q: %v", name, err)
		return nil, err
	}
	pod, err := p.milpaToK8sPod(milpaPod)
	if err != nil {
		klog.Errorf("GetPodStatus %q: %v", name, err)
		return nil, err
	}
	return &pod.Status, nil
}

func (p *InstanceProvider) GetPods(ctx context.Context) ([]*v1.Pod, error) {
	ctx, span := trace.StartSpan(ctx, "GetPods")
	defer span.End()
	klog.V(5).Infof("GetPods")
	podRegistry := p.getPodRegistry()
	milpaPods, err := podRegistry.ListPods(func(pod *api.Pod) bool {
		return true
	})
	if err != nil {
		klog.Errorf("GetPods: %v", err)
		return nil, err
	}
	pods := make([]*v1.Pod, len(milpaPods.Items))
	for i, milpaPod := range milpaPods.Items {
		pods[i], err = p.milpaToK8sPod(milpaPod)
		if err != nil {
			klog.Errorf("GetPods: %v", err)
			return nil, err
		}
	}
	return pods, nil
}

func (p *InstanceProvider) ConfigureNode(ctx context.Context, n *v1.Node) {
	ctx, span := trace.StartSpan(ctx, "ConfigureNode")
	defer span.End()
	klog.V(5).Infof("ConfigureNode")
	n.Status.Capacity = p.capacity()
	n.Status.Allocatable = p.capacity()
	n.Status.Conditions = p.nodeConditions()
	n.Status.Addresses = p.nodeAddresses()
	n.Status.DaemonEndpoints = p.nodeDaemonEndpoints()
	n.Status.NodeInfo.OperatingSystem = "Linux"
	n.Status.NodeInfo.Architecture = "amd64"
}

func (p *InstanceProvider) capacity() v1.ResourceList {
	return v1.ResourceList{
		"cpu":    p.kubeletConfig.CPU,
		"memory": p.kubeletConfig.Memory,
		"pods":   p.kubeletConfig.Pods,
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
