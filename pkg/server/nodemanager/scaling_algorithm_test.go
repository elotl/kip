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
	"sort"
	"testing"

	"github.com/elotl/kip/pkg/api"
	"github.com/stretchr/testify/assert"
)

type ScalingTestEntryInput struct {
	pods  []*api.Pod
	nodes []*api.Node
}

type ScalingTestEntry struct {
	input ScalingTestEntryInput
	start []*api.Node
	stop  []*api.Node
}

func mkin(vals ...interface{}) ScalingTestEntryInput {
	pods := make([]*api.Pod, 0)
	nodes := make([]*api.Node, 0)
	for _, v := range vals {
		switch v.(type) {
		case *api.Pod:
			pods = append(pods, v.(*api.Pod))
		case *api.Node:
			nodes = append(nodes, v.(*api.Node))
		}
	}
	return ScalingTestEntryInput{
		pods:  pods,
		nodes: nodes,
	}
}
func sortNodes(nodes []*api.Node) {
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].Spec.BootImage != nodes[j].Spec.BootImage {
			return nodes[i].Spec.BootImage < nodes[j].Spec.BootImage
		} else if nodes[i].Spec.InstanceType != nodes[j].Spec.InstanceType {
			return nodes[i].Spec.InstanceType < nodes[j].Spec.InstanceType
		} else if nodes[i].Spec.Spot != nodes[j].Spec.Spot {
			return nodes[i].Spec.Spot
		} else if nodes[i].Spec.Resources.PrivateIPOnly != nodes[j].Spec.Resources.PrivateIPOnly {
			return !nodes[i].Spec.Resources.PrivateIPOnly
		}
		return false
	})
}

func nodesEqual(t *testing.T, output, expected []*api.Node, msgArgs ...interface{}) {
	sortNodes(output)
	sortNodes(expected)
	assert.Equal(t, len(expected), len(output), msgArgs...)
	for i := 0; i < len(output); i++ {
		assert.Equal(t, output[i].Spec.InstanceType, expected[i].Spec.InstanceType, msgArgs...)
		assert.Equal(t, output[i].Spec.Resources.PrivateIPOnly, expected[i].Spec.Resources.PrivateIPOnly, msgArgs...)
		assert.Equal(t, output[i].Spec.Spot, expected[i].Spec.Spot, msgArgs...)
	}
}

func verify(t *testing.T, scaler ScalingAlgorithm, entry ScalingTestEntry, i int) {
	startNodes, stopNodes, _ := scaler.Compute(entry.input.nodes, entry.input.pods)
	nodesEqual(t, startNodes, entry.start, "test %d start nodes failed", i)
	nodesEqual(t, stopNodes, entry.stop, "test %d stop nodes failed", i)
}

func mkPod(instance string, privateIP bool) *api.Pod {
	p := api.GetFakePod()
	p.Spec.Resources.PrivateIPOnly = privateIP
	p.Spec.InstanceType = instance
	return p
}

func mkNode(instance string, privateIP bool, spot bool) *api.Node {
	n := api.GetFakeNode()
	n.Spec.InstanceType = instance
	n.Spec.Resources.PrivateIPOnly = privateIP
	n.Spec.Spot = spot
	return n
}

func mkout(vals ...*api.Node) []*api.Node {
	nodes := make([]*api.Node, 0)
	return append(nodes, vals...)
}

func TestOnDemandScalingNoStandby(t *testing.T) {
	t.Parallel()
	noNodes := []*api.Node{}
	instanceA := "t2.A"
	instanceB := "t2.B"
	pa1 := mkPod(instanceA, false)
	pa2 := mkPod(instanceA, false)
	pb1 := mkPod(instanceB, true)
	na1 := mkNode(instanceA, false, false)
	na2 := mkNode(instanceA, false, false)
	nb1 := mkNode(instanceB, true, false)
	table := []ScalingTestEntry{
		// no pods, no nodes
		{mkin(), noNodes, noNodes},
		// 1 pod, no nodes -- 1 node
		{mkin(pa1), mkout(na1), noNodes},
		// 2 pods (different images), no nodes -- 2 nodes
		{mkin(pa1, pb1), mkout(na1, nb1), noNodes},
		// 2 different pods, 1 node of input pod type -- 1 node
		{mkin(pa1, pb1, na1), mkout(nb1), noNodes},
		// 2 pods (different images), 1 node with same image -- 1 node
		{mkin(pa1, pa2, na1), mkout(na1), noNodes},
		// 1 node -- delete 1 node
		{mkin(na1), noNodes, mkout(na1)},
		// 2 nodes identical types -- delete 2 nodes
		{mkin(na1, na2), noNodes, mkout(na1, na2)},
		// 2 nodes different -- delete 2 nodes
		{mkin(na1, nb1), noNodes, mkout(na1, nb1)},
		// 1 pod, 2 nodes (one same, one different)
		{mkin(pa1, pa2, na1, nb1), mkout(na1), mkout(nb1)},
	}
	scaler := &BindingNodeScaler{
		nodeRegistry: &FakeNodeStatusUpdater{},
		bootLimiter:  NewInstanceBootLimiter(),
		standbyNodes: nil,
		getArchitecture: getArchitecture,
	}
	for i, entry := range table {
		fmt.Println("test entry", i)
		verify(t, scaler, entry, i)
	}
}

func TestOnDemandScalingWithStandby(t *testing.T) {
	t.Parallel()
	noNodes := []*api.Node{}
	instanceA := "t2.A"
	instanceB := "t2.B"
	pa1 := mkPod(instanceA, false)
	pb1 := mkPod(instanceB, true)
	na1 := mkNode(instanceA, false, false)
	nb1 := mkNode(instanceB, true, false)
	table := []ScalingTestEntry{
		{mkin(), mkout(na1, na1), noNodes},
		{mkin(pa1, na1, na1), mkout(na1), noNodes},
		{mkin(na1, na1, na1), noNodes, mkout(na1)},
		{mkin(pb1, na1), mkout(na1, nb1), noNodes},
	}
	standbyNodes := []StandbyNodeSpec{
		{
			InstanceType: instanceA,
			Count:        2,
			Spot:         false,
		},
		{
			// this tests for a bug where we buffered a node
			// for a pool with a count of 0
			InstanceType: instanceA,
			Count:        0,
			Spot:         false,
		},
	}
	scaler := &BindingNodeScaler{
		nodeRegistry: &FakeNodeStatusUpdater{},
		bootLimiter:  NewInstanceBootLimiter(),
		standbyNodes: standbyNodes,
		getArchitecture: getArchitecture,
	}
	for i, entry := range table {
		fmt.Println("test entry", i)
		verify(t, scaler, entry, i)
	}
}

func TestSpotScaling(t *testing.T) {
	noNodes := []*api.Node{}
	instanceA := "t2.A"
	instanceB := "t2.B"
	instanceC := "t2.C"
	pa1 := mkPod(instanceA, false)
	pa1.Spec.Spot.Policy = api.SpotAlways
	pa2 := mkPod(instanceA, false)

	pa3 := mkPod(instanceA, false)
	pa3.Spec.Spot.Policy = api.SpotAlways

	pa4 := mkPod(instanceA, false)
	pa4.Spec.Phase = api.PodFailed
	pa4.Status.Phase = api.PodFailed

	pa5 := mkPod(instanceA, false)
	pa5.Spec.Spot.Policy = api.SpotAlways

	pb1 := mkPod(instanceB, false)
	pb1.Spec.Spot.Policy = api.SpotAlways
	pc1 := mkPod(instanceC, false)
	pc1.Spec.Spot.Policy = api.SpotAlways
	na1 := mkNode(instanceA, false, true)
	na2 := mkNode(instanceA, false, false)
	na3 := mkNode(instanceA, false, true)
	nb1 := mkNode(instanceB, false, true)
	table := []ScalingTestEntry{
		// no pods, no nodes
		{mkin(), noNodes, noNodes},
		// 1 spot pod, no nodes -- 1 node
		{mkin(pa1), mkout(na1), noNodes},
		// 2 pods, one spot, the other not, no nodes -- 1 spot node
		{mkin(pa1, pa2), mkout(na1, na2), noNodes},
		// 1 spot pod with a corresponding spot node
		{mkin(pa1, na1), noNodes, noNodes},
		// 1 spot pod with spot node and another node
		{mkin(pa1, na2), mkout(na1), mkout(na2)},
		// 2 pods, both spot with corresponding spot nodes
		{mkin(pa1, pa3, na1, na3), noNodes, noNodes},
		// make sure a running pod doesn't influece pods we need
		{mkin(pa4), noNodes, noNodes},
		// spot nodes without pods should be deleted
		{mkin(na1), noNodes, mkout(na1)},
		{mkin(pa1, na1, na3), noNodes, mkout(na3)},

		{mkin(pa1, pb1), mkout(na1, nb1), noNodes},

		// make sure that we don't try to boot nodes for unavailable spot pods
		{mkin(pc1), noNodes, noNodes},
	}

	bootLimiter := NewInstanceBootLimiter()
	bootLimiter.AddUnavailableInstance(pc1.Spec.InstanceType, true)
	scaler := &BindingNodeScaler{
		nodeRegistry:    &FakeNodeStatusUpdater{},
		bootLimiter:     bootLimiter,
		standbyNodes:    nil,
		getArchitecture: getArchitecture,
	}

	for i, entry := range table {
		fmt.Println("test entry", i)
		verify(t, scaler, entry, i)
	}
}
