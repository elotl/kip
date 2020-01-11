package server

import (
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/stretchr/testify/assert"
)

func TestMetricsControllerClean(t *testing.T) {
	podReg, closer := registry.SetupTestPodRegistry()
	defer closer()
	metricsReg := registry.NewMetricsRegistry(100)
	ctl := MetricsController{
		metricsRegistry: metricsReg,
		podLister:       podReg,
	}
	p1 := api.GetFakePod()
	_, err := podReg.CreatePod(p1)
	assert.NoError(t, err)
	p2 := api.GetFakePod()
	p2, err = podReg.CreatePod(p2)
	assert.NoError(t, err)
	now := api.Now()
	m := api.ResourceMetrics{"cpu": 55.0}
	metricsReg.Insert(p2.Name, now, m)
	metricsReg.Insert("oldPodName", now, m)
	metrics := metricsReg.GetLatestMetrics()
	assert.Len(t, metrics.Items, 2)
	err = ctl.cleanOldMetrics()
	assert.NoError(t, err)
	metrics = metricsReg.GetLatestMetrics()
	assert.Len(t, metrics.Items, 1)
	assert.Equal(t, p2.Name, metrics.Items[0].Name)
}
