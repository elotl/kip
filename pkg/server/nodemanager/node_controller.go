/*
Copyright 2020 Elotl Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nodemanager

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/docker/libkv/store"
	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/certs"
	"github.com/elotl/kip/pkg/nodeclient"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/server/events"
	"github.com/elotl/kip/pkg/server/registry"
	"github.com/elotl/kip/pkg/util"
	"github.com/elotl/kip/pkg/util/cloudinitfile"
	"github.com/elotl/kip/pkg/util/stats"
	"github.com/elotl/kip/pkg/util/timeoutmap"
	"k8s.io/klog"
)

// Making these vars makes it easier testing
// non-const timeouts were endorsed by Mitchell Hashimoto
var (
	BootTimeout         time.Duration = 300 * time.Second
	HealthyTimeout      time.Duration = 90 * time.Second
	HealthcheckPause    time.Duration = 5 * time.Second
	SpotRequestPause    time.Duration = 60 * time.Second
	BootImages          []cloud.Image = []cloud.Image{}
	MaxBootPerIteration int           = 10
	itzoDir             string        = "/tmp/itzo"
)

// when configuring these intervals we want the following constraints
// to be satisfied:
//
// 1. The pool interval should be longer than the heartbeat interval
// 2. The heartbeat interval should be longer than the heartbeat
// client timeout.
type NodeControllerConfig struct {
	PoolInterval           time.Duration
	HeartbeatInterval      time.Duration
	ReaperInterval         time.Duration
	ItzoVersion            string
	ItzoURL                string
	CellConfig             map[string]string
	UseCloudParameterStore bool
	DefaultIAMPermissions  string
}

type NodeController struct {
	Config             NodeControllerConfig
	NodeRegistry       *registry.NodeRegistry
	LogRegistry        *registry.LogRegistry
	PodReader          registry.PodLister
	NodeDispenser      *NodeDispenser
	NodeScaler         ScalingAlgorithm
	CloudClient        cloud.CloudClient
	NodeClientFactory  nodeclient.ItzoClientFactoryer
	Events             *events.EventSystem
	PoolLoopTimer      *stats.LoopTimer
	ImageIdCache       *timeoutmap.TimeoutMap
	CloudInitFile      *cloudinitfile.File
	CertificateFactory *certs.CertificateFactory
	BootLimiter        *InstanceBootLimiter
	BootImageSpec      cloud.BootImageSpec
}

func (c *NodeController) Start(quit <-chan struct{}, wg *sync.WaitGroup) {
	c.PoolLoopTimer = &stats.LoopTimer{}
	c.StopCreatingNodes()
	go c.ResumeWaits()
	go c.runHeartbeatsLoop(quit, wg)
	go c.reaperLoop(quit, wg)
	nodeBindingsUpdate := make(chan map[string]string)
	go c.updateBufferedNodesLoop(quit, wg, nodeBindingsUpdate)
	go c.dispatchNodesLoop(quit, wg, nodeBindingsUpdate)
	go c.ImageIdCache.Start(30 * time.Second)
}

func (c *NodeController) Dump() []byte {
	t := c.PoolLoopTimer.Copy()
	b, err := json.MarshalIndent(*t, "", "    ")
	if err != nil {
		klog.Errorln("Error dumping data from NodeController", err)
		return nil
	}
	return b
}

func (c *NodeController) updateBufferedNodesLoop(quit <-chan struct{}, wg *sync.WaitGroup, nodeBindingsUpdate chan map[string]string) {
	wg.Add(1)
	defer wg.Done()

	ticker := time.NewTicker(c.Config.PoolInterval)
	defer ticker.Stop()
	for {
		select {
		case <-quit:
			return
		default:
			select {
			case <-ticker.C:
				c.PoolLoopTimer.StartLoop()
				updatedBindings, err := c.doPoolsCalculation()
				if err != nil {
					klog.Errorln("Error adjusting node pools", err.Error())
				} else {
					nodeBindingsUpdate <- updatedBindings
				}
				c.PoolLoopTimer.EndLoop()
			case <-quit:
				return
			}
		}
	}
}

// Ensure that 2 instances of this don't run concurrently
// also ensure the node creation in the registry happens
// in this goroutine, otherwise our node scaling calculations
// will get thrown off
func (c *NodeController) doPoolsCalculation() (map[string]string, error) {
	nodes, err := c.NodeRegistry.ListNodes(registry.MatchAllNodes)
	if err != nil {
		return nil, util.WrapError(err, "Couldn't list nodes for pool calculation")
	}

	pods, err := c.PodReader.ListPods(func(p *api.Pod) bool {
		return registry.MatchAllLivePods(p)
	})
	if err != nil {
		return nil, util.WrapError(err, "Couldn't list pods for pool calculation")
	}

	// If we can't get the boot image, just use the old value for the image
	// need to update this function first so we start returning a slice of
	// images based on the number of image filters within the bootspec
	newBootImages, err := c.imageSpecToImage(c.BootImageSpec)
	if err != nil {
		if len(newBootImages) < len(c.BootImageSpec)-1 {
			return nil, util.WrapError(err, "Could not get latest boot images")
		}
		for _, bootImg := range BootImages {
			if bootImg.ID == "" {
				return nil, util.WrapError(err, "Could not get latest boot image")
			} else {
				klog.Warningf("Could not get latest boot image: %s, using stored value for boot image: %v", err, bootImg)
				newBootImages = append(newBootImages, bootImg)
			}
		}
	}
	BootImages = newBootImages

	for _, bootImg := range BootImages {
		if bootImg.ID == "" {
			return nil, fmt.Errorf("can not create create new nodes: empty value for machine image.  Please ensure boot image spec maps to a machine image: %v", c.BootImageSpec)
		}
	}
	startNodes, stopNodes, podNodeMap := c.NodeScaler.Compute(nodes.Items, pods.Items)
	if podNodeMap == nil {
		return nil, fmt.Errorf("Error computing new node pools, this is likely a problem with the DB. Not updating pod-node bindings")
	}
	c.startNodes(startNodes, BootImages)
	for _, node := range stopNodes {
		err := c.stopSingleNode(node)
		if err != nil {
			klog.Warningln("Error stopping single node", err)
			continue
		}
	}
	return podNodeMap, nil
}

func (c *NodeController) getInstanceCloudInit() error {
	c.CloudInitFile.ResetInstanceData()
	if c.Config.UseCloudParameterStore {
		return nil
	}

	// Use instance metadata to distribute parameters and configuration to
	// instances.
	params, err := getInstanceParameters(c.CertificateFactory, InstanceConfig{
		ItzoURL:     c.Config.ItzoURL,
		ItzoVersion: c.Config.ItzoVersion,
		CellConfig:  c.Config.CellConfig,
	})
	if err != nil {
		return util.WrapError(err, "getInstanceParameters() for instance metadata failed")
	}
	for key, value := range params {
		c.CloudInitFile.AddKipFile(value, key, "0400")
	}
	return nil
}

func (c *NodeController) getCloudInitContents() (string, error) {
	err := c.getInstanceCloudInit()
	if err != nil {
		return "", util.WrapError(
			err, "Error creating Kip instance keys for cloud-init data")
	}
	cloudInitData, err := c.CloudInitFile.Contents()
	if err != nil {
		return "", util.WrapError(err, "Error creating Kip cloud-init contents")
	}
	metadata := base64.StdEncoding.EncodeToString(cloudInitData)
	return metadata, nil
}

func (c *NodeController) startNodes(nodes []*api.Node, images []cloud.Image) {
	if len(nodes) <= 0 {
		return
	}
	metadata, err := c.getCloudInitContents()
	if err != nil {
		klog.Errorf("Error creating node metadata: %s", err)
		return
	}
	// Randomize boot order to prevent getting stuck with 10 nodes at
	// the start of the boot list that can't be booted for some reason
	if len(nodes) > MaxBootPerIteration {
		for i := range nodes {
			j := rand.Intn(i + 1)
			nodes[i], nodes[j] = nodes[j], nodes[i]
		}
	}
	for i, newNode := range nodes {
		if i >= MaxBootPerIteration {
			klog.V(2).Infof("Rate limiting start requests to %d per iteration", MaxBootPerIteration)
			break
		}
		newNode, err := c.NodeRegistry.CreateNode(newNode)
		if err != nil {
			klog.Errorf("Error creating node in registry: %v", err)
			continue
		}
		image, found := getImageForInstance(newNode.Spec.InstanceType, images)
		if !found {
			klog.Errorf("Error finding image for instance type: %s", newNode.Spec.InstanceType)
			return
		}
		go c.startSingleNode(newNode, image, metadata)
	}
}

// TODO create a map[string]string of instances where the key will be the instance
// type and the value will be its corresponding architecture. This will allow us
// to match the instance with the proper image for instantiation
func getImageForInstance(instType string, images []cloud.Image) (cloud.Image, bool) {
	for _, img := range images {
		isMacInst := strings.HasPrefix(instType, "mac")
		if isMacInst && img.Architecture == cloud.Arch_x86_64_mac {
			return img, true
		}
		if !isMacInst && img.Architecture == cloud.Arch_x86_64 {
			return img, true
		}
	}
	return cloud.Image{}, false
}

func (c *NodeController) handleStartNodeError(node *api.Node, err error, isSpot bool) {
	switch err.(type) {
	case *cloud.NoCapacityError:
		c.BootLimiter.AddUnavailableInstance(node.Spec.InstanceType, isSpot)
	case *cloud.UnsupportedInstanceError:
		// It's possible we should eventually kill the pod associated
		// with this but I hesitate to do that, instead lets push that
		// off to the operator for now.
		c.BootLimiter.AddUnavailableInstance(node.Spec.InstanceType, isSpot)
	}
}

func (c *NodeController) startSingleNode(node *api.Node, image cloud.Image, cloudInitData string) error {
	var (
		instanceID string
		err        error
	)
	if node.Spec.Spot {
		instanceID, err = c.CloudClient.StartSpotNode(node, image, cloudInitData, c.Config.DefaultIAMPermissions)
	} else {
		instanceID, err = c.CloudClient.StartNode(node, image, cloudInitData, c.Config.DefaultIAMPermissions)
	}
	if err != nil {
		c.handleStartNodeError(node, err, false)
		klog.Errorf("Error in node start: %v", err)
		_, regError := c.NodeRegistry.PurgeNode(node)
		if regError != nil {
			klog.Errorf("Error marking node %s terminated after failed start: %s",
				node.Name, regError.Error())
		}
		return util.WrapError(err, "Error starting node")
	}
	node.Status.InstanceID = instanceID
	return c.finishNodeStart(node)
}

func (c *NodeController) finishNodeStart(node *api.Node) error {
	instanceID := node.Status.InstanceID

	if c.Config.UseCloudParameterStore {
		params, err := getMarshalledInstanceParameters(c.CertificateFactory, InstanceConfig{
			ItzoURL:     c.Config.ItzoURL,
			ItzoVersion: c.Config.ItzoVersion,
			CellConfig:  c.Config.CellConfig,
		})
		if err != nil {
			_ = c.stopSingleNode(node)
			wrapErr := fmt.Errorf("getMarshalledInstanceParameters() for %s: %v", instanceID, err)
			klog.Errorf("%v", wrapErr)
			return wrapErr
		}
		// We could add parameters one by one, but there are limits on the
		// number of parameters per account in AWS. Let's keep it compact, and
		// pass everything in one go, letting the cloud backend split it into
		// chunks if necessary.
		err = c.CloudClient.AddInstanceParameter(instanceID, "config", params, true)
		if err != nil {
			_ = c.stopSingleNode(node)
			wrapErr := fmt.Errorf("AddInstanceParameter() for %s: %v", instanceID, err)
			klog.Errorf("%v", wrapErr)
			return wrapErr
		}
	}

	node.Status.Phase = api.NodeCreated
	_, _ = c.NodeRegistry.UpdateStatus(node)
	c.Events.Emit(events.NodeCreated, "node-created", node, "")
	// todo: we know the instance is running, we could just do
	// a describe instance here...
	addresses, err := c.CloudClient.WaitForRunning(node)
	if err != nil {
		klog.V(2).Infof("Unhealthy wait for running, terminating node: %s", node.Name)
		_ = c.stopSingleNode(node)
		return util.WrapError(err, "Error waiting for node to be running")
	}
	c.Events.Emit(events.NodeRunning, "node-controller", node, "")
	node.Status.Addresses = addresses
	_, _ = c.NodeRegistry.UpdateStatus(node)
	return c.waitForAvailableOrTerminate(node, BootTimeout)
}

func (c *NodeController) stopSingleNode(node *api.Node) error {
	// to keep counts in sync, don't move this inside the goroutine
	klog.V(2).Infof("Stopping node: %s", node.Name)

	node.Status.Phase = api.NodeTerminating
	_, err := c.NodeRegistry.UpdateStatus(node)
	if err != nil {
		msg := fmt.Sprintf("Error stopping node: Could not set node phase to Terminating: %v", node)
		err = util.WrapError(err, msg)
		return err
	}
	c.NodeClientFactory.DeleteClient(node.Status.Addresses)
	go func(n *api.Node) {
		err = c.CloudClient.StopInstance(n.Status.InstanceID)
		if err != nil {
			klog.Warningf("stopping instance %s: %v", n.Status.InstanceID, err)
		}
		err = c.CloudClient.DeleteInstanceParameter(n.Status.InstanceID, "")
		if err != nil {
			klog.Warningf("deleting parameters for instance %s: %v", n.Status.InstanceID, err)
		}
		_, err := c.NodeRegistry.PurgeNode(node)
		if err != nil {
			klog.Errorf("Could not mark node %s as terminated: %v", n.Name, err)
		}
	}(node)
	return nil
}

func (c *NodeController) runHeartbeatsLoop(quit <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	ticker := time.NewTicker(c.Config.HeartbeatInterval)
	defer ticker.Stop()
	heartbeats := make(chan string)
	LastHeartbeat := make(map[string]time.Time)
	for {
		// Might want to add some timing jitter

		// Note: node health doesn't have to to be written to
		// the datastore.  When a node is deemed unhealthy, we
		// should terminate the node, it's likely we should
		// do something to reschedule the pod on that node as well
		select {
		case <-ticker.C:
			allNodes, err := c.NodeRegistry.ListNodes(registry.MatchAllNodes)
			if err != nil {
				klog.Errorf("Error listing nodes for heartbeat: %s", err.Error())
				// Hack attack....
				// The period of this ticker is pretty quick, if our
				// ListNodes times out then we will not see the quit
				// for a long time.  See if we need to quit.
				select {
				case <-quit:
					return
				default:
					continue
				}
			}
			if err := c.sendOutHeartbeats(allNodes, heartbeats); err != nil {
				klog.Error(err.Error())
			}
			if err := c.markUnhealthyNodes(allNodes, LastHeartbeat); err != nil {
				klog.Error(err.Error())
			}
			pruneHeartbeats(allNodes, LastHeartbeat)
		case nodeName := <-heartbeats:
			LastHeartbeat[nodeName] = time.Now().UTC()
		case <-quit:
			return
		}
	}
}

func (c *NodeController) sendOutHeartbeats(allNodes *api.NodeList, heartbeats chan string) error {
	// Once a node is claimed, its status will be monitored by the pod
	// controller.
	nodes := filterNodes(allNodes, func(n *api.Node) bool {
		// We purposefully wait on created AND wait for that to be running
		// as well.
		return ((n.Status.Phase == api.NodeCreated &&
			n.Status.Addresses != nil) ||
			n.Status.Phase == api.NodeAvailable)
	})
	for _, n := range nodes {
		client := c.NodeClientFactory.GetClient(n.Status.Addresses)
		// todo, add jitter here
		go singleNodeHeartbeat(n, client, heartbeats)
	}
	return nil
}

// If the controller was shut down while creating a node, it will
// remain in creating indeffinately since we don't have an instanceID
// for the node.  Kill it here.
func (c *NodeController) StopCreatingNodes() {
	nodes, err := c.NodeRegistry.ListNodes(func(n *api.Node) bool {
		return n.Status.Phase == api.NodeCreating
	})
	if err != nil {
		klog.Errorf("Could not list nodes to check for creating nodes")
	}
	for _, node := range nodes.Items {
		klog.V(2).Infof("Terminating creating node %s: it was likely lost at restart",
			node.Name)
		go func(n *api.Node) {
			err := c.stopSingleNode(n)
			if err != nil {
				klog.Errorln("Error stopping creating node", err)
			}
		}(node)
	}
}

// When we restart the server, we had old nodes that we were starting
// or restarting and we were polling them in order to know when to
// change their state to available.  We need to restart those polls.
func (c *NodeController) ResumeWaits() {
	// for each node in Creating or Created or Restarting,
	// get the ip address, store in the node, start waiting
	nodes, err := c.NodeRegistry.ListNodes(func(n *api.Node) bool {
		return (n.Status.Phase == api.NodeCreated ||
			n.Status.Phase == api.NodeCleaning)
	})
	if err != nil {
		klog.Errorf("Could not list nodes for resuming waits")
	}
	klog.V(2).Infof("Resume waiting on healty from %d instances", len(nodes.Items))
	for _, node := range nodes.Items {
		go func(node *api.Node) error {
			if len(node.Status.Addresses) == 0 {
				addresses, err := c.CloudClient.WaitForRunning(node)
				c.Events.Emit(events.NodeRunning, "node-controller", node, "")
				if err != nil {
					klog.V(2).Infof("Unhealthy wait for running, terminating node: %s",
						node.Name)
					_ = c.stopSingleNode(node)
					return util.WrapError(
						err, "Error waiting for node to be running")
				}
				if !reflect.DeepEqual(node.Status.Addresses, addresses) {
					node.Status.Addresses = addresses
					_, _ = c.NodeRegistry.UpdateStatus(node)
				}
			}
			return c.waitForAvailableOrTerminate(node, BootTimeout)
		}(node)
	}
}

func (c *NodeController) markUnhealthyNodes(allNodes *api.NodeList, LastHeartbeat map[string]time.Time) error {
	nodes := filterNodes(allNodes, func(n *api.Node) bool {
		return n.Status.Phase == api.NodeAvailable
	})
	// for each node, go through and get time in LastHeartbeat
	// if it's too old, terminate the node
	// if it's not in the loop, add NOW as the time of the last heartbeat
	// so that we start the clock on the node's healthyness
	now := time.Now().UTC()
	for _, node := range nodes {
		last, exists := LastHeartbeat[node.Name]
		if !exists {
			LastHeartbeat[node.Name] = now
			continue
		}
		if now.Sub(last) < HealthyTimeout {
			continue
		}
		klog.Warningf("No heartbeats from node %s. Set to terminate.", node.Name)
		node, err := c.NodeRegistry.MarkForTermination(node)
		if err != nil {
			klog.Errorf("Error marking node %s for termination", node.Name)
		}
	}
	return nil
}

// go through and remove any heartbeat records for nodes
// that no longer exist.
func pruneHeartbeats(allNodes *api.NodeList, lastHeartbeat map[string]time.Time) {
	nodeSet := make(map[string]bool)
	for _, n := range allNodes.Items {
		if n.Status.Phase != api.NodeTerminated {
			nodeSet[n.Name] = true
		}
	}
	// According to the internet, deletes over range is safe
	// Also, dogs can't look up.
	for nodeName, _ := range lastHeartbeat {
		if !nodeSet[nodeName] {
			delete(lastHeartbeat, nodeName)
		}
	}
}

func singleNodeHeartbeat(node *api.Node, client nodeclient.NodeClient, healthyReply chan string) {
	err := client.Healthcheck()
	if err != nil {
		klog.Warningf("Heartbeat error from node %s: %s", node.Name, err.Error())
		return
	}
	healthyReply <- node.Name
}

func (c *NodeController) waitForAvailableOrTerminate(node *api.Node, timeout time.Duration) error {
	if len(node.Status.Addresses) == 0 {
		err := fmt.Errorf("No IP address stored for node %s", node.Name)
		klog.Errorf(err.Error())
		_ = c.stopSingleNode(node)
		return err
	}
	klog.V(2).Infof("Waiting for available on node %s", node.Name)
	client := c.NodeClientFactory.GetClient(node.Status.Addresses)
	err := waitForHealthy(node, client, timeout)
	if err != nil {
		klog.Errorf("Error in node start: node unresponsive for %s seconds", timeout)
		klog.V(2).Infof("Terminating node: %s", node.Name)
		_ = c.stopSingleNode(node)
		return util.WrapError(err, "Error waiting for healthy node")
	}
	node.Status.Phase = api.NodeAvailable
	_, err = c.NodeRegistry.UpdateStatus(node)
	if err != nil {
		klog.Errorf("Error setting node %s to available,", node.Name)
		klog.V(2).Infof("Terminating node: %s", node.Name)
		_ = c.stopSingleNode(node)
		return util.WrapError(err, "Error waiting for healthy node")
	}
	return nil
}

func waitForHealthy(node *api.Node, client nodeclient.NodeClient, timeout time.Duration) error {
	success := false
	giveUp := time.Now().UTC().Add(timeout)
	for time.Now().UTC().Before(giveUp) {
		err := client.Healthcheck()
		if err == nil {
			success = true
			break
		}
		time.Sleep(HealthcheckPause)
	}
	if !success {
		return fmt.Errorf("Not Healthy")
	}
	return nil
}

func filterNodes(allNodes *api.NodeList, pred func(*api.Node) bool) []*api.Node {
	nodes := make([]*api.Node, 0)
	for _, n := range allNodes.Items {
		if pred(n) {
			nodes = append(nodes, n)
		}
	}
	return nodes
}

// Goes through and stops nodes that we have asked to be terminated
func (c NodeController) reaperLoop(quit <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	ticker := time.NewTicker(c.Config.ReaperInterval)
	for {
		select {
		case <-ticker.C:
			nodes, err := c.NodeRegistry.ListNodes(func(n *api.Node) bool {
				return (n.Spec.Terminate &&
					n.Status.Phase != api.NodeTerminated)
			})
			if err != nil {
				klog.Errorf("Error listing nodes for reaper loop: %s", err.Error())
				continue
			}
			for _, node := range nodes.Items {
				if node.Status.BoundPodName != "" {
					c.removePodFromNode(node)
				}
				go c.stopSingleNode(node)
			}
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func (c *NodeController) removePodFromNode(node *api.Node) {
	pod, err := c.PodReader.GetPod(node.Status.BoundPodName)
	if err != nil {
		klog.Warningf("Could not find pod %s that was reported to be on node %s",
			node.Status.BoundPodName, node.Name)
		return
	}
	if pod.Status.BoundNodeName != node.Name {
		klog.Warningf("Pod %s no longer on node %s", node.Status.BoundPodName, node.Name)
		return
	}
	c.Events.Emit(
		events.PodEjected, "node-controller", pod,
		"Pod %s was found on terminating node %s. Ejecting pod", pod.Name, node.Name)

	node.Status.BoundPodName = ""
	_, err = c.NodeRegistry.UpdateStatus(node)
	if err != nil {
		klog.Errorf("Error deleting bound pod on failed node %s: %v", node.Name, err.Error())
	}
}

// There are 2 ways to implement this: either we have one loop responsible
// for buffering nodes and dispatching nodes or we could break it into 2
// loops that communicate with a channel. I thought the 2 loops communicating
// would be easier to test and validate.
//
// What's happening here is the updateBufferedNodesLoop periodically
// recomputes pod-node bindings and sends a map of those updated
// bindings to this function (through nodeBindingUpdate).  Those
// bindings are used to speed up dispatching (this helps ensure we
// don't need to cycle through all nodes when a pod requests a node).
func (c *NodeController) dispatchNodesLoop(quit <-chan struct{}, wg *sync.WaitGroup, nodeBindingsUpdate <-chan map[string]string) {
	wg.Add(1)
	defer wg.Done()
	poolTicker := time.NewTicker(c.Config.PoolInterval)
	defer poolTicker.Stop()

	podNodeMap := <-nodeBindingsUpdate
	for {
		select {
		case <-quit:
			return
		default:
			select {
			case updatedBindings := <-nodeBindingsUpdate:
				podNodeMap = updatedBindings
			case nodeReq := <-c.NodeDispenser.NodeRequestChan:
				nodeReq.ReplyChan <- c.requestNode(nodeReq, podNodeMap)
			case returnedNodeMsg := <-c.NodeDispenser.NodeReturnChan:
				klog.V(2).Infof("Got node %s back", returnedNodeMsg.NodeName)
				if returnedNodeMsg.Unused {
					go c.cleanUnusedNode(returnedNodeMsg.NodeName)
				} else {
					go c.cleanUsedNode(returnedNodeMsg.NodeName)
				}
			case <-quit:
				return
			}
		}
	}
}

// If the pod_controller didn't use a node, it can let us know that
// the node can be reused, here we just mark the node as available and
// wipe any info that might have been set.
func (c *NodeController) cleanUnusedNode(name string) {
	klog.V(2).Infof("Node %s is unused, returning to pool", name)
	node, err := c.NodeRegistry.GetNode(name)
	if err != nil {
		klog.Errorln("Error retrieving unused node from registry", name)
		return
	}
	node.Status.Phase = api.NodeAvailable
	node.Status.BoundPodName = ""
	_, err = c.NodeRegistry.UpdateStatus(node)
	if err != nil {
		klog.Errorf("Error updating node %s status for cleaning unused node: %v",
			name, err)
		// if things went wrong when putting it back into available, try to
		// clean it.
		go c.cleanUsedNode(name)
	}
}

func (c *NodeController) requestNode(nodeReq NodeRequest, podNodeMapping map[string]string) NodeReply {
	// look up node
	boundNodeName := podNodeMapping[nodeReq.requestingPod.Name]
	if boundNodeName == "" {
		return NodeReply{NoBinding: true}
	}
	boundNode, err := c.NodeRegistry.GetNode(boundNodeName)
	if err != nil {
		if err != store.ErrKeyNotFound {
			klog.Errorln("Could not list nodes for dispensing to pod:", err)
		}
		return NodeReply{}

	} else if boundNode.Status.Phase != api.NodeAvailable {
		return NodeReply{}
	}
	err = c.bindNodeToPod(&nodeReq.requestingPod, boundNode)
	if err != nil {
		klog.Errorln("Error binding pod to available node", err)
		return NodeReply{}
	}
	return NodeReply{
		Node: boundNode,
	}
}

// Since we are supporting multiple cpu architectures now this function must return
// a slice of images based on the number of filters we have within the bootspec
func (c *NodeController) imageSpecToImage(spec cloud.BootImageSpec) ([]cloud.Image, error) {
	var images []cloud.Image
	for _, ispec := range c.CloudClient.Extend(spec) {
		var img cloud.Image
		obj, exists := c.ImageIdCache.Get(ispec.String())
		if obj != nil {
			img = obj.(cloud.Image)
			images = append(images, img)
		}
		if !exists || img.ID == "" {
			var err error
			img, err = c.CloudClient.GetImage(ispec)
			if err != nil {
				klog.Errorf("resolving image spec %v to image ID: %v",
					ispec, err)
				return images, err
			}
			c.ImageIdCache.Add(ispec.String(), img, 5*time.Minute,
				func(obj interface{}) {
					_, _ = c.imageSpecToImage(ispec)
				})
			klog.V(2).Infof("latest image for spec %v: %v", ispec, img)
		}
	}
	return images, nil
}

func (c *NodeController) bindNodeToPod(pod *api.Pod, node *api.Node) error {
	node.Status.Phase = api.NodeClaimed
	node.Status.BoundPodName = pod.Name
	_, err := c.NodeRegistry.UpdateStatus(node)
	return err
}

func (c *NodeController) saveNodeLogs(node *api.Node) {
	klog.V(2).Infof("Saving node logs")
	filename := "/var/log/itzo/itzo.log"
	client := c.NodeClientFactory.GetClient(node.Status.Addresses)
	data, err := client.GetFile(filename, 0, nodeclient.SAVE_LOG_BYTES)
	if err != nil {
		klog.Errorf("Error saving node %s log: %s", node.Name, err.Error())
		return
	}
	log := api.NewLogFile()
	log.Name = filename
	log.ParentObject = api.ToObjectReference(node)
	log.Content = string(data)
	_, err = c.LogRegistry.CreateLog(log)
	if err != nil {
		klog.Errorf("Error saving node %s log to registry: %s",
			node.Name, err.Error())
	}
}

func (c *NodeController) cleanUsedNode(name string) error {
	node, err := c.NodeRegistry.GetNode(name)
	if err != nil {
		err = util.WrapError(
			err, "Error retrieving node %s for cleaning", name)
		klog.Errorf(err.Error())
		return err
	}

	// Since we're now terminating nodes, the cleaning
	// phase is useless...  Should we get rid of it entirely?
	node.Status.Phase = api.NodeCleaning
	// event MUST contain node with boundpodname still set (makes
	// services DNS work)
	c.Events.Emit(events.NodeCleaning, "node-controller", *node)
	node.Status.BoundPodName = ""
	_, err = c.NodeRegistry.UpdateStatus(node)
	if err != nil {
		// todo: figure out what should happen in this case if we
		// can't put the node into cleaning, then what should we do?
		// Should we continue and hope for the best?
		err = util.WrapError(err, "Error updating node to cleaning status")
		klog.Errorf(err.Error())
	}
	c.saveNodeLogs(node)
	// We've decided to skip cleaning and just terminate.  if you
	// decide to remove the node-cleaning phase entirely then please
	// make sure to double check and make sure that any nodes returned
	// to the node_controller emit a NodeCleaning event (or re-write the
	// consumers of that event)
	if err = c.stopSingleNode(node); err != nil {
		klog.Errorf("Error in cleaning: could not terinate: %s", err.Error())
		return err
	}
	return nil
}
