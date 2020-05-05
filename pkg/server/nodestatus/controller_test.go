package nodestatus

import (
	"testing"

	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

//func (n *NodeStatusController) setNodeStatus()
func TestCheckCloud(t *testing.T) {
	cloud := &cloud.MockCloudClient{
		AvailabilityChecker: func() (bool, error) {
			return true, nil
		},
	}
	ctrl := &NodeStatusController{
		cloudClient: cloud,
	}

	wasSet := ctrl.setNodeStatus()
	assert.False(t, wasSet)
	ctrl.cb = func(*corev1.Node) {}
	wasSet = ctrl.setNodeStatus()
	assert.False(t, wasSet)

	ctrl.node = &corev1.Node{}
	wasSet = ctrl.setNodeStatus()
	assert.True(t, wasSet)
	assert.True(t, ctrl.nodeReady)
	assert.False(t, ctrl.networkUnavailable)
	cloud.AvailabilityChecker = func() (bool, error) {
		return false, nil
	}
	wasSet = ctrl.setNodeStatus()
	assert.True(t, wasSet)
	assert.False(t, ctrl.nodeReady)
	assert.True(t, ctrl.networkUnavailable)
}

//func (n *NodeStatusController) GetNodeStatus() corev1.NodeStatus
func TestGetNodeStatus(t *testing.T) {
	cloud := &cloud.MockCloudClient{
		AvailabilityChecker: func() (bool, error) {
			return true, nil
		},
	}
	capacity := corev1.ResourceList{
		"cpu":    resource.MustParse("100m"),
		"memory": resource.MustParse("2Gi"),
	}
	ctrl := &NodeStatusController{
		cloudClient:        cloud,
		nodeReady:          false,
		networkUnavailable: true,
		internalIP:         "10.20.30.40",
		daemonEndpointPort: 12345,
		kubeletCapacity:    capacity,
		cb:                 func(*corev1.Node) {},
		node:               &corev1.Node{},
	}
	status := ctrl.GetNodeStatus()
	assert.Equal(t, "10.20.30.40", status.Addresses[0].Address)
	assert.Equal(t, v1.NodeInternalIP, status.Addresses[0].Type)
	assert.Equal(t, int32(12345), status.DaemonEndpoints.KubeletEndpoint.Port)
	assert.Equal(t, capacity, status.Capacity)
	for _, cond := range status.Conditions {
		if cond.Type == v1.NodeReady {
			assert.Equal(t, v1.ConditionFalse, cond.Status)
		}
		if cond.Type == v1.NodeNetworkUnavailable {
			assert.Equal(t, v1.ConditionTrue, cond.Status)
		}
	}
	ctrl.setNodeStatus()
	status = ctrl.GetNodeStatus()
	for _, cond := range status.Conditions {
		if cond.Type == v1.NodeReady {
			assert.Equal(t, v1.ConditionTrue, cond.Status)
		}
		if cond.Type == v1.NodeNetworkUnavailable {
			assert.Equal(t, v1.ConditionFalse, cond.Status)
		}
	}
	cloud.AvailabilityChecker = func() (bool, error) {
		return false, nil
	}
	ctrl.setNodeStatus()
	status = ctrl.GetNodeStatus()
	for _, cond := range status.Conditions {
		if cond.Type == v1.NodeReady {
			assert.Equal(t, v1.ConditionFalse, cond.Status)
		}
		if cond.Type == v1.NodeNetworkUnavailable {
			assert.Equal(t, v1.ConditionTrue, cond.Status)
		}
	}
}
