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

package instanceselector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/util"
	"github.com/ryanuber/go-glob"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog"
)

const t2UnlimitedPrice float32 = 0.05

// have the data in there
type InstanceData struct {
	InstanceType      string         `json:"instanceType"`
	Price             float32        `json:"price"`
	GPU               int            `json:"gpu"`
	SupportedGPUTypes map[string]int `json:"supportedGPUTypes"`
	Memory            float32        `json:"memory"`
	CPU               float32        `json:"cpu"`
	Burstable         bool           `json:"burstable"`
	Baseline          float32        `json:"baseline"`
}

// CustomInstanceData holds instance type information for custom sized
// instances.
type CustomInstanceData struct {
	InstanceFamily       string         `json:"instanceFamily"`
	BaseMemoryUnit       float32        `json:"baseMemoryUnit"`
	PricePerCPU          float32        `json:"pricePerCPU"`
	PricePerGBOfMemory   float32        `json:"pricePerGBOfMemory"`
	PossibleNumberOfCPUs []float32      `json:"possibleNumberOfCPUs"`
	MinimumMemoryPerCPU  float32        `json:"minimumMemoryPerCPU"`
	MaximumMemoryPerCPU  float32        `json:"maximumMemoryPerCPU"`
	SupportedGPUTypes    map[string]int `json:"supportedGPUTypes"`
}

type instanceSelector struct {
	defaultInstanceType  string
	instanceData         []InstanceData
	customInstanceData   []CustomInstanceData
	unsupportedInstances sets.String
	sustainedCPUSupport  bool
	// AWS uses GiB while Google uses GB for memory.  We'll make the
	// memory spec parser vary for each cloud I imagine we'll
	// eventually need to make the GPU spec vary as well
	memorySpecParser          func(resource.Quantity) float32
	containerInstanceSelector func(*api.ResourceSpec) (int64, int64, error)
}

var selector *instanceSelector

func Setup(cloud, region, zone, defaultInstanceType, instanceDataPath string) error {
	switch cloud {
	case "aws":
		data, err := getSelectorData(awsInstanceJson, region, instanceDataPath)
		if err != nil {
			return err
		}
		selector = &instanceSelector{
			defaultInstanceType: defaultInstanceType,
			instanceData:        data,
			unsupportedInstances: sets.NewString([]string{
				"t1", // TODO: should we support previous generation families?
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
		data, err := getSelectorData(azureInstanceJson, region, instanceDataPath)
		if err != nil {
			return err
		}
		selector = &instanceSelector{
			defaultInstanceType:  defaultInstanceType,
			instanceData:         data,
			unsupportedInstances: sets.NewString(),
			sustainedCPUSupport:  false,
			memorySpecParser: func(q resource.Quantity) float32 {
				return util.ToGiBFloat32(&q)
			},
			containerInstanceSelector: AzureContainenrInstanceSelector,
		}
	case "gce":
		data, err := getSelectorData(gceInstanceJson, zone, instanceDataPath)
		if err != nil {
			return err
		}
		customData, err := getSelectorCustomData(gceCustomInstanceJson, zone)
		if err != nil {
			return err
		}
		selector = &instanceSelector{
			defaultInstanceType:  defaultInstanceType,
			instanceData:         data,
			customInstanceData:   customData,
			unsupportedInstances: sets.NewString(),
			sustainedCPUSupport:  false,
			memorySpecParser: func(q resource.Quantity) float32 {
				return util.ToGiBFloat32(&q)
			},
			containerInstanceSelector: GCEContainenrInstanceSelector,
		}
		klog.Infof("custom instances in %s: %+v", zone, customData)
	default:
		return fmt.Errorf("unknown cloud for instanceselector setup: %s", cloud)
	}
	return nil
}

func getSelectorData(data, regionOrZone, filepath string) ([]InstanceData, error) {
	d := make(map[string][]InstanceData)
	fileData, err := getSelectorDataFromFile(filepath, regionOrZone)
	if err == nil {
		// If loading from file is successful, return data from there,
		// if not, fallback to data from instanceselector package
		return fileData, nil
	} else {
		klog.Warningf("failed to load instance data from path %s: %v , falling back to data baked in binary", filepath, err)
	}
	err = json.Unmarshal([]byte(data), &d)
	if err != nil {
		return nil, err
	}
	regionData, exists := d[regionOrZone]
	if !exists {
		return nil, fmt.Errorf("could not find instance data for cloud region/zone: %s", regionOrZone)
	}
	return regionData, nil
}

func getSelectorDataFromFile(filepath, regionOrZone string) ([]InstanceData, error) {
	data, err := ioutil.ReadFile(filepath)
	d := make(map[string][]InstanceData)
	err = json.Unmarshal(data, &d)
	if err != nil {
		return nil, err
	}
	regionData, exists := d[regionOrZone]
	if !exists {
		return nil, fmt.Errorf("could not load instance data for cloud region/zone: %s from file %s", regionOrZone, filepath)
	}
	return regionData, nil

}

func getSelectorCustomData(data, regionOrZone string) ([]CustomInstanceData, error) {
	d := make(map[string][]CustomInstanceData)
	err := json.Unmarshal([]byte(data), &d)
	if err != nil {
		return nil, err
	}
	regionData, exists := d[regionOrZone]
	if !exists {
		return nil, fmt.Errorf("could not find custom instance data for cloud region/zone: %s", regionOrZone)
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

func parseGPUSpec(gpuSpec string) (int, string, error) {
	if gpuSpec == "" {
		return 0, "", nil
	}
	parts := strings.Fields(gpuSpec)
	count, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", err
	}
	typ := strings.ToLower(strings.Join(parts[1:], " "))
	return count, typ, nil
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

type CustomInstanceParameters struct {
	Price  float32
	CPUs   float32
	Memory float32
}

func cheapestCustomInstanceSizeForCPUAndMemory(cid CustomInstanceData, memoryRequirement, cpuRequirement float32) *CustomInstanceParameters {
	customPrice := float32(math.MaxFloat32)
	customCPUs := float32(math.MaxFloat32)
	customMemory := float32(math.MaxFloat32)
	memCeil := math.Ceil(float64(memoryRequirement / cid.BaseMemoryUnit))
	baseMemSize := cid.BaseMemoryUnit * float32(memCeil)
	for _, cpu := range cid.PossibleNumberOfCPUs {
		if cpu >= cpuRequirement &&
			cpu < customCPUs &&
			baseMemSize <= cid.MaximumMemoryPerCPU*cpu {
			memory := baseMemSize
			if memory < cid.MinimumMemoryPerCPU*cpu {
				memory = cid.MinimumMemoryPerCPU * cpu
			}
			ceil := math.Ceil(float64(memory / cid.BaseMemoryUnit))
			memory = float32(ceil) * cid.BaseMemoryUnit
			price := memory*cid.PricePerGBOfMemory + cpu*cid.PricePerCPU
			if price < customPrice {
				customPrice = price
				customCPUs = cpu
				customMemory = memory
			}
		}
	}
	if customPrice == math.MaxFloat32 {
		return nil
	}
	return &CustomInstanceParameters{
		Price:  customPrice,
		CPUs:   customCPUs,
		Memory: customMemory,
	}
}

func toInstanceData(data []CustomInstanceData, memoryRequirement, cpuRequirement float32) []InstanceData {
	instanceData := make([]InstanceData, 0, len(data))
	for _, cid := range data {
		if cid.BaseMemoryUnit == 0.0 || len(cid.PossibleNumberOfCPUs) < 1 {
			continue
		}
		customParams := cheapestCustomInstanceSizeForCPUAndMemory(cid, memoryRequirement, cpuRequirement)
		if customParams == nil {
			// This instance family doesn't satisfy CPU and/or memory
			// requirements.
			continue
		}
		maxGPUs := 0
		for _, gpu := range cid.SupportedGPUTypes {
			if gpu > maxGPUs {
				maxGPUs = gpu
			}
		}
		// This naming is GCE-specific. If we ever get another cloud provider
		// with custom instances, we'll have to set up a callback similar to
		// memorySpecParser.
		instanceType := fmt.Sprintf(
			"%s-custom-%d-%d", cid.InstanceFamily, int(customParams.CPUs), int(1024*customParams.Memory))
		// On GCE, only non-burstable types are supported for custom instances.
		burstable := false
		baseline := customParams.CPUs
		instanceData = append(instanceData, InstanceData{
			InstanceType:      instanceType,
			Price:             customParams.Price,
			GPU:               maxGPUs,
			SupportedGPUTypes: cid.SupportedGPUTypes,
			Memory:            customParams.Memory,
			CPU:               customParams.CPUs,
			Burstable:         burstable,
			Baseline:          baseline,
		})
	}
	return instanceData
}

// The instance selector tries to find the minimum cost instance that
// satisfies all constraints in the resource spec.  This gets a bit
// tricky to figure out the easiest way to satisfy constraints with
// the t2.Unlimited option from AWS. For T2 instances, we try to
// figure out what percentage of a CPU a user will likely use and
// use that to compute t2.Unlimited cost.
func (instSel *instanceSelector) getInstanceFromResources(rs api.ResourceSpec, additionalFilter func(inst InstanceData) bool) (string, bool) {
	memoryRequirement, err := instSel.parseMemorySpec(rs.Memory)
	if err != nil {
		klog.Errorf("Error parsing memory spec: %s", err)
	}
	cpuRequirements, err := parseCPUSpec(rs.CPU)
	if err != nil {
		klog.Errorf("Error parsing CPU spec: %s", err)
	}
	gpuCountRequirements, gpuTypeRequirements, err := parseGPUSpec(rs.GPU)
	if err != nil {
		klog.Errorf("Error parsing GPU spec: %s", err)
	}

	matches := filterInstanceData(instSel.instanceData, func(inst InstanceData) bool {
		return !IsUnsupportedInstance(inst.InstanceType)
	})

	// Memory
	matches = filterInstanceData(matches, func(inst InstanceData) bool {
		return memoryRequirement == 0.0 || inst.Memory >= memoryRequirement
	})

	matches = append(matches, toInstanceData(instSel.customInstanceData, memoryRequirement, cpuRequirements)...)

	// This allows you to match instance family (e.g. c*), exclude some instances and basically
	// do any additional filtering that is needed
	matches = filterInstanceData(matches, additionalFilter)

	// GPU
	matches = filterInstanceData(matches, func(inst InstanceData) bool {
		if gpuTypeRequirements == "" {
			return inst.GPU >= gpuCountRequirements
		}
		available := inst.SupportedGPUTypes[gpuTypeRequirements]
		return available >= gpuCountRequirements
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
	klog.Infof("chose instance %+v", cheapestInstance)
	return cheapestInstance, cheapestIsSustained
}

func noResourceSpecified(ps *api.PodSpec) bool {
	return ps.InstanceType == "" &&
		ps.Resources.CPU == "" &&
		ps.Resources.Memory == "" &&
		ps.Resources.GPU == ""
}

// Used by validation code in Kip
func IsUnsupportedInstance(instanceType string) bool {
	if len(instanceType) < 2 {
		return true
	}
	prefix := instanceType[:2]
	// We have prefixes and instance names in the sets so test both
	return selector.unsupportedInstances.Has(prefix) ||
		selector.unsupportedInstances.Has(instanceType)
}

func instanceTypeSpecified(instanceType string) bool {
	return instanceType != "" && !strings.ContainsRune(instanceType, '*')
}

func ResourcesToInstanceType(ps *api.PodSpec) (string, *bool, error) {
	if ps.Resources.ContainerInstance != nil && *ps.Resources.ContainerInstance {
		return api.ContainerInstanceType, nil, nil
	}
	if instanceTypeSpecified(ps.InstanceType) {
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
	if ps.InstanceType == "" && noResourceSpecified(ps) {
		return selector.defaultInstanceType, nil, nil
	}

	instanceType, needsSustainedCPU := selector.getInstanceFromResources(ps.Resources, makeInstanceTypeGlobberFunc(ps.InstanceType))
	if instanceType == "" {
		msg := "could not compute instance type from Spec.Resources. It's likely that the Pod.Spec.Resources specify an instance that doesnt exist in the cloud"
		return "", nil, fmt.Errorf(msg)
	}
	return instanceType, &needsSustainedCPU, nil
}

func makeInstanceTypeGlobberFunc(instanceType string) func(inst InstanceData) bool {
	return func(inst InstanceData) bool {
		if instanceType == "" {
			return true
		}
		return glob.Glob(instanceType, inst.InstanceType)
	}
}

func ResourcesToContainerInstance(rs *api.ResourceSpec) (int64, int64, error) {
	return selector.containerInstanceSelector(rs)
}

func GetInstanceFromResources(rs api.ResourceSpec, additionalFilter func(inst InstanceData) bool) (string, bool) {
	return selector.getInstanceFromResources(rs, additionalFilter)
}
