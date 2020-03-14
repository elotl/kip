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
	"reflect"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/events"
	"github.com/elotl/kip/pkg/server/registry"
	"k8s.io/klog"
)

// The routines in here are shared between pod_controller and
// pod_controller_utils.

const (
	allowedStartFailures int = 2
)

// // Update the phase to running and the (init)unitStatuses as waiting
// // so the user and anything reading the pod statuses knows what's
// // going on.
func setPodRunning(pod *api.Pod, nodeName string, podRegistry *registry.PodRegistry, eventSystem *events.EventSystem) error {
	pod.Status.Phase = api.PodRunning
	pod.Status.InitUnitStatuses = make([]api.UnitStatus, len(pod.Spec.InitUnits))
	for i := range pod.Spec.InitUnits {
		pod.Status.InitUnitStatuses[i] = api.UnitStatus{
			Name: pod.Spec.InitUnits[i].Name,
			State: api.UnitState{
				Waiting: &api.UnitStateWaiting{Reason: "PodInitializing"},
			},
			Image: pod.Spec.InitUnits[i].Image,
		}
	}
	pod.Status.UnitStatuses = make([]api.UnitStatus, len(pod.Spec.Units))
	for i := range pod.Spec.Units {
		pod.Status.UnitStatuses[i] = api.UnitStatus{
			Name: pod.Spec.Units[i].Name,
			State: api.UnitState{
				Waiting: &api.UnitStateWaiting{Reason: "PodInitializing"},
			},
			Image: pod.Spec.Units[i].Image,
		}
	}
	msg := fmt.Sprintf("pod %s running on node %s", pod.Name, nodeName)
	eventSystem.Emit(events.PodRunning, "pod-controller", pod, msg)
	_, err := podRegistry.UpdatePodStatus(pod, "Pod is running")
	return err
}

func computePodPhase(policy api.RestartPolicy, unitstatus []api.UnitStatus, podName string) (phase api.PodPhase, failMsg string) {
	// Note: we need to treat the "created" unit state the same way as the
	// "running" state. Itzo will set the status of units to "created" right
	// after creating them, and only to "running" once the application is
	// started.
	phase = api.PodRunning
	// We need to handle 0 units as "Running" for kiyot. If we no
	// longer need kiyot, remove this and use validation to ensure
	// pods always have > 1 unit
	if len(unitstatus) == 0 {
		return phase, failMsg
	}

	// If we ever fail to launch a unit, that fails the whole pod
	for _, us := range unitstatus {
		if us.State.Waiting != nil && us.State.Waiting.StartFailure {
			failMsg = fmt.Sprintf("Start failure for unit %s: %s", us.Name, us.State.Waiting.Reason)
			return api.PodFailed, failMsg
		}
	}

	switch policy {
	case api.RestartPolicyAlways:
		// Itzo should restart everything
		// Pods that are waiting to be restarted will be in a terminated state
		phase = api.PodRunning
	case api.RestartPolicyNever:
		// If all units have succeeded, pod has succeeded.
		// If at least one unit is waiting or still running, pod is running
		phase = api.PodSucceeded
		for _, us := range unitstatus {
			if us.State.Waiting != nil || us.State.Running != nil {
				failMsg = ""
				phase = api.PodRunning
				break
			}
			if us.State.Terminated != nil && us.State.Terminated.ExitCode != 0 {
				failMsg = fmt.Sprintf("Unit %s terminated with non-zero exit code %d", us.Name, us.State.Terminated.ExitCode)
				phase = api.PodFailed
			}
		}
	case api.RestartPolicyOnFailure:
		// If all units have succeeded, pod has succeeded.
		// If at least one unit is still running, pod is running.
		// Failed is not a valid status with this restart policy.
		phase = api.PodSucceeded
		for _, us := range unitstatus {
			if us.State.Running != nil ||
				us.State.Waiting != nil ||
				(us.State.Terminated != nil && us.State.Terminated.ExitCode != 0) {
				failMsg = ""
				phase = api.PodRunning
				break
			}
		}
	}
	return phase, failMsg
}

func remedyFailedPod(pod *api.Pod, podRegistry *registry.PodRegistry) {
	if pod.Status.StartFailures <= allowedStartFailures &&
		pod.Spec.RestartPolicy != api.RestartPolicyNever {
		msg := fmt.Sprintf("Pod %s has failed to start %d times, retrying",
			pod.Name, pod.Status.StartFailures)
		klog.Warningf("%s", msg)
		// reset most everything in the status
		pod.Status = api.PodStatus{
			Phase:         api.PodWaiting,
			StartFailures: pod.Status.StartFailures,
		}
		podRegistry.UpdatePodStatus(pod, msg)
	} else {
		klog.Errorf("pod %s has failed to start %d times. Not trying again, pod has failed",
			pod.Name, pod.Status.StartFailures)
		podRegistry.TerminatePod(pod, api.PodFailed,
			"Pod failed: too many start failures")
	}
}

func updatePodWithStatus(pod *api.Pod, reply FullPodStatus) (changed, startFailure bool, failMsg string) {
	policy := pod.Spec.RestartPolicy
	podPhase, unitFailMsg := computePodPhase(policy, reply.UnitStatuses, pod.Name)

	if policy == api.RestartPolicyAlways {
		policy = api.RestartPolicyOnFailure
	}
	initPodPhase, initUnitFailMsg := computePodPhase(policy, reply.InitUnitStatuses, pod.Name)
	failMsg = initUnitFailMsg + unitFailMsg
	if initPodPhase == api.PodFailed {
		podPhase = api.PodFailed
	}

	startFailure = false
	if podPhase == api.PodFailed {
		startFailure = wasStartFailure(reply.UnitStatuses) ||
			wasStartFailure(reply.InitUnitStatuses)
	}

	// Reset StartFailures if all units on the pod get out of the
	// waiting state
	resetStartFailures := false
	if pod.Status.StartFailures > 0 &&
		allUnitsStarted(reply.UnitStatuses) {
		if len(reply.InitUnitStatuses) == 0 || allUnitsStarted(reply.InitUnitStatuses) {
			resetStartFailures = true
		}
	}

	// Performance: could be faster with a manual comparison but
	// that's painful to maintain
	statusSame := reflect.DeepEqual(pod.Status.UnitStatuses, reply.UnitStatuses) && reflect.DeepEqual(pod.Status.InitUnitStatuses, reply.InitUnitStatuses)
	if podPhase == pod.Status.Phase &&
		statusSame &&
		!resetStartFailures {
		// No change.
		return false, false, ""
	}

	if !statusSame {
		pod.Status.UnitStatuses = reply.UnitStatuses
		pod.Status.InitUnitStatuses = reply.InitUnitStatuses
	}
	if resetStartFailures {
		pod.Status.StartFailures = 0
	}

	if podPhase != pod.Status.Phase {
		klog.V(2).Infof("Changing pod %s phase %s -> %s",
			pod.Name, pod.Status.Phase, podPhase)
		pod.Status.Phase = podPhase
		if podPhase == api.PodFailed {
			failMsg = "Failure: " + failMsg
		}
		if (podPhase == api.PodFailed || podPhase == api.PodSucceeded) &&
			pod.Status.BoundNodeName == "" {
			klog.Errorf("Programming error: unbound pod %s is %s",
				pod.Name, podPhase)
		}
	}
	return true, startFailure, failMsg
}

func wasStartFailure(unitstatus []api.UnitStatus) bool {
	for _, us := range unitstatus {
		if us.State.Waiting != nil && us.State.Waiting.StartFailure {
			return true
		}
	}
	return false
}

func allUnitsStarted(unitstatus []api.UnitStatus) bool {
	// if we have no unitsstatuses then we're likely waiting on units
	// to be sent from the controller to itzo. In that case, we say
	// that not all units have started. (If we're just looking at
	// InitUnits, that isn't the case)
	if len(unitstatus) == 0 {
		return false
	}
	for _, us := range unitstatus {
		if us.State.Waiting != nil {
			return false
		}
	}
	return true
}
