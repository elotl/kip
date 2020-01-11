package server

import (
	"fmt"
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/stretchr/testify/assert"
)

func TestRemedyFailedPod(t *testing.T) {
	tests := []struct {
		startFails    int
		restartPolicy api.RestartPolicy
		expectedPhase api.PodPhase
	}{
		{
			startFails:    1,
			restartPolicy: api.RestartPolicyAlways,
			expectedPhase: api.PodWaiting,
		},
		{
			startFails:    allowedStartFailures,
			restartPolicy: api.RestartPolicyAlways,
			expectedPhase: api.PodWaiting,
		},
		{
			startFails:    allowedStartFailures + 1,
			restartPolicy: api.RestartPolicyAlways,
			expectedPhase: api.PodFailed,
		},
		{
			startFails:    0,
			restartPolicy: api.RestartPolicyNever,
			expectedPhase: api.PodFailed,
		},
	}
	podReg, closer := registry.SetupTestPodRegistry()
	defer closer()
	for i, tc := range tests {
		pod := api.GetFakePod()
		pod.Status.Phase = api.PodFailed
		pod.Spec.RestartPolicy = tc.restartPolicy
		pod.Status.StartFailures = tc.startFails
		_, err := podReg.CreatePod(pod)
		assert.NoError(t, err)
		remedyFailedPod(pod, podReg)
		p, err := podReg.GetPod(pod.Name)
		assert.NoError(t, err)
		msg := fmt.Sprintf("test %d", i)
		assert.Equal(t, tc.expectedPhase, p.Status.Phase, msg)
		assert.Equal(t, tc.startFails, p.Status.StartFailures, msg)
		if tc.expectedPhase == api.PodFailed {
			assert.Equal(t, api.PodFailed, p.Spec.Phase, msg)
		}
	}
}

func TestComputePodPhase(t *testing.T) {
	t.Parallel()
	inputs := []podPhaseInput{
		// RestartPolicyAlways
		podPhaseInput{
			restartPolicy: api.RestartPolicyAlways,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitRunning("bar"),
			},
			phase:   api.PodRunning,
			isValid: true,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyAlways,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitWaiting("bar"),
			},
			phase:   api.PodRunning,
			isValid: true,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyAlways,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitFailed("bar"),
			},
			isValid: false,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyAlways,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitSucceeded("bar"),
			},
			isValid: false,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyAlways,
			units: []api.UnitStatus{
				MakeUnitFailed("bar"),
				MakeUnitRunning("foo"),
			},
			isValid: false,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyAlways,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitSucceeded("bar"),
			},
			phase:   api.PodRunning,
			isValid: false,
		},
		// RestartPolicyNever
		podPhaseInput{
			restartPolicy: api.RestartPolicyNever,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitRunning("bar"),
			},
			phase:   api.PodRunning,
			isValid: true,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyNever,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitWaiting("bar"),
			},
			phase:   api.PodRunning,
			isValid: true,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyNever,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitFailed("bar"),
			},
			phase:   api.PodRunning,
			isValid: true,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyNever,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitSucceeded("bar"),
			},
			phase:   api.PodRunning,
			isValid: true,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyNever,
			units: []api.UnitStatus{
				MakeUnitSucceeded("foo"),
				MakeUnitSucceeded("bar"),
			},
			phase:   api.PodSucceeded,
			isValid: true,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyNever,
			units: []api.UnitStatus{
				MakeUnitSucceeded("bar"),
				MakeUnitSucceeded("foo"),
			},
			phase:   api.PodSucceeded,
			isValid: true,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyNever,
			units: []api.UnitStatus{
				MakeUnitWaiting("bar"),
				MakeUnitFailed("foo"),
			},
			phase:   api.PodFailed,
			isValid: true,
		},
		// RestartPolicyOnFailure
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitWaiting("bar"),
			},
			phase:   api.PodRunning,
			isValid: true,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitRunning("bar"),
			},
			phase:   api.PodRunning,
			isValid: true,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitFailed("bar"),
			},
			phase:   api.PodRunning,
			isValid: true,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitSucceeded("bar"),
			},
			phase:   api.PodRunning,
			isValid: true,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitSucceeded("foo"),
				MakeUnitSucceeded("bar"),
			},
			phase:   api.PodSucceeded,
			isValid: true,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitSucceeded("foo"),
				MakeUnitFailed("bar"),
			},
			isValid: false,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitSucceeded("foo"),
				MakeUnitStartFailure("bar"),
			},
			phase:   api.PodFailed,
			isValid: true,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitFailed("bar"),
				MakeUnitSucceeded("foo"),
			},
			isValid: false,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitStartFailure("bar"),
				MakeUnitSucceeded("foo"),
			},
			phase:   api.PodFailed,
			isValid: true,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitSucceeded("bar"),
				MakeUnitSucceeded("foo"),
			},
			phase:   api.PodSucceeded,
			isValid: true,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitSucceeded("foo"),
			},
			phase:   api.PodSucceeded,
			isValid: true,
		},
	}
	for _, inp := range inputs {
		phase, failMsg := computePodPhase(inp.restartPolicy, inp.units, "testpod")
		if !inp.isValid || inp.phase == api.PodFailed {
			assert.NotEmpty(t, failMsg)
		} else {
			assert.Empty(t, failMsg)
			assert.Equal(t, inp.phase, phase)
		}

	}
}

func TestRunningMaxLicensePods(t *testing.T) {
	podRegistry, closer := registry.SetupTestPodRegistry()
	defer closer()

	v := runningMaxLicensePods(podRegistry, 0)
	assert.True(t, v)

	p := api.GetFakePod()
	p.Status.Phase = api.PodDispatching
	_, err := podRegistry.CreatePod(p)
	assert.NoError(t, err)

	v = runningMaxLicensePods(podRegistry, 1)
	assert.True(t, v)
	v = runningMaxLicensePods(podRegistry, 2)
	assert.False(t, v)
}

func TestAllUnitsStarted(t *testing.T) {
	notStarted := [][]api.UnitStatus{
		[]api.UnitStatus{},
		[]api.UnitStatus{
			{State: api.UnitState{Waiting: &api.UnitStateWaiting{}}},
		},
		[]api.UnitStatus{
			{State: api.UnitState{Running: &api.UnitStateRunning{}}},
			{State: api.UnitState{Waiting: &api.UnitStateWaiting{}}},
		},
		[]api.UnitStatus{
			{State: api.UnitState{Terminated: &api.UnitStateTerminated{}}},
			{State: api.UnitState{Waiting: &api.UnitStateWaiting{}}},
		},
	}
	started := [][]api.UnitStatus{
		[]api.UnitStatus{
			{State: api.UnitState{Running: &api.UnitStateRunning{}}},
		},
		[]api.UnitStatus{
			{State: api.UnitState{Running: &api.UnitStateRunning{}}},
			{State: api.UnitState{Terminated: &api.UnitStateTerminated{}}},
		},
	}

	for i, s := range notStarted {
		assert.False(t, allUnitsStarted(s), "not started test %d", i)
	}
	for i, s := range started {
		assert.True(t, allUnitsStarted(s), "started test %d", i)
	}
}

func TestPodIsReady(t *testing.T) {
	tests := []struct {
		description string
		podmod      func(*api.Pod)
		ready       bool
	}{
		{
			description: "pod is not in running phase",
			podmod: func(pod *api.Pod) {
				pod.Status.Phase = api.PodFailed
			},
			ready: false,
		},
		{
			description: "empty pod should be ready",
			podmod:      func(pod *api.Pod) {},
			ready:       true,
		},
		{
			description: "pod with initUnits but no status is not ready",
			podmod: func(pod *api.Pod) {
				pod.Spec.InitUnits = []api.Unit{api.Unit{}}
			},
			ready: false,
		},
		{
			description: "pod with units but no status is not ready",
			podmod: func(pod *api.Pod) {
				pod.Spec.Units = []api.Unit{api.Unit{}}
			},
			ready: false,
		},
		{
			description: "pod with units and running status is ready",
			podmod: func(pod *api.Pod) {
				pod.Spec.Units = []api.Unit{api.Unit{}}
				pod.Status.UnitStatuses = []api.UnitStatus{api.UnitStatus{}}
				pod.Status.UnitStatuses[0].State.Running = &api.UnitStateRunning{}
			},
			ready: true,
		},
		{
			description: "pod with units and waiting status is not ready",
			podmod: func(pod *api.Pod) {
				pod.Spec.Units = []api.Unit{api.Unit{}}
				pod.Status.UnitStatuses = []api.UnitStatus{api.UnitStatus{}}
				pod.Status.UnitStatuses[0].State.Waiting = &api.UnitStateWaiting{}
			},
			ready: false,
		},
		{
			description: "pod with waiting initunitsis not ready",
			podmod: func(pod *api.Pod) {
				pod.Spec.InitUnits = []api.Unit{api.Unit{}}
				pod.Status.InitUnitStatuses = []api.UnitStatus{api.UnitStatus{}}
				pod.Status.InitUnitStatuses[0].State.Waiting = &api.UnitStateWaiting{}
			},
			ready: false,
		},
		{
			description: "pod with failed initunits is not ready",
			podmod: func(pod *api.Pod) {
				pod.Spec.InitUnits = []api.Unit{api.Unit{}}
				pod.Status.InitUnitStatuses = []api.UnitStatus{api.UnitStatus{}}
				pod.Status.InitUnitStatuses[0].State.Terminated = &api.UnitStateTerminated{ExitCode: 2}
			},
			ready: false,
		},
		{
			description: "pod with running initunits is not ready",
			podmod: func(pod *api.Pod) {
				pod.Spec.InitUnits = []api.Unit{api.Unit{}}
				pod.Status.InitUnitStatuses = []api.UnitStatus{api.UnitStatus{}}
				pod.Status.InitUnitStatuses[0].State.Running = &api.UnitStateRunning{}
			},
			ready: false,
		},
		{
			description: "pod with finished initunits and waiting pods is not ready",
			podmod: func(pod *api.Pod) {
				pod.Spec.InitUnits = []api.Unit{api.Unit{}}
				pod.Status.InitUnitStatuses = []api.UnitStatus{api.UnitStatus{}}
				pod.Status.InitUnitStatuses[0].State.Terminated = &api.UnitStateTerminated{ExitCode: 0}
				pod.Spec.Units = []api.Unit{api.Unit{}}
				pod.Status.UnitStatuses = []api.UnitStatus{api.UnitStatus{}}
				pod.Status.UnitStatuses[0].State.Waiting = &api.UnitStateWaiting{}
			},
			ready: false,
		},
		{
			description: "pod with finished initunits and running pods is ready",
			podmod: func(pod *api.Pod) {
				pod.Spec.InitUnits = []api.Unit{api.Unit{}}
				pod.Status.InitUnitStatuses = []api.UnitStatus{api.UnitStatus{}}
				pod.Status.InitUnitStatuses[0].State.Terminated = &api.UnitStateTerminated{ExitCode: 0}
				pod.Spec.Units = []api.Unit{api.Unit{}}
				pod.Status.UnitStatuses = []api.UnitStatus{api.UnitStatus{}}
				pod.Status.UnitStatuses[0].State.Running = &api.UnitStateRunning{}
			},
			ready: true,
		},
		{
			description: "pod with one waiting and one running pods is not ready",
			podmod: func(pod *api.Pod) {
				pod.Spec.Units = []api.Unit{api.Unit{}, api.Unit{}}
				pod.Status.UnitStatuses = []api.UnitStatus{api.UnitStatus{}, api.UnitStatus{}}
				pod.Status.UnitStatuses[0].State.Running = &api.UnitStateRunning{}
				pod.Status.UnitStatuses[1].State.Waiting = &api.UnitStateWaiting{}
			},
			ready: false,
		},
		{
			description: "pod with one waiting and one termianted pod (RestartPolicyAlways) is not ready",
			podmod: func(pod *api.Pod) {
				pod.Spec.Units = []api.Unit{api.Unit{}, api.Unit{}}
				pod.Status.UnitStatuses = []api.UnitStatus{api.UnitStatus{}, api.UnitStatus{}}
				pod.Status.UnitStatuses[0].State.Running = &api.UnitStateRunning{}
				pod.Status.UnitStatuses[1].State.Terminated = &api.UnitStateTerminated{ExitCode: 0}
			},
			ready: false,
		},
		{
			description: "pod with one waiting and one termianted pod (RestartPolicyOnFailure) is ready",
			podmod: func(pod *api.Pod) {
				pod.Spec.RestartPolicy = api.RestartPolicyOnFailure
				pod.Spec.Units = []api.Unit{api.Unit{}, api.Unit{}}
				pod.Status.UnitStatuses = []api.UnitStatus{api.UnitStatus{}, api.UnitStatus{}}
				pod.Status.UnitStatuses[0].State.Running = &api.UnitStateRunning{}
				pod.Status.UnitStatuses[1].State.Terminated = &api.UnitStateTerminated{ExitCode: 0}
			},
			ready: true,
		},
		{
			description: "pod with one waiting and one termianted pod (RestartPolicyNever) is ready",
			podmod: func(pod *api.Pod) {
				pod.Spec.RestartPolicy = api.RestartPolicyNever
				pod.Spec.Units = []api.Unit{api.Unit{}, api.Unit{}}
				pod.Status.UnitStatuses = []api.UnitStatus{api.UnitStatus{}, api.UnitStatus{}}
				pod.Status.UnitStatuses[0].State.Running = &api.UnitStateRunning{}
				pod.Status.UnitStatuses[1].State.Terminated = &api.UnitStateTerminated{ExitCode: 1}
			},
			ready: true,
		},
	}
	for i, tc := range tests {
		pod := api.NewPod()
		pod.Status.Phase = api.PodRunning
		tc.podmod(pod)
		isReady := podIsReady(pod)
		if tc.ready != isReady {
			msg := fmt.Sprintf("test PodIsReady %d: %s failed", i, tc.description)
			assert.Fail(t, msg)
		}
	}
}
