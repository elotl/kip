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

// This is taken care of by the PodController
func (chc *cloudAPIHealthCheck) checkPods(lastStatusTime *conmap.StringTimeTime) error {
	// list instances, put in set
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
			klog.Warningf("cloud instance health check found pod with empty BoundInstanceID: %s/%s", pod.Namespace, pod.Name)
			continue
		}
		if instIDs.Has(podInstID) {
			lastStatusTime.Set(pod.Name, now)
		}
	}
	return nil
}

// We could send one last query to the cloud to see if the instance
// is alive but, that might risk flooding the cloud with lots of
// individual queries for instances
func (chc *cloudAPIHealthCheck) podHasFailed(pod *api.Pod) bool {
	return true
}
