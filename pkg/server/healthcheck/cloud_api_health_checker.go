package healthcheck

import (
	"time"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/server/registry"
	"github.com/elotl/kip/pkg/util"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog"
)

type CloudAPIHealthCheck struct {
	podLister     registry.PodLister
	cloudClient   cloud.CloudClient
	terminateChan chan *api.Pod
}

func NewCloudAPIHealthCheck(podLister registry.PodLister, cloudClient cloud.CloudClient) *CloudAPIHealthCheck {
	return &CloudAPIHealthCheck{
		podLister:     podLister,
		cloudClient:   cloudClient,
		terminateChan: make(chan *api.Pod, terminateChanSize),
	}
}

func (chc *CloudAPIHealthCheck) Start() {
	for x := range time.Tick(checkPeriod) {
		if err := chc.checkPodInstances(); err != nil {
			continue
		}
		// Check on timedout instances
	}
}

func (chc *CloudAPIHealthCheck) checkPodInstances() error {
	// list instances, put in set
	instances, err := chc.cloudClient.ListInstances()
	if err != nil {
		return util.WrapError(err, "error listing cloud instances for cloud API health check")
	}
	instIDs := sets.NewString()
	for i := range instances {
		instIDs.Insert(instances[i].ID)
	}
	podList, err := chc.podLister.ListPods(func(p *api.Pod) bool {
		return p.Status.Phase == api.PodRunning
	})
	if err != nil {
		return util.WrapError(err, "error listing pods for cloud API health check")
	}
	for _, pod := range podList.Items {
		podInstID := pod.Status.BoundInstanceID
		if podInstID == "" {
			klog.Warningf("cloud instance health check found pod with empty BoundInstanceID: %s/%s", pod.Namespace, pod.Name)
			continue
		}
		if !instIDs.Has(podInstID) {
			chc.handleMissingInstance(pod)
		}
	}
	return nil
}

func (chc *CloudAPIHealthCheck) handleMissingInstance(pod *api.Pod) {
	// TODO!
	// The API might have made a mistake or dropped an instance, how
	// long can we allow the health check to fail?  How long should
	// we keep track of these things?
}

func (chc *CloudAPIHealthCheck) TerminatePodsChan() <-chan *api.Pod {
	return chc.terminateChan
}
