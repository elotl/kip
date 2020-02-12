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
			phase: api.PodRunning,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyAlways,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitWaiting("bar"),
			},
			phase: api.PodRunning,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyAlways,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitFailed("bar"),
			},
			phase: api.PodRunning,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyAlways,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitSucceeded("bar"),
			},
			phase: api.PodRunning,
		},
		// RestartPolicyNever
		podPhaseInput{
			restartPolicy: api.RestartPolicyNever,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitRunning("bar"),
			},
			phase: api.PodRunning,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyNever,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitWaiting("bar"),
			},
			phase: api.PodRunning,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyNever,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitFailed("bar"),
			},
			phase: api.PodRunning,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyNever,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitSucceeded("bar"),
			},
			phase: api.PodRunning,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyNever,
			units: []api.UnitStatus{
				MakeUnitSucceeded("foo"),
				MakeUnitSucceeded("bar"),
			},
			phase: api.PodSucceeded,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyNever,
			units: []api.UnitStatus{
				MakeUnitSucceeded("bar"),
				MakeUnitSucceeded("foo"),
			},
			phase: api.PodSucceeded,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyNever,
			units: []api.UnitStatus{
				MakeUnitWaiting("bar"),
				MakeUnitFailed("foo"),
			},
			phase: api.PodRunning,
		},
		// RestartPolicyOnFailure
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitWaiting("bar"),
			},
			phase: api.PodRunning,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitRunning("bar"),
			},
			phase: api.PodRunning,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitFailed("bar"),
			},
			phase: api.PodRunning,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitRunning("foo"),
				MakeUnitSucceeded("bar"),
			},
			phase: api.PodRunning,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitSucceeded("foo"),
				MakeUnitSucceeded("bar"),
			},
			phase: api.PodSucceeded,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitSucceeded("foo"),
				MakeUnitFailed("bar"),
			},
			phase: api.PodRunning,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitFailed("bar"),
				MakeUnitSucceeded("foo"),
			},
			phase: api.PodRunning,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitStartFailure("bar"),
				MakeUnitSucceeded("foo"),
			},
			phase: api.PodFailed,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitSucceeded("bar"),
				MakeUnitSucceeded("foo"),
			},
			phase: api.PodSucceeded,
		},
		podPhaseInput{
			restartPolicy: api.RestartPolicyOnFailure,
			units: []api.UnitStatus{
				MakeUnitSucceeded("foo"),
			},
			phase: api.PodSucceeded,
		},
	}
	for i, inp := range inputs {
		msg := fmt.Sprintf("test %d", i)
		phase, failMsg := computePodPhase(inp.restartPolicy, inp.units, "testpod")
		if inp.phase == api.PodFailed {
			assert.NotEmpty(t, failMsg, msg)
		} else {
			assert.Empty(t, failMsg, msg)
			assert.Equal(t, inp.phase, phase, msg)
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
