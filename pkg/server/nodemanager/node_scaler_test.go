package nodemanager

import (
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/stretchr/testify/assert"
)

type FakeNodeStatusUpdater struct{}

func (f *FakeNodeStatusUpdater) UpdateStatus(n *api.Node) (*api.Node, error) {
	return nil, nil
}

func MakeNodeScaler() (*BindingNodeScaler, func()) {
	nodeRegistry, closer := registry.SetupTestNodeRegistry()
	cloudStatus, _ := cloud.NewLinkedAZSubnetStatus(cloud.NewMockClient())
	return &BindingNodeScaler{
		nodeRegistry:      nodeRegistry,
		cloudStatus:       cloudStatus,
		defaultVolumeSize: "2G",
	}, closer
}

func TestSpotMatches(t *testing.T) {
	ns := &BindingNodeScaler{}
	tests := []struct {
		result          bool
		spotPolicy      api.SpotPolicy
		nodeSpot        bool
		spotUnavailable bool
	}{
		{true, api.SpotNever, false, false},
		{false, api.SpotNever, true, false},
		{true, api.SpotAlways, true, false},
		{false, api.SpotAlways, false, false},
	}
	for i, tc := range tests {
		pod := api.GetFakePod()
		pod.Spec.Spot.Policy = tc.spotPolicy
		node := api.GetFakeNode()
		node.Spec.Spot = tc.nodeSpot
		cloudStatus, _ := cloud.NewLinkedAZSubnetStatus(cloud.NewMockClient())
		if tc.spotUnavailable {
			cloudStatus.AddUnavailableInstance(pod.Spec.InstanceType, true)
		}
		ns.cloudStatus = cloudStatus
		matches := ns.spotMatches(pod, node)
		assert.Equal(t, tc.result, matches, "error on case %d", i)
	}
}

func TestPodMatchesNode(t *testing.T) {
	cloudStatus, _ := cloud.NewLinkedAZSubnetStatus(cloud.NewMockClient())
	ns := BindingNodeScaler{cloudStatus: cloudStatus, defaultVolumeSize: "5G"}
	pod := api.GetFakePod()
	node := ns.createNodeForPod(pod)
	assert.True(t, ns.podMatchesNode(pod, node))
	p2 := *pod
	p2.Spec.Placement.AvailabilityZone = "us-east-1a"
	assert.False(t, ns.podMatchesNode(&p2, node))
	p2 = *pod
	p2.Spec.InstanceType = "p3.xlarge"
	assert.False(t, ns.podMatchesNode(&p2, node))
	p2 = *pod
	p2.Spec.Resources.PrivateIPOnly = true
	assert.False(t, ns.podMatchesNode(&p2, node))
	p2 = *pod
	p2.Spec.Spot.Policy = api.SpotAlways
	assert.False(t, ns.podMatchesNode(&p2, node))
}

func TestCreateNodeForPodUnavailable(t *testing.T) {
	cloudStatus, _ := cloud.NewLinkedAZSubnetStatus(cloud.NewMockClient())
	ns := BindingNodeScaler{cloudStatus: cloudStatus, defaultVolumeSize: "5G"}

	// if not a spot pod
	pod := api.GetFakePod()
	node := ns.createNodeForPod(pod)
	assert.False(t, node.Spec.Spot)
	assert.Equal(t, pod.Name, node.Status.BoundPodName)

	// now with spot
	pod.Spec.Spot.Policy = api.SpotAlways
	node = ns.createNodeForPod(pod)
	assert.True(t, node.Spec.Spot)

	// Now it's unavailable...
	ns.cloudStatus.AddUnavailableInstance(pod.Spec.InstanceType, true)
	node = ns.createNodeForPod(pod)
	assert.Nil(t, node)

	// Now it's totally unavailable we shouldn't get a node back
	ns.cloudStatus.AddUnavailableInstance(pod.Spec.InstanceType, true)
	ns.cloudStatus.AddUnavailableInstance(pod.Spec.InstanceType, false)
	node = ns.createNodeForPod(pod)
	assert.Nil(t, node)
}

func TestCreateNodeForPodVolumeSize(t *testing.T) {
	cloudStatus, _ := cloud.NewLinkedAZSubnetStatus(cloud.NewMockClient())
	defaultVolumeSize := "5G"
	ns := BindingNodeScaler{
		cloudStatus:       cloudStatus,
		defaultVolumeSize: defaultVolumeSize,
	}

	pod := api.GetFakePod()
	node := ns.createNodeForPod(pod)
	assert.Equal(t, defaultVolumeSize, node.Spec.Resources.VolumeSize)

	ns.fixedSizeVolume = true
	node = ns.createNodeForPod(pod)
	assert.Equal(t, defaultVolumeSize, node.Spec.Resources.VolumeSize)

	largerVolume := "10G"
	pod.Spec.Resources.VolumeSize = largerVolume
	node = ns.createNodeForPod(pod)
	assert.Equal(t, largerVolume, node.Spec.Resources.VolumeSize)
}

func TestPlacementMatches(t *testing.T) {
	node := api.GetFakeNode()
	node.Spec.Placement.AvailabilityZone = "us-east-1a"
	pod := api.GetFakePod()
	assert.True(t, placementMatches(pod, node))
	pod.Spec.Placement.AvailabilityZone = "us-east-1a"
	assert.True(t, placementMatches(pod, node))
	pod.Spec.Placement.AvailabilityZone = "us-west-1a"
	assert.False(t, placementMatches(pod, node))
}

func TestCreateNodeForStandbySpec(t *testing.T) {
	cloudStatus, _ := cloud.NewLinkedAZSubnetStatus(cloud.NewMockClient())
	ns := BindingNodeScaler{cloudStatus: cloudStatus, defaultVolumeSize: "5G"}

	sb := &StandbyNodeSpec{
		InstanceType: "t2.nano",
		Spot:         false,
	}
	node := ns.createNodeForStandbySpec(sb)
	assert.True(t, ns.nodeMatchesStandbySpec(node, sb))
	sb = &StandbyNodeSpec{
		InstanceType: "t3.large",
		Spot:         true,
	}
	node = ns.createNodeForStandbySpec(sb)
	assert.True(t, ns.nodeMatchesStandbySpec(node, sb))
}

func TestNodeMatchesStandbySpec(t *testing.T) {
	cloudStatus, _ := cloud.NewLinkedAZSubnetStatus(cloud.NewMockClient())
	ns := BindingNodeScaler{cloudStatus: cloudStatus, defaultVolumeSize: "5G"}

	sb := &StandbyNodeSpec{
		InstanceType: "t2.nano",
		Spot:         false,
	}
	node := api.Node{}
	node.Spec.InstanceType = sb.InstanceType
	node.Spec.Resources.VolumeSize = "5G"
	assert.True(t, ns.nodeMatchesStandbySpec(&node, sb))
	node.Spec.Resources.VolumeSize = "10G"
	assert.False(t, ns.nodeMatchesStandbySpec(&node, sb))
	node.Spec.Resources.VolumeSize = "5G"

	n2 := node
	n2.Spec.Spot = true
	assert.False(t, ns.nodeMatchesStandbySpec(&n2, sb))
	n3 := node
	n3.Spec.InstanceType = "m3.large"
	assert.False(t, ns.nodeMatchesStandbySpec(&n3, sb))
}

func makeNodeScaler() (*BindingNodeScaler, func()) {
	nodeRegistry, closer := registry.SetupTestNodeRegistry()
	cloudStatus, _ := cloud.NewLinkedAZSubnetStatus(cloud.NewMockClient())
	ns := &BindingNodeScaler{
		nodeRegistry:      nodeRegistry,
		standbyNodes:      []StandbyNodeSpec{},
		cloudStatus:       cloudStatus,
		defaultVolumeSize: "5G",
	}
	return ns, closer
}

func TestNodeScalerUnbindsAndRebinds(t *testing.T) {
	ns, closer := makeNodeScaler()
	defer closer()

	pod := api.GetFakePod()
	node := ns.createNodeForPod(pod)
	nodeReg := ns.nodeRegistry.(*registry.NodeRegistry)
	_, _ = nodeReg.CreateNode(node)
	node.Status.BoundPodName = "doesn't exist"
	start, stop, bindings := ns.Compute([]*api.Node{node}, []*api.Pod{pod})
	assert.Len(t, start, 0)
	assert.Len(t, stop, 0)
	assert.Equal(t, node.Name, bindings[pod.Name])
	n, err := nodeReg.GetNode(node.Name)
	assert.NoError(t, err)
	assert.Equal(t, pod.Name, n.Status.BoundPodName)
}

func TestMachingPodToStandbyPool(t *testing.T) {
	ns, closer := makeNodeScaler()
	defer closer()
	pod := api.GetFakePod()

	standbySpec := StandbyNodeSpec{
		InstanceType: pod.Spec.InstanceType,
		Spot:         false,
		Count:        1,
	}
	ns.standbyNodes = []StandbyNodeSpec{standbySpec}
	node := ns.createNodeForStandbySpec(&standbySpec)
	nodeReg := ns.nodeRegistry.(*registry.NodeRegistry)
	_, _ = nodeReg.CreateNode(node)
	start, stop, bindings := ns.Compute([]*api.Node{node}, []*api.Pod{pod})
	assert.Len(t, start, 1)
	assert.Len(t, stop, 0)
	assert.Equal(t, node.Name, bindings[pod.Name])
}

func TestNodeScalerDiskMatches(t *testing.T) {
	defaultVolumeSize := "5G"
	tests := []struct {
		podVolSize  string
		nodeVolSize string
		result      bool
	}{
		{podVolSize: defaultVolumeSize, nodeVolSize: defaultVolumeSize, result: true},
		{podVolSize: "", nodeVolSize: defaultVolumeSize, result: true},
		{podVolSize: "6G", nodeVolSize: defaultVolumeSize, result: false},
	}
	for i, tc := range tests {
		ns := BindingNodeScaler{
			defaultVolumeSize: defaultVolumeSize,
			fixedSizeVolume:   true,
		}
		pod := api.GetFakePod()
		pod.Spec.Resources.VolumeSize = tc.podVolSize
		node := api.GetFakeNode()
		node.Spec.Resources.VolumeSize = tc.nodeVolSize
		matches := ns.diskMatches(pod, node)
		assert.Equal(t, tc.result, matches, "error on case %d", i)
		ns.fixedSizeVolume = false
		matches = ns.diskMatches(pod, node)
		// resizable volumes always match
		assert.Equal(t, true, matches, "error on fixedSizeVolume = false case %d", i)
	}
}
