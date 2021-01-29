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
	"testing"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/server/registry"
	"github.com/stretchr/testify/assert"
)

type FakeNodeStatusUpdater struct{}

func (f *FakeNodeStatusUpdater) UpdateStatus(n *api.Node) (*api.Node, error) {
	return nil, nil
}

func MakeNodeScaler() (*BindingNodeScaler, func()) {
	nodeRegistry, closer := registry.SetupTestNodeRegistry()
	bootLimiter := NewInstanceBootLimiter()
	return &BindingNodeScaler{
		nodeRegistry:      nodeRegistry,
		bootLimiter:       bootLimiter,
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
		bootLimiter := NewInstanceBootLimiter()
		if tc.spotUnavailable {
			bootLimiter.AddUnavailableInstance(pod.Spec.InstanceType, true)
		}
		ns.bootLimiter = bootLimiter
		matches := ns.spotMatches(pod, node)
		assert.Equal(t, tc.result, matches, "error on case %d", i)
	}
}

func TestPodMatchesNode(t *testing.T) {
	// XXX: BootImages is a global variable. We should certainely find a
	// a more elegant way to test this.
	BootImages = map[api.Architecture]cloud.Image{api.ArchX8664: cloud.Image{}}
	bootLimiter := NewInstanceBootLimiter()
	ns := BindingNodeScaler{
		bootLimiter:       bootLimiter,
		defaultVolumeSize: "5G",
	}
	pod := api.GetFakePod()
	node := ns.createNodeForPod(pod)
	assert.True(t, ns.podMatchesNode(pod, node))
	p2 := *pod
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
	bootLimiter := NewInstanceBootLimiter()
	ns := BindingNodeScaler{
		bootLimiter:       bootLimiter,
		defaultVolumeSize: "5G",
	}

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
	ns.bootLimiter.AddUnavailableInstance(pod.Spec.InstanceType, true)
	node = ns.createNodeForPod(pod)
	assert.Nil(t, node)

	// Now it's totally unavailable we shouldn't get a node back
	ns.bootLimiter.AddUnavailableInstance(pod.Spec.InstanceType, true)
	ns.bootLimiter.AddUnavailableInstance(pod.Spec.InstanceType, false)
	node = ns.createNodeForPod(pod)
	assert.Nil(t, node)
}

func TestCreateNodeForPodVolumeSize(t *testing.T) {
	bootLimiter := NewInstanceBootLimiter()
	defaultVolumeSize := "5G"
	ns := BindingNodeScaler{
		bootLimiter:       bootLimiter,
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

func TestCreateNodeForStandbySpec(t *testing.T) {
	bootLimiter := NewInstanceBootLimiter()
	ns := BindingNodeScaler{
		bootLimiter:       bootLimiter,
		defaultVolumeSize: "5G",
	}

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
	bootLimiter := NewInstanceBootLimiter()
	ns := BindingNodeScaler{
		bootLimiter:       bootLimiter,
		defaultVolumeSize: "5G",
	}

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
	ns := &BindingNodeScaler{
		nodeRegistry:      nodeRegistry,
		standbyNodes:      []StandbyNodeSpec{},
		bootLimiter:       NewInstanceBootLimiter(),
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

func TestFullStandbyPool(t *testing.T) {
	ns, closer := makeNodeScaler()
	defer closer()
	standbySpec := StandbyNodeSpec{
		InstanceType: "t3.nano",
		Spot:         false,
		Count:        2,
	}
	ns.standbyNodes = []StandbyNodeSpec{standbySpec}
	nodeReg := ns.nodeRegistry.(*registry.NodeRegistry)
	n1 := ns.createNodeForStandbySpec(&standbySpec)
	_, err := nodeReg.CreateNode(n1)
	assert.NoError(t, err)
	n2 := ns.createNodeForStandbySpec(&standbySpec)
	_, err = nodeReg.CreateNode(n2)
	assert.NoError(t, err)
	start, stop, _ := ns.Compute([]*api.Node{n1, n2}, []*api.Pod{})
	assert.Len(t, start, 0)
	assert.Len(t, stop, 0)
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
