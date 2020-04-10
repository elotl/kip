package healthcheck

import (
	"fmt"
	"testing"
	"time"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/api/annotations"
	"github.com/elotl/kip/pkg/server/registry"
	"github.com/elotl/kip/pkg/util/conmap"
	"github.com/stretchr/testify/assert"
)

type mockHealthChecker struct {
	podChecker func(*conmap.StringTimeTime) error
	podFailer  func(*api.Pod) bool
}

func (m *mockHealthChecker) checkPods(cm *conmap.StringTimeTime) error {
	return m.podChecker(cm)
}

func (m *mockHealthChecker) podHasFailed(pod *api.Pod) bool {
	return m.podFailer(pod)
}

func makeTestHealthChecker() (*HealthCheckController, func()) {
	podReg, closer := registry.SetupTestPodRegistry()
	hc := HealthCheckController{
		podLister:               podReg,
		lastStatusTime:          conmap.NewStringTimeTime(),
		checkPeriod:             1 * time.Second,
		defaultUnhealthyTimeout: 10 * time.Second,
		terminateChan:           make(chan *api.Pod, terminateChanSize),
		checker: &mockHealthChecker{
			podFailer: func(p *api.Pod) bool { return true },
		},
	}
	return &hc, closer
}

func podWasFailed(c <-chan *api.Pod) bool {
	select {
	case <-c:
		fmt.Println("returning true")
		return true
	default:
		fmt.Println("returning false")
		return false
	}
}

func TestHandlePodTimeouts(t *testing.T) {
	tests := []struct {
		name        string
		podMod      func(p *api.Pod)
		lastCheckin time.Time
		shouldFail  bool
	}{
		{
			name:       "pod with no entry sets last check entry",
			shouldFail: false,
		},
		{
			name:        "expired check fails the pod",
			lastCheckin: time.Now().Add(-1 * time.Hour),
			shouldFail:  true,
		},
		{
			name:        "unexpired pod doesn't fail",
			lastCheckin: time.Now(),
			shouldFail:  false,
		},
		{
			name: "expired but annotated pod doesn't fail",
			podMod: func(p *api.Pod) {
				p.Annotations = map[string]string{
					annotations.PodHealthcheckTimeout: "1000000",
				}
			},
			lastCheckin: time.Now().Add(-1 * time.Hour),
			shouldFail:  false,
		},
		{
			name: "not running pod does not fail",
			// We'll hack this a bit and say the not running pod
			// hasn't checked in for a long time
			podMod: func(p *api.Pod) {
				p.Status.Phase = api.PodWaiting
			},
			lastCheckin: time.Now().Add(-1 * time.Hour),
			shouldFail:  false,
		},
	}
	for _, tc := range tests {
		ctl, closer := makeTestHealthChecker()
		defer closer()
		podReg := ctl.podLister.(*registry.PodRegistry)

		p := api.GetFakePod()
		p.Spec.Phase = api.PodRunning
		p.Status.Phase = api.PodRunning
		if tc.podMod != nil {
			tc.podMod(p)
		}
		p, err := podReg.CreatePod(p)
		assert.NoError(t, err)
		if !tc.lastCheckin.IsZero() {
			ctl.lastStatusTime.Set(p.Name, tc.lastCheckin)
		}

		ctl.handlePodTimeouts()

		assert.Equal(t, tc.shouldFail, podWasFailed(ctl.TerminatePodsChan()), tc.name)
		// Make sure that any pods that didn't have a lastStatusTime
		// set, get lastStatusTime set
		if tc.lastCheckin.IsZero() {
			tm, ok := ctl.lastStatusTime.GetOK(p.Name)
			assert.True(t, ok, tc.name)
			assert.False(t, tm.IsZero(), tc.name)
		}
	}
}

func TestMaybeFailUnresponsivePod(t *testing.T) {
	ctl, closer := makeTestHealthChecker()
	defer closer()
	p := api.GetFakePod()
	ctl.checker = &mockHealthChecker{
		podFailer: func(p *api.Pod) bool { return true },
	}
	ctl.maybeFailUnresponsivePod(p)
	assert.True(t, podWasFailed(ctl.TerminatePodsChan()))
	ctl.checker = &mockHealthChecker{
		podFailer: func(p *api.Pod) bool { return false },
	}
	ctl.maybeFailUnresponsivePod(p)
	assert.False(t, podWasFailed(ctl.TerminatePodsChan()))
	_, ok := ctl.lastStatusTime.GetOK(p.Name)
	assert.True(t, ok)
}

func TestCleanupLastStatusTimes(t *testing.T) {
	t.Parallel()
	ctl, closer := makeTestHealthChecker()
	defer closer()
	podReg := ctl.podLister.(*registry.PodRegistry)
	p1 := api.GetFakePod()
	p1.Name = "pod1"
	p1.Spec.Phase = api.PodRunning
	p1.Status.Phase = api.PodRunning
	p1, err := podReg.CreatePod(p1)
	assert.Nil(t, err)
	p2 := api.GetFakePod()
	p2.Name = "pod2"
	p2.Spec.Phase = api.PodTerminated
	p2.Status.Phase = api.PodTerminated
	p2, err = podReg.CreatePod(p2)
	assert.Nil(t, err)
	ctl.lastStatusTime.Set(p1.Name, time.Now().UTC())
	ctl.lastStatusTime.Set(p2.Name, time.Now().UTC())
	ctl.cleanupLastStatusTimes()
	_, exists := ctl.lastStatusTime.GetOK(p1.Name)
	assert.True(t, exists)
	_, exists = ctl.lastStatusTime.GetOK(p2.Name)
	assert.False(t, exists)
}
