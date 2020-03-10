/*
Copyright 2020 Elotl Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

// Used externally in provider.yml to specify a buffered node
type StandbyNodeSpec struct {
	InstanceType string `json:"instanceType"`
	Count        int    `json:"count"`
	Spot         bool   `json:"spot"`
	// for now, standby nodes don't get public IPs and can't have GPUs
}
