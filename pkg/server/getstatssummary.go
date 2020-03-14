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

type usageMetrics struct {
	UsageNanoCores  uint64
	UsageBytes      uint64
	WorkingSetBytes uint64
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
		if len(podMetricsItems) < 1 {
			klog.V(2).Infof("no metrics found for pod %s", pod.Name)
			continue
		}
		// First metrics sample from the pod.
		firstSample := podMetricsItems[0]
		startTime := metav1.NewTime(firstSample.Timestamp.Time)
		// Last sample from the pod, with the latest metrics.
		lastSample := podMetricsItems[len(podMetricsItems)-1]
		namespace, name := util.SplitNamespaceAndName(pod.Name)
		timestamp := metav1.NewTime(lastSample.Timestamp.Time)
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
		}
		var podUsage usageMetrics
		podUsage, pss.Containers = getContainerStats(startTime, lastSample, pod.Spec.Units)
		pss.CPU.UsageNanoCores = &podUsage.UsageNanoCores
		pss.Memory.UsageBytes = &podUsage.UsageBytes
		pss.Memory.WorkingSetBytes = &podUsage.WorkingSetBytes
		res.Pods = append(res.Pods, pss)
	}
	klog.V(5).Infof("GetStatsSummary() %+v", res)
	return res, nil
}

func getContainerStats(startTime metav1.Time, podMetrics *api.Metrics, units []api.Unit) (usageMetrics, []stats.ContainerStats) {
	timestamp := metav1.NewTime(podMetrics.Timestamp.Time)
	unitUsageMap := make(map[string]*usageMetrics)
	for k, v := range podMetrics.ResourceUsage {
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
			usage.UsageNanoCores = uint64(v)
		}
		if metric == "memoryUsage" {
			usage.UsageBytes = uint64(v)
			if usage.WorkingSetBytes == 0 {
				// Old itzo versions don't support this metric.
				usage.WorkingSetBytes = usage.UsageBytes
			}
		}
		if metric == "memoryWorkingSet" {
			usage.WorkingSetBytes = uint64(v)
		}
	}
	podUsage := usageMetrics{}
	containerStats := make([]stats.ContainerStats, len(units))
	for i, unit := range units {
		usage, ok := unitUsageMap[unit.Name]
		if !ok {
			usage = &usageMetrics{}
		}
		containerStats[i] = stats.ContainerStats{
			Name:      unit.Name,
			StartTime: startTime,
			CPU: &stats.CPUStats{
				Time:           timestamp,
				UsageNanoCores: &usage.UsageNanoCores,
			},
			Memory: &stats.MemoryStats{
				Time:            timestamp,
				UsageBytes:      &usage.UsageBytes,
				WorkingSetBytes: &usage.WorkingSetBytes,
			},
		}
		podUsage.UsageNanoCores += usage.UsageNanoCores
		podUsage.UsageBytes += usage.UsageBytes
		podUsage.WorkingSetBytes += usage.WorkingSetBytes
	}
	if podUsage.WorkingSetBytes == 0 {
		podUsage.WorkingSetBytes = podUsage.UsageBytes
	}
	return podUsage, containerStats
}
