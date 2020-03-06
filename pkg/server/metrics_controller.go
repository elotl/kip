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

package server

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog"
)

const (
	cleanMetricsInterval = 30 * time.Second
)

type MetricsController struct {
	metricsRegistry *registry.MetricsRegistry
	podLister       registry.PodLister
}

func (c *MetricsController) Dump() []byte {
	dumpStruct := struct {
		NumPods int
	}{
		NumPods: len(c.metricsRegistry.ListPods()),
	}
	b, err := json.MarshalIndent(dumpStruct, "", "    ")
	if err != nil {
		klog.Errorln("Error dumping data from metrics controller", err)
		return nil
	}
	return b
}

func (c *MetricsController) Start(quit <-chan struct{}, wg *sync.WaitGroup) {
	go c.runSyncLoop(quit, wg)
}

func (c *MetricsController) runSyncLoop(quit <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	cleanTicker := time.NewTicker(cleanMetricsInterval)
	for {
		select {
		case <-cleanTicker.C:
			err := c.cleanOldMetrics()
			if err != nil {
				klog.Errorf("Error cleaning old metrics: %s", err)
			}
		case <-quit:
			cleanTicker.Stop()
			klog.V(2).Info("Exiting MetricsController Sync Loop")
			return
		}
	}
}

func (c *MetricsController) cleanOldMetrics() error {
	pods, err := c.podLister.ListPods(registry.MatchAllLivePods)
	if err != nil {
		return err
	}
	podNameSet := sets.NewString()
	for i := range pods.Items {
		podNameSet.Insert(pods.Items[i].Name)
	}
	metricPodNames := c.metricsRegistry.ListPods()
	metricNameSet := sets.NewString(metricPodNames...)
	oldMetrics := metricNameSet.Difference(podNameSet)
	c.metricsRegistry.DeletePods(oldMetrics.List()...)
	return nil
}
