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

package healthcheck

import (
	"strconv"
	"time"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/api/annotations"
	"github.com/elotl/kip/pkg/nodeclient"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/server/registry"
	"github.com/elotl/kip/pkg/util/conmap"
	"k8s.io/klog"
)

const (
	terminateChanSize = 1000
)

type healthChecker interface {
	// Goes through pods and updates the time in the conmap if the pod
	// is healthy.  This is a bit of a kluge since, for the
	// statusHealthCheck the PodController actually updates the conmap
	checkPods(*conmap.StringTimeTime) error

	// Fail the pod after one last check.  This got pushed into the
	// interface since it needs to be a goroutine for
	// statusHealthCheck but should be syncronous for tests and for
	// the cloudAPIHealthCheck
	maybeFailUnresponsivePod(pod *api.Pod, terminateChan chan *api.Pod)
}

type HealthCheckController struct {
	podLister        registry.PodLister
	lastStatusTime   *conmap.StringTimeTime
	checkInterval    time.Duration
	unhealthyTimeout time.Duration
	terminateChan    chan *api.Pod
	checker          healthChecker
}

func NewStatusHealthChecker(
	podLister registry.PodLister,
	nodeLister registry.NodeLister,
	nodeClientFactory nodeclient.ItzoClientFactoryer,
	checkInterval time.Duration,
	unhealthyTimeout time.Duration) *HealthCheckController {
	lastStatusTime := conmap.NewStringTimeTime()
	return &HealthCheckController{
		podLister:        podLister,
		lastStatusTime:   lastStatusTime,
		checkInterval:    checkInterval,
		unhealthyTimeout: unhealthyTimeout,
		terminateChan:    make(chan *api.Pod, terminateChanSize),
		checker: &statusHealthCheck{
			nodeLister:        nodeLister,
			nodeClientFactory: nodeClientFactory,
			lastStatusTime:    lastStatusTime,
		},
	}
}

func NewCloudAPIHealthChecker(
	podLister registry.PodLister,
	cloudClient cloud.CloudClient,
	checkInterval time.Duration,
	unhealthyTimeout time.Duration) *HealthCheckController {
	return &HealthCheckController{
		podLister:        podLister,
		lastStatusTime:   conmap.NewStringTimeTime(),
		checkInterval:    checkInterval,
		unhealthyTimeout: unhealthyTimeout,
		terminateChan:    make(chan *api.Pod, terminateChanSize),
		checker: &cloudAPIHealthCheck{
			podLister:   podLister,
			cloudClient: cloudClient,
		},
	}
}

func (c *HealthCheckController) Start() {
	for range time.Tick(c.checkInterval) {
		if err := c.checker.checkPods(c.lastStatusTime); err != nil {
			// If our pod check fails, don't try to terminate pods
			// pods should only be terminated if we know they're not
			// running. We don't want a failure in the cloud API to
			// kill pods when they're still functioning OK.
			klog.Errorf("Error checking on status of pods: %v", err)
			continue
		}
		c.handlePodTimeouts()
		c.cleanupLastStatusTimes()
	}
}

func (c *HealthCheckController) getPodHealthCheckTimeout(pod *api.Pod) time.Duration {
	unhealthyTimeout := c.unhealthyTimeout
	if val, ok := pod.Annotations[annotations.PodHealthcheckHealthyTimeout]; ok {
		t, err := strconv.ParseFloat(val, 64)
		if err == nil {
			unhealthyTimeout = time.Duration(t) * time.Second
		}
	}
	return unhealthyTimeout
}

// Check if a pod has timed out, if so, call checker.maybeFailUnresponsivePod
// to terminate the unresponsive pod
func (c *HealthCheckController) handlePodTimeouts() {
	podList, err := c.podLister.ListPods(func(p *api.Pod) bool {
		return p.Status.Phase == api.PodRunning
	})
	if err != nil {
		klog.Errorf("Error getting list of pods from registry")
		return
	}
	now := time.Now().UTC()
	for _, pod := range podList.Items {
		last, exists := c.lastStatusTime.GetOK(pod.Name)
		if !exists {
			c.lastStatusTime.Set(pod.Name, now)
			continue
		}
		unhealthyTimeout := c.getPodHealthCheckTimeout(pod)
		if unhealthyTimeout <= 0 {
			continue
		}
		if now.Sub(last) >= unhealthyTimeout {
			klog.Warningf("pod %s has failed healthchecks for %ds pod has likely failed", pod.Name, int(unhealthyTimeout.Seconds()))
			c.checker.maybeFailUnresponsivePod(pod, c.terminateChan)
		}
	}
}

func (c *HealthCheckController) cleanupLastStatusTimes() {
	runningPods := make(map[string]bool)
	_, err := c.podLister.ListPods(func(p *api.Pod) bool {
		if p.Status.Phase == api.PodRunning {
			runningPods[p.Name] = true
		}
		return false
	})
	if err != nil {
		klog.Errorf("Error getting list of pods from registry")
		return
	}
	for _, item := range c.lastStatusTime.Items() {
		podName := item.Key
		_, exists := runningPods[podName]
		if !exists {
			c.lastStatusTime.Delete(podName)
		}
	}
}

func (c *HealthCheckController) TerminatePodsChan() <-chan *api.Pod {
	return c.terminateChan
}

func (c *HealthCheckController) ClearLastStatusTime(podName string) {
	c.lastStatusTime.Delete(podName)
}

func (c *HealthCheckController) SetLastStatusTime(podName string) {
	c.lastStatusTime.Set(podName, time.Now().UTC())
}

func (c *HealthCheckController) LastStatusTime(podName string) (time.Time, bool) {
	return c.lastStatusTime.GetOK(podName)
}
