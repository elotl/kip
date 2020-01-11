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
