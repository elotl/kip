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
	"time"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/util"
	"k8s.io/klog"
)

type StatusUpdater interface {
	UpdateStatus(*api.Node) (*api.Node, error)
}

type BindingNodeScaler struct {
	nodeRegistry      StatusUpdater
	standbyNodes      []StandbyNodeSpec
	bootLimiter       *InstanceBootLimiter
	defaultVolumeSize string
	fixedSizeVolume   bool
}

func NewBindingNodeScaler(nodeReg StatusUpdater, standbyNodes []StandbyNodeSpec, bootLimiter *InstanceBootLimiter, defaultVolumeSize string, fixedSizeVolume bool) *BindingNodeScaler {
	return &BindingNodeScaler{
		nodeRegistry:      nodeReg,
		standbyNodes:      standbyNodes,
		bootLimiter:       bootLimiter,
		defaultVolumeSize: defaultVolumeSize,
		fixedSizeVolume:   fixedSizeVolume,
	}
}

// We try to match spotAlways and spotPreferred to any spot nodes.
// but if we're spotPreferred and we have unavailability, we allow the
// pod to match to a non-spot node.
func (s *BindingNodeScaler) spotMatches(pod *api.Pod, node *api.Node) bool {
	if (pod.Spec.Spot.Policy == api.SpotNever && !node.Spec.Spot) ||
		(pod.Spec.Spot.Policy == api.SpotAlways && node.Spec.Spot) {
		return true
	}
	return false
}

func (s *BindingNodeScaler) podMatchesNode(pod *api.Pod, node *api.Node) bool {
	return node.Spec.InstanceType == pod.Spec.InstanceType &&
		node.Spec.Resources.PrivateIPOnly == pod.Spec.Resources.PrivateIPOnly &&
		node.Spec.Resources.GPU == pod.Spec.Resources.GPU &&
		s.spotMatches(pod, node) &&
		s.diskMatches(pod, node)
}

func (s *BindingNodeScaler) diskMatches(pod *api.Pod, node *api.Node) bool {
	if s.fixedSizeVolume {
		// when we can't resize a disk, you must match on disk size
		// and to keep things simple, you have to match exactly
		return (pod.Spec.Resources.VolumeSize == node.Spec.Resources.VolumeSize ||
			(pod.Spec.Resources.VolumeSize == "" && node.Spec.Resources.VolumeSize == s.defaultVolumeSize))
	} else {
		return true
	}
}

func (s *BindingNodeScaler) createNodeForPod(pod *api.Pod) *api.Node {
	isSpotPod := false
	if pod.Spec.Spot.Policy == api.SpotAlways {
		isSpotPod = true
	}
	if s.bootLimiter.IsUnavailableInstance(pod.Spec.InstanceType, isSpotPod) {
		return nil
	}
	// XXX henry: we need the instance’s architecture to be able to pick the right image.
	// This is currently handled by CloudClient, which isn’t accessible from this object.
	// We assume x84_64 because everybody still uses it.
	image, found := BootImages[api.ArchX8664]
	if !found {
		klog.Errorf("Error could not find image for instance type: %s", pod.Spec.InstanceType)
		return nil
	}
	node := api.NewNode()
	node.Spec.InstanceType = pod.Spec.InstanceType
	node.Spec.BootImage = image.ID
	node.Spec.Spot = isSpotPod
	node.Spec.Resources = pod.Spec.Resources
	// If we can resize, keep things simple and never enlarge the disk
	// until dispatch (just launch with the default size), otherwise,
	// if we have a fixed size disk of non-default size, size the disk
	// here and also take care of matching disk sizes
	node.Spec.Resources.VolumeSize = s.defaultVolumeSize
	if s.fixedSizeVolume && pod.Spec.Resources.VolumeSize != "" {
		node.Spec.Resources.VolumeSize = pod.Spec.Resources.VolumeSize
	}
	node.Status.BoundPodName = pod.Name
	return node
}

func (s *BindingNodeScaler) createNodeForStandbySpec(spec *StandbyNodeSpec) *api.Node {
	// XXX henry: we need the instance’s architecture to be able to pick the right image.
	// This is currently handled by CloudClient, which isn’t accessible from this object.
	// We assume x84_64 because everybody still uses it.
	image, found := BootImages[api.ArchX8664]
	if !found {
		klog.Errorf("Error could not find image for instance type: %s", spec.InstanceType)
		return nil
	}
	node := api.NewNode()
	node.Spec.InstanceType = spec.InstanceType
	node.Spec.BootImage = image.ID
	node.Spec.Spot = spec.Spot
	node.Spec.Resources.VolumeSize = s.defaultVolumeSize
	return node
}

func (s *BindingNodeScaler) nodeMatchesStandbySpec(node *api.Node, spec *StandbyNodeSpec) bool {
	return node.Spec.Spot == spec.Spot &&
		node.Spec.InstanceType == spec.InstanceType &&
		node.Spec.Resources.VolumeSize == s.defaultVolumeSize
}

// A brief summary of how we figure out what nodes need to be started
// and what nodes need to be shut down:
//
// 1. We only care about watiting pods and available or creat(ing|ed)
// nodes.
//
// 2. Re-generate the podNodeBinding map by looking at the existing
// bindings from nodes to pods in each node's
// node.Status.BoundPodName. Also make sure that any pods listed in
// there are actually still waiting (because the user might have
// killed a pod). Along the way we keep track of unbound pods and
// nodes.
//
// 3. Match any unbound pods to unbound nodes.  Before doing that, we
// need to ensure that we order our pods and nodes so that we choose
// the most specific matches for our pods and nodes).
//
// 4. Any remaining unbound pods that haven't been matched will get a
// node booted for them with the exception of node requests that we
// know cannot be fulfilled due to unavailability in the cloud.
//
// 5. Finally, make sure that we have enough nodes to satisfy our
// standby pools of nodes.
//
// At the end of this process, return the nodes that we should start,
// the nodes that need to be shut down and the current bindings map
// (so that the dispatcher can be fast).
func (s *BindingNodeScaler) Compute(nodes []*api.Node, pods []*api.Pod) ([]*api.Node, []*api.Node, map[string]string) {
	// we only care about nodes in availableOrBaking...
	// remove bound nodes and pods, leaving only needyPods and unboundNodes
	newNodes := make([]*api.Node, 0)
	dirtyNodes := make(map[string]*api.Node)

	waitingPods := make(map[string]*api.Pod)
	for _, pod := range pods {
		if pod.Status.Phase == api.PodWaiting {
			waitingPods[pod.Name] = pod
		}
	}

	podNodeBinding := make(map[string]string)
	unboundNodes := make([]*api.Node, 0, len(nodes))
	for _, node := range nodes {
		if availableOrBaking(node) {
			boundPodName := node.Status.BoundPodName
			if boundPodName != "" && waitingPods[boundPodName] != nil {
				podNodeBinding[node.Status.BoundPodName] = node.Name
			} else {
				// There's no waiting pod for this node, unbind it
				if boundPodName != "" {
					node.Status.BoundPodName = ""
					dirtyNodes[node.Name] = node
				}
				unboundNodes = append(unboundNodes, node)
			}
		}
	}
	unboundPods := make([]*api.Pod, 0, len(waitingPods))
	for _, pod := range waitingPods {
		if val, exists := podNodeBinding[pod.Name]; !exists || val == "" {
			unboundPods = append(unboundPods, pod)
		}
	}
	// Prioritize matching to spot nodes by putting them at the front
	// of the slice.  We do this because we might have unavilability
	// in spot nodes but still have some sitting around that we've
	// already booted (maybe we wanted 10 spot nodes and only got
	// 5...)
	util.PartitionSlice(unboundNodes, func(i int) bool {
		return unboundNodes[i].Spec.Spot
	})

	// match needyPods to any unboundNodes O(n**2) but with any luck,
	// the unboundNodes pool should be the size of the buffered nodes pool
	// which is typically small...
	needyPods := make([]*api.Pod, 0, len(unboundPods))
	for _, pod := range unboundPods {
		matched := false
		for i, node := range unboundNodes {
			if s.podMatchesNode(pod, node) {
				node.Status.BoundPodName = pod.Name
				podNodeBinding[pod.Name] = node.Name
				unboundNodes = append(unboundNodes[:i], unboundNodes[i+1:]...)
				dirtyNodes[node.Name] = node
				matched = true
				break
			}
		}
		if !matched {
			needyPods = append(needyPods, pod)
		}
	}

	// for all remaining pods, create new nodes for those and add a
	// BoundPodName to them as well
	for _, pod := range needyPods {
		newNode := s.createNodeForPod(pod)
		if newNode != nil {
			podNodeBinding[pod.Name] = newNode.Name
			newNodes = append(newNodes, newNode)
		}
	}
	// for all nodes that remain, match them to the buffered pool spec.
	// Keep track of the number of nodes each pool needs and keep
	// unbound nodes up to date by
	standbyToNumNeeded := make(map[StandbyNodeSpec]int)
	for _, standbySpec := range s.standbyNodes {
		neededNodes := standbySpec.Count
		for i := 0; i < len(unboundNodes); i++ {
			node := unboundNodes[i]
			if neededNodes == 0 {
				break
			}
			if s.nodeMatchesStandbySpec(node, &standbySpec) {
				neededNodes -= 1
				unboundNodes = append(unboundNodes[:i], unboundNodes[i+1:]...)
				i--
			}
		}
		standbyToNumNeeded[standbySpec] = neededNodes
	}
	// Create nodes to keep the buffered pool up to date
	for spec, numNeeded := range standbyToNumNeeded {
		for i := 0; i < numNeeded; i++ {
			newNode := s.createNodeForStandbySpec(&spec)
			newNodes = append(newNodes, newNode)
		}
	}
	// update bindings for any nodes that we updated here. Note that
	// we might be updating a ton of nodes and if the DB goes away
	// this will take 10s per node to timeout. A context would be
	// better but isn't part of the registry interface....
	deadline := time.Now().Add(time.Second * 15)
	for _, node := range dirtyNodes {
		if time.Now().After(deadline) {
			return nil, nil, nil
		}
		_, err := s.nodeRegistry.UpdateStatus(node)
		if err != nil {
			klog.Errorf("Error updating node %s with pod bindings: %s",
				node.Name, err)
		}
	}

	return newNodes, unboundNodes, podNodeBinding
}
