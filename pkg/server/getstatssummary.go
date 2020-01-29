package server

import (
	"context"
	"strings"
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/golang/glog"
	"github.com/virtual-kubelet/virtual-kubelet/trace"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	stats "k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1"
)

func (p *InstanceProvider) GetStatsSummary(ctx context.Context) (*stats.Summary, error) {
	var span trace.Span
	ctx, span = trace.StartSpan(ctx, "GetStatsSummary")
	defer span.End()
	glog.Infof("GetStatsSummary()")
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
		glog.Errorf("listing pods for stats: %v", err)
		return nil, util.WrapError(err, "listing pods for stats")
	}
	metricsRegistry := p.getMetricsRegistry()
	for _, pod := range pods.Items {
		podMetricsList := metricsRegistry.GetPodMetrics(pod.Name)
		podMetricsItems := podMetricsList.Items
		if len(podMetricsItems) < 1 {
			glog.V(3).Infof("no metrics found for pod %s", pod.Name)
			continue
		}
		// First metrics sample from the pod.
		podMetrics := podMetricsItems[0]
		startTime := metav1.NewTime(podMetrics.Timestamp.Time)
		// Last sample from the pod, with the latest metrics.
		podMetrics = podMetricsItems[len(podMetricsItems)-1]
		podUsageNanoCores := uint64(0)
		podUsageBytes := uint64(0)
		podWorkingSetBytes := uint64(0)
		namespace, name := util.SplitNamespaceAndName(pod.Name)
		timestamp := metav1.NewTime(podMetrics.Timestamp.Time)
		pss := stats.PodStats{
			PodRef: stats.PodReference{
				Name:      name,
				Namespace: namespace,
				UID:       pod.UID,
			},
			StartTime: startTime,
			CPU: &stats.CPUStats{
				Time:           timestamp,
				UsageNanoCores: &podUsageNanoCores,
			},
			Memory: &stats.MemoryStats{
				Time:            timestamp,
				UsageBytes:      &podUsageBytes,
				WorkingSetBytes: &podWorkingSetBytes,
			},
		}
		for _, unit := range pod.Spec.Units {
			usageNanoCores := uint64(0)
			usageBytes := uint64(0)
			workingSetBytes := uint64(0)
			for k, v := range podMetrics.ResourceUsage {
				if !strings.HasPrefix(k, unit.Name+".") {
					continue
				}
				parts := strings.Split(k, ".")
				if len(parts) != 2 {
					continue
				}
				if parts[1] == "cpuUsage" {
					usageNanoCores = uint64(v)
				}
				if parts[1] == "memoryUsage" {
					usageBytes = uint64(v)
					if workingSetBytes == 0 {
						// Old itzo versions don't support this metric.
						workingSetBytes = uint64(v)
					}
				}
				if parts[1] == "memoryWorkingSet" {
					workingSetBytes = uint64(v)
				}
			}
			pss.Containers = append(pss.Containers, stats.ContainerStats{
				Name:      unit.Name,
				StartTime: startTime,
				CPU: &stats.CPUStats{
					Time:           timestamp,
					UsageNanoCores: &usageNanoCores,
				},
				Memory: &stats.MemoryStats{
					Time:            timestamp,
					UsageBytes:      &usageBytes,
					WorkingSetBytes: &workingSetBytes,
				},
			})
			podUsageNanoCores += usageNanoCores
			podUsageBytes += usageBytes
			podWorkingSetBytes += workingSetBytes
		}
		if podWorkingSetBytes == 0 {
			podWorkingSetBytes = podUsageBytes
		}
		res.Pods = append(res.Pods, pss)
	}
	glog.Infof("GetStatsSummary() %+v", res)
	return res, nil
}
