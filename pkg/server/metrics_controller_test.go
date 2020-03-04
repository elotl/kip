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
