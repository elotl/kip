package server

import (
	"fmt"
	"reflect"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/server/events"
	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/golang/glog"
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
				Waiting: &api.UnitStateWaiting{Reason: "Initializing"},
			},
			Image: pod.Spec.InitUnits[i].Image,
		}
	}
	pod.Status.UnitStatuses = make([]api.UnitStatus, len(pod.Spec.Units))
	for i := range pod.Spec.Units {
		pod.Status.UnitStatuses[i] = api.UnitStatus{
			Name: pod.Spec.Units[i].Name,
			State: api.UnitState{
				Waiting: &api.UnitStateWaiting{Reason: "Initializing"},
			},
			Image: pod.Spec.Units[i].Image,
		}
	}
	msg := fmt.Sprintf("pod %s running on node %s", pod.Name, nodeName)
	eventSystem.Emit(events.PodRunning, "pod-controller", pod, msg)
	_, err := podRegistry.UpdatePodStatus(pod, "Pod is running")
	return err
}

// Notes for PR:
// Updated this to simplify return values (and handling those return values
// later on in updatePodStatus.  If a pod has an invalid status, we always
// fail the pod.  We can do that at the end of this function and return
// That simplifies code that depends on this function
func computePodPhase(policy api.RestartPolicy, unitstatus []api.UnitStatus, podName string) (api.PodPhase, string) {
	// Note: we need to treat the "created" unit state the same way as the
	// "running" state. Itzo will set the status of units to "created" right
	// after creating them, and only to "running" once the application is
	// started.
	valid := true
	phase := api.PodRunning
	var failMsg string
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
		// All unit status should be either running or waiting.
		phase = api.PodRunning
		for _, us := range unitstatus {
			if us.State.Waiting != nil || us.State.Running != nil {
				continue
			}
			// This should not happen - itzo will keep restarting units
			// when RestartPolicy is always, so they all should be in the
			// running state (or in the created state, if the helper in
			// itzo is a bit slow to start up).
			failMsg = fmt.Sprintf("Invalid unit status for unit %s", us.Name)
			glog.Warningln(failMsg)
			valid = false
		}
	case api.RestartPolicyNever:
		// If all units have succeeded, pod has succeeded.
		// If at least one unit is still running, pod is running.
		// If at least one unit has failed, and no units are running, pod has
		// failed.
		phase = api.PodSucceeded
		for _, us := range unitstatus {
			if us.State.Running != nil {
				failMsg = ""
				phase = api.PodRunning
				break
			}
			if us.State.Waiting != nil && phase != api.PodFailed {
				failMsg = ""
				phase = api.PodRunning
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
			if us.State.Running != nil || us.State.Waiting != nil {
				failMsg = ""
				phase = api.PodRunning
				break
			}
			if us.State.Terminated != nil && us.State.Terminated.ExitCode != 0 {
				failMsg = fmt.Sprintf("Invalid unit status for unit %s", us.Name)
				glog.Warningln(failMsg)
				valid = false
			}
		}
	}
	if !valid {
		glog.Warningf("Invalid unit state for pod %s. Setting pod phase to Failed", podName)
		phase = api.PodFailed
	}
	return phase, failMsg
}

// Requirements for being ready:
// All initUnit and all units have a unit status reported for them
// All initUnits must have completed successfully
// All units are either running or terminated in a way in which we consider
// successful:
//
// Requirements for successful termination of a unit:
// 1. RestartPolicyOnFailure and exitCode == 0
// 2. RestartPolicyNever
//
// The way the tests are coded makes the assumption that the status
// fields of the units are not invalid (all unit statuses have one
// valid State field)
func podIsReady(pod *api.Pod) bool {
	if pod.Status.Phase != api.PodRunning {
		return false
	}
	if len(pod.Spec.Units) > 0 {
		if len(pod.Status.UnitStatuses) != len(pod.Spec.Units) {
			return false
		}
		for i := range pod.Status.UnitStatuses {
			if pod.Status.UnitStatuses[i].State.Waiting != nil {
				return false
			} else if pod.Status.UnitStatuses[i].State.Terminated != nil {
				if pod.Spec.RestartPolicy == api.RestartPolicyAlways ||
					(pod.Spec.RestartPolicy == api.RestartPolicyOnFailure &&
						pod.Status.UnitStatuses[i].State.Terminated.ExitCode != 0) {
					return false
				}
			}
		}
	} else if len(pod.Spec.InitUnits) > 0 {
		if len(pod.Status.InitUnitStatuses) != len(pod.Spec.InitUnits) {
			return false
		}
		for i := range pod.Status.InitUnitStatuses {
			if pod.Status.InitUnitStatuses[i].State.Terminated == nil ||
				pod.Status.InitUnitStatuses[i].State.Terminated.ExitCode != 0 {
				return false
			}
		}
	}
	return true
}

func runningMaxLicensePods(podRegistry *registry.PodRegistry, maxResources int) bool {
	pods, err := podRegistry.ListPods(func(p *api.Pod) bool {
		return (p.Status.Phase == api.PodDispatching ||
			p.Status.Phase == api.PodRunning)
	})
	if err != nil {
		glog.Errorf("Error listing pods for checking license limits: %s", err.Error())
		return true
	}
	if len(pods.Items) >= maxResources {
		return true
	}
	return false
}

func remedyFailedPod(pod *api.Pod, podRegistry *registry.PodRegistry) {
	if pod.Status.StartFailures <= allowedStartFailures &&
		pod.Spec.RestartPolicy != api.RestartPolicyNever {
		msg := fmt.Sprintf("Pod %s has failed to start %d times, retrying",
			pod.Name, pod.Status.StartFailures)
		glog.Warningf("%s", msg)
		// reset most everything in the status
		pod.Status = api.PodStatus{
			Phase:         api.PodWaiting,
			StartFailures: pod.Status.StartFailures,
		}
		podRegistry.UpdatePodStatus(pod, msg)
	} else {
		glog.Errorf("pod %s has failed to start %d times. Not trying again, pod has failed",
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

	updatedReadyTime := false
	if pod.Status.ReadyTime == nil && podIsReady(pod) {
		now := api.Now()
		pod.Status.ReadyTime = &now
		updatedReadyTime = true
	}

	// Performance: could be faster with a manual comparison but
	// that's painful to maintain
	statusSame := reflect.DeepEqual(pod.Status.UnitStatuses, reply.UnitStatuses) && reflect.DeepEqual(pod.Status.InitUnitStatuses, reply.InitUnitStatuses)
	if podPhase == pod.Status.Phase &&
		statusSame &&
		!resetStartFailures &&
		!updatedReadyTime {
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
		glog.Infof("Changing pod %s phase %s -> %s",
			pod.Name, pod.Status.Phase, podPhase)
		pod.Status.Phase = podPhase
		if podPhase == api.PodFailed {
			failMsg = "Failure: " + failMsg
		}
		if (podPhase == api.PodFailed || podPhase == api.PodSucceeded) &&
			pod.Status.BoundNodeName == "" {
			glog.Errorf("Programming error: unbound pod %s is %s",
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
