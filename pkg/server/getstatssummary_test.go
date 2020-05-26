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

	"github.com/elotl/kip/pkg/api"
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
		sampleTimeDiff time.Duration
		podMetrics     []*api.Metrics
		units          []api.Unit
		podUsage       usageMetrics
		containerStats []stats.ContainerStats
	}{
		{
			// Two units.
			start:          metav1.NewTime(time.Now().Add(-5 * time.Minute)),
			timestamp:      metav1.NewTime(time.Now()),
			sampleTimeDiff: 10000 * time.Nanosecond,
			podMetrics: []*api.Metrics{
				{
					ResourceUsage: api.ResourceMetrics{
						"unit1.cpuUsage":         120,
						"unit1.memoryUsage":      789,
						"unit1.memoryWorkingSet": 456,
						"unit1.memoryAvailable":  1<<62 + 1,
						"unit1.memoryRSS":        88,
						"unit2.cpuUsage":         100,
						"unit2.memoryUsage":      222,
						"unit2.memoryWorkingSet": 333,
						"unit2.memoryAvailable":  999,
						"unit2.memoryRSS":        88,
					},
				},
				{
					ResourceUsage: api.ResourceMetrics{
						"unit1.cpuUsage":         123,
						"unit1.memoryUsage":      789,
						"unit1.memoryWorkingSet": 456,
						"unit1.memoryAvailable":  888,
						"unit1.memoryRSS":        88,
						"unit2.cpuUsage":         111,
						"unit2.memoryUsage":      222,
						"unit2.memoryWorkingSet": 333,
						"unit2.memoryAvailable":  999,
						"unit2.memoryRSS":        88,
					},
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
				UsageCoreNanoSeconds: 123 + 111,
				UsageNanoCores:       1400000,
				AvailableBytes:       888 + 999,
				UsageBytes:           789 + 222,
				WorkingSetBytes:      456 + 333,
				RSSBytes:             88 + 88,
			},
			containerStats: []stats.ContainerStats{
				{
					Name: "unit1",
					CPU: &stats.CPUStats{
						UsageCoreNanoSeconds: makeUint64Ptr(123),
						UsageNanoCores:       makeUint64Ptr(300000),
					},
					Memory: &stats.MemoryStats{
						UsageBytes:      makeUint64Ptr(789),
						WorkingSetBytes: makeUint64Ptr(456),
						RSSBytes:        makeUint64Ptr(88),
						AvailableBytes:  makeUint64Ptr(888),
					},
				},
				{
					Name: "unit2",
					CPU: &stats.CPUStats{
						UsageCoreNanoSeconds: makeUint64Ptr(111),
						UsageNanoCores:       makeUint64Ptr(1100000),
					},
					Memory: &stats.MemoryStats{
						UsageBytes:      makeUint64Ptr(222),
						WorkingSetBytes: makeUint64Ptr(333),
						RSSBytes:        makeUint64Ptr(88),
						AvailableBytes:  makeUint64Ptr(999),
					},
				},
			},
		},
		{
			// Two units, only one has reported metrics.
			start:          metav1.NewTime(time.Now().Add(-2 * time.Minute)),
			timestamp:      metav1.NewTime(time.Now()),
			sampleTimeDiff: 5000 * time.Nanosecond,
			podMetrics: []*api.Metrics{
				{
					ResourceUsage: api.ResourceMetrics{
						"unit1.cpuUsage":         12345,
						"unit1.memoryUsage":      67890,
						"unit1.memoryWorkingSet": 66666,
					},
				},
				{
					ResourceUsage: api.ResourceMetrics{
						"unit1.cpuUsage":         12345,
						"unit1.memoryUsage":      67890,
						"unit1.memoryWorkingSet": 66666,
					},
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
				UsageCoreNanoSeconds: 12345,
				UsageBytes:           67890,
				WorkingSetBytes:      66666,
			},
			containerStats: []stats.ContainerStats{
				{
					Name: "unit1",
					CPU: &stats.CPUStats{
						UsageCoreNanoSeconds: makeUint64Ptr(12345),
					},
					Memory: &stats.MemoryStats{
						UsageBytes:      makeUint64Ptr(67890),
						WorkingSetBytes: makeUint64Ptr(66666),
					},
				},
				{
					Name:   "unit2",
					CPU:    &stats.CPUStats{},
					Memory: &stats.MemoryStats{},
				},
			},
		},
		{
			// One unit, spurious CPU sample (timestamp decreased).
			start:          metav1.NewTime(time.Now().Add(-2 * time.Minute)),
			timestamp:      metav1.NewTime(time.Now()),
			sampleTimeDiff: -1 * time.Nanosecond,
			podMetrics: []*api.Metrics{
				{
					ResourceUsage: api.ResourceMetrics{
						"unit1.cpuUsage":         12345,
						"unit1.memoryUsage":      22222,
						"unit1.memoryWorkingSet": 11111,
					},
				},
				{
					ResourceUsage: api.ResourceMetrics{
						"unit1.cpuUsage":         12346,
						"unit1.memoryUsage":      22222,
						"unit1.memoryWorkingSet": 11111,
					},
				},
			},
			units: []api.Unit{
				{
					Name: "unit1",
				},
			},
			podUsage: usageMetrics{
				UsageCoreNanoSeconds: 12346,
				UsageBytes:           22222,
				WorkingSetBytes:      11111,
			},
			containerStats: []stats.ContainerStats{
				{
					Name: "unit1",
					CPU: &stats.CPUStats{
						UsageCoreNanoSeconds: makeUint64Ptr(12346),
					},
					Memory: &stats.MemoryStats{
						UsageBytes:      makeUint64Ptr(22222),
						WorkingSetBytes: makeUint64Ptr(11111),
					},
				},
			},
		},
		{
			// Pod disk metrics.
			start:          metav1.NewTime(time.Now().Add(-2 * time.Minute)),
			timestamp:      metav1.NewTime(time.Now()),
			sampleTimeDiff: -1 * time.Nanosecond,
			podMetrics: []*api.Metrics{
				{
					ResourceUsage: api.ResourceMetrics{
						"fsAvailable":  11111,
						"fsUsed":       22222,
						"fsCapacity":   33333,
						"fsInodes":     99999,
						"fsInodesFree": 44444,
						"fsInodesUsed": 55555,
					},
				},
				{
					ResourceUsage: api.ResourceMetrics{
						"fsAvailable":  11111,
						"fsUsed":       22222,
						"fsCapacity":   33333,
						"fsInodes":     99999,
						"fsInodesFree": 44444,
						"fsInodesUsed": 55555,
					},
				},
			},
			units: []api.Unit{
				{
					Name: "unit1",
				},
			},
			podUsage: usageMetrics{
				FSAvailableBytes: 11111,
				FSUsedBytes:      22222,
				FSCapacityBytes:  33333,
				FSInodes:         99999,
				FSInodesFree:     44444,
				FSInodesUsed:     55555,
			},
			containerStats: []stats.ContainerStats{
				{
					Name:   "unit1",
					CPU:    &stats.CPUStats{},
					Memory: &stats.MemoryStats{},
				},
			},
		},
		{
			// Pod network metrics.
			start:          metav1.NewTime(time.Now().Add(-2 * time.Minute)),
			timestamp:      metav1.NewTime(time.Now()),
			sampleTimeDiff: -1 * time.Nanosecond,
			podMetrics: []*api.Metrics{
				{
					ResourceUsage: api.ResourceMetrics{
						"netRx":       11111,
						"netRxErrors": 1,
						"netTx":       33333,
						"netTxErrors": 3,
					},
				},
				{
					ResourceUsage: api.ResourceMetrics{
						"netRx":       11111,
						"netRxErrors": 1,
						"netTx":       33333,
						"netTxErrors": 3,
					},
				},
			},
			units: []api.Unit{
				{
					Name: "unit1",
				},
			},
			podUsage: usageMetrics{
				NetRxBytes:  11111,
				NetRxErrors: 1,
				NetTxBytes:  33333,
				NetTxErrors: 3,
			},
			containerStats: []stats.ContainerStats{
				{
					Name:   "unit1",
					CPU:    &stats.CPUStats{},
					Memory: &stats.MemoryStats{},
				},
			},
		},
	}
	for _, tc := range testCases {
		tc.podMetrics[1].Timestamp = api.Time{
			Time: tc.timestamp.Time,
		}
		tc.podMetrics[0].Timestamp = api.Time{
			Time: tc.timestamp.Add(-tc.sampleTimeDiff),
		}
		for i, _ := range tc.containerStats {
			tc.containerStats[i].StartTime = tc.start
			tc.containerStats[i].CPU.Time = tc.timestamp
			tc.containerStats[i].Memory.Time = tc.timestamp
		}
		podUsage, containerStats := getStats(tc.start, tc.podMetrics, tc.units)
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
