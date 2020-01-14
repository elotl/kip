package server

import (
	"fmt"
	"io"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/containerd/containerd/log"
	"github.com/docker/libkv/store"
	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/api/validation"
	"github.com/elotl/cloud-instance-provider/pkg/certs"
	"github.com/elotl/cloud-instance-provider/pkg/etcd"
	"github.com/elotl/cloud-instance-provider/pkg/nodeclient"
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
	"github.com/golang/glog"
	vkapi "github.com/virtual-kubelet/virtual-kubelet/node/api"
	"github.com/virtual-kubelet/virtual-kubelet/trace"
	"golang.org/x/net/context"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	stats "k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1"
)

const (
	// Values used in tracing as attribute keys.
	namespaceKey          = "namespace"
	nameKey               = "name"
	containerNameKey      = "containerName"
	etcdClusterRegionPath = "milpa/cluster/region"
	kubernetesPodKey      = "elotl.co/kubernetes-pod"
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
}

func validateWriteToEtcd(client *etcd.SimpleEtcd) error {
	glog.Info("Validating write access to etcd (will block until we can connect)")
	wo := &store.WriteOptions{
		IsDir: false,
		TTL:   2 * time.Second,
	}

	err := client.PutNoTimeout("/milpa/startup", []byte("OK"), wo)
	if err != nil {
		return err
	}
	glog.Info("Write to etcd successful")
	return nil
}

func setupEtcd(configFile, dataDir string, quit <-chan struct{}, wg *sync.WaitGroup) (*etcd.SimpleEtcd, error) {
	// if we have client endpoints, don't start the server. This could
	// change in the future if we want the embedded server to join
	// existing etcd server, but, for now just don't start it.
	var client *etcd.SimpleEtcd
	glog.Infof("Starting Internal Etcd")
	etcdServer := etcd.EtcdServer{
		ConfigFile: configFile,
		DataDir:    dataDir,
	}
	err := etcdServer.Start(quit, wg)
	if err != nil {
		return nil, util.WrapError(
			err, "Error creating internal etcd storage backend")
	}
	client = etcdServer.Client
	err = validateWriteToEtcd(client)
	if err != nil {
		return nil, util.WrapError(err, "Fatal Error: Could not write to etcd")
	}
	return client, err
}

func ensureRegionUnchanged(etcdClient *etcd.SimpleEtcd, region string) error {
	glog.Infof("Ensuring region has not changed")
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
				"Please delete all cluster resources and rename your cluster.",
			savedRegion, region)
	}
	return nil
}

// InstanceProvider should implement node.PodLifecycleHandler
func NewInstanceProvider(configFilePath, nodeName, internalIP string, daemonEndpointPort int32, systemQuit <-chan struct{}) (*InstanceProvider, error) {
	systemWG := &sync.WaitGroup{}

	serverConfigFile, err := ParseConfig(configFilePath)
	if err != nil {
		glog.Errorf("Loading Config file (%s) failed with error: %s",
			configFilePath, err.Error())
		os.Exit(1)
	}
	errs := validateServerConfigFile(serverConfigFile)
	if len(errs) > 0 {
		return nil, fmt.Errorf("invalid server config: %v", errs.ToAggregate())
	}

	// todo: systemQuit should get passed in...
	//systemQuit, systemWG := SetupSignalHandler()
	etcdClient, err := setupEtcd(
		serverConfigFile.Etcd.Internal.ConfigFile,
		serverConfigFile.Etcd.Internal.DataDir,
		systemQuit,
		systemWG,
	)
	if err != nil {
		glog.Errorf("Etcd error: %s", err)
		os.Exit(1)
	}
	controllerID, err := getControllerID(etcdClient)
	if err != nil {
		glog.Errorf("Controller ID error: %s", err)
		os.Exit(1)
	}
	if serverConfigFile.Testing.ControllerID != "" {
		controllerID = serverConfigFile.Testing.ControllerID
	}
	nametag := serverConfigFile.Cells.Nametag
	if nametag == "" {
		nametag = controllerID
	}

	glog.Infof("ControllerID: %s", controllerID)

	certFactory, err := certs.New(etcdClient)
	if err != nil {
		glog.Errorf("Error setting up certificate factory: %v", err)
		os.Exit(1)
	}

	cloudClient, err := ConfigureCloud(serverConfigFile, controllerID, nametag)
	if err != nil {
		glog.Errorln("Error configuring cloud client", err)
		os.Exit(1)
	}
	cloudRegion := cloudClient.GetAttributes().Region
	err = ensureRegionUnchanged(etcdClient, cloudRegion)
	if err != nil {
		glog.Errorln("Error ensuring Milpa region is unchanged", err)
		os.Exit(1)
	}
	clientCert, err := certFactory.CreateClientCert()
	if err != nil {
		glog.Errorf("Error creating node client certificate: %v", err)
		os.Exit(1)
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
		glog.Errorf("Error setting up instance selector %s", err)
		os.Exit(1)
	}
	// Ugly: need to do validation of this field after we have setup
	// the instanceselector
	errs = validation.ValidateInstanceType(serverConfigFile.Cells.DefaultInstanceType, field.NewPath("nodes.defaultInstanceType"))
	if len(errs) > 0 {
		glog.Errorf("Error validating server.yml: %v", errs.ToAggregate())
		os.Exit(1)
	}

	glog.Infof("Setting up events")
	eventSystem := events.NewEventSystem(systemQuit, systemWG)

	glog.Infof("Setting up registry")
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
		podRegistry:       podRegistry,
		logRegistry:       logRegistry,
		metricsRegistry:   metricsRegistry,
		nodeLister:        nodeRegistry,
		nodeDispenser:     nodeDispenser,
		nodeClientFactory: itzoClientFactory,
		events:            eventSystem,
		cloudClient:       cloudClient,
		controllerID:      controllerID,
		nametag:           nametag,
		lastStatusReply:   conmap.NewStringTimeTime(),
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
	}

	go controllerManager.Start()
	go controllerManager.WaitForShutdown(systemQuit, systemWG)

	controllerManager.StartControllers()

	if ctrl, ok := controllers["ImageController"]; ok {
		azureImageController := ctrl.(*azure.ImageController)
		glog.Infof("Downloading Milpa node image to local Azure subscription (this could take a few minutes)")
		azureImageController.WaitForAvailable()
	}

	err = validateBootImageTags(
		serverConfigFile.Cells.BootImageTags, cloudClient)
	if err != nil {
		glog.Errorf("Failed to validate boot image tags.")
		os.Exit(1)
	}

	return s, err
}

func (ip *InstanceProvider) Stop() {
	quitTimeout := time.Duration(10)
	waitGroupDone := make(chan struct{})
	go waitForWaitGroup(ip.SystemWaitGroup, waitGroupDone)
	select {
	case <-waitGroupDone:
		return
	case <-time.After(time.Second * quitTimeout):
		glog.Errorf(
			"Loops were still running after %d seconds, forcing exit",
			quitTimeout)
		return
	}
}

func waitForWaitGroup(wg *sync.WaitGroup, waitGroupDone chan struct{}) {
	wg.Wait()
	glog.Info("All controllers have exited")
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

func ContainerToUnit(container v1.Container) api.Unit {
	unit := api.Unit{
		Name:       container.Name,
		Image:      container.Image,
		Command:    container.Command,
		Args:       container.Args,
		WorkingDir: container.WorkingDir,
	}
	unit.Env = make([]api.EnvVar, len(container.Env))
	for i, e := range container.Env {
		unit.Env[i] = api.EnvVar{
			Name:  e.Name,
			Value: e.Value,
		}
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
	//container.VolumeMounts,
	//container.EnvFrom,
	//container.Ports,
	return unit
}

func UnitToContainer(unit api.Unit, container *v1.Container) v1.Container {
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
	return *container
}

func K8sToMilpaPod(pod *v1.Pod) (*api.Pod, error) {
	milpapod := api.NewPod()
	milpapod.Name = util.WithNamespace(pod.Namespace, pod.Name)
	milpapod.Namespace = pod.Namespace
	milpapod.Labels = pod.Labels
	milpapod.Annotations = pod.Annotations
	podBytes, err := pod.Marshal()
	if err != nil {
		return nil, err
	}
	milpapod.Annotations[kubernetesPodKey] = string(podBytes)
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
		initUnit := ContainerToUnit(initContainer)
		milpapod.Spec.InitUnits = append(milpapod.Spec.InitUnits, initUnit)
	}
	for _, container := range pod.Spec.Containers {
		unit := ContainerToUnit(container)
		milpapod.Spec.Units = append(milpapod.Spec.Units, unit)
	}
	return milpapod, nil
}

func MilpaToK8sPod(milpaPod *api.Pod) (*v1.Pod, error) {
	pod := &v1.Pod{}
	annotations := milpaPod.Annotations
	if podStr, exists := annotations[kubernetesPodKey]; exists {
		err := pod.Unmarshal([]byte(podStr))
		if err != nil {
			return nil, err
		}
	}
	name, namespace := util.SplitNamespaceAndName(milpaPod.Name)
	pod.Name = name
	pod.Namespace = namespace
	pod.Labels = milpaPod.Labels
	pod.Annotations = milpaPod.Annotations
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
		initContainer = UnitToContainer(initUnit, ptr)
		pod.Spec.InitContainers = append(pod.Spec.InitContainers, initContainer)
	}
	for _, unit := range milpaPod.Spec.Units {
		container, exists := containerMap[unit.Name]
		ptr := &container
		if !exists {
			ptr = nil
		}
		container = UnitToContainer(unit, ptr)
		pod.Spec.Containers = append(pod.Spec.Containers, container)
	}
	return pod, nil
}

func (p *InstanceProvider) getPodRegistry() *registry.PodRegistry {
	reg := p.Registries["Pod"]
	return reg.(*registry.PodRegistry)
}

func (p *InstanceProvider) CreatePod(ctx context.Context, pod *v1.Pod) error {
	ctx, span := trace.StartSpan(ctx, "CreatePod")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, pod.Namespace, nameKey, pod.Name)
	log.G(ctx).Infof("CreatePod %q", pod.Name)
	milpaPod, err := K8sToMilpaPod(pod)
	if err != nil {
		log.G(ctx).Errorf("CreatePod %q: %v", pod.Name, err)
		return err
	}
	podRegistry := p.getPodRegistry()
	_, err = podRegistry.CreatePod(milpaPod)
	if err != nil {
		log.G(ctx).Errorf("CreatePod %q: %v", pod.Name, err)
		return err
	}
	p.notifier(pod)
	return nil
}

func (p *InstanceProvider) UpdatePod(ctx context.Context, pod *v1.Pod) error {
	ctx, span := trace.StartSpan(ctx, "UpdatePod")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, pod.Namespace, nameKey, pod.Name)
	log.G(ctx).Infof("UpdatePod %q", pod.Name)
	milpaPod, err := K8sToMilpaPod(pod)
	if err != nil {
		log.G(ctx).Errorf("UpdatePod %q: %v", pod.Name, err)
		return err
	}
	podRegistry := p.getPodRegistry()
	_, err = podRegistry.UpdatePodSpecAndLabels(milpaPod)
	if err != nil {
		log.G(ctx).Errorf("UpdatePod %q: %v", pod.Name, err)
		return err
	}
	p.notifier(pod)
	return nil
}

func (p *InstanceProvider) DeletePod(ctx context.Context, pod *v1.Pod) (err error) {
	ctx, span := trace.StartSpan(ctx, "DeletePod")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, pod.Namespace, nameKey, pod.Name)
	log.G(ctx).Infof("DeletePod %q", pod.Name)
	milpaPod, err := K8sToMilpaPod(pod)
	if err != nil {
		log.G(ctx).Errorf("DeletePod %q: %v", pod.Name, err)
		return err
	}
	podRegistry := p.getPodRegistry()
	_, err = podRegistry.Delete(milpaPod.Name)
	if err != nil {
		log.G(ctx).Errorf("DeletePod %q: %v", pod.Name, err)
		return err
	}
	p.notifier(pod)
	return nil
}

func (p *InstanceProvider) GetPod(ctx context.Context, namespace, name string) (*v1.Pod, error) {
	ctx, span := trace.StartSpan(ctx, "GetPod")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, namespace, nameKey, name)
	log.G(ctx).Infof("GetPod %q", name)
	podRegistry := p.getPodRegistry()
	milpaPod, err := podRegistry.GetPod(util.WithNamespace(namespace, name))
	if err != nil {
		log.G(ctx).Errorf("GetPod %q: %v", name, err)
		return nil, err
	}
	pod, err := MilpaToK8sPod(milpaPod)
	if err != nil {
		log.G(ctx).Errorf("GetPod %q: %v", name, err)
		return nil, err
	}
	return pod, nil
}

func (p *InstanceProvider) GetContainerLogs(ctx context.Context, namespace, podName, containerName string, opts vkapi.ContainerLogOpts) (io.ReadCloser, error) {
	ctx, span := trace.StartSpan(ctx, "GetContainerLogs")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, namespace, nameKey, podName, containerNameKey, containerName)
	log.G(ctx).Infof("GetContainerLogs %q", podName)
	return nil, fmt.Errorf("not implemented")
}

func (p *InstanceProvider) RunInContainer(ctx context.Context, namespace, podName, containerName string, cmd []string, attach vkapi.AttachIO) error {
	ctx, span := trace.StartSpan(ctx, "RunInContainer")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, namespace, nameKey, podName, containerNameKey, containerName)
	log.G(ctx).Infof("RunInContainer %q %v", podName, cmd)
	return fmt.Errorf("not implemented")
}

func (p *InstanceProvider) GetPodStatus(ctx context.Context, namespace, name string) (*v1.PodStatus, error) {
	ctx, span := trace.StartSpan(ctx, "GetPodStatus")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, namespace, nameKey, name)
	log.G(ctx).Infof("GetPodStatus %q", name)
	podRegistry := p.getPodRegistry()
	milpaPod, err := podRegistry.GetPod(util.WithNamespace(namespace, name))
	if err != nil {
		log.G(ctx).Errorf("GetPodStatus %q: %v", name, err)
		return nil, err
	}
	pod, err := MilpaToK8sPod(milpaPod)
	if err != nil {
		log.G(ctx).Errorf("GetPodStatus %q: %v", name, err)
		return nil, err
	}
	return &pod.Status, nil
}

func (p *InstanceProvider) GetPods(ctx context.Context) ([]*v1.Pod, error) {
	ctx, span := trace.StartSpan(ctx, "GetPods")
	defer span.End()
	log.G(ctx).Infof("GetPods")
	podRegistry := p.getPodRegistry()
	milpaPods, err := podRegistry.ListPods(func(pod *api.Pod) bool {
		if pod.Status.Phase == api.PodRunning {
			return true
		}
		return false
	})
	if err != nil {
		log.G(ctx).Errorf("GetPods: %v", err)
		return nil, err
	}
	pods := make([]*v1.Pod, len(milpaPods.Items))
	for i, milpaPod := range milpaPods.Items {
		pods[i], err = MilpaToK8sPod(milpaPod)
		if err != nil {
			log.G(ctx).Errorf("GetPods: %v", err)
			return nil, err
		}
	}
	return pods, nil
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
