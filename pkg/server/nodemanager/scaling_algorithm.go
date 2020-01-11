package nodemanager

import (
	"github.com/elotl/cloud-instance-provider/pkg/api"
)

type ScalingAlgorithm interface {
	// todo, figure out what we really need to pass in
	// and return value will likely get much more complex
	//Compute(nodes []*api.Node, pods []*api.Pod) ([]*api.Node, int)
	Compute(nodes []*api.Node, pods []*api.Pod) ([]*api.Node, []*api.Node, map[string]string)
}

// Used externally in server.yml to specify a buffered node
type StandbyNodeSpec struct {
	InstanceType string `json:"instanceType"`
	Count        int    `json:"count"`
	Spot         bool   `json:"spot"`
	// for now, standby nodes don't get public IPs and can't have GPUs
}
