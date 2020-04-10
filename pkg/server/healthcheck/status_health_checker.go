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

// type statusHealthCheck struct {
// 	nodeLister        registry.NodeLister
// 	nodeClientFactory nodeclient.ItzoClientFactoryer
// }

// func NewStatusHealthCheck(
// 	podLister registry.PodLister,
// 	nodeLister registry.NodeLister,
// 	lastStatusReply *conmap.StringTimeTime,
// 	nodeClientFactory nodeclient.ItzoClientFactoryer,
// ) *StatusHealthCheck {
// 	return &StatusHealthCheck{
// 		podLister:          podLister,
// 		nodeLister:         nodeLister,
// 		lastStatusReply:    lastStatusReply,
// 		nodeClientFactory:  nodeClientFactory,
// 		statusReplyTimeout: 90 * time.Second,
// 		terminateChan:      make(chan *api.Pod, terminateChanSize),
// 	}
// }

// func (shc *StatusHealthCheck) pruneLastStatusReplies() {
// 	runningPods := make(map[string]bool)
// 	_, err := shc.podLister.ListPods(func(p *api.Pod) bool {
// 		if p.Status.Phase == api.PodRunning {
// 			runningPods[p.Name] = true
// 		}
// 		return false
// 	})
// 	if err != nil {
// 		klog.Errorf("Error getting list of pods from registry")
// 		return
// 	}
// 	for _, replyItem := range shc.lastStatusReply.Items() {
// 		replyPodName := replyItem.Key
// 		_, exists := runningPods[replyPodName]
// 		if !exists {
// 			shc.lastStatusReply.Delete(replyPodName)
// 		}
// 	}
// }

// func (shc *StatusHealthCheck) handleReplyTimeouts() {
// 	podList, err := shc.podLister.ListPods(func(p *api.Pod) bool {
// 		return p.Status.Phase == api.PodRunning
// 	})
// 	if err != nil {
// 		klog.Errorf("Error getting list of pods from registry")
// 		return
// 	}
// 	now := time.Now().UTC()
// 	for _, pod := range podList.Items {
// 		last, exists := shc.lastStatusReply.GetOK(pod.Name)
// 		if !exists {
// 			shc.lastStatusReply.Set(pod.Name, now)
// 			continue
// 		}
// 		if now.Sub(last) < shc.statusReplyTimeout {
// 			continue
// 		}
// 		go shc.maybeFailUnresponsivePod(pod)
// 	}
// }

// func (shc *StatusHealthCheck) maybeFailUnresponsivePod(pod *api.Pod) {
// 	node, err := shc.nodeLister.GetNode(pod.Status.BoundNodeName)
// 	if err != nil {
// 		klog.Warningf("No node found for pod %s", pod.Name)
// 		shc.terminateChan <- pod
// 		return
// 	}
// 	client := shc.nodeClientFactory.GetClient(node.Status.Addresses)
// 	_, err = client.GetStatus()
// 	if err != nil {
// 		klog.Warningf("No status reply from pod %s/%s in %ds failing pod",
// 			pod.Namespace, pod.Name, int(shc.statusReplyTimeout.Seconds()))
// 		shc.terminateChan <- pod
// 	} else {
// 		klog.Warningf("Last chance healthcheck for pod %s saved the pod from failure. Pod status is possibly out of date", pod.Name)
// 		shc.lastStatusReply.Set(pod.Name, time.Now().UTC())
// 	}
// }

// func (shc *StatusHealthCheck) TerminatePodsChan() <-chan *api.Pod {
// 	return shc.terminateChan
// }
