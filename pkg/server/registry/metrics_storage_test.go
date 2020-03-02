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

package registry

import (
	"testing"
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestPodMetrics(t *testing.T) {
	numDatapoints := 100
	pm := NewPodMetrics("p1", numDatapoints)
	start := api.Now()
	n := 1000
	factor := float64(0.1)
	for i := 0; i < n; i++ {
		amount := float64(i+1) * factor
		u := api.ResourceMetrics{"cpu": amount}
		pm.insert(start.Add(time.Duration(i)*time.Second), u)
	}
	assert.Equal(t, int64(n), pm.count)
	latest := pm.getLatest()
	assert.Equal(t, float64(100.0), latest.ResourceUsage["cpu"])
	endTime := start.Add(time.Duration(time.Duration(n-1) * time.Second))
	assert.Equal(t, endTime, latest.Timestamp)

	allMetrics := pm.listAll()
	assert.Len(t, allMetrics.Items, numDatapoints)
	startMetric := allMetrics.Items[0]
	for i, m := range allMetrics.Items {
		assert.Equal(t, api.Duration{1 * time.Second}, m.Window)
		assert.InEpsilon(t, startMetric.ResourceUsage["cpu"]+factor*float64(i), m.ResourceUsage["cpu"], 0.000001)
	}
}

func TestMetricsStore(t *testing.T) {
	ms := NewMetricsStore(100)
	t1 := api.Now()
	u1 := api.ResourceMetrics{"cpu": 0.0}
	u2 := api.ResourceMetrics{"cpu": 10.0}
	u3 := api.ResourceMetrics{"cpu": 20.0}
	u4 := api.ResourceMetrics{"cpu": 30.0}
	ms.Insert("p1", t1, u1)
	ms.Insert("p2", t1, u2)
	t2 := t1.Add(1 * time.Second)
	ms.Insert("p2", t2, u3)
	ms.Insert("p3", t2, u4)
	podNames := ms.ListPods()
	assert.Len(t, podNames, 3)
	assert.ElementsMatch(t, podNames, []string{"p1", "p2", "p3"})
	assert.Len(t, ms.GetLatestMetrics().Items, 3)
	podMetrics := ms.GetPodMetrics("p2")
	assert.Len(t, podMetrics.Items, 2)
	assert.Equal(t, u2, podMetrics.Items[0].ResourceUsage)
	assert.Equal(t, u3, podMetrics.Items[1].ResourceUsage)
	ms.DeletePods("p2")
	assert.Len(t, ms.GetLatestMetrics().Items, 2)
	podNames = ms.ListPods()
	assert.ElementsMatch(t, podNames, []string{"p1", "p3"})
}
