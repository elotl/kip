package server

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/nodeclient"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/elotl/cloud-instance-provider/pkg/server/events"
	"github.com/elotl/cloud-instance-provider/pkg/server/nodemanager"
	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/elotl/cloud-instance-provider/pkg/util/conmap"
	"github.com/elotl/cloud-instance-provider/pkg/util/k8s/eventrecorder"
	"github.com/kubernetes/kubernetes/pkg/kubelet/network/dns"
	"github.com/stretchr/testify/assert"
)

func createPodController(c nodeclient.ItzoClientFactoryer) (*PodController, func()) {
	quit := make(chan struct{})
	wg := &sync.WaitGroup{}
	podRegistry, closer1 := registry.SetupTestPodRegistry()
	logRegistry, closer2 := registry.SetupTestLogRegistry()
	nodeRegistry, closer3 := registry.SetupTestNodeRegistry()
	closer := func() { closer1(); closer2(); closer3() }
	dispenser := nodemanager.NewNodeDispenser()
	controller := &PodController{
		podRegistry:       podRegistry,
		logRegistry:       logRegistry,
		metricsRegistry:   registry.NewMetricsRegistry(100),
		nodeLister:        nodeRegistry,
		nodeDispenser:     dispenser,
		nodeClientFactory: c,
		events:            events.NewEventSystem(quit, wg),
		cloudClient:       cloud.NewMockClient(),
		lastStatusReply:   conmap.NewStringTimeTime(),
	}
	controller.dnsConfigurer = dns.NewConfigurer(
		eventrecorder.NewLoggingEventRecorder(5),
		nil,
		nil,
		nil,
		"",
		"",
	)
	return controller, closer
}

func waitForPodInState(t *testing.T, ctl *PodController, podName string, state api.PodPhase) {
	var pod *api.Pod
	var err error
	for i := 0; i < 30; i++ {
		pod, err = ctl.podRegistry.GetPod(podName)
		assert.NoError(t, err)
		if pod.Status.Phase == state {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	assert.Equal(t, string(state), string(pod.Status.Phase))
}

func TestDispatchPodToNodeHappy(t *testing.T) {
	t.Parallel()
	client := nodeclient.NewMockItzoClientFactory()
	ctl, closer := createPodController(client)
	defer closer()
	pod := api.GetFakePod()
	pod, err := ctl.podRegistry.CreatePod(pod)
	assert.NoError(t, err)
	pod.Status.Phase = api.PodDispatching
	ctl.podRegistry.UpdatePodStatus(pod, "")
	node := bindPodToANode(t, pod, ctl)
	ctl.dispatchPodToNode(pod, node)
	if pod.Status.Phase != api.PodRunning {
		t.Errorf("Pod should be running it's phase is %s", pod.Status.Phase)
	}
}

func schedulePodHelper(t *testing.T, ctl *PodController, pod *api.Pod) {
	go func() {
		node := api.GetFakeNode()
		nodeReg := ctl.nodeLister.(*registry.NodeRegistry)
		nodeReg.CreateNode(node)
		req := <-ctl.nodeDispenser.NodeRequestChan
		req.ReplyChan <- nodemanager.NodeReply{Node: node}
	}()
	ctl.schedulePod(pod)
}

func TestCheckClaimedNodesSimple(t *testing.T) {
	t.Parallel()
	client := nodeclient.NewMockItzoClientFactory()
	ctl, closer := createPodController(client)
	defer closer()
	n := api.GetFakeNode()
	n.Status.BoundPodName = ""
	n.Status.Phase = api.NodeClaimed
	nodeReg := ctl.nodeLister.(*registry.NodeRegistry)
	nodeReg.CreateNode(n)
	ctl.checkClaimedNodes()
	assert.Equal(t, 0, len(ctl.nodeDispenser.NodeReturnChan))
	ctl.checkClaimedNodes()
	assert.Equal(t, 1, len(ctl.nodeDispenser.NodeReturnChan))
}

func MakeUnitWaiting(name string) api.UnitStatus {
	return api.UnitStatus{
		Name: name,
		State: api.UnitState{
			Waiting: &api.UnitStateWaiting{}},
	}
}

func MakeUnitStartFailure(name string) api.UnitStatus {
	return api.UnitStatus{
		Name: name,
		State: api.UnitState{
			Waiting: &api.UnitStateWaiting{StartFailure: true}},
	}
}

func MakeUnitRunning(name string) api.UnitStatus {
	return api.UnitStatus{
		Name: name,
		State: api.UnitState{
			Running: &api.UnitStateRunning{}},
	}
}

func MakeUnitSucceeded(name string) api.UnitStatus {
	return api.UnitStatus{
		Name: name,
		State: api.UnitState{
			Terminated: &api.UnitStateTerminated{ExitCode: 0}},
	}
}

func MakeUnitFailed(name string) api.UnitStatus {
	return api.UnitStatus{
		Name: name,
		State: api.UnitState{
			Terminated: &api.UnitStateTerminated{ExitCode: 1}},
	}
}

type podPhaseInput struct {
	restartPolicy api.RestartPolicy
	units         []api.UnitStatus
	phase         api.PodPhase
	isValid       bool
	failMsg       string
}

func TestUpdatePodStatus(t *testing.T) {
	t.Parallel()
	client := nodeclient.NewMockItzoClientFactory()
	ctl, closer := createPodController(client)
	defer closer()

	tests := []struct {
		name          string
		modifyPod     func(*api.Pod)
		statuses      []api.UnitStatus
		expectedPhase api.PodPhase
		startFailures int
	}{
		{
			modifyPod: func(p *api.Pod) { p.Status.Phase = api.PodDispatching },
			statuses: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitRunning("bar"),
			},
			expectedPhase: api.PodRunning,
		},
		{
			modifyPod: func(p *api.Pod) { p.Status.Phase = api.PodDispatching },
			statuses: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitFailed("bar"),
			},
			expectedPhase: api.PodRunning,
		},
		{
			modifyPod: func(p *api.Pod) { p.Status.Phase = api.PodDispatching },
			statuses: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitSucceeded("bar"),
			},
			expectedPhase: api.PodRunning,
		},
		{
			modifyPod: func(p *api.Pod) { p.Status.Phase = api.PodDispatching },
			statuses: []api.UnitStatus{
				MakeUnitSucceeded("foo"),
				MakeUnitFailed("bar"),
			},
			expectedPhase: api.PodFailed,
		},
		{
			modifyPod: func(p *api.Pod) { p.Status.Phase = api.PodDispatching },
			statuses: []api.UnitStatus{
				MakeUnitFailed("foo"),
				MakeUnitFailed("bar"),
			},
			expectedPhase: api.PodFailed,
		},
		{
			modifyPod: func(p *api.Pod) {},
			statuses: []api.UnitStatus{
				MakeUnitSucceeded("foo"),
				MakeUnitSucceeded("bar"),
			},
			expectedPhase: api.PodSucceeded,
		},
		{
			modifyPod: func(p *api.Pod) {},
			statuses: []api.UnitStatus{
				MakeUnitSucceeded("foo"),
				MakeUnitSucceeded("bar"),
			},
			expectedPhase: api.PodSucceeded,
		},
		{
			name:      "launch failure increments start failure",
			modifyPod: func(p *api.Pod) { p.Status.StartFailures = 1 },
			statuses: []api.UnitStatus{
				MakeUnitStartFailure("foo"),
				MakeUnitSucceeded("bar"),
			},
			expectedPhase: api.PodFailed,
			startFailures: 2,
		},
		{
			name:      "no pods waiting resets start failure",
			modifyPod: func(p *api.Pod) { p.Status.StartFailures = 1 },
			statuses: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitSucceeded("bar"),
			},
			expectedPhase: api.PodRunning,
			startFailures: 0,
		},
		{
			name:          "no units in pod doesnt update all units started",
			modifyPod:     func(p *api.Pod) { p.Status.StartFailures = 1 },
			statuses:      []api.UnitStatus{},
			expectedPhase: api.PodRunning,
			startFailures: 1,
		},
	}
	for i, test := range tests {
		p := api.GetFakePod()
		p.Spec.RestartPolicy = api.RestartPolicyNever
		p.Spec.Phase = api.PodRunning
		p.Status.Phase = api.PodRunning
		p.Status.Addresses = api.NewNetworkAddresses("1.2.3.4", "")
		p.Status.BoundNodeName = "mynode"
		test.modifyPod(p)
		p, err := ctl.podRegistry.CreatePod(p)
		assert.Nil(t, err)
		reply := FullPodStatus{
			Name:         p.Name,
			UnitStatuses: test.statuses,
			PodIP:        "1.2.3.4",
		}
		ctl.handlePodStatusReply(reply)
		p, err = ctl.podRegistry.GetPod(p.Name)
		assert.Nil(t, err)
		msg := fmt.Sprintf("failed test %d: %s", i, test.name)
		assert.Equal(t, test.expectedPhase, p.Status.Phase, msg)
		assert.Equal(t, test.startFailures, p.Status.StartFailures, msg)
	}
}

func TestCheckRunningPods(t *testing.T) {
	t.Parallel()
	client := nodeclient.NewMockItzoClientFactory()
	ctl, closer := createPodController(client)
	defer closer()
	p := api.GetFakePod()
	p, err := ctl.podRegistry.CreatePod(p)
	assert.Nil(t, err)
	p.Status.Phase = api.PodDispatching
	_, err = ctl.podRegistry.UpdatePodStatus(p, "")
	assert.Nil(t, err)
	p.Status.Phase = api.PodRunning
	p.Status.BoundNodeName = ""
	_, err = ctl.podRegistry.UpdatePodStatus(p, "")
	assert.Nil(t, err)
	ctl.checkRunningPods()
	pods, err := ctl.podRegistry.ListPods(registry.MatchAllPods)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(pods.Items))
	ctl.checkRunningPods()
	pods, err = ctl.podRegistry.ListPods(registry.MatchAllPods)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(pods.Items))
	assert.Equal(t, api.PodFailed, pods.Items[0].Status.Phase)
}

func TestCheckPodStatusRunning(t *testing.T) {
	t.Parallel()
	client := nodeclient.NewMockItzoClientFactory()
	ctl, closer := createPodController(client)
	defer closer()
	p := api.GetFakePod()
	p.Name = "foopod"
	p.Spec.Phase = api.PodRunning
	units := []api.Unit{
		api.Unit{
			Name:    "foo",
			Image:   "myimage",
			Command: []string{"foo"},
		},
	}
	p.Spec.Units = units
	p.Status.Phase = api.PodRunning
	p, err := ctl.podRegistry.CreatePod(p)
	assert.Nil(t, err)
	bindPodToANode(t, p, ctl)
	reply := ctl.queryPodStatus(p)
	assert.Nil(t, reply.Error)
	assert.Equal(t, p.Name, reply.Name)
	_, exists := ctl.lastStatusReply.GetOK(p.Name)
	assert.True(t, exists)
}

func TestFailingToStartPod(t *testing.T) {
	t.Parallel()
	client := nodeclient.NewMockItzoClientFactory()
	ctl, closer := createPodController(client)
	defer closer()

	p := api.GetFakePod()
	p.Status.Phase = api.PodRunning
	p, err := ctl.podRegistry.CreatePod(p)
	assert.NoError(t, err)
	// Go through and test a pod that continually fails to launch
	for i := 0; i <= allowedStartFailures; i++ {
		ctl.markFailedPod(p, true, "")
		remedyFailedPod(p, ctl.podRegistry)
		p, err := ctl.podRegistry.GetPod(p.Name)
		assert.NoError(t, err)
		assert.Equal(t, i+1, p.Status.StartFailures)
		if i == allowedStartFailures {
			assert.Equal(t, api.PodFailed, p.Spec.Phase)
			assert.Equal(t, api.PodFailed, p.Status.Phase)
			break
		} else {
			assert.Equal(t, api.PodRunning, p.Spec.Phase)
			assert.Equal(t, api.PodWaiting, p.Status.Phase)
		}
		p.Status.Phase = api.PodDispatching
		p, err = ctl.podRegistry.UpdatePodStatus(p, "")
		assert.NoError(t, err)
	}
}

func FailStatus() (*api.PodStatusReply, error) {
	return nil, fmt.Errorf("Status failed")
}

func TestCheckPodStatusError(t *testing.T) {
	t.Parallel()
	client := nodeclient.NewMockItzoClientFactory()
	client.Status = FailStatus
	ctl, closer := createPodController(client)
	defer closer()
	p := api.GetFakePod()
	p.Name = "foopod"
	p.Spec.Phase = api.PodRunning
	units := []api.Unit{
		api.Unit{
			Name:    "foo",
			Image:   "myimage",
			Command: []string{"foo"},
		},
	}
	p.Spec.Units = units
	p.Status.Phase = api.PodRunning
	p.Status.BoundNodeName = ""
	p, err := ctl.podRegistry.CreatePod(p)
	assert.Nil(t, err)
	reply := ctl.queryPodStatus(p)
	assert.NotNil(t, reply.Error)
	_, exists := ctl.lastStatusReply.GetOK(p.Name)
	assert.False(t, exists)
}

func TestPruneLastStatusReplies(t *testing.T) {
	t.Parallel()
	client := nodeclient.NewMockItzoClientFactory()
	client.Status = FailStatus
	ctl, closer := createPodController(client)
	defer closer()
	p1 := api.GetFakePod()
	p1.Name = "pod1"
	p1.Spec.Phase = api.PodRunning
	p1.Status.Phase = api.PodRunning
	p1, err := ctl.podRegistry.CreatePod(p1)
	assert.Nil(t, err)
	p2 := api.GetFakePod()
	p2.Name = "pod2"
	p2.Spec.Phase = api.PodTerminated
	p2.Status.Phase = api.PodTerminated
	p2, err = ctl.podRegistry.CreatePod(p2)
	assert.Nil(t, err)
	ctl.lastStatusReply.Set(p1.Name, time.Now().UTC())
	ctl.lastStatusReply.Set(p2.Name, time.Now().UTC())
	ctl.pruneLastStatusReplies()
	_, exists := ctl.lastStatusReply.GetOK(p1.Name)
	assert.True(t, exists)
	_, exists = ctl.lastStatusReply.GetOK(p2.Name)
	assert.False(t, exists)
}

func TestHandleReplyTimeouts(t *testing.T) {
	t.Parallel()
	client := nodeclient.NewMockItzoClientFactory()
	client.Status = FailStatus
	ctl, closer := createPodController(client)
	defer closer()
	p1 := api.GetFakePod()
	p1.Name = "pod1"
	p1.Spec.Phase = api.PodRunning
	p1.Status.Phase = api.PodRunning
	p1, err := ctl.podRegistry.CreatePod(p1)
	assert.Nil(t, err)
	p2 := api.GetFakePod()
	p2.Name = "pod2"
	p2.Spec.Phase = api.PodTerminated
	p2.Status.Phase = api.PodTerminated
	p2, err = ctl.podRegistry.CreatePod(p2)
	assert.Nil(t, err)
	ctl.lastStatusReply.Set(p1.Name, time.Now().UTC().Add(-2*statusReplyTimeout))
	ctl.lastStatusReply.Set(p2.Name, time.Now().UTC().Add(-2*statusReplyTimeout))
	ctl.handleReplyTimeouts()
	pod, err := ctl.podRegistry.GetPod(p1.Name)
	assert.Nil(t, err)
	//assert.Equal(t, api.PodFailed, pod.Status.Phase)
	waitForPodInState(t, ctl, pod.Name, api.PodFailed)
	pod, err = ctl.podRegistry.GetPod(p2.Name)
	assert.Nil(t, err)
	//assert.Equal(t, api.PodTerminated, pod.Status.Phase)
	waitForPodInState(t, ctl, pod.Name, api.PodTerminated)
}

func TestQueryPodStatus(t *testing.T) {
	// Tests out the case where the pod has a unit that
	// hasn't been run on the node yet (bit of a race)
	// that pod should be reported as waiting
	t.Parallel()
	client := nodeclient.NewMockItzoClientFactory()
	client.Status = func() (*api.PodStatusReply, error) {
		reply := &api.PodStatusReply{
			UnitStatuses: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitRunning("bar"),
				MakeUnitRunning("oldUnit"),
			},
		}
		return reply, nil
	}
	ctl, closer := createPodController(client)
	defer closer()
	pod := api.GetFakePod()
	pod.Spec.Units = []api.Unit{
		{
			Name:  "foo",
			Image: "fooimage",
		},
		{
			Name:  "bar",
			Image: "barimage",
		},
		{
			Name:  "new",
			Image: "newImage",
		},
	}
	pod, err := ctl.podRegistry.CreatePod(pod)
	assert.NoError(t, err)
	bindPodToANode(t, pod, ctl)
	s := ctl.queryPodStatus(pod)
	us := s.UnitStatuses
	assert.Equal(t, len(pod.Spec.Units), len(us))
	for i := 0; i < len(pod.Spec.Units); i++ {
		assert.Equal(t, pod.Spec.Units[i].Name, us[i].Name)
	}
	// Make sure the new pod is in the waiting state
	assert.NotNil(t, us[2].State.Waiting)
}

func bindPodToANode(t *testing.T, pod *api.Pod, ctl *PodController) *api.Node {
	node := api.GetFakeNode()
	nodeRegistry := ctl.nodeLister.(*registry.NodeRegistry)
	node, err := nodeRegistry.CreateNode(node)
	assert.NoError(t, err)
	pod.Status.BoundNodeName = node.Name
	pod.Status.BoundInstanceID = node.Status.InstanceID
	_, err = ctl.podRegistry.Update(pod)
	assert.NoError(t, err)
	return node
}

func TestPCMaybeFailUnresponsivePod(t *testing.T) {
	client := nodeclient.NewMockItzoClientFactory()
	ctl, closer := createPodController(client)
	defer closer()
	pod := api.GetFakePod()
	pod.Status.Phase = api.PodRunning
	pod, err := ctl.podRegistry.CreatePod(pod)
	assert.NoError(t, err)
	bindPodToANode(t, pod, ctl)
	ctl.maybeFailUnresponsivePod(pod)
	_, exists := ctl.lastStatusReply.GetOK(pod.Name)
	assert.True(t, exists)
	// now fail the status check
	ctl.lastStatusReply.Delete(pod.Name)
	client.Status = FailStatus
	ctl.maybeFailUnresponsivePod(pod)
	pod, err = ctl.podRegistry.GetPod(pod.Name)
	assert.NoError(t, err)
	assert.Equal(t, api.PodFailed, pod.Status.Phase)
	_, exists = ctl.lastStatusReply.GetOK(pod.Name)
	assert.False(t, exists)
}

func TestSetPodDispatchingParams(t *testing.T) {
	client := nodeclient.NewMockItzoClientFactory()
	ctl, closer := createPodController(client)
	defer closer()

	pod := api.GetFakePod()
	pod, err := ctl.podRegistry.CreatePod(pod)
	assert.NoError(t, err)
	node := api.GetFakeNode()
	instid := "abc"
	node.Status.InstanceID = instid
	_, err = ctl.setPodDispatchingParams(pod, node)
	assert.NoError(t, err)
	pod, err = ctl.podRegistry.GetPod(pod.Name)
	assert.NoError(t, err)
	assert.Equal(t, api.PodDispatching, pod.Status.Phase)
	assert.Equal(t, instid, pod.Status.BoundInstanceID)
	assert.Equal(t, 0, len(pod.Status.Addresses))
}
