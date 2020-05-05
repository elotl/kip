package nodestatus

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/util/stats"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

type NodeStatusController struct {
	nodeReady          bool
	networkUnavailable bool
	internalIP         string
	daemonEndpointPort int32
	kubeletCapacity    corev1.ResourceList
	cidrs              []string
	node               *corev1.Node
	cloudClient        cloud.CloudClient
	controlLoopTimer   stats.LoopTimer
	ping               chan interface{}
	cb                 func(*corev1.Node)
}

func NewNodeStatusController(
	cli cloud.CloudClient,
	internalIP string,
	daemonEndpointPort int32,
	kubeletCapacity corev1.ResourceList,
) *NodeStatusController {
	return &NodeStatusController{
		nodeReady:          false,
		networkUnavailable: true,
		cloudClient:        cli,
		internalIP:         internalIP,
		daemonEndpointPort: daemonEndpointPort,
		kubeletCapacity:    kubeletCapacity,
		ping:               make(chan interface{}),
	}
}

func (n *NodeStatusController) Start(quit <-chan struct{}, wg *sync.WaitGroup) {
	go n.controlLoop(quit, wg)
}

func (n *NodeStatusController) controlLoop(quit <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	ticker := time.NewTicker(59 * time.Second)
	defer ticker.Stop()
	for {
		wasSet := n.setNodeStatus()
		if wasSet {
			break
		}
		time.Sleep(5 * time.Second)
	}
	for {
		select {
		case <-ticker.C:
			n.controlLoopTimer.StartLoop()
			n.setNodeStatus()
			n.controlLoopTimer.EndLoop()
		case <-n.ping:
			n.ping <- "pong"
		case <-quit:
			klog.V(2).Infof("stopping NodeStatusController")
			return
		}
	}
}

func (n *NodeStatusController) setNodeStatus() bool {
	available, err := n.cloudClient.IsAvailable()
	if err != nil {
		klog.Errorf("checking cloud available status: %v", err)
		return false
	}
	nodeReady := available
	networkUnavailable := !available
	if nodeReady == n.nodeReady && networkUnavailable == n.networkUnavailable {
		klog.V(5).Infof("node status unchanged")
		return false
	}
	node := n.createNode()
	if node == nil || n.cb == nil {
		return false
	}
	klog.V(5).Infof("updating node spec %+v status %+v", node.Spec, node.Status)
	n.cb(node)
	n.nodeReady = nodeReady
	n.networkUnavailable = networkUnavailable
	klog.V(2).Infof("node status changed ready: %v network unavailabe: %v",
		n.nodeReady, n.networkUnavailable)
	return true
}

func (n *NodeStatusController) createNode() *corev1.Node {
	if n.node == nil {
		klog.Warningf("UpdateNode() has not been called")
		return nil
	}
	node := n.node.DeepCopy()
	node.Status = n.GetNodeStatus()
	return node
}

func (n *NodeStatusController) Dump() []byte {
	var nodeStatus corev1.NodeStatus
	if n.node != nil {
		nodeStatus = n.node.Status
	}
	dumpStruct := struct {
		NodeReady          bool
		NetworkUnavailable bool
		LoopTimerCount     int64
		LoopTimerAverage   time.Duration
		LoopTimerLastLoop  time.Duration
		InternalIP         string
		DaemonEndpointPort int32
		KubeletCapacity    corev1.ResourceList
		CIDRs              []string
		NodeStatus         corev1.NodeStatus
	}{
		NodeReady:          n.nodeReady,
		NetworkUnavailable: n.networkUnavailable,
		InternalIP:         n.internalIP,
		DaemonEndpointPort: n.daemonEndpointPort,
		KubeletCapacity:    n.kubeletCapacity,
		CIDRs:              n.cidrs,
		NodeStatus:         nodeStatus,
		LoopTimerCount:     n.controlLoopTimer.Count,
		LoopTimerAverage:   n.controlLoopTimer.Average,
		LoopTimerLastLoop:  n.controlLoopTimer.LastLoop,
	}
	b, err := json.MarshalIndent(dumpStruct, "", "    ")
	if err != nil {
		klog.Errorf("dumping state of NodeStatusController: %v", err)
		return nil
	}
	return b
}

func (n *NodeStatusController) Ping(ctx context.Context) error {
	klog.V(5).Infof("node ping")
	n.ping <- "ping"
	select {
	case <-n.ping:
		break
	case <-ctx.Done():
		return fmt.Errorf("timeout pinging NodeStatusController")
	}
	klog.V(5).Infof("node pong")
	return nil
}

func (n *NodeStatusController) NotifyNodeStatus(ctx context.Context, cb func(*corev1.Node)) {
	n.cb = cb
	klog.V(5).Infof("registered node status callback")
}

func (n *NodeStatusController) UpdateNode(node *corev1.Node) {
	if len(n.cidrs) < 1 {
		n.cidrs = n.cloudClient.GetVPCCIDRs()
		klog.V(5).Infof("setting pod CIDRs to %v", n.cidrs)
	}
	node.Status = n.GetNodeStatus()
	// Save node metadata and spec.
	n.node = node.DeepCopy()
}

func (n *NodeStatusController) GetNodeStatus() corev1.NodeStatus {
	return corev1.NodeStatus{
		Addresses:       n.nodeAddresses(),
		Capacity:        n.nodeCapacity(),
		Allocatable:     n.nodeCapacity(),
		Conditions:      n.nodeConditions(),
		DaemonEndpoints: n.nodeDaemonEndpoints(),
		NodeInfo: corev1.NodeSystemInfo{
			OperatingSystem: "Linux",
			Architecture:    "amd64",
		},
	}
}

func (n *NodeStatusController) nodeCapacity() corev1.ResourceList {
	return n.kubeletCapacity
}

func (n *NodeStatusController) nodeConditions() []corev1.NodeCondition {
	readyCondition := corev1.ConditionTrue
	if !n.nodeReady {
		readyCondition = corev1.ConditionFalse
	}
	networkUnavailableCondition := corev1.ConditionFalse
	if n.networkUnavailable {
		networkUnavailableCondition = corev1.ConditionTrue
	}
	return []corev1.NodeCondition{
		{
			Type:               "Ready",
			Status:             readyCondition,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "KubeletReady",
			Message:            "kubelet is ready",
		},
		{
			Type:               "NetworkUnavailable",
			Status:             networkUnavailableCondition,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "RouteCreated",
			Message:            "RouteController created a route",
		},
		{
			Type:               "OutOfDisk",
			Status:             corev1.ConditionFalse,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "KubeletHasSufficientDisk",
			Message:            "kubelet has sufficient disk space available",
		},
		{
			Type:               "MemoryPressure",
			Status:             corev1.ConditionFalse,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "KubeletHasSufficientMemory",
			Message:            "kubelet has sufficient memory available",
		},
		{
			Type:               "DiskPressure",
			Status:             corev1.ConditionFalse,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "KubeletHasNoDiskPressure",
			Message:            "kubelet has no disk pressure",
		},
	}
}

func (n *NodeStatusController) nodeAddresses() []corev1.NodeAddress {
	return []corev1.NodeAddress{
		{
			Type:    "InternalIP",
			Address: n.internalIP,
		},
	}
}

func (n *NodeStatusController) nodeDaemonEndpoints() corev1.NodeDaemonEndpoints {
	return corev1.NodeDaemonEndpoints{
		KubeletEndpoint: corev1.DaemonEndpoint{
			Port: n.daemonEndpointPort,
		},
	}
}
