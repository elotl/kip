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
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/server/registry"
	"github.com/elotl/kip/pkg/util"
	"github.com/elotl/kip/pkg/util/conmap"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog"
)

type cloudAPIHealthCheck struct {
	podLister   registry.PodLister
	cloudClient cloud.CloudClient
}

func (chc *cloudAPIHealthCheck) checkPods(lastStatusTime *conmap.StringTimeTime) error {
	instances, err := chc.cloudClient.ListInstances()
	if err != nil {
		return util.WrapError(err, "error listing cloud instances for cloud API health check")
	}
	instIDs := sets.NewString()
	for i := range instances {
		instIDs.Insert(instances[i].ID)
	}
	podList, err := chc.podLister.ListPods(func(p *api.Pod) bool {
		return p.Status.Phase == api.PodRunning
	})
	if err != nil {
		return util.WrapError(err, "error listing pods for cloud API health check")
	}
	now := time.Now().UTC()
	for _, pod := range podList.Items {
		podInstID := pod.Status.BoundInstanceID
		if podInstID == "" {
			klog.Warningf("cloud instance health check found running pod with empty BoundInstanceID: %s", pod.Name)
			continue
		}
		if instIDs.Has(podInstID) {
			lastStatusTime.Set(pod.Name, now)
		}
	}
	return nil
}

// The cloudAPI healthchecker just fails pods without another check
func (chc *cloudAPIHealthCheck) maybeFailUnresponsivePod(pod *api.Pod, terminateChan chan *api.Pod) {
	terminateChan <- pod
}
