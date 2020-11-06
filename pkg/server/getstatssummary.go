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
	"context"
	"strings"
	"time"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/util"
	"github.com/virtual-kubelet/virtual-kubelet/trace"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
	stats "k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1"
)

const (
	nanosecondsPerSecond = float64(time.Second / time.Nanosecond)
)

func (p *InstanceProvider) GetStatsSummary(ctx context.Context) (*stats.Summary, error) {
	var span trace.Span
	ctx, span = trace.StartSpan(ctx, "GetStatsSummary")
	defer span.End()
	klog.V(5).Infof("GetStatsSummary()")
	zero := uint64(0)
	now := metav1.NewTime(time.Now())
	res := &stats.Summary{}
	res.Node = stats.NodeStats{
		NodeName:  p.nodeName,
		StartTime: metav1.NewTime(p.startTime),
		CPU: &stats.CPUStats{
			Time:           now,
			UsageNanoCores: &zero,
		},
		Memory: &stats.MemoryStats{
			Time:            now,
			UsageBytes:      &zero,
			WorkingSetBytes: &zero,
		},
	}
	podRegistry := p.getPodRegistry()
	pods, err := podRegistry.ListPods(func(pod *api.Pod) bool {
		if pod.Status.Phase == api.PodRunning {
			return true
		}
		return false
	})
	if err != nil {
		klog.Errorf("listing pods for stats: %v", err)
		return nil, util.WrapError(err, "listing pods for stats")
	}
	metricsRegistry := p.getMetricsRegistry()
	for _, pod := range pods.Items {
		podMetricsList := metricsRegistry.GetPodMetrics(pod.Name)
		podMetricsItems := podMetricsList.Items
		if len(podMetricsItems) < 2 {
			klog.V(2).Infof("not enough metrics yet for pod %s", pod.Name)
			continue
		}
		// The docs say PodStats.StartTime is:
		//	 The time at which data collection for the pod-scoped (e.g.
		//	 network) stats was (re)started.
		// That is, when the pod or container was started.
		startTime := metav1.NewTime(pod.CreationTimestamp.Time)
		// Last two samples from the pod, with the latest metrics.
		currentSample := podMetricsItems[len(podMetricsItems)-1]
		previousSample := podMetricsItems[len(podMetricsItems)-2]
		// Extract pod and container metrics.
		ps := getStats(
			startTime,
			[]*api.Metrics{
				previousSample,
				currentSample,
			},
			pod.Spec.Units,
		)
		namespace, name := util.SplitNamespaceAndName(pod.Name)
		ps.PodRef = stats.PodReference{
			Name:      name,
			Namespace: namespace,
			UID:       pod.UID,
		}
		res.Pods = append(res.Pods, ps)
	}
	klog.V(5).Infof("GetStatsSummary() %+v", res)
	return res, nil
}

func float64ToUint64Ptr(f float64, exists bool) *uint64 {
	if !exists {
		return nil
	}
	v := uint64(f)
	return &v
}

func getStats(startTime metav1.Time, podMetrics []*api.Metrics, units []api.Unit) stats.PodStats {
	timestamp := metav1.NewTime(podMetrics[1].Timestamp.Time)
	prevTimestamp := metav1.NewTime(podMetrics[0].Timestamp.Time)
	nanoseconds := timestamp.UnixNano() - prevTimestamp.UnixNano()
	ps := stats.PodStats{
		StartTime: startTime,
	}
	statsMap := make(map[string]stats.ContainerStats)
	for k, v := range podMetrics[1].ResourceUsage {
		value := uint64(v)
		// Some metrics are calculated as a diff between the previous and
		// current values.
		prev, prevExists := podMetrics[0].ResourceUsage[k]
		prevValue := float64ToUint64Ptr(prev, prevExists)
		// First update per-pod metrics.
		updatePodNetworkStats(&ps, timestamp, k, value)
		updatePodVolumeStats(&ps, timestamp, k, value)
		// Update per-container metrics.
		parts := strings.Split(k, ".")
		if len(parts) != 2 {
			continue
		}
		unitName := parts[0]
		metric := parts[1]
		cstats, ok := statsMap[unitName]
		if !ok {
			cstats = stats.ContainerStats{
				Name:      unitName,
				StartTime: startTime,
			}
		}
		updateContainerCPUStats(&cstats, timestamp, metric, value, prevValue, nanoseconds)
		updateContainerMemoryStats(&cstats, timestamp, metric, value)
		statsMap[unitName] = cstats
	}
	ps.Containers = make([]stats.ContainerStats, len(units))
	for i, unit := range units {
		cstats, ok := statsMap[unit.Name]
		if !ok {
			klog.Warningf("container %s is missing from stats map", unit.Name)
			continue
		}
		ps.Containers[i] = cstats
		updatePodCPUStats(&ps, &cstats, timestamp)
		updatePodMemoryStats(&ps, &cstats, timestamp)
	}
	return ps
}

func ensurePodCPUStats(ps *stats.PodStats, timestamp metav1.Time) {
	if ps.CPU == nil {
		ps.CPU = &stats.CPUStats{
			Time: timestamp,
		}
	}
}

func ensurePodMemoryStats(ps *stats.PodStats, timestamp metav1.Time) {
	if ps.Memory == nil {
		ps.Memory = &stats.MemoryStats{
			Time: timestamp,
		}
	}
}

func ensurePodNetworkStats(ps *stats.PodStats, timestamp metav1.Time) {
	if ps.Network == nil {
		ps.Network = &stats.NetworkStats{
			Time: timestamp,
			InterfaceStats: stats.InterfaceStats{
				Name: "eth0",
			},
		}
		ps.Network.Interfaces = []stats.InterfaceStats{
			{},
		}
	}
}

func ensurePodVolumeStats(ps *stats.PodStats, timestamp metav1.Time) {
	if len(ps.VolumeStats) == 0 {
		ps.VolumeStats = []stats.VolumeStats{
			{
				Name: "/",
				FsStats: stats.FsStats{
					Time: timestamp,
				},
			},
		}
	}
}

func ensureContainerCPUStats(cstats *stats.ContainerStats, timestamp metav1.Time) {
	if cstats.CPU == nil {
		cstats.CPU = &stats.CPUStats{
			Time: timestamp,
		}
	}
}

func ensureContainerMemoryStats(cstats *stats.ContainerStats, timestamp metav1.Time) {
	if cstats.Memory == nil {
		cstats.Memory = &stats.MemoryStats{
			Time: timestamp,
		}
	}
}

func addToUint64Ptr(orig, add *uint64) *uint64 {
	if add == nil {
		return orig
	}
	v := *add
	if orig == nil {
		return &v
	}
	v += *orig
	return &v
}

func updatePodCPUStats(ps *stats.PodStats, cstats *stats.ContainerStats, timestamp metav1.Time) {
	if cstats.CPU == nil {
		return
	}
	ensurePodCPUStats(ps, timestamp)
	ps.CPU.UsageNanoCores = addToUint64Ptr(ps.CPU.UsageNanoCores, cstats.CPU.UsageNanoCores)
	ps.CPU.UsageCoreNanoSeconds = addToUint64Ptr(ps.CPU.UsageCoreNanoSeconds, cstats.CPU.UsageCoreNanoSeconds)
}

func updatePodMemoryStats(ps *stats.PodStats, cstats *stats.ContainerStats, timestamp metav1.Time) {
	if cstats.Memory == nil {
		return
	}
	ensurePodMemoryStats(ps, timestamp)
	ps.Memory.UsageBytes = addToUint64Ptr(ps.Memory.UsageBytes, cstats.Memory.UsageBytes)
	ps.Memory.WorkingSetBytes = addToUint64Ptr(ps.Memory.WorkingSetBytes, cstats.Memory.WorkingSetBytes)
	ps.Memory.AvailableBytes = addToUint64Ptr(ps.Memory.AvailableBytes, cstats.Memory.AvailableBytes)
	ps.Memory.RSSBytes = addToUint64Ptr(ps.Memory.RSSBytes, cstats.Memory.RSSBytes)
	ps.Memory.PageFaults = addToUint64Ptr(ps.Memory.PageFaults, cstats.Memory.PageFaults)
	ps.Memory.MajorPageFaults = addToUint64Ptr(ps.Memory.MajorPageFaults, cstats.Memory.MajorPageFaults)
}

func updatePodNetworkStats(ps *stats.PodStats, timestamp metav1.Time, k string, v uint64) {
	switch k {
	case "netRx":
		ensurePodNetworkStats(ps, timestamp)
		ps.Network.InterfaceStats.RxBytes = &v
		ps.Network.Interfaces[0] = ps.Network.InterfaceStats
	case "netRxErrors":
		ensurePodNetworkStats(ps, timestamp)
		ps.Network.InterfaceStats.RxErrors = &v
		ps.Network.Interfaces[0] = ps.Network.InterfaceStats
	case "netTx":
		ensurePodNetworkStats(ps, timestamp)
		ps.Network.InterfaceStats.TxBytes = &v
		ps.Network.Interfaces[0] = ps.Network.InterfaceStats
	case "netTxErrors":
		ensurePodNetworkStats(ps, timestamp)
		ps.Network.InterfaceStats.TxErrors = &v
		ps.Network.Interfaces[0] = ps.Network.InterfaceStats
	}
}

func updatePodVolumeStats(ps *stats.PodStats, timestamp metav1.Time, k string, v uint64) {
	switch k {
	case "fsAvailable":
		ensurePodVolumeStats(ps, timestamp)
		ps.VolumeStats[0].AvailableBytes = &v
	case "fsCapacity":
		ensurePodVolumeStats(ps, timestamp)
		ps.VolumeStats[0].CapacityBytes = &v
	case "fsUsed":
		ensurePodVolumeStats(ps, timestamp)
		ps.VolumeStats[0].UsedBytes = &v
	case "fsInodesFree":
		ensurePodVolumeStats(ps, timestamp)
		ps.VolumeStats[0].InodesFree = &v
	case "fsInodes":
		ensurePodVolumeStats(ps, timestamp)
		ps.VolumeStats[0].Inodes = &v
	case "fsInodesUsed":
		ensurePodVolumeStats(ps, timestamp)
		ps.VolumeStats[0].InodesUsed = &v
	}
}

func updateContainerCPUStats(cstats *stats.ContainerStats, timestamp metav1.Time, k string, value uint64, prevValue *uint64, nanoseconds int64) {
	switch k {
	case "cpuUsage":
		ensureContainerCPUStats(cstats, timestamp)
		cstats.CPU.UsageCoreNanoSeconds = &value
		if prevValue != nil && nanoseconds > 0 && value >= *prevValue {
			diff := float64(value - *prevValue)
			cores := uint64(diff / float64(nanoseconds) * nanosecondsPerSecond)
			cstats.CPU.UsageNanoCores = &cores
		}
		klog.V(5).Infof("container %s cpustats %+v", cstats.Name, *cstats.CPU)
	}
}

func updateContainerMemoryStats(cstats *stats.ContainerStats, timestamp metav1.Time, k string, value uint64) {
	switch k {
	case "memoryUsage":
		ensureContainerMemoryStats(cstats, timestamp)
		cstats.Memory.UsageBytes = &value
	case "memoryWorkingSet":
		ensureContainerMemoryStats(cstats, timestamp)
		cstats.Memory.WorkingSetBytes = &value
	case "memoryAvailable":
		ensureContainerMemoryStats(cstats, timestamp)
		cstats.Memory.AvailableBytes = &value
	case "memoryRSS":
		ensureContainerMemoryStats(cstats, timestamp)
		cstats.Memory.RSSBytes = &value
	case "memoryPageFaults":
		ensureContainerMemoryStats(cstats, timestamp)
		cstats.Memory.PageFaults = &value
	case "memoryMajorPageFaults":
		ensureContainerMemoryStats(cstats, timestamp)
		cstats.Memory.MajorPageFaults = &value
	}
}
