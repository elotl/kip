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

package healthcheck

import (
	"time"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/nodeclient"
	"github.com/elotl/kip/pkg/server/registry"
	"github.com/elotl/kip/pkg/util/conmap"
	"k8s.io/klog"
)

type statusHealthCheck struct {
	nodeLister        registry.NodeLister
	nodeClientFactory nodeclient.ItzoClientFactoryer
	lastStatusTime    *conmap.StringTimeTime
}

// This is taken care of by the PodController
func (shc *statusHealthCheck) checkPods(lastStatusReply *conmap.StringTimeTime) error {
	return nil
}

func (shc *statusHealthCheck) hasPodFailed(pod *api.Pod) bool {
	node, err := shc.nodeLister.GetNode(pod.Status.BoundNodeName)
	if err != nil {
		klog.Warningf("No node found for pod %s", pod.Name)
		return true
	}
	client := shc.nodeClientFactory.GetClient(node.Status.Addresses)
	_, err = client.GetStatus()
	if err != nil {
		return true
	}

	klog.Warningf("Last chance healthcheck for pod %s saved the pod from failure. Pod status is possibly out of date", pod.Name)
	return false
}

// This tries one last time to check the status of the pod and then
// fails it if that doesn't work.  The one last check (hasPodFailed)
// comes from working with an overloaded system that couldn't process
// all the status replies coming into the system.
func (shc *statusHealthCheck) maybeFailUnresponsivePod(pod *api.Pod, terminateChan chan *api.Pod) {
	go func() {
		if shc.hasPodFailed(pod) {
			// We'll run this syncronously to ensure that, if nothing is
			// processing failed pods, more pods don't get failed. We use
			// a buffered channel to allow up to a full iteration through
			// the list of running pods.
			klog.Errorf("failing pod %s", pod.Name)
			terminateChan <- pod
		} else {
			// We've gotten a good status back from the pod, lets
			// reset the lastStatusTime and come back to this pod if
			// it fails again.
			shc.lastStatusTime.Set(pod.Name, time.Now().UTC())
		}
	}()
}
