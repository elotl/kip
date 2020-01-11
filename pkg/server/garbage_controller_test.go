package server

import (
	"fmt"
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/elotl/cloud-instance-provider/pkg/util/sets"
	"github.com/stretchr/testify/assert"
)

func createGarbageController() (*GarbageController, func()) {
	// quit := make(chan struct{})
	// wg := &sync.WaitGroup{}
	podRegistry, closer1 := registry.SetupTestPodRegistry()
	nodeRegistry, closer2 := registry.SetupTestNodeRegistry()
	closer := func() { closer1(); closer2() }
	ctl := &GarbageController{
		podRegistry:  podRegistry,
		nodeRegistry: nodeRegistry,
	}
	return ctl, closer
}

type MockResourcer struct {
	groups sets.String
}

func (r *MockResourcer) ListNodeResourceGroups() ([]string, error) {
	return r.groups.List(), nil
}

func (r *MockResourcer) DeleteResourceGroup(name string) error {
	if r.groups.Has(name) {
		r.groups.Delete(name)
		return nil
	}
	return fmt.Errorf("Could not find resource")
}

func TestCleanAzureResourceGroupsHelper(t *testing.T) {
	ctl, closer := createGarbageController()
	defer closer()

	// Resource Groups
	groups := []string{"milpa-testcluster-1", "milpa-testcluster-2", "milpa-testcluster-3"}
	r := &MockResourcer{
		groups: sets.NewString(groups...),
	}
	// Node
	n := api.GetFakeNode()
	n.Name = "3"
	n.Status.InstanceID = "milpa-testcluster-3"
	_, err := ctl.nodeRegistry.CreateNode(n)
	assert.NoError(t, err)

	// Run it twice, and ensure we delete things on the second time through
	err = ctl.CleanAzureResourceGroupsHelper(r)
	assert.NoError(t, err)
	assert.Equal(t, 3, r.groups.Len())
	assert.Equal(t, 2, ctl.lastOrphanedAzureGroups.Len())
	err = ctl.CleanAzureResourceGroupsHelper(r)
	assert.NoError(t, err)
	assert.Equal(t, 1, r.groups.Len())
	assert.True(t, r.groups.Has("milpa-testcluster-3"))
	assert.Equal(t, 0, ctl.lastOrphanedAzureGroups.Len())
}

func TestCleanUnknownContainerInstance(t *testing.T) {
	ctl, closer := createGarbageController()
	defer closer()

	ctl.cloudClient = cloud.NewMockClient()
	p1 := api.GetFakePod()
	p1.Status.Phase = api.PodRunning
	tv := true
	p1.Spec.Resources.ContainerInstance = &tv
	p1.Status.BoundInstanceID = "1234"
	ctl.cloudClient.StartContainerInstance(p1)

	p2 := api.GetFakePod()
	p2.Status.Phase = api.PodRunning
	p2.Spec.Resources.ContainerInstance = &tv
	p2.Spec.InstanceType = ""
	p2.Status.BoundInstanceID = "5678"
	_, err := ctl.podRegistry.CreatePod(p2)
	assert.NoError(t, err)
	_, err = ctl.cloudClient.StartContainerInstance(p2)

	assert.NoError(t, err)
	cloudInsts, err := ctl.cloudClient.ListContainerInstances()
	assert.NoError(t, err)
	assert.Len(t, cloudInsts, 2)
	ctl.cleanUnknownContainerInstances()
	cloudInsts, err = ctl.cloudClient.ListContainerInstances()
	assert.NoError(t, err)
	assert.Len(t, cloudInsts, 2)
	ctl.cleanUnknownContainerInstances()
	cloudInsts, err = ctl.cloudClient.ListContainerInstances()
	assert.NoError(t, err)
	assert.Len(t, cloudInsts, 1)
	if len(cloudInsts) == 1 {
		assert.Equal(t, "5678", cloudInsts[0].ID)
	}
}

func TestGetOutdatedTaskDefinitions(t *testing.T) {
	tests := []struct {
		taskARNs []string
		podNames []string
		expected []string
	}{
		{
			taskARNs: []string{},
			podNames: []string{"foo", "bar"},
			expected: []string{},
		},
		{
			taskARNs: []string{"arn:aws:ecs:501:td/milpa-tc_foo:5", "arn:aws:ecs:501:td/milpa-tc_bar:1"},
			podNames: []string{"foo", "bar"},
			expected: []string{},
		},
		{
			taskARNs: []string{"arn:aws:ecs:501:td/milpa-tc_foo:5", "arn:aws:ecs:501:td/milpa-tc_foo:2", "arn:aws:ecs:501:td/milpa-tc_foo:1"},
			podNames: []string{"foo", "bar"},
			expected: []string{"arn:aws:ecs:501:td/milpa-tc_foo:2", "arn:aws:ecs:501:td/milpa-tc_foo:1"},
		},
		{
			taskARNs: []string{"arn:aws:ecs:501:td/milpa-tc_baz:1"},
			podNames: []string{"foo", "bar"},
			expected: []string{"arn:aws:ecs:501:td/milpa-tc_baz:1"},
		},
	}
	controllerID := "tc"
	for i, tc := range tests {
		podNamesSet := sets.NewString(tc.podNames...)
		oldTaskDefs := getOutdatedTaskDefinitions(tc.taskARNs, podNamesSet, controllerID)
		if !oldTaskDefs.Equal(sets.NewString(tc.expected...)) {
			t.Errorf("test case %d failed, expected equal: %v %v", i, tc.expected, oldTaskDefs.List())
		}
	}
}
