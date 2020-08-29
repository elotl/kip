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

package cloud

import (
	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/util"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog"
)

// Service port definition. This is a TCP or UDP port that a Service uses.
type InstancePort struct {
	// Name of the Service port.
	Name string `json:"name"`
	// Protocol. Can be "TCP", "UDP" or "ICMP".
	Protocol api.Protocol `json:"protocol"`
	// Port number. Not used for "ICMP".
	Port int `json:"port"`
	// portRangeSize is the contiguous ports number that are exposed
	// by this service. Given port = 80 and portRangeSize = 100, the
	// InstancePort will represent a range of ports from 80-179 (100
	// ports in total). In this case, port means the starting port of
	// a range.
	PortRangeSize int `json:"portRangeSize,omitempty"`
}

//Allow ports to be sorted
type SortableSliceOfPorts []InstancePort

func (p SortableSliceOfPorts) Len() int           { return len(p) }
func (p SortableSliceOfPorts) Less(i, j int) bool { return lessPorts(p[i], p[j]) }
func (p SortableSliceOfPorts) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func lessPorts(lhs, rhs InstancePort) bool {
	if lhs.Port != rhs.Port {
		return lhs.Port < rhs.Port
	} else if lhs.Protocol < rhs.Protocol {
		return lhs.Protocol < rhs.Protocol
	} else if lhs.PortRangeSize != rhs.PortRangeSize {
		return lhs.PortRangeSize < rhs.PortRangeSize
	} else {
		return lhs.Name < rhs.Name
	}
}

// Diffing rules got a little nasty... We take the cross of the
// service ports and sourceRanges and each value becomes an
// IngressRule
type IngressRule struct {
	Port          int
	PortRangeSize int
	Protocol      api.Protocol
	Source        string
}

// We do simple set differences on the keys here
func ingressRuleDiff(spec, status map[IngressRule]struct{}) ([]IngressRule, []IngressRule) {
	// delete is anything in status that isn't in the spec
	delete := make([]IngressRule, 0)
	for k := range status {
		_, exists := spec[k]
		if !exists {
			delete = append(delete, k)
		}
	}
	// add is anything in the spec that isn't in the status
	add := make([]IngressRule, 0)
	for k := range spec {
		_, exists := status[k]
		if !exists {
			add = append(add, k)
		}
	}
	return add, delete
}

func makeIngressRulesMap(ports []InstancePort, sourceRanges []string) map[IngressRule]struct{} {
	rules := make(map[IngressRule]struct{})
	for _, port := range ports {
		for _, source := range sourceRanges {
			rules[NewIngressRule(port, source)] = struct{}{}
		}
	}
	return rules
}

func MakeIngressRules(ports []InstancePort, sourceRanges []string) []IngressRule {
	rules := make([]IngressRule, 0, len(ports)*len(sourceRanges))
	for _, port := range ports {
		for _, source := range sourceRanges {
			rules = append(rules, NewIngressRule(port, source))
		}
	}
	return rules
}

func NewIngressRule(port InstancePort, source string) IngressRule {
	return IngressRule{
		Port:          port.Port,
		PortRangeSize: port.PortRangeSize,
		Protocol:      port.Protocol,
		Source:        source,
	}
}

func MergeSecurityGroups(cloudSG SecurityGroup, specPorts []InstancePort, specSourceRanges []string) ([]IngressRule, []IngressRule) {
	// Explode the cross of ports and sources into IngressRules
	// Do a diff of those and use that for updating rules
	status := makeIngressRulesMap(cloudSG.Ports, cloudSG.SourceRanges)
	spec := makeIngressRulesMap(specPorts, specSourceRanges)
	add, delete := ingressRuleDiff(spec, status)
	return add, delete
}

func ToSaneVolumeSize(volSizeSpec string, image Image) int32 {
	size, _ := resource.ParseQuantity(volSizeSpec)
	volSizeGiB := util.ToGiBRoundUp(&size)
	if volSizeGiB == 0 {
		// This should never happen but handle it with grace. It would
		// be nice to set volSizeGiB to the default volume size but
		// we're not carrying that values around anywhere.  We could
		// somehow make that value a global var but it seemed like I
		// would start abusing that out of lazyness.
		klog.Errorln("Empty volume size found in resource spec, setting to reasonable value")
		volSizeGiB = 8
	}
	if image.VolumeDiskSize != nil {
		if *image.VolumeDiskSize > volSizeGiB {
			return *image.VolumeDiskSize
		}
	}
	return volSizeGiB
}
