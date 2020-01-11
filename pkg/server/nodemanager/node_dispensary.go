// not sure if this is a good pattern for decoupling the
// pod_controller from the node controller...  Going to give it a try.
package nodemanager

import (
	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/golang/glog"
)

type NodeReply struct {
	Node *api.Node
	// When there's no binding for a pod, that either means the
	// pod is new or something might have gone wrong with the pod
	// spec, possibly it was created by a replicaSet and we can't
	// satisfy the placement spec of the pod.  We use NoBinding
	// to signal that we can't currently create a node for the pod.
	// if a pod remains unbound for too long, we can act accordingly
	// (e.g. for a replicaSet pod, we kill the pod).
	NoBinding bool
}

type NodeRequest struct {
	requestingPod api.Pod
	ReplyChan     chan NodeReply
}

type NodeReturn struct {
	NodeName string
	Unused   bool
}

type NodeDispenser struct {
	NodeRequestChan chan NodeRequest
	NodeReturnChan  chan NodeReturn
}

func NewNodeDispenser() *NodeDispenser {
	return &NodeDispenser{
		NodeRequestChan: make(chan NodeRequest),
		NodeReturnChan:  make(chan NodeReturn, 1),
	}
}

// we pass in a copy of the requesting pod for safety reasons.
func (e *NodeDispenser) RequestNode(requestingPod api.Pod) NodeReply {
	replyChan := make(chan NodeReply)
	if e.NodeRequestChan == nil {
		glog.Errorf("NodeRequestChan is nil!")
		return NodeReply{
			Node: nil,
		}
	}
	e.NodeRequestChan <- NodeRequest{requestingPod, replyChan}
	nodeReply := <-replyChan
	return nodeReply
}

func (e *NodeDispenser) ReturnNode(nodeName string, unused bool) {
	if nodeName == "" {
		glog.Warningf("Got empty node name in ReturnNode")
		return
	}
	e.NodeReturnChan <- NodeReturn{
		NodeName: nodeName,
		Unused:   unused,
	}
}
