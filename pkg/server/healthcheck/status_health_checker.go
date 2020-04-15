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
	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/nodeclient"
	"github.com/elotl/kip/pkg/server/registry"
	"github.com/elotl/kip/pkg/util/conmap"
	"k8s.io/klog"
)

type statusHealthCheck struct {
	nodeLister        registry.NodeLister
	nodeClientFactory nodeclient.ItzoClientFactoryer
}

// This is taken care of by the PodController
func (shc *statusHealthCheck) checkPods(lastStatusReply *conmap.StringTimeTime) error {
	return nil
}

func (shc *statusHealthCheck) podHasFailed(pod *api.Pod) bool {
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
