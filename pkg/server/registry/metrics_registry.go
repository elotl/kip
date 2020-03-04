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

import "github.com/elotl/cloud-instance-provider/pkg/api"

type MetricsRegistry struct {
	MetricsStore
}

func NewMetricsRegistry(numDatapoints int) *MetricsRegistry {
	return &MetricsRegistry{
		MetricsStore: MetricsStore{
			numDatapoints: numDatapoints,
			pods:          make(map[string]*PodMetrics),
		},
	}
}

func (reg *MetricsRegistry) New() api.MilpaObject {
	return api.NewMetrics()
}

func (reg *MetricsRegistry) Create(obj api.MilpaObject) (api.MilpaObject, error) {
	metrics := obj.(*api.Metrics)
	reg.insertMetric(metrics)
	return metrics, nil
}

func (reg *MetricsRegistry) Update(obj api.MilpaObject) (api.MilpaObject, error) {
	return reg.Create(obj)
}

func (reg *MetricsRegistry) Get(name string) (api.MilpaObject, error) {
	m := reg.GetPodMetrics(name)
	return m, nil
}

func (reg *MetricsRegistry) List() (api.MilpaObject, error) {
	m := reg.GetLatestMetrics()
	return m, nil
}

func (reg *MetricsRegistry) Delete(name string) (api.MilpaObject, error) {
	reg.DeletePods(name)
	return api.MetricsList{}, nil
}

func (reg *MetricsRegistry) insertMetric(metric *api.Metrics) {
	reg.Insert(metric.Name, metric.Timestamp, metric.ResourceUsage)
}
