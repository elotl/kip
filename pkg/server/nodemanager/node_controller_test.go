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
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/certs"
	"github.com/elotl/kip/pkg/nodeclient"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/server/events"
	"github.com/elotl/kip/pkg/server/registry"
	"github.com/elotl/kip/pkg/util/cloudinitfile"
	"github.com/elotl/kip/pkg/util/stats"
	"github.com/elotl/kip/pkg/util/timeoutmap"
	"github.com/stretchr/testify/assert"
)

var (
	defaultInstanceType  = "t2.nano"
	defaultBootImageID   = "ami-elotl"
	defaultBootImageSpec = cloud.BootImageSpec{
		"owners":  "689494258501",
		"filters": "name=elotl-kip-*",
	}
)

type StartStopFunc func(node *api.Node) error

func StringStringReturnNil(cidr, iid string) error {
	return nil
}

func ReturnNil(iid string) error {
	return nil
}

func FakeLister() ([]cloud.CloudInstance, error) {
	return nil, nil
}

func StartReturnsOK(node *api.Node, image cloud.Image, metadata, iamProfile string) (string, error) {
	result := "instID"
	return result, nil
}

func StartFails(node *api.Node, image cloud.Image, metadata, iamProfile string) (string, error) {
	return "", fmt.Errorf("Testing, purposefully returning error")
}

func ReturnAddresses(node *api.Node) ([]api.NetworkAddress, error) {
	return api.NewNetworkAddresses("instIP", "instDNS"), nil
}

func ReturnError(node *api.Node) error {
	return fmt.Errorf("Testing, purposefully returning error")
}

func Panics(node *api.Node) (string, error) {
	return "instID", nil
}

func AddInstanceParameterReturnsOK(instanceID, key, value string, isSecret bool) error {
	return nil
}

func AddInstanceParameterFails(instanceID, key, value string, isSecret bool) error {
	return fmt.Errorf("testing AddInstanceParameter() failure path")
}

func DeleteInstanceParameterReturnsOK(instanceID, key string) error {
	return nil
}

func DeleteInstanceParameterFails(instanceID, key string) error {
	return fmt.Errorf("testing DeleteInstanceParameter() failure path")
}

func MakeNodeController() (*NodeController, func()) {
	quit := make(chan struct{})
	wg := &sync.WaitGroup{}
	nodeRegistry, closer1 := registry.SetupTestNodeRegistry()
	logRegistry, closer2 := registry.SetupTestLogRegistry()
	podRegistry, closer3 := registry.SetupTestPodRegistry()
	closer := func() { closer1(); closer2(); closer3() }
	cloudClient := &cloud.MockCloudClient{
		Starter:                  StartReturnsOK,
		SpotStarter:              StartReturnsOK,
		Stopper:                  ReturnNil,
		Waiter:                   ReturnAddresses,
		RouteRemover:             StringStringReturnNil,
		InstanceParameterAdder:   AddInstanceParameterReturnsOK,
		InstanceParameterRemover: DeleteInstanceParameterReturnsOK,
	}
	defaultBootImage := cloud.Image{
		ID: defaultBootImageID,
	}
	imageIdCache := timeoutmap.New(false, make(chan struct{}))
	imageIdCache.Add(defaultBootImageSpec.String(), defaultBootImage, 5*time.Minute, timeoutmap.Noop)
	fakeCertFactory, _ := certs.NewFake()
	bootLimiter := NewInstanceBootLimiter()
	ciFile, _ := cloudinitfile.New("")
	nc := &NodeController{
		Config: NodeControllerConfig{
			PoolInterval:      1 * time.Second,
			HeartbeatInterval: 20 * time.Second,
			ReaperInterval:    20 * time.Second,
		},
		NodeRegistry:  nodeRegistry,
		LogRegistry:   logRegistry,
		PodReader:     podRegistry,
		NodeDispenser: NewNodeDispenser(),
		NodeScaler: &BindingNodeScaler{
			nodeRegistry: nodeRegistry,
			standbyNodes: nil,
			bootLimiter:  bootLimiter,
		},
		CloudClient:        cloudClient,
		NodeClientFactory:  nodeclient.NewMockItzoClientFactory(),
		Events:             events.NewEventSystem(quit, wg),
		PoolLoopTimer:      &stats.LoopTimer{},
		ImageIdCache:       imageIdCache,
		CertificateFactory: fakeCertFactory,
		CloudInitFile:      ciFile,
		BootLimiter:        bootLimiter,
		BootImageSpec:      defaultBootImageSpec,
	}
	return nc, closer
}

func makeTestNode(t *testing.T, ctl *NodeController, phase api.NodePhase, spot bool) *api.Node {
	n := api.GetFakeNode()
	n.Spec.BootImage = defaultBootImageID
	n.Spec.InstanceType = defaultInstanceType
	n.Spec.Spot = spot
	n, err := ctl.NodeRegistry.CreateNode(n)
	assert.Nil(t, err)
	switch phase {
	case api.NodeCreating, api.NodeCreated, api.NodeAvailable, api.NodeTerminating, api.NodeTerminated:
		n.Status.Phase = phase
		n, err = ctl.NodeRegistry.UpdateStatus(n)
		assert.Nil(t, err)
	case api.NodeClaimed, api.NodeCleaning:
		n.Status.Phase = api.NodeAvailable
		n, err = ctl.NodeRegistry.UpdateStatus(n)
		assert.Nil(t, err)
		n.Status.Phase = phase
		n, err = ctl.NodeRegistry.UpdateStatus(n)
		assert.Nil(t, err)
	}
	return n
}

func makeTestOndemandNode(t *testing.T, ctl *NodeController, phase api.NodePhase) *api.Node {
	return makeTestNode(t, ctl, phase, false)
}

func makeTestSpotNode(t *testing.T, ctl *NodeController, phase api.NodePhase) *api.Node {
	return makeTestNode(t, ctl, phase, true)
}

func TestStopSingleNode(t *testing.T) {
	t.Parallel()
	ctl, closer := MakeNodeController()
	defer closer()
	n := api.GetFakeNode()
	n, err := ctl.NodeRegistry.CreateNode(n)
	assert.Nil(t, err)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = ctl.stopSingleNode(n)
	assert.Nil(t, err)
	time.Sleep(1 * time.Second)
	nodes, err := ctl.NodeRegistry.ListAllNodes(registry.MatchAllNodes)
	assert.Nil(t, err)
	assert.Len(t, nodes.Items, 1)
	assert.Equal(t, api.NodeTerminated, nodes.Items[0].Status.Phase)
}

func StartAFewNodes(t *testing.T, numNodes int, spotNode bool) {
	ctl, closer := MakeNodeController()
	defer closer()
	startNodes := make([]*api.Node, 0, numNodes)
	for i := 0; i < numNodes; i++ {
		node := api.GetFakeNode()
		if spotNode {
			node.Spec.Spot = true
		}
		startNodes = append(startNodes, node)
	}
	ctl.startNodes(startNodes, cloud.Image{})
	// starting happens in a goroutine so we'll sleep here
	time.Sleep(1 * time.Second)
	nodes, err := ctl.NodeRegistry.ListNodes(registry.MatchAllNodes)
	assert.Nil(t, err)
	assert.Equal(t, len(nodes.Items), numNodes)
	for i := 0; i < numNodes; i++ {
		assert.Equal(t, api.NodeAvailable, nodes.Items[i].Status.Phase)
	}
}

func TestStartNodes(t *testing.T) {
	t.Parallel()
	HealthyTimeout = 3000 * time.Millisecond
	HealthcheckPause = 100 * time.Millisecond
	StartAFewNodes(t, 1, false)
	StartAFewNodes(t, 2, false)
	StartAFewNodes(t, 1, true)
}

func TestStartNodeHealthcheckFails(t *testing.T) {
	t.Parallel()
	HealthyTimeout = 500 * time.Millisecond
	HealthcheckPause = 100 * time.Millisecond
	BootTimeout = 500 * time.Millisecond
	ctl, closer := MakeNodeController()
	defer closer()
	ctl.NodeClientFactory.(*nodeclient.MockItzoClientFactory).Health = func() error {
		return fmt.Errorf("fail")
	}
	ctl.startNodes([]*api.Node{api.GetFakeNode()}, cloud.Image{})
	time.Sleep(1 * time.Second)
	nodes, err := ctl.NodeRegistry.ListAllNodes(registry.MatchAllNodes)
	assert.Nil(t, err)
	assert.Equal(t, len(nodes.Items), 1)
	assert.Equal(t, api.NodeTerminated, nodes.Items[0].Status.Phase)
}

func TestStartNodeFails(t *testing.T) {
	t.Parallel()
	HealthyTimeout = 500 * time.Millisecond
	HealthcheckPause = 100 * time.Millisecond
	ctl, closer := MakeNodeController()
	defer closer()
	ctl.CloudClient = &cloud.MockCloudClient{
		Starter: StartFails,
	}
	ctl.startNodes([]*api.Node{api.GetFakeNode()}, cloud.Image{})
	time.Sleep(1 * time.Second)
	nodes, err := ctl.NodeRegistry.ListAllNodes(registry.MatchAllNodes)
	assert.Nil(t, err)
	assert.Equal(t, len(nodes.Items), 1)
	assert.Equal(t, api.NodeTerminated, nodes.Items[0].Status.Phase)
}

func OKResponse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func TestSingleHeartbeat(t *testing.T) {
	//create an http endpoint and really test it out
	t.Parallel()
	s := httptest.NewServer(http.HandlerFunc(OKResponse))
	c := nodeclient.NewItzoWithClient(s.URL, s.Client())
	replyChan := make(chan string, 1)
	n := api.GetFakeNode()
	singleNodeHeartbeat(n, c, replyChan)
	select {
	case gotReply := <-replyChan:
		assert.Equal(t, n.Name, gotReply)
	case <-time.After(1 * time.Second):
		t.Error("Should have gotten a reply")
	}
}

func TestDispenseNoNodes(t *testing.T) {
	t.Parallel()
	ctl, closer := MakeNodeController()
	defer closer()
	quit := make(chan struct{})
	wg := &sync.WaitGroup{}
	updateChan := make(chan map[string]string)
	go ctl.dispatchNodesLoop(quit, wg, updateChan)
	availableNodes := map[string]string{}
	updateChan <- availableNodes
	pod := api.GetFakePod()
	resp := ctl.NodeDispenser.RequestNode(*pod)
	assert.Nil(t, resp.Node)
	assert.True(t, resp.NoBinding)
}

func TestDispenseNoAvailableNodes(t *testing.T) {
	t.Parallel()
	ctl, closer := MakeNodeController()
	defer closer()
	n := api.GetFakeNode()
	n, _ = ctl.NodeRegistry.CreateNode(n)
	n.Status.Phase = api.NodeAvailable
	n, _ = ctl.NodeRegistry.UpdateStatus(n)
	n.Status.Phase = api.NodeClaimed
	_, _ = ctl.NodeRegistry.UpdateStatus(n)
	quit := make(chan struct{})
	wg := &sync.WaitGroup{}
	updateChan := make(chan map[string]string)
	go ctl.dispatchNodesLoop(quit, wg, updateChan)
	availableNodes := map[string]string{}
	updateChan <- availableNodes
	pod := api.GetFakePod()
	resp := ctl.NodeDispenser.RequestNode(*pod)
	assert.Nil(t, resp.Node)
}

func TestDispenseAvailableNode(t *testing.T) {
	t.Parallel()
	ctl, closer := MakeNodeController()
	defer closer()
	n := api.GetFakeNode()
	pod := api.GetFakePod()
	pod.Spec.InstanceType = n.Spec.InstanceType
	n, _ = ctl.NodeRegistry.CreateNode(n)
	n.Status.Phase = api.NodeAvailable
	n.Status.BoundPodName = pod.Name
	_, _ = ctl.NodeRegistry.UpdateStatus(n)
	nodeBindingsUpdate := make(chan map[string]string)
	quit := make(chan struct{})
	wg := &sync.WaitGroup{}
	go ctl.dispatchNodesLoop(quit, wg, nodeBindingsUpdate)
	nodeBindingsUpdate <- map[string]string{pod.Name: n.Name}

	resp := ctl.NodeDispenser.RequestNode(*pod)
	node := resp.Node
	assert.NotNil(t, node)
	assert.Equal(t, api.NodeClaimed, node.Status.Phase)
}

func TestBufferingAndDispatchingTogether(t *testing.T) {
	t.Parallel()
	ctl, closer := MakeNodeController()
	defer closer()
	ctl.CloudClient = &cloud.MockCloudClient{
		Starter:     StartReturnsOK,
		SpotStarter: StartReturnsOK,
		Stopper:     ReturnNil,
		Waiter:      ReturnAddresses,
		ImageGetter: func(spec cloud.BootImageSpec) (cloud.Image, error) {
			return cloud.Image{}, nil
		},
	}
	pod := api.GetFakePod()
	podReg := ctl.PodReader.(*registry.PodRegistry)
	pod, err := podReg.CreatePod(pod)
	assert.NoError(t, err)

	quit := make(chan struct{})
	wg := &sync.WaitGroup{}
	nodeBindingsUpdate := make(chan map[string]string)
	go ctl.updateBufferedNodesLoop(quit, wg, nodeBindingsUpdate)
	go ctl.dispatchNodesLoop(quit, wg, nodeBindingsUpdate)
	var node *api.Node
	// Wait for the node to go through and become available
	for i := 0; i < 30; i++ {
		resp := ctl.NodeDispenser.RequestNode(*pod)
		node = resp.Node
		if node != nil {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
	assert.NotNil(t, node)
	if node != nil {
		assert.Equal(t, api.NodeClaimed, node.Status.Phase)
	}
	close(quit)
}

func TestCleanUsedNode(t *testing.T) {
	t.Parallel()
	ctl, closer := MakeNodeController()
	defer closer()
	var eventCleanedNode *api.Node
	ctl.Events.RegisterHandlerFunc(events.NodeCleaning, func(e events.Event) error {
		node := e.Object.(*api.Node)
		eventCleanedNode = node
		return nil
	})
	n := api.GetFakeNode()
	n, _ = ctl.NodeRegistry.CreateNode(n)
	n.Status.Phase = api.NodeAvailable
	n.Status.Addresses = api.NewNetworkAddresses("1.2.3.4", "")
	n, _ = ctl.NodeRegistry.UpdateStatus(n)
	n.Status.Phase = api.NodeClaimed
	boundPod := "testpod"
	n.Status.BoundPodName = boundPod
	n, _ = ctl.NodeRegistry.UpdateStatus(n)
	_, err := ctl.NodeRegistry.GetNode(n.Name)
	assert.Nil(t, err)
	go func() {
		err := ctl.cleanUsedNode(n.Name)
		assert.Nil(t, err)
	}()
	time.Sleep(1 * time.Second)
	// at this point, the node has been purged so we shouldn't get
	// a result back
	_, err = ctl.NodeRegistry.GetNode(n.Name)
	assert.NotNil(t, err)
	// Make sure that the cleaned node has boundPodName set when
	// we fired the event (if it doesn't DNS services will fail)
	assert.Equal(t, eventCleanedNode.Status.BoundPodName, boundPod)
}

func TestSendOutHeartbeats(t *testing.T) {
	t.Parallel()
	ctl, closer := MakeNodeController()
	defer closer()
	heartbeats := make(chan string)
	n := makeTestOndemandNode(t, ctl, api.NodeAvailable)
	nodes, err := ctl.NodeRegistry.ListNodes(registry.MatchAllNodes)
	assert.Nil(t, err)
	err = ctl.sendOutHeartbeats(nodes, heartbeats)
	assert.Nil(t, err)
	select {
	case d := <-heartbeats:
		assert.Equal(t, n.Name, d)
	case <-time.After(time.Second * 2):
		t.Error("Timed out waiting for heartbeat")
	}
}

func TestNoHeartbeatsForClaimed(t *testing.T) {
	t.Parallel()
	ctl, closer := MakeNodeController()
	defer closer()
	heartbeats := make(chan string)
	_ = makeTestOndemandNode(t, ctl, api.NodeClaimed)
	nodes, err := ctl.NodeRegistry.ListNodes(registry.MatchAllNodes)
	assert.Nil(t, err)
	err = ctl.sendOutHeartbeats(nodes, heartbeats)
	assert.Nil(t, err)
	select {
	case _ = <-heartbeats:
		assert.Fail(t, "Got heartbeat from claimed node")
	case <-time.After(time.Second * 1):
		// No heartbeats for claimed nodes.
	}
}

func TestMarkUnhealthyNodesClaimedIgnored(t *testing.T) {
	lastHeartbeat := make(map[string]time.Time)
	ctl, closer := MakeNodeController()
	defer closer()
	n1 := makeTestOndemandNode(t, ctl, api.NodeAvailable)
	lastHeartbeat[n1.Name] = time.Now().UTC()
	n2 := makeTestOndemandNode(t, ctl, api.NodeClaimed)
	lastHeartbeat[n2.Name] = time.Now().UTC().Add(-999999 * time.Second)
	nodes, err := ctl.NodeRegistry.ListNodes(registry.MatchAllNodes)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(nodes.Items))
	err = ctl.markUnhealthyNodes(nodes, lastHeartbeat)
	assert.Nil(t, err)
	n, err := ctl.NodeRegistry.GetNode(n1.Name)
	assert.Nil(t, err)
	assert.False(t, n.Spec.Terminate)
	n, err = ctl.NodeRegistry.GetNode(n2.Name)
	assert.Nil(t, err)
	assert.False(t, n.Spec.Terminate)
}

func TestMarkUnhealthyNodesAllHealthy(t *testing.T) {
	lastHeartbeat := make(map[string]time.Time)
	ctl, closer := MakeNodeController()
	defer closer()
	n := makeTestOndemandNode(t, ctl, api.NodeAvailable)
	lastHeartbeat[n.Name] = time.Now().UTC()
	nodes, err := ctl.NodeRegistry.ListNodes(registry.MatchAllNodes)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(nodes.Items))
	_ = ctl.markUnhealthyNodes(nodes, lastHeartbeat)
	assert.Nil(t, err)
	n, _ = ctl.NodeRegistry.GetNode(n.Name)
	assert.False(t, n.Spec.Terminate)
}

func TestMarkUnhealthyNodesUnhealthyAndMissing(t *testing.T) {
	lastHeartbeat := make(map[string]time.Time)
	ctl, closer := MakeNodeController()
	defer closer()
	n1 := makeTestOndemandNode(t, ctl, api.NodeAvailable)
	lastHeartbeat[n1.Name] = time.Now().UTC().Add(-5 * time.Minute)
	n2 := makeTestOndemandNode(t, ctl, api.NodeAvailable)
	nodes, err := ctl.NodeRegistry.ListNodes(registry.MatchAllNodes)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(nodes.Items))
	err = ctl.markUnhealthyNodes(nodes, lastHeartbeat)
	assert.Nil(t, err)
	n1, _ = ctl.NodeRegistry.GetNode(n1.Name)
	n2, _ = ctl.NodeRegistry.GetNode(n2.Name)
	assert.True(t, n1.Spec.Terminate)
	assert.False(t, n2.Spec.Terminate)
	_, exists := lastHeartbeat[n2.Name]
	assert.True(t, exists)
}

func TestPruneHeartbeats(t *testing.T) {
	lastHeartbeat := make(map[string]time.Time)
	ctl, closer := MakeNodeController()
	defer closer()
	n1 := makeTestOndemandNode(t, ctl, api.NodeAvailable)
	lastHeartbeat[n1.Name] = time.Now().UTC()
	lastHeartbeat["NoLongerInReg"] = time.Now().UTC()
	nodes, err := ctl.NodeRegistry.ListNodes(registry.MatchAllNodes)
	assert.Nil(t, err)
	pruneHeartbeats(nodes, lastHeartbeat)
	assert.Equal(t, 1, len(lastHeartbeat))
}

func TestWaitForAvailableOrTerminateWorks(t *testing.T) {
	HealthyTimeout = 500 * time.Millisecond
	HealthcheckPause = 1000 * time.Millisecond
	ctl, closer := MakeNodeController()
	defer closer()
	n := makeTestOndemandNode(t, ctl, api.NodeCreated)
	n.Status.Addresses = api.NewNetworkAddresses("1.2.3.4", "")
	err := ctl.waitForAvailableOrTerminate(n, HealthyTimeout)
	assert.Nil(t, err)
	assert.Equal(t, api.NodeAvailable, n.Status.Phase)
}

func TestWaitForAvailableOrTerminateFails(t *testing.T) {
	HealthyTimeout = 500 * time.Millisecond
	HealthcheckPause = 1000 * time.Millisecond
	ctl, closer := MakeNodeController()
	defer closer()
	ctl.NodeClientFactory.(*nodeclient.MockItzoClientFactory).Health = func() error {
		return fmt.Errorf("fail")
	}
	n := makeTestOndemandNode(t, ctl, api.NodeCreated)
	n.Status.Addresses = api.NewNetworkAddresses("1.2.3.4", "")
	err := ctl.waitForAvailableOrTerminate(n, HealthyTimeout)
	assert.NotNil(t, err)
	assert.Contains(t, []api.NodePhase{api.NodeTerminating, api.NodeTerminated}, n.Status.Phase)
}

func TestRemovePodFromNode(t *testing.T) {
	//todo
}

func TestImageSpecToImage(t *testing.T) {
	ctl, closer := MakeNodeController()
	defer closer()
	ctl.CloudClient = &cloud.MockCloudClient{
		Starter:     StartReturnsOK,
		SpotStarter: StartReturnsOK,
		Stopper:     ReturnNil,
		Waiter:      ReturnAddresses,
		ImageGetter: func(spec cloud.BootImageSpec) (cloud.Image, error) {
			return cloud.Image{
				ID:         "my-image-id",
				Name:       "my-image-name",
				RootDevice: "/dev/rootdev0",
			}, nil
		},
	}
	img, err := ctl.imageSpecToImage(defaultBootImageSpec)
	assert.Nil(t, err)
	assert.Equal(t, defaultBootImageID, img.ID)
	spec := cloud.BootImageSpec{
		"name": "my-name-*",
	}
	img, err = ctl.imageSpecToImage(spec)
	assert.Nil(t, err)
	assert.Equal(t, "my-image-id", img.ID)
}

func TestImageSpecToImageFailure(t *testing.T) {
	t.Parallel()
	ctl, closer := MakeNodeController()
	defer closer()
	ctl.CloudClient = &cloud.MockCloudClient{
		Starter:     StartReturnsOK,
		SpotStarter: StartReturnsOK,
		Stopper:     ReturnNil,
		Waiter:      ReturnAddresses,
		ImageGetter: func(spec cloud.BootImageSpec) (cloud.Image, error) {
			return cloud.Image{}, fmt.Errorf("Testing GetImage() failure")
		},
	}
	spec := cloud.BootImageSpec{
		"name": "my-name-*",
	}
	_, err := ctl.imageSpecToImage(spec)
	assert.NotNil(t, err)
}

func TestRequestNode(t *testing.T) {
	ctl, closer := MakeNodeController()
	defer closer()

	pod := api.GetFakePod()
	req := NodeRequest{requestingPod: *pod}
	mapping := map[string]string{}
	reply := ctl.requestNode(req, mapping)
	assert.Nil(t, reply.Node)
	mapping = map[string]string{pod.Name: "somethign else"}
	reply = ctl.requestNode(req, mapping)
	assert.Nil(t, reply.Node)
	node := api.GetFakeNode()
	node, _ = ctl.NodeRegistry.CreateNode(node)
	mapping = map[string]string{pod.Name: node.Name}
	reply = ctl.requestNode(req, mapping)
	assert.Nil(t, reply.Node)
	node.Status.Phase = api.NodeAvailable
	node.Status.BoundPodName = pod.Name
	node, _ = ctl.NodeRegistry.UpdateStatus(node)
	reply = ctl.requestNode(req, mapping)
	assert.NotNil(t, reply.Node)
	if reply.Node != nil {
		assert.Equal(t, node.Name, reply.Node.Name)
	}
	node, err := ctl.NodeRegistry.GetNode(node.Name)
	assert.NoError(t, err)
	assert.Equal(t, api.NodeClaimed, node.Status.Phase)
	assert.Equal(t, node.Name, reply.Node.Name)
}

func TestDoPoolsCalculation(t *testing.T) {
	t.Parallel()
	ctl, closer := MakeNodeController()
	defer closer()
	ctl.CloudClient = &cloud.MockCloudClient{
		Starter:      StartReturnsOK,
		SpotStarter:  StartReturnsOK,
		Stopper:      ReturnNil,
		Waiter:       ReturnAddresses,
		RouteRemover: StringStringReturnNil,
		ImageGetter: func(spec cloud.BootImageSpec) (cloud.Image, error) {
			return cloud.Image{}, nil
		},
		InstanceParameterRemover: DeleteInstanceParameterReturnsOK,
	}
	// we create a new pod that needs a node and a node that
	// doesn't match, make sure the pod gets a new node and that the
	// node we started with is marked for termination.
	pod := api.GetFakePod()
	pod.Spec.InstanceType = "t2000.nano"
	podReg := ctl.PodReader.(*registry.PodRegistry)
	pod, err := podReg.CreatePod(pod)
	assert.NoError(t, err)
	node := api.GetFakeNode()
	node, err = ctl.NodeRegistry.CreateNode(node)
	assert.NoError(t, err)
	mapping, err := ctl.doPoolsCalculation()
	assert.NoError(t, err)
	boundNodeName := mapping[pod.Name]
	assert.True(t, boundNodeName != "")
	assert.NotEqual(t, node.Name, boundNodeName)

	startedNode, err := ctl.NodeRegistry.GetNode(boundNodeName)
	assert.NoError(t, err)
	assert.Equal(t, pod.Spec.InstanceType, startedNode.Spec.InstanceType)
	// We'll retry this up to 2 seconds
	for i := 0; i < 20; i++ {
		nodes, err := ctl.NodeRegistry.ListAllNodes(func(n *api.Node) bool {
			return n.Name == node.Name
		})
		assert.NoError(t, err)
		// This is a strange flake and I can't seem to figure out what
		// exactly is going on...  We sometimes pull 2 identically named
		// pods back from the registry.
		if len(nodes.Items) == 0 {
			// There's a race between deleting the node and creating
			// the trashed node in storage
			fmt.Println("expected to find a terminated node in the registry", i)
			time.Sleep(100 * time.Millisecond)
			continue
		} else {
			stoppedNode := nodes.Items[0]
			possiblePhases := []api.NodePhase{api.NodeTerminating, api.NodeTerminated}
			assert.Contains(t, possiblePhases, stoppedNode.Status.Phase)
			return
		}
	}
	assert.Fail(t, "Did not find terminated node in the registry")
}

type ErroringNodeScaler struct{}

func (ns *ErroringNodeScaler) Compute(nodes []*api.Node, pods []*api.Pod) ([]*api.Node, []*api.Node, map[string]string) {
	return nil, nil, nil
}

func TestDoPoolsCalculationComputeFails(t *testing.T) {
	t.Parallel()
	ctl, closer := MakeNodeController()
	defer closer()

	ctl.NodeScaler = &ErroringNodeScaler{}
	_, err := ctl.doPoolsCalculation()
	assert.Error(t, err)
}

func TestAddInstanceParameter(t *testing.T) {
	t.Parallel()
	ctl, closer := MakeNodeController()
	defer closer()
	ctl.Config.UseCloudParameterStore = true
	called := false
	ctl.CloudClient = &cloud.MockCloudClient{
		Starter: StartReturnsOK,
		Waiter:  ReturnAddresses,
		InstanceParameterAdder: func(instanceID, key, value string, isSecret bool) error {
			called = true
			return nil
		},
	}
	node := api.GetFakeNode()
	ctl.startNodes([]*api.Node{node}, cloud.Image{})
	time.Sleep(1 * time.Second)
	nodes, err := ctl.NodeRegistry.ListNodes(registry.MatchAllNodes)
	assert.Nil(t, err)
	assert.Equal(t, len(nodes.Items), 1)
	assert.Equal(t, api.NodeAvailable, nodes.Items[0].Status.Phase)
	assert.True(t, called)
}

func TestAddInstanceParameterFails(t *testing.T) {
	t.Parallel()
	HealthyTimeout = 500 * time.Millisecond
	HealthcheckPause = 100 * time.Millisecond
	ctl, closer := MakeNodeController()
	defer closer()
	ctl.Config.UseCloudParameterStore = true
	ctl.CloudClient = &cloud.MockCloudClient{
		Starter:                  StartReturnsOK,
		Stopper:                  ReturnNil,
		Waiter:                   ReturnAddresses,
		InstanceParameterAdder:   AddInstanceParameterFails,
		InstanceParameterRemover: DeleteInstanceParameterReturnsOK,
	}
	ctl.startNodes([]*api.Node{api.GetFakeNode()}, cloud.Image{})
	time.Sleep(1 * time.Second)
	nodes, err := ctl.NodeRegistry.ListAllNodes(registry.MatchAllNodes)
	assert.Nil(t, err)
	assert.Equal(t, len(nodes.Items), 1)
	assert.Equal(t, api.NodeTerminated, nodes.Items[0].Status.Phase)
}

func TestDeleteInstanceParameter(t *testing.T) {
	t.Parallel()
	ctl, closer := MakeNodeController()
	defer closer()
	called := false
	ctl.CloudClient = &cloud.MockCloudClient{
		Stopper: ReturnNil,
		InstanceParameterRemover: func(instanceID, key string) error {
			called = true
			return nil
		},
	}
	n := api.GetFakeNode()
	n, err := ctl.NodeRegistry.CreateNode(n)
	assert.Nil(t, err)
	err = ctl.stopSingleNode(n)
	assert.Nil(t, err)
	time.Sleep(1 * time.Second)
	nodes, err := ctl.NodeRegistry.ListAllNodes(registry.MatchAllNodes)
	assert.Nil(t, err)
	assert.Len(t, nodes.Items, 1)
	assert.Equal(t, api.NodeTerminated, nodes.Items[0].Status.Phase)
	assert.True(t, called)
}

func TestDeleteInstanceParameterFails(t *testing.T) {
	t.Parallel()
	ctl, closer := MakeNodeController()
	defer closer()
	called := false
	ctl.CloudClient = &cloud.MockCloudClient{
		Stopper: ReturnNil,
		InstanceParameterRemover: func(instanceID, key string) error {
			called = true
			return fmt.Errorf("testing DeleteInstanceParameter() failure path")
		},
	}
	n := api.GetFakeNode()
	n, err := ctl.NodeRegistry.CreateNode(n)
	assert.Nil(t, err)
	err = ctl.stopSingleNode(n)
	assert.Nil(t, err)
	time.Sleep(1 * time.Second)
	nodes, err := ctl.NodeRegistry.ListAllNodes(registry.MatchAllNodes)
	assert.Nil(t, err)
	assert.Len(t, nodes.Items, 1)
	assert.Equal(t, api.NodeTerminated, nodes.Items[0].Status.Phase)
	assert.True(t, called)
}
