package server

import (
	"fmt"
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/api"
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
