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
	"sync"

	"github.com/elotl/kip/pkg/api"
)

// Writing Metrics happens frequently (at most, every time we poll
// pods for their status) but reading metrics can be fairly
// infrequent. We'll aim to make metric storage lightweight at the
// expense of storing non-api objects. We'll convert the storage
// format to our api objects when reading metrics back out of here
//
// If this design doesnt suit us, it's easy to upgrade our
// representation since there's no persistence.

type timestampedMetrics struct {
	Timestamp api.Time
	// We need to store the window or else we can't get the window of
	// the last entry (we've overwritten the datapoint before the last
	// datapoint)
	Window api.Duration
	api.ResourceMetrics
}

type MetricsStore struct {
	sync.RWMutex
	numDatapoints int
	pods          map[string]*PodMetrics
}

func NewMetricsStore(numDatapoints int) *MetricsStore {
	return &MetricsStore{
		numDatapoints: numDatapoints,
		pods:          make(map[string]*PodMetrics),
	}
}

// todo: fix this to take a pointer to a metric and calculate the window
// lightweight: t, ResourceMetrics
func (m *MetricsStore) Insert(podName string, t api.Time, r api.ResourceMetrics) {
	m.Lock()
	defer m.Unlock()
	podMetrics, exists := m.pods[podName]
	if !exists {
		podMetrics = NewPodMetrics(podName, m.numDatapoints)
	}
	podMetrics.insert(t, r)
	m.pods[podName] = podMetrics
}

// get the most recent metrics from all pods
func (m *MetricsStore) GetLatestMetrics() *api.MetricsList {
	m.RLock()
	defer m.RUnlock()
	metricsSlice := make([]*api.Metrics, 0, len(m.pods))
	for _, v := range m.pods {
		m := v.getLatest()
		if m != nil {
			metricsSlice = append(metricsSlice, m)
		}
	}
	metricsList := api.NewMetricsList()
	metricsList.Items = metricsSlice
	return metricsList
}

// Get all the metrics for a pod
// Todo: update this to return not found
func (m *MetricsStore) GetPodMetrics(podName string) *api.MetricsList {
	m.RLock()
	defer m.RUnlock()
	podMetrics, exists := m.pods[podName]
	if !exists {
		return api.NewMetricsList()
	}
	return podMetrics.listAll()
}

func (m *MetricsStore) DeletePods(podNames ...string) {
	m.Lock()
	defer m.Unlock()
	for _, podName := range podNames {
		delete(m.pods, podName)
	}
}

func (m *MetricsStore) ListPods() []string {
	m.RLock()
	defer m.RUnlock()
	podNames := make([]string, 0, len(m.pods))
	for k := range m.pods {
		podNames = append(podNames, k)
	}
	return podNames
}

//////////////////////////////////////////////////////////////////////
type PodMetrics struct {
	name    string
	count   int64
	metrics []timestampedMetrics
}

func NewPodMetrics(name string, numDatpoints int) *PodMetrics {
	return &PodMetrics{
		name:    name,
		metrics: make([]timestampedMetrics, numDatpoints),
	}
}

func (h *PodMetrics) insert(t api.Time, r api.ResourceMetrics) {
	var window api.Duration
	if h.count > 0 {
		prev := (h.count - 1) % int64(len(h.metrics))
		window = api.Duration{t.Sub(h.metrics[prev].Timestamp)}
	}
	idx := h.count % int64(len(h.metrics))
	h.metrics[idx] = timestampedMetrics{
		Timestamp:       t,
		Window:          window,
		ResourceMetrics: r,
	}
	h.count++
}

func toMetrics(name string, tsMetrics timestampedMetrics) *api.Metrics {
	metrics := api.NewMetrics()
	metrics.Name = name
	metrics.Timestamp = tsMetrics.Timestamp
	metrics.Window = tsMetrics.Window
	metrics.ResourceUsage = tsMetrics.ResourceMetrics
	return metrics
}

func (h *PodMetrics) getLatest() *api.Metrics {
	if h.count == 0 {
		return nil
	}
	idx := (h.count - 1) % int64(len(h.metrics))
	tsMetrics := h.metrics[idx]
	return toMetrics(h.name, tsMetrics)
}

func (h *PodMetrics) listAll() *api.MetricsList {
	start := h.count - int64(len(h.metrics))
	if start < 0 {
		start = 0
	}

	metricsSlice := make([]*api.Metrics, 0, len(h.metrics))
	for i := start; i < h.count; i++ {
		idx := i % int64(len(h.metrics))
		tsMetric := h.metrics[idx]
		metricsSlice = append(metricsSlice, toMetrics(h.name, tsMetric))
	}
	metricsList := api.NewMetricsList()
	metricsList.Items = metricsSlice
	return metricsList
}
