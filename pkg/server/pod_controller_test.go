package server

import (
	"fmt"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/manager"
	"github.com/elotl/cloud-instance-provider/pkg/nodeclient"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/elotl/cloud-instance-provider/pkg/server/events"
	"github.com/elotl/cloud-instance-provider/pkg/server/nodemanager"
	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/elotl/cloud-instance-provider/pkg/util/conmap"
	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
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

type envs []api.EnvVar

func (e envs) Len() int {
	return len(e)
}

func (e envs) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

func (e envs) Less(i, j int) bool { return e[i].Name < e[j].Name }

func TestMakeEnvironmentVariables(t *testing.T) {
	trueVal := true
	testCases := []struct {
		name            string        // the name of the test case
		ns              string        // the namespace to generate environment for
		unit            *api.Unit     // the container to use
		masterServiceNs string        // the namespace to read master service info from
		nilLister       bool          // whether the lister should be nil
		configMap       *v1.ConfigMap // an optional ConfigMap to pull from
		secret          *v1.Secret    // an optional Secret to pull from
		expectedEnvs    []api.EnvVar  // a set of expected environment vars
		expectedError   bool          // does the test fail
		expectedEvent   string        // does the test emit an event
	}{
		{
			name: "env vars",
			ns:   "test1",
			unit: &api.Unit{
				Env: []api.EnvVar{
					{Name: "FOO", Value: "BAZ"},
				},
			},
			masterServiceNs: metav1.NamespaceDefault,
			nilLister:       false,
			expectedEnvs: []api.EnvVar{
				{Name: "FOO", Value: "BAZ"},
			},
		},
		{
			name: "env expansion",
			ns:   "test1",
			unit: &api.Unit{
				Env: []api.EnvVar{
					{
						Name:  "TEST_LITERAL",
						Value: "test-test-test",
					},
					{
						Name:  "OUT_OF_ORDER_TEST",
						Value: "$(OUT_OF_ORDER_TARGET)",
					},
					{
						Name:  "OUT_OF_ORDER_TARGET",
						Value: "FOO",
					},
					{
						Name: "EMPTY_VAR",
					},
					{
						Name:  "EMPTY_TEST",
						Value: "foo-$(EMPTY_VAR)",
					},
					{
						Name:  "LITERAL_TEST",
						Value: "literal-$(TEST_LITERAL)",
					},
					{
						Name:  "TEST_UNDEFINED",
						Value: "$(UNDEFINED_VAR)",
					},
				},
			},
			masterServiceNs: "nothing",
			nilLister:       false,
			expectedEnvs: []api.EnvVar{
				{
					Name:  "TEST_LITERAL",
					Value: "test-test-test",
				},
				{
					Name:  "LITERAL_TEST",
					Value: "literal-test-test-test",
				},
				{
					Name:  "OUT_OF_ORDER_TEST",
					Value: "$(OUT_OF_ORDER_TARGET)",
				},
				{
					Name:  "OUT_OF_ORDER_TARGET",
					Value: "FOO",
				},
				{
					Name:  "TEST_UNDEFINED",
					Value: "$(UNDEFINED_VAR)",
				},
				{
					Name: "EMPTY_VAR",
				},
				{
					Name:  "EMPTY_TEST",
					Value: "foo-",
				},
			},
		},
		{
			name: "configmapkeyref_missing_optional",
			ns:   "test",
			unit: &api.Unit{
				Env: []api.EnvVar{
					{
						Name: "POD_NAME",
						ValueFrom: &api.EnvVarSource{
							ConfigMapKeyRef: &api.ConfigMapKeySelector{
								LocalObjectReference: api.LocalObjectReference{Name: "missing-config-map"},
								Key:                  "key",
								Optional:             &trueVal,
							},
						},
					},
				},
			},
			masterServiceNs: "nothing",
			expectedEnvs:    nil,
		},
		{
			name: "configmapkeyref_missing_key_optional",
			ns:   "test",
			unit: &api.Unit{
				Env: []api.EnvVar{
					{
						Name: "POD_NAME",
						ValueFrom: &api.EnvVarSource{
							ConfigMapKeyRef: &api.ConfigMapKeySelector{
								LocalObjectReference: api.LocalObjectReference{Name: "test-config-map"},
								Key:                  "key",
								Optional:             &trueVal,
							},
						},
					},
				},
			},
			masterServiceNs: "nothing",
			nilLister:       true,
			configMap: &v1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test1",
					Name:      "test-configmap",
				},
				Data: map[string]string{
					"a": "b",
				},
			},
			expectedEnvs: nil,
		},
		{
			name: "secretkeyref_missing_optional",
			ns:   "test",
			unit: &api.Unit{
				Env: []api.EnvVar{
					{
						Name: "POD_NAME",
						ValueFrom: &api.EnvVarSource{
							SecretKeyRef: &api.SecretKeySelector{
								LocalObjectReference: api.LocalObjectReference{Name: "missing-secret"},
								Key:                  "key",
								Optional:             &trueVal,
							},
						},
					},
				},
			},
			masterServiceNs: "nothing",
			expectedEnvs:    nil,
		},
		{
			name: "secretkeyref_missing_key_optional",
			ns:   "test",
			unit: &api.Unit{
				Env: []api.EnvVar{
					{
						Name: "POD_NAME",
						ValueFrom: &api.EnvVarSource{
							SecretKeyRef: &api.SecretKeySelector{
								LocalObjectReference: api.LocalObjectReference{Name: "test-secret"},
								Key:                  "key",
								Optional:             &trueVal,
							},
						},
					},
				},
			},
			masterServiceNs: "nothing",
			nilLister:       true,
			secret: &v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test1",
					Name:      "test-secret",
				},
				Data: map[string][]byte{
					"a": []byte("b"),
				},
			},
			expectedEnvs: nil,
		},
		{
			name: "configmap",
			ns:   "test1",
			unit: &api.Unit{
				EnvFrom: []api.EnvFromSource{
					{
						ConfigMapRef: &api.ConfigMapEnvSource{LocalObjectReference: api.LocalObjectReference{Name: "test-config-map"}},
					},
					{
						Prefix:       "p_",
						ConfigMapRef: &api.ConfigMapEnvSource{LocalObjectReference: api.LocalObjectReference{Name: "test-config-map"}},
					},
				},
				Env: []api.EnvVar{
					{
						Name:  "TEST_LITERAL",
						Value: "test-test-test",
					},
					{
						Name:  "EXPANSION_TEST",
						Value: "$(REPLACE_ME)",
					},
					{
						Name:  "DUPE_TEST",
						Value: "ENV_VAR",
					},
				},
			},
			masterServiceNs: "nothing",
			nilLister:       false,
			configMap: &v1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test1",
					Name:      "test-config-map",
				},
				Data: map[string]string{
					"REPLACE_ME": "FROM_CONFIG_MAP",
					"DUPE_TEST":  "CONFIG_MAP",
				},
			},
			expectedEnvs: []api.EnvVar{
				{
					Name:  "TEST_LITERAL",
					Value: "test-test-test",
				},
				{
					Name:  "REPLACE_ME",
					Value: "FROM_CONFIG_MAP",
				},
				{
					Name:  "EXPANSION_TEST",
					Value: "FROM_CONFIG_MAP",
				},
				{
					Name:  "DUPE_TEST",
					Value: "ENV_VAR",
				},
				{
					Name:  "p_REPLACE_ME",
					Value: "FROM_CONFIG_MAP",
				},
				{
					Name:  "p_DUPE_TEST",
					Value: "CONFIG_MAP",
				},
			},
		},
		{
			name: "configmap_missing",
			ns:   "test1",
			unit: &api.Unit{
				EnvFrom: []api.EnvFromSource{
					{ConfigMapRef: &api.ConfigMapEnvSource{LocalObjectReference: api.LocalObjectReference{Name: "test-config-map"}}},
				},
			},
			masterServiceNs: "nothing",
			expectedError:   true,
		},
		{
			name: "configmap_missing_optional",
			ns:   "test",
			unit: &api.Unit{
				EnvFrom: []api.EnvFromSource{
					{ConfigMapRef: &api.ConfigMapEnvSource{
						Optional:             &trueVal,
						LocalObjectReference: api.LocalObjectReference{Name: "missing-config-map"}}},
				},
			},
			masterServiceNs: "nothing",
			expectedEnvs:    nil,
		},
		{
			name: "configmap_invalid_keys",
			ns:   "test1",
			unit: &api.Unit{
				EnvFrom: []api.EnvFromSource{
					{ConfigMapRef: &api.ConfigMapEnvSource{LocalObjectReference: api.LocalObjectReference{Name: "test-config-map"}}},
				},
			},
			masterServiceNs: "nothing",
			configMap: &v1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test1",
					Name:      "test-config-map",
				},
				Data: map[string]string{
					"1234": "abc",
					"1z":   "abc",
					"key":  "value",
				},
			},
			expectedEnvs: []api.EnvVar{
				{
					Name:  "key",
					Value: "value",
				},
			},
			expectedEvent: "Warning InvalidEnvironmentVariableNames Keys [1234, 1z] from the EnvFrom configMap test/test-config-map were skipped since they are considered invalid environment variable names.",
		},
		{
			name: "configmap_invalid_keys_valid",
			ns:   "test1",
			unit: &api.Unit{
				EnvFrom: []api.EnvFromSource{
					{
						Prefix:       "p_",
						ConfigMapRef: &api.ConfigMapEnvSource{LocalObjectReference: api.LocalObjectReference{Name: "test-config-map"}},
					},
				},
			},
			masterServiceNs: "",
			configMap: &v1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test1",
					Name:      "test-config-map",
				},
				Data: map[string]string{
					"1234": "abc",
				},
			},
			expectedEnvs: []api.EnvVar{
				{
					Name:  "p_1234",
					Value: "abc",
				},
			},
		},
		{
			name: "secret",
			ns:   "test1",
			unit: &api.Unit{
				EnvFrom: []api.EnvFromSource{
					{
						SecretRef: &api.SecretEnvSource{LocalObjectReference: api.LocalObjectReference{Name: "test-secret"}},
					},
					{
						Prefix:    "p_",
						SecretRef: &api.SecretEnvSource{LocalObjectReference: api.LocalObjectReference{Name: "test-secret"}},
					},
				},
				Env: []api.EnvVar{
					{
						Name:  "TEST_LITERAL",
						Value: "test-test-test",
					},
					{
						Name:  "EXPANSION_TEST",
						Value: "$(REPLACE_ME)",
					},
					{
						Name:  "DUPE_TEST",
						Value: "ENV_VAR",
					},
				},
			},
			masterServiceNs: "nothing",
			nilLister:       false,
			secret: &v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test1",
					Name:      "test-secret",
				},
				Data: map[string][]byte{
					"REPLACE_ME": []byte("FROM_SECRET"),
					"DUPE_TEST":  []byte("SECRET"),
				},
			},
			expectedEnvs: []api.EnvVar{
				{
					Name:  "TEST_LITERAL",
					Value: "test-test-test",
				},
				{
					Name:  "REPLACE_ME",
					Value: "FROM_SECRET",
				},
				{
					Name:  "EXPANSION_TEST",
					Value: "FROM_SECRET",
				},
				{
					Name:  "DUPE_TEST",
					Value: "ENV_VAR",
				},
				{
					Name:  "p_REPLACE_ME",
					Value: "FROM_SECRET",
				},
				{
					Name:  "p_DUPE_TEST",
					Value: "SECRET",
				},
			},
		},
		{
			name: "secret_missing",
			ns:   "test1",
			unit: &api.Unit{
				EnvFrom: []api.EnvFromSource{
					{SecretRef: &api.SecretEnvSource{LocalObjectReference: api.LocalObjectReference{Name: "test-secret"}}},
				},
			},
			masterServiceNs: "nothing",
			expectedError:   true,
		},
		{
			name: "secret_missing_optional",
			ns:   "test",
			unit: &api.Unit{
				EnvFrom: []api.EnvFromSource{
					{SecretRef: &api.SecretEnvSource{
						LocalObjectReference: api.LocalObjectReference{Name: "missing-secret"},
						Optional:             &trueVal}},
				},
			},
			masterServiceNs: "nothing",
			expectedEnvs:    nil,
		},
		{
			name: "secret_invalid_keys",
			ns:   "test1",
			unit: &api.Unit{
				EnvFrom: []api.EnvFromSource{
					{SecretRef: &api.SecretEnvSource{LocalObjectReference: api.LocalObjectReference{Name: "test-secret"}}},
				},
			},
			masterServiceNs: "nothing",
			secret: &v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test1",
					Name:      "test-secret",
				},
				Data: map[string][]byte{
					"1234":  []byte("abc"),
					"1z":    []byte("abc"),
					"key.1": []byte("value"),
				},
			},
			expectedEnvs: []api.EnvVar{
				{
					Name:  "key.1",
					Value: "value",
				},
			},
			expectedEvent: "Warning InvalidEnvironmentVariableNames Keys [1234, 1z] from the EnvFrom secret test/test-secret were skipped since they are considered invalid environment variable names.",
		},
		{
			name: "secret_invalid_keys_valid",
			ns:   "test1",
			unit: &api.Unit{
				EnvFrom: []api.EnvFromSource{
					{
						Prefix:    "p_",
						SecretRef: &api.SecretEnvSource{LocalObjectReference: api.LocalObjectReference{Name: "test-secret"}},
					},
				},
			},
			masterServiceNs: "",
			secret: &v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test1",
					Name:      "test-secret",
				},
				Data: map[string][]byte{
					"1234.name": []byte("abc"),
				},
			},
			expectedEnvs: []api.EnvVar{
				{
					Name:  "p_1234.name",
					Value: "abc",
				},
			},
		},
	}

	client := nodeclient.NewMockItzoClientFactory()
	ctl, closer := createPodController(client)
	defer closer()
	for _, tc := range testCases {
		// create a resource manager
		// add our clients into there...

		//fakeRecorder := record.NewFakeRecorder(1)
		// testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
		// testKubelet.kubelet.recorder = fakeRecorder
		// defer testKubelet.Cleanup()
		// kl := testKubelet.kubelet
		//kl.masterServiceNamespace = tc.masterServiceNs
		// if tc.nilLister {
		// 	kl.serviceLister = nil
		// } else {
		// 	kl.serviceLister = testServiceLister{services}
		// }

		indexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
		if tc.configMap != nil {
			assert.Nil(t, indexer.Add(tc.configMap))
		}
		configMapLister := corev1listers.NewConfigMapLister(indexer)
		if tc.secret != nil {
			assert.Nil(t, indexer.Add(tc.secret))
		}
		secretLister := corev1listers.NewSecretLister(indexer)
		rm, err := manager.NewResourceManager(nil, secretLister, configMapLister, nil)
		if err != nil {
			t.Fatal(err)
		}
		ctl.resourceManager = rm
		// testKubelet.fakeKubeClient.AddReactor("get", "configmaps", func(action core.Action) (bool, runtime.Object, error) {
		// 	var err error
		// 	if tc.configMap == nil {
		// 		err = apierrors.NewNotFound(action.GetResource().GroupResource(), "configmap-name")
		// 	}
		// 	return true, tc.configMap, err
		// })
		// testKubelet.fakeKubeClient.AddReactor("get", "secrets", func(action core.Action) (bool, runtime.Object, error) {
		// 	var err error
		// 	if tc.secret == nil {
		// 		err = apierrors.NewNotFound(action.GetResource().GroupResource(), "secret-name")
		// 	}
		// 	return true, tc.secret, err
		// })

		// testKubelet.fakeKubeClient.AddReactor("get", "secrets", func(action core.Action) (bool, runtime.Object, error) {
		// 	var err error
		// 	if tc.secret == nil {
		// 		err = errors.New("no secret defined")
		// 	}
		// 	return true, tc.secret, err
		// })

		testPod := &api.Pod{
			ObjectMeta: api.ObjectMeta{
				Namespace: tc.ns,
				Name:      "dapi-test-pod-name",
			},
			Spec: api.PodSpec{},
			//ServiceAccountName: "special",
			// NodeName:           "node-name",
			//EnableServiceLinks: tc.enableServiceLinks,
			// },
		}
		result, err := ctl.makeUnitEnvVars(testPod, tc.unit)
		if tc.expectedError {
			assert.Error(t, err, tc.name)
		} else {
			assert.NoError(t, err, "[%s]", tc.name)
			sort.Sort(envs(result))
			sort.Sort(envs(tc.expectedEnvs))
			assert.Equal(t, tc.expectedEnvs, result, "[%s] env entries", tc.name)
		}
	}
}
