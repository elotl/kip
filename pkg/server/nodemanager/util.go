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

import "github.com/elotl/kip/pkg/api"

func isSpotPod(pod *api.Pod) bool {
	return pod.Spec.Spot.Policy == api.SpotAlways

}

func podNeedsNode(pod *api.Pod) bool {
	return pod.Status.Phase == api.PodWaiting

	// I commented this out since we have a race condition where a pod
	// fails, the pool algorithm runs and we think that the Failed pod
	// needs a node. However, if the pod either has restartPolicy ==
	// Never or the pod has failed too many times in a row that pod
	// doesn't need a node.  So we booted a node that we'll end up
	// deleting.  Lets delay things a bit and not boot nodes for
	// failed pods until they're put back into waiting.

	// || (pod.Status.Phase == api.PodFailed && pod.Spec.Phase == api.PodRunning)
}

func availableOrBaking(node *api.Node) bool {
	return (node.Status.Phase == api.NodeCreating ||
		node.Status.Phase == api.NodeCreated ||
		node.Status.Phase == api.NodeAvailable)
}
