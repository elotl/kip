package instanceselector

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/elotl/cloud-instance-provider/pkg/util/sets"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog"
)

const t2UnlimitedPrice float32 = 0.05

// have the data in there
type InstanceData struct {
	InstanceType string  `json:"instanceType"`
	Price        float32 `json:"price"`
	GPU          int     `json:"gpu"`
	Memory       float32 `json:"memory"`
	CPU          float32 `json:"cpu"`
	Burstable    bool    `json:"burstable"`
	Baseline     float32 `json:"baseline"`
}

type instanceSelector struct {
	defaultInstanceType  string
	data                 []InstanceData
	unsupportedInstances sets.String
	sustainedCPUSupport  bool
	// AWS uses GiB while Google uses GB for memory.  We'll make the
	// memory spec parser vary for each cloud I imagine we'll
	// eventually need to make the GPU spec vary as well
	memorySpecParser          func(resource.Quantity) float32
	containerInstanceSelector func(*api.ResourceSpec) (int64, int64, error)
}

var selector *instanceSelector

func Setup(cloud, region, defaultInstanceType string) error {
	switch cloud {
	case "aws":
		d, err := getSelectorData(awsInstanceJson, region)
		if err != nil {
			return err
		}
		selector = &instanceSelector{
			defaultInstanceType:  defaultInstanceType,
			data:                 d,
			unsupportedInstances: sets.NewString([]string{
			// "c5",
			// "i3",
			// "m5",
			// "m4.16xlarge",
			}...),
			sustainedCPUSupport: true,
			memorySpecParser: func(q resource.Quantity) float32 {
				return util.ToGiBFloat32(&q)
			},
			containerInstanceSelector: FargateInstanceSelector,
		}
	case "azure":
		d, err := getSelectorData(azureInstanceJson, region)
		if err != nil {
			return err
		}
		selector = &instanceSelector{
			defaultInstanceType:  defaultInstanceType,
			data:                 d,
			unsupportedInstances: sets.NewString(),
			sustainedCPUSupport:  false,
			memorySpecParser: func(q resource.Quantity) float32 {
				return util.ToGiBFloat32(&q)
			},
			containerInstanceSelector: AzureContainenrInstanceSelector,
		}

	default:
		return fmt.Errorf("unknown cloud for instanceselector setup: %s", cloud)
	}
	return nil
}

func getSelectorData(data, region string) ([]InstanceData, error) {
	d := make(map[string][]InstanceData)
	err := json.Unmarshal([]byte(data), &d)
	if err != nil {
		return nil, err
	}
	regionData, exists := d[region]
	if !exists {
		return nil, fmt.Errorf("could not find instance data for cloud region: %s", region)
	}
	return regionData, nil
}

func (instSel *instanceSelector) parseMemorySpec(memSpec string) (float32, error) {
	if memSpec == "" {
		return 0.0, nil
	}
	memQuantity, err := resource.ParseQuantity(memSpec)
	if err != nil {
		return 0.0, err
	}
	return instSel.memorySpecParser(memQuantity), nil
}

func parseGPUSpec(gpuSpec string) (int, error) {
	if gpuSpec == "" {
		return 0, nil
	}
	return strconv.Atoi(gpuSpec)
}

func parseCPUSpec(cpuSpec string) (float32, error) {
	if cpuSpec == "" {
		return 0.0, nil
	}
	cpuQuantity, err := resource.ParseQuantity(cpuSpec)
	if err != nil {
		return 0.0, err
	}
	return util.CPUCoresFraction(&cpuQuantity), nil
}

func (instSel *instanceSelector) priceForCPUSpec(cpu float32, inst InstanceData) (float32, bool) {
	if !inst.Burstable || !instSel.sustainedCPUSupport {
		return inst.Price, false
	} else if cpu <= inst.Baseline {
		return inst.Price, false
	} else {
		cpuNeeded := cpu - inst.Baseline
		extraCPUCost := cpuNeeded * t2UnlimitedPrice
		cost := inst.Price + extraCPUCost
		return cost, true
	}
}

func filterInstanceData(instances []InstanceData, predicate func(i InstanceData) bool) []InstanceData {
	filtered := make([]InstanceData, 0, len(instances))
	for _, inst := range instances {
		if predicate(inst) {
			filtered = append(filtered, inst)
		}
	}
	return filtered
}

func findCheapestInstance(matches []InstanceData) string {
	lowestPrice := float32(math.MaxFloat32)
	cheapestInstance := ""
	for _, inst := range matches {
		if inst.Price > 0.0 && inst.Price < lowestPrice {
			lowestPrice = inst.Price
			cheapestInstance = inst.InstanceType
		}
	}
	return cheapestInstance
}

//The instance selector tries to find the minimum cost instance that
// satisfies all constraints in the resource spec.  This gets a bit
// tricky to figure out the easiest way to satisfy constraints with
// the t2.Unlimited option from AWS. For T2 instances, we try to
// figure out what percentage of a CPU a user will likely use and
// use that to compute t2.Unlimited cost.
func (instSel *instanceSelector) getInstanceFromResources(rs api.ResourceSpec) (string, bool) {
	memoryRequirement, err := instSel.parseMemorySpec(rs.Memory)
	if err != nil {
		klog.Errorf("Error parsing memory spec: %s", err)
	}
	cpuRequirements, err := parseCPUSpec(rs.CPU)
	if err != nil {
		klog.Errorf("Error parsing CPU spec: %s", err)
	}
	gpuRequirements, err := parseGPUSpec(rs.GPU)
	if err != nil {
		klog.Errorf("Error parsing GPU spec: %s", err)
	}

	matches := filterInstanceData(instSel.data, func(inst InstanceData) bool {
		return !IsUnsupportedInstance(inst.InstanceType)
	})

	// Memory
	matches = filterInstanceData(matches, func(inst InstanceData) bool {
		return memoryRequirement == 0.0 || inst.Memory >= memoryRequirement
	})

	// GPU
	matches = filterInstanceData(matches, func(inst InstanceData) bool {
		return inst.GPU >= gpuRequirements
	})

	// CPU
	cheapestInstance := ""
	cheapestIsSustained := false
	if rs.DedicatedCPU {
		matches = filterInstanceData(matches, func(inst InstanceData) bool {
			return !inst.Burstable
		})
		cheapestInstance = findCheapestInstance(matches)
	} else if (rs.SustainedCPU != nil && *rs.SustainedCPU == false) ||
		!instSel.sustainedCPUSupport {
		// In this case, we don't have to worry about T2.unlimited so
		// we just match the CPU requirements
		matches = filterInstanceData(matches, func(inst InstanceData) bool {
			if inst.Burstable {
				return inst.Baseline >= cpuRequirements
			} else {
				return inst.CPU >= cpuRequirements
			}
		})
		cheapestInstance = findCheapestInstance(matches)
	} else {
		// Here we do work to find the cheapest instance while taking
		// T2.unlimited into account.  We duplicate
		// findCheapestInstance because we need to compute the
		// priceForCpu and know whether that includes sustainedCPU.
		lowestPrice := float32(math.MaxFloat32)
		for _, inst := range matches {
			if inst.CPU < cpuRequirements {
				continue
			}
			price, sustainedCPU := instSel.priceForCPUSpec(cpuRequirements, inst)
			if price > 0.0 && price < lowestPrice {
				lowestPrice = price
				cheapestInstance = inst.InstanceType
				cheapestIsSustained = sustainedCPU
			}
		}
	}
	return cheapestInstance, cheapestIsSustained
}

func noResourceSpecified(ps *api.PodSpec) bool {
	return ps.InstanceType == "" &&
		ps.Resources.CPU == "" &&
		ps.Resources.Memory == "" &&
		ps.Resources.GPU == ""
}

// Used by validation code in Milpa
func IsUnsupportedInstance(instanceType string) bool {
	if len(instanceType) < 2 {
		return true
	}
	prefix := instanceType[:2]
	// We have prefixes and instance names in the sets so test both
	return selector.unsupportedInstances.Has(prefix) ||
		selector.unsupportedInstances.Has(instanceType)
}

func ResourcesToInstanceType(ps *api.PodSpec) (string, *bool, error) {
	if ps.Resources.ContainerInstance != nil && *ps.Resources.ContainerInstance {
		return api.ContainerInstanceType, nil, nil
	}
	if ps.InstanceType != "" {
		var sustainedCPU *bool
		if ps.Resources.SustainedCPU != nil {
			sustainedCPU = ps.Resources.SustainedCPU
		}
		return ps.InstanceType, sustainedCPU, nil
	}
	if selector == nil {
		msg := "fatal: instanceselector has not been initialized"
		klog.Errorf(msg)
		return "", nil, fmt.Errorf(msg)
	}
	if noResourceSpecified(ps) {
		return selector.defaultInstanceType, nil, nil
	}

	instanceType, needsSustainedCPU := selector.getInstanceFromResources(ps.Resources)
	if instanceType == "" {
		msg := "could not compute instance type from Spec.Resources. It's likely that the Pod.Spec.Resources specify an instance that doesnt exist in the cloud"
		return "", nil, fmt.Errorf(msg)
	}
	return instanceType, &needsSustainedCPU, nil
}

func ResourcesToContainerInstance(rs *api.ResourceSpec) (int64, int64, error) {
	return selector.containerInstanceSelector(rs)
}
