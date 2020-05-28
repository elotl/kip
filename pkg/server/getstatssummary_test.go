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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	stats "k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1"
)

func makeUint64Ptr(i int) *uint64 {
	ret := uint64(i)
	return &ret
}

//func float64ToUint64Ptr(f float64, exists bool) *uint64
func TestFloat64ToUint64Ptr(t *testing.T) {
	testCases := []struct {
		f      float64
		exists bool
		result *uint64
	}{
		{
			f:      1.0,
			exists: true,
			result: makeUint64Ptr(1),
		},
		{
			f:      0.0,
			exists: true,
			result: makeUint64Ptr(0),
		},
		{
			f:      0.0,
			exists: false,
			result: nil,
		},
	}
	for _, tc := range testCases {
		result := float64ToUint64Ptr(tc.f, tc.exists)
		assert.Equal(t, tc.result, result)
	}
}

//func addToUint64Ptr(orig, add *uint64) *uint64
func TestAddToUint64Ptr(t *testing.T) {
	testCases := []struct {
		orig   *uint64
		add    *uint64
		result *uint64
	}{
		{
			orig:   makeUint64Ptr(1),
			add:    makeUint64Ptr(2),
			result: makeUint64Ptr(3),
		},
		{
			orig:   nil,
			add:    makeUint64Ptr(2),
			result: makeUint64Ptr(2),
		},
		{
			orig:   makeUint64Ptr(3),
			add:    nil,
			result: makeUint64Ptr(3),
		},
		{
			orig:   nil,
			add:    nil,
			result: nil,
		},
	}
	for _, tc := range testCases {
		result := addToUint64Ptr(tc.orig, tc.add)
		assert.Equal(t, tc.result, result)
	}
}

//func ensurePodCPUStats(ps *stats.PodStats, timestamp metav1.Time)
func TestEnsurePodCPUStats(t *testing.T) {
	ts := metav1.Now()
	ps := stats.PodStats{}
	ensurePodCPUStats(&ps, ts)
	assert.NotNil(t, ps.CPU)
	assert.Equal(t, ts, ps.CPU.Time)
	ensurePodCPUStats(&ps, ts)
	assert.NotNil(t, ps.CPU)
	assert.Equal(t, ts, ps.CPU.Time)
}

//func ensurePodMemoryStats(ps *stats.PodStats, timestamp metav1.Time)
func TestEnsurePodMemoryStats(t *testing.T) {
	ts := metav1.Now()
	ps := stats.PodStats{}
	ensurePodMemoryStats(&ps, ts)
	assert.NotNil(t, ps.Memory)
	assert.Equal(t, ts, ps.Memory.Time)
	ensurePodMemoryStats(&ps, ts)
	assert.NotNil(t, ps.Memory)
	assert.Equal(t, ts, ps.Memory.Time)
}

//func ensurePodNetworkStats(ps stats.PodStats, timestamp metav1.Time)
func TestEnsurePodNetworkStats(t *testing.T) {
	ts := metav1.Now()
	ps := stats.PodStats{}
	ensurePodNetworkStats(&ps, ts)
	assert.NotNil(t, ps.Network)
	assert.Equal(t, ts, ps.Network.Time)
	assert.Len(t, ps.Network.Interfaces, 1)
	ensurePodNetworkStats(&ps, ts)
	assert.NotNil(t, ps.Network)
	assert.Equal(t, ts, ps.Network.Time)
	assert.Len(t, ps.Network.Interfaces, 1)
}

//func ensurePodVolumeStats(ps stats.PodStats, timestamp metav1.Time)
func TestEnsurePodVolumeStats(t *testing.T) {
	ts := metav1.Now()
	ps := stats.PodStats{}
	ensurePodVolumeStats(&ps, ts)
	assert.Len(t, ps.VolumeStats, 1)
	assert.Equal(t, ts, ps.VolumeStats[0].FsStats.Time)
	ensurePodVolumeStats(&ps, ts)
	assert.Len(t, ps.VolumeStats, 1)
	assert.Equal(t, ts, ps.VolumeStats[0].FsStats.Time)
}

//func ensureContainerCPUStats(cstats stats.ContainerStats, timestamp metav1.Time)
func TestEnsureContainerCPUStats(t *testing.T) {
	ts := metav1.Now()
	cs := stats.ContainerStats{}
	ensureContainerCPUStats(&cs, ts)
	assert.NotNil(t, cs.CPU)
	assert.Equal(t, ts, cs.CPU.Time)
	ensureContainerCPUStats(&cs, ts)
	assert.NotNil(t, cs.CPU)
	assert.Equal(t, ts, cs.CPU.Time)
}

//func ensureContainerMemoryStats(cstats stats.ContainerStats, timestamp metav1.Time)
func TestEnsureContainerMemoryStats(t *testing.T) {
	ts := metav1.Now()
	cs := stats.ContainerStats{}
	ensureContainerMemoryStats(&cs, ts)
	assert.NotNil(t, cs.Memory)
	assert.Equal(t, ts, cs.Memory.Time)
	ensureContainerMemoryStats(&cs, ts)
	assert.NotNil(t, cs.Memory)
	assert.Equal(t, ts, cs.Memory.Time)
}

//func updatePodCPUStats(ps stats.PodStats, cstats stats.ContainerStats, timestamp metav1.Time)
func TestUpdatePodCPUStats(t *testing.T) {
	ts := metav1.Now()
	cs := stats.ContainerStats{
		CPU: &stats.CPUStats{
			UsageCoreNanoSeconds: makeUint64Ptr(9999),
			UsageNanoCores:       makeUint64Ptr(1234),
		},
	}
	ps := stats.PodStats{}
	updatePodCPUStats(&ps, &cs, ts)
	assert.NotNil(t, ps.CPU)
	assert.Equal(t, makeUint64Ptr(9999), ps.CPU.UsageCoreNanoSeconds)
	assert.Equal(t, makeUint64Ptr(1234), ps.CPU.UsageNanoCores)
	cs = stats.ContainerStats{
		CPU: &stats.CPUStats{
			UsageCoreNanoSeconds: makeUint64Ptr(1000),
			UsageNanoCores:       makeUint64Ptr(3000),
		},
	}
	updatePodCPUStats(&ps, &cs, ts)
	assert.NotNil(t, ps.CPU)
	assert.Equal(t, makeUint64Ptr(10999), ps.CPU.UsageCoreNanoSeconds)
	assert.Equal(t, makeUint64Ptr(4234), ps.CPU.UsageNanoCores)
}

//func updatePodMemoryStats(ps stats.PodStats, cstats stats.ContainerStats, timestamp metav1.Time)
func TestUpdatePodMemoryStats(t *testing.T) {
	ts := metav1.Now()
	cs := stats.ContainerStats{
		Memory: &stats.MemoryStats{
			AvailableBytes:  makeUint64Ptr(1),
			MajorPageFaults: makeUint64Ptr(2),
			PageFaults:      makeUint64Ptr(3),
			RSSBytes:        makeUint64Ptr(4),
			UsageBytes:      makeUint64Ptr(5),
			WorkingSetBytes: makeUint64Ptr(6),
		},
	}
	ps := stats.PodStats{}
	updatePodMemoryStats(&ps, &cs, ts)
	assert.NotNil(t, ps.Memory)
	assert.Equal(t, makeUint64Ptr(1), ps.Memory.AvailableBytes)
	assert.Equal(t, makeUint64Ptr(2), ps.Memory.MajorPageFaults)
	assert.Equal(t, makeUint64Ptr(3), ps.Memory.PageFaults)
	assert.Equal(t, makeUint64Ptr(4), ps.Memory.RSSBytes)
	assert.Equal(t, makeUint64Ptr(5), ps.Memory.UsageBytes)
	assert.Equal(t, makeUint64Ptr(6), ps.Memory.WorkingSetBytes)
	cs = stats.ContainerStats{
		Memory: &stats.MemoryStats{
			AvailableBytes:  makeUint64Ptr(11),
			MajorPageFaults: makeUint64Ptr(22),
			PageFaults:      makeUint64Ptr(33),
			RSSBytes:        makeUint64Ptr(44),
			UsageBytes:      makeUint64Ptr(55),
			WorkingSetBytes: makeUint64Ptr(66),
		},
	}
	updatePodMemoryStats(&ps, &cs, ts)
	assert.NotNil(t, ps.Memory)
	assert.Equal(t, makeUint64Ptr(12), ps.Memory.AvailableBytes)
	assert.Equal(t, makeUint64Ptr(24), ps.Memory.MajorPageFaults)
	assert.Equal(t, makeUint64Ptr(36), ps.Memory.PageFaults)
	assert.Equal(t, makeUint64Ptr(48), ps.Memory.RSSBytes)
	assert.Equal(t, makeUint64Ptr(60), ps.Memory.UsageBytes)
	assert.Equal(t, makeUint64Ptr(72), ps.Memory.WorkingSetBytes)
}

//func updatePodNetworkStats(ps stats.PodStats, timestamp metav1.Time, k string, v uint64)
func TestUpdatePodNetworkStats(t *testing.T) {
	ts := metav1.Now()
	ps := stats.PodStats{}
	updatePodNetworkStats(&ps, ts, "netRx", 123)
	updatePodNetworkStats(&ps, ts, "netRxErrors", 10)
	updatePodNetworkStats(&ps, ts, "netTx", 456)
	updatePodNetworkStats(&ps, ts, "netTxErrors", 20)
	assert.NotNil(t, ps.Network)
	assert.Equal(t, makeUint64Ptr(123), ps.Network.InterfaceStats.RxBytes)
	assert.Equal(t, makeUint64Ptr(10), ps.Network.InterfaceStats.RxErrors)
	assert.Equal(t, makeUint64Ptr(456), ps.Network.InterfaceStats.TxBytes)
	assert.Equal(t, makeUint64Ptr(20), ps.Network.InterfaceStats.TxErrors)
}

//func updatePodVolumeStats(ps stats.PodStats, timestamp metav1.Time, k string, v uint64)
func TestUpdatePodVolumeStats(t *testing.T) {
	ts := metav1.Now()
	ps := stats.PodStats{}
	updatePodVolumeStats(&ps, ts, "fsAvailable", 11)
	updatePodVolumeStats(&ps, ts, "fsCapacity", 22)
	updatePodVolumeStats(&ps, ts, "fsUsed", 33)
	updatePodVolumeStats(&ps, ts, "fsInodesFree", 44)
	updatePodVolumeStats(&ps, ts, "fsInodes", 55)
	updatePodVolumeStats(&ps, ts, "fsInodesUsed", 66)
	assert.Len(t, ps.VolumeStats, 1)
	assert.Equal(t, makeUint64Ptr(11), ps.VolumeStats[0].AvailableBytes)
	assert.Equal(t, makeUint64Ptr(22), ps.VolumeStats[0].CapacityBytes)
	assert.Equal(t, makeUint64Ptr(33), ps.VolumeStats[0].UsedBytes)
	assert.Equal(t, makeUint64Ptr(44), ps.VolumeStats[0].InodesFree)
	assert.Equal(t, makeUint64Ptr(55), ps.VolumeStats[0].Inodes)
	assert.Equal(t, makeUint64Ptr(66), ps.VolumeStats[0].InodesUsed)
}

//func updateContainerMemoryStats(cstats stats.ContainerStats, timestamp metav1.Time, k string, value uint64)
func TestUpdateContainerMemoryStats(t *testing.T) {
	ts := metav1.Now()
	cs := stats.ContainerStats{}
	updateContainerMemoryStats(&cs, ts, "memoryAvailable", 111)
	updateContainerMemoryStats(&cs, ts, "memoryMajorPageFaults", 222)
	updateContainerMemoryStats(&cs, ts, "memoryPageFaults", 333)
	updateContainerMemoryStats(&cs, ts, "memoryRSS", 444)
	updateContainerMemoryStats(&cs, ts, "memoryUsage", 555)
	updateContainerMemoryStats(&cs, ts, "memoryWorkingSet", 666)
	assert.NotNil(t, cs.Memory)
	assert.Equal(t, makeUint64Ptr(111), cs.Memory.AvailableBytes)
	assert.Equal(t, makeUint64Ptr(222), cs.Memory.MajorPageFaults)
	assert.Equal(t, makeUint64Ptr(333), cs.Memory.PageFaults)
	assert.Equal(t, makeUint64Ptr(444), cs.Memory.RSSBytes)
	assert.Equal(t, makeUint64Ptr(555), cs.Memory.UsageBytes)
	assert.Equal(t, makeUint64Ptr(666), cs.Memory.WorkingSetBytes)
}

//func updateContainerCPUStats(cstats stats.ContainerStats, timestamp metav1.Time, k string, value uint64, prevValue *uint64, nanoseconds int64)
func TestUpdateContainerCPUStats(t *testing.T) {
	testCases := []struct {
		value          uint64
		prevValue      *uint64
		nanoseconds    int64
		usageNanoCores *uint64
	}{
		{
			// Basic case.
			value:          1234,
			prevValue:      makeUint64Ptr(333),
			nanoseconds:    10000000,
			usageNanoCores: makeUint64Ptr(90100),
		},
		{
			// Time difference between samples is negative.
			value:          1234,
			prevValue:      makeUint64Ptr(333),
			nanoseconds:    -1000,
			usageNanoCores: nil,
		},
		{
			// No previous sample.
			value:          1234,
			prevValue:      nil,
			nanoseconds:    1000,
			usageNanoCores: nil,
		},
		{
			// Negative CPU usage diff.
			value:          1234,
			prevValue:      makeUint64Ptr(2000),
			nanoseconds:    1000,
			usageNanoCores: nil,
		},
	}
	for i, tc := range testCases {
		ts := metav1.Now()
		cs := stats.ContainerStats{}
		msg := fmt.Sprintf("test case #%d %+v failed", i+1, tc)
		updateContainerCPUStats(&cs, ts, "cpuUsage", tc.value, tc.prevValue, tc.nanoseconds)
		assert.NotNil(t, cs.CPU)
		assert.Equal(t, makeUint64Ptr(int(tc.value)), cs.CPU.UsageCoreNanoSeconds)
		assert.Equal(t, tc.usageNanoCores, cs.CPU.UsageNanoCores, msg)
	}
}
