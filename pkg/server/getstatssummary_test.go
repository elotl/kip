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
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	stats "k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1"
)

func makeUint64Ptr(i int) *uint64 {
	ret := uint64(i)
	return &ret
}

//func getContainerStats(startTime metav1.Time, podMetrics *api.Metrics, units []api.Unit) (usageMetrics, []stats.ContainerStats)
func TestGetContainerStats(t *testing.T) {
	testCases := []struct {
		start          metav1.Time
		timestamp      metav1.Time
		podMetrics     *api.Metrics
		units          []api.Unit
		podUsage       usageMetrics
		containerStats []stats.ContainerStats
	}{
		{
			// Two units.
			start:     metav1.NewTime(time.Now().Add(-5 * time.Minute)),
			timestamp: metav1.NewTime(time.Now()),
			podMetrics: &api.Metrics{
				ResourceUsage: api.ResourceMetrics{
					"unit1.cpuUsage":         123,
					"unit1.memoryUsage":      789,
					"unit1.memoryWorkingSet": 456,
					"unit2.cpuUsage":         111,
					"unit2.memoryUsage":      222,
					"unit2.memoryWorkingSet": 333,
				},
			},
			units: []api.Unit{
				{
					Name: "unit1",
				},
				{
					Name: "unit2",
				},
			},
			podUsage: usageMetrics{
				UsageNanoCores:  123 + 111,
				UsageBytes:      789 + 222,
				WorkingSetBytes: 456 + 333,
			},
			containerStats: []stats.ContainerStats{
				{
					Name: "unit1",
					CPU: &stats.CPUStats{
						UsageNanoCores: makeUint64Ptr(123),
					},
					Memory: &stats.MemoryStats{
						UsageBytes:      makeUint64Ptr(789),
						WorkingSetBytes: makeUint64Ptr(456),
					},
				},
				{
					Name: "unit2",
					CPU: &stats.CPUStats{
						UsageNanoCores: makeUint64Ptr(111),
					},
					Memory: &stats.MemoryStats{
						UsageBytes:      makeUint64Ptr(222),
						WorkingSetBytes: makeUint64Ptr(333),
					},
				},
			},
		},
		{
			// One unit, not WorkingSetBytes metric.
			start:     metav1.NewTime(time.Now().Add(-1 * time.Minute)),
			timestamp: metav1.NewTime(time.Now()),
			podMetrics: &api.Metrics{
				ResourceUsage: api.ResourceMetrics{
					"unit1.cpuUsage":    12345,
					"unit1.memoryUsage": 67890,
				},
			},
			units: []api.Unit{
				{
					Name: "unit1",
				},
			},
			podUsage: usageMetrics{
				UsageNanoCores:  12345,
				UsageBytes:      67890,
				WorkingSetBytes: 67890,
			},
			containerStats: []stats.ContainerStats{
				{
					Name: "unit1",
					CPU: &stats.CPUStats{
						UsageNanoCores: makeUint64Ptr(12345),
					},
					Memory: &stats.MemoryStats{
						UsageBytes:      makeUint64Ptr(67890),
						WorkingSetBytes: makeUint64Ptr(67890),
					},
				},
			},
		},
		{
			// Two units, only one has reported metrics
			start:     metav1.NewTime(time.Now().Add(-2 * time.Minute)),
			timestamp: metav1.NewTime(time.Now()),
			podMetrics: &api.Metrics{
				ResourceUsage: api.ResourceMetrics{
					"unit1.cpuUsage":         12345,
					"unit1.memoryUsage":      67890,
					"unit1.memoryWorkingSet": 66666,
				},
			},
			units: []api.Unit{
				{
					Name: "unit1",
				},
				{
					Name: "unit2",
				},
			},
			podUsage: usageMetrics{
				UsageNanoCores:  12345,
				UsageBytes:      67890,
				WorkingSetBytes: 66666,
			},
			containerStats: []stats.ContainerStats{
				{
					Name: "unit1",
					CPU: &stats.CPUStats{
						UsageNanoCores: makeUint64Ptr(12345),
					},
					Memory: &stats.MemoryStats{
						UsageBytes:      makeUint64Ptr(67890),
						WorkingSetBytes: makeUint64Ptr(66666),
					},
				},
				{
					Name: "unit2",
					CPU: &stats.CPUStats{
						UsageNanoCores: makeUint64Ptr(0),
					},
					Memory: &stats.MemoryStats{
						UsageBytes:      makeUint64Ptr(0),
						WorkingSetBytes: makeUint64Ptr(0),
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		tc.podMetrics.Timestamp = api.Time{Time: tc.timestamp.Time}
		for i, _ := range tc.containerStats {
			tc.containerStats[i].StartTime = tc.start
			tc.containerStats[i].CPU.Time = tc.timestamp
			tc.containerStats[i].Memory.Time = tc.timestamp
		}
		podUsage, containerStats := getContainerStats(
			tc.start, tc.podMetrics, tc.units)
		assert.Equal(t, tc.podUsage, podUsage)
		assert.Equal(t, tc.containerStats, containerStats)
		for _, cstat := range tc.containerStats {
			assert.Equal(t, tc.start, cstat.StartTime)
			if cstat.CPU != nil {
				assert.Equal(t, tc.timestamp, cstat.CPU.Time)
			}
			if cstat.Memory != nil {
				assert.Equal(t, tc.timestamp, cstat.Memory.Time)
			}
		}
	}
}
