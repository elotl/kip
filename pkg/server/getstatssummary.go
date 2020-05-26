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

type usageMetrics struct {
	UsageNanoCores       uint64
	UsageCoreNanoSeconds uint64
	AvailableBytes       uint64
	UsageBytes           uint64
	WorkingSetBytes      uint64
	RSSBytes             uint64
	PageFaults           uint64
	MajorPageFaults      uint64
	NetRxBytes           uint64
	NetRxErrors          uint64
	NetTxBytes           uint64
	NetTxErrors          uint64
	FSAvailableBytes     uint64
	FSCapacityBytes      uint64
	FSUsedBytes          uint64
	FSInodesFree         uint64
	FSInodes             uint64
	FSInodesUsed         uint64
}

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
		// First metrics sample from the pod.
		firstSample := podMetricsItems[0]
		startTime := metav1.NewTime(firstSample.Timestamp.Time)
		// Last two samples from the pod, with the latest metrics.
		currentSample := podMetricsItems[len(podMetricsItems)-1]
		timestamp := metav1.NewTime(currentSample.Timestamp.Time)
		previousSample := podMetricsItems[len(podMetricsItems)-2]
		namespace, name := util.SplitNamespaceAndName(pod.Name)
		pss := stats.PodStats{
			PodRef: stats.PodReference{
				Name:      name,
				Namespace: namespace,
				UID:       pod.UID,
			},
			StartTime: startTime,
			CPU: &stats.CPUStats{
				Time: timestamp,
			},
			Memory: &stats.MemoryStats{
				Time: timestamp,
			},
			Network: &stats.NetworkStats{
				Time: timestamp,
			},
			VolumeStats: []stats.VolumeStats{
				{
					FsStats: stats.FsStats{
						Time: timestamp,
					},
				},
			},
		}
		var podUsage usageMetrics
		podUsage, pss.Containers = getStats(
			startTime,
			[]*api.Metrics{
				previousSample,
				currentSample,
			},
			pod.Spec.Units,
		)
		pss.CPU.UsageNanoCores = valueOrNil(podUsage.UsageNanoCores)
		pss.CPU.UsageCoreNanoSeconds = valueOrNil(podUsage.UsageCoreNanoSeconds)
		pss.Memory.UsageBytes = valueOrNil(podUsage.UsageBytes)
		pss.Memory.WorkingSetBytes = valueOrNil(podUsage.WorkingSetBytes)
		pss.Memory.AvailableBytes = valueOrNil(podUsage.AvailableBytes)
		pss.Memory.RSSBytes = valueOrNil(podUsage.RSSBytes)
		pss.Memory.PageFaults = valueOrNil(podUsage.PageFaults)
		pss.Memory.MajorPageFaults = valueOrNil(podUsage.MajorPageFaults)
		pss.VolumeStats[0].Name = "/"
		pss.VolumeStats[0].AvailableBytes = valueOrNil(podUsage.FSAvailableBytes)
		pss.VolumeStats[0].CapacityBytes = valueOrNil(podUsage.FSCapacityBytes)
		pss.VolumeStats[0].UsedBytes = valueOrNil(podUsage.FSUsedBytes)
		pss.VolumeStats[0].InodesFree = valueOrNil(podUsage.FSInodesFree)
		pss.VolumeStats[0].Inodes = valueOrNil(podUsage.FSInodes)
		pss.VolumeStats[0].InodesUsed = valueOrNil(podUsage.FSInodesUsed)
		pss.Network.InterfaceStats.Name = "eth0"
		pss.Network.InterfaceStats.RxBytes = valueOrNil(podUsage.NetRxBytes)
		pss.Network.InterfaceStats.RxErrors = valueOrNil(podUsage.NetRxErrors)
		pss.Network.InterfaceStats.TxBytes = valueOrNil(podUsage.NetTxBytes)
		pss.Network.InterfaceStats.TxErrors = valueOrNil(podUsage.NetTxErrors)
		pss.Network.Interfaces = []stats.InterfaceStats{
			pss.Network.InterfaceStats,
		}
		res.Pods = append(res.Pods, pss)
	}
	klog.V(5).Infof("GetStatsSummary() %+v", res)
	return res, nil
}

func getStats(startTime metav1.Time, podMetrics []*api.Metrics, units []api.Unit) (usageMetrics, []stats.ContainerStats) {
	timestamp := metav1.NewTime(podMetrics[1].Timestamp.Time)
	prevTimestamp := metav1.NewTime(podMetrics[0].Timestamp.Time)
	nanoseconds := timestamp.UnixNano() - prevTimestamp.UnixNano()
	unitUsageMap := make(map[string]*usageMetrics)
	podUsage := usageMetrics{}
	for k, v := range podMetrics[1].ResourceUsage {
		if k == "netRx" {
			podUsage.NetRxBytes = uint64(v)
		}
		if k == "netRxErrors" {
			podUsage.NetRxErrors = uint64(v)
		}
		if k == "netTx" {
			podUsage.NetTxBytes = uint64(v)
		}
		if k == "netTxErrors" {
			podUsage.NetTxErrors = uint64(v)
		}
		if k == "fsAvailable" {
			podUsage.FSAvailableBytes = uint64(v)
		}
		if k == "fsCapacity" {
			podUsage.FSCapacityBytes = uint64(v)
		}
		if k == "fsUsed" {
			podUsage.FSUsedBytes = uint64(v)
		}
		if k == "fsInodesFree" {
			podUsage.FSInodesFree = uint64(v)
		}
		if k == "fsInodes" {
			podUsage.FSInodes = uint64(v)
		}
		if k == "fsInodesUsed" {
			podUsage.FSInodesUsed = uint64(v)
		}
		parts := strings.Split(k, ".")
		if len(parts) != 2 {
			continue
		}
		unitName := parts[0]
		metric := parts[1]
		usage, ok := unitUsageMap[unitName]
		if !ok {
			usage = &usageMetrics{}
			unitUsageMap[unitName] = usage
		}
		if metric == "cpuUsage" {
			value := uint64(v)
			usage.UsageCoreNanoSeconds = value
			prevValue, ok := podMetrics[0].ResourceUsage[k]
			if ok {
				if nanoseconds > 0 {
					usage.UsageNanoCores = uint64(
						float64(value-uint64(prevValue)) /
							float64(nanoseconds) * nanosecondsPerSecond)
				}
			}
		}
		if metric == "memoryUsage" {
			usage.UsageBytes = uint64(v)
		}
		if metric == "memoryWorkingSet" {
			usage.WorkingSetBytes = uint64(v)
		}
		if metric == "memoryAvailable" {
			usage.AvailableBytes = uint64(v)
		}
		if metric == "memoryRSS" {
			usage.RSSBytes = uint64(v)
		}
		if metric == "memoryPageFault" {
			usage.PageFaults = uint64(v)
		}
		if metric == "memoryMajorPageFaults" {
			usage.MajorPageFaults = uint64(v)
		}
	}
	containerStats := make([]stats.ContainerStats, len(units))
	for i, unit := range units {
		usage, ok := unitUsageMap[unit.Name]
		if !ok {
			usage = &usageMetrics{}
		}
		cstats := stats.ContainerStats{
			Name:      unit.Name,
			StartTime: startTime,
			CPU: &stats.CPUStats{
				Time: timestamp,
			},
			Memory: &stats.MemoryStats{
				Time: timestamp,
			},
		}
		cstats.CPU.UsageNanoCores = valueOrNil(usage.UsageNanoCores)
		cstats.CPU.UsageCoreNanoSeconds = valueOrNil(usage.UsageCoreNanoSeconds)
		cstats.Memory.UsageBytes = valueOrNil(usage.UsageBytes)
		cstats.Memory.WorkingSetBytes = valueOrNil(usage.WorkingSetBytes)
		cstats.Memory.AvailableBytes = valueOrNil(usage.AvailableBytes)
		cstats.Memory.RSSBytes = valueOrNil(usage.RSSBytes)
		cstats.Memory.PageFaults = valueOrNil(usage.PageFaults)
		cstats.Memory.MajorPageFaults = valueOrNil(usage.MajorPageFaults)
		containerStats[i] = cstats
		podUsage.UsageNanoCores += usage.UsageNanoCores
		podUsage.UsageCoreNanoSeconds += usage.UsageCoreNanoSeconds
		podUsage.UsageBytes += usage.UsageBytes
		podUsage.WorkingSetBytes += usage.WorkingSetBytes
		podUsage.AvailableBytes += usage.AvailableBytes
		podUsage.RSSBytes += usage.RSSBytes
		podUsage.PageFaults += usage.PageFaults
		podUsage.MajorPageFaults += usage.MajorPageFaults
	}
	return podUsage, containerStats
}

func valueOrNil(value uint64) *uint64 {
	if value > 0 {
		return &value
	}
	return nil
}
