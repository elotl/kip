package server

import (
	"fmt"
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/util/rand"
	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
)

func fakeInstanceProvider() *InstanceProvider {
	ipStr := fmt.Sprintf(
		"%d.%d.%d.%d",
		rand.Intn(255),
		rand.Intn(255),
		rand.Intn(255),
		rand.Intn(255))
	return &InstanceProvider{
		nodeName:   rand.String(8),
		internalIP: ipStr,
	}
}

//func (p *InstanceProvider) getStatus(milpaPod *api.Pod, pod *v1.Pod) v1.PodStatus
func TestGetStatus(t *testing.T) {
	p := fakeInstanceProvider()
	milpaPod := api.GetFakePod()
	pod := &v1.Pod{}
	testCases := []struct {
		milpaPodPhase api.PodPhase
		k8sPodPhase   v1.PodPhase
	}{
		{
			milpaPodPhase: api.PodDispatching,
			k8sPodPhase:   v1.PodPending,
		},
		{
			milpaPodPhase: api.PodFailed,
			k8sPodPhase:   v1.PodFailed,
		},
		{
			milpaPodPhase: api.PodRunning,
			k8sPodPhase:   v1.PodRunning,
		},
		{
			milpaPodPhase: api.PodSucceeded,
			k8sPodPhase:   v1.PodSucceeded,
		},
		{
			milpaPodPhase: api.PodTerminated,
			k8sPodPhase:   v1.PodFailed,
		},
		{
			milpaPodPhase: api.PodWaiting,
			k8sPodPhase:   v1.PodPending,
		},
	}
	for _, tc := range testCases {
		milpaPod.Status.Phase = tc.milpaPodPhase
		podStatus := p.getStatus(milpaPod, pod)
		assert.Equal(t, podStatus.Phase, tc.k8sPodPhase)
	}
}

//func unitToContainerStatus(st api.UnitStatus) v1.ContainerStatus

//func containerToUnit(container v1.Container) api.Unit

//func unitToContainer(unit api.Unit, container *v1.Container) v1.Container

//func k8sToMilpaVolume(vol v1.Volume) *api.Volume

//func milpaToK8sVolume(vol api.Volume) *v1.Volume

//func (p *InstanceProvider) k8sToMilpaPod(pod *v1.Pod) (*api.Pod, error)

//func (p *InstanceProvider) milpaToK8sPod(milpaPod *api.Pod) (*v1.Pod, error)
