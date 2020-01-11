package cloud

import (
	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/api/resource"
)

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

func makeIngressRulesMap(ports []api.ServicePort, sourceRanges []string) map[IngressRule]struct{} {
	rules := make(map[IngressRule]struct{})
	for _, port := range ports {
		for _, source := range sourceRanges {
			rules[NewIngressRule(port, source)] = struct{}{}
		}
	}
	return rules
}

func MakeIngressRules(ports []api.ServicePort, sourceRanges []string) []IngressRule {
	rules := make([]IngressRule, 0, len(ports)*len(sourceRanges))
	for _, port := range ports {
		for _, source := range sourceRanges {
			rules = append(rules, NewIngressRule(port, source))
		}
	}
	return rules
}

func NewIngressRule(port api.ServicePort, source string) IngressRule {
	return IngressRule{
		Port:          port.Port,
		PortRangeSize: port.PortRangeSize,
		Protocol:      port.Protocol,
		Source:        source,
	}
}

func MergeSecurityGroups(cloudSG SecurityGroup, specPorts []api.ServicePort, specSourceRanges []string) ([]IngressRule, []IngressRule) {
	// Explode the cross of ports and sources into IngressRules
	// Do a diff of those and use that for updating rules
	status := makeIngressRulesMap(cloudSG.Ports, cloudSG.SourceRanges)
	spec := makeIngressRulesMap(specPorts, specSourceRanges)
	add, delete := ingressRuleDiff(spec, status)
	return add, delete
}

func ToSaneVolumeSize(volSizeSpec string) int32 {
	size, _ := resource.ParseQuantity(volSizeSpec)
	volSizeGiB := util.ToGiBRoundUp(&size)
	if volSizeGiB == 0 {
		// This should never happen but handle it with grace. It would
		// be nice to set volSizeGiB to the default volume size but
		// we're not carrying that values around anywhere.  We could
		// somehow make that value a global var but it seemed like I
		// would start abusing that out of lazyness.
		glog.Errorln("Empty volume size found in resource spec, setting to reasonable value")
		volSizeGiB = 8
	}
	return volSizeGiB
}
