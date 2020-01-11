package instanceselector

import (
	"fmt"
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestSetupInstanceSelector(t *testing.T) {
	defaultInstanceType := "t2.nano"
	err := Setup("aws", "us-east-1", defaultInstanceType)
	assert.NoError(t, err)
}

func TestHappy(t *testing.T) {
	defaultInstanceType := "t2.nano"
	_ = Setup("aws", "us-east-1", defaultInstanceType)
	ps := api.PodSpec{}
	ps.Resources.CPU = "1"
	ps.Resources.Memory = "1Gi"
	ps.Resources.DedicatedCPU = true
	inst, sustainedCPU, err := ResourcesToInstanceType(&ps)
	assert.NoError(t, err)
	assert.Equal(t, "c5.large", inst)
	assert.False(t, sustainedCPU)
	ps.Resources = api.ResourceSpec{}
	inst, sustainedCPU, err = ResourcesToInstanceType(&ps)
	assert.NoError(t, err)
	assert.Equal(t, inst, defaultInstanceType)
	assert.False(t, sustainedCPU)
}

func TestAWSGPUInstance(t *testing.T) {
	defaultInstanceType := "t2.nano"
	_ = Setup("aws", "us-east-1", defaultInstanceType)
	ps := api.PodSpec{}
	ps.Resources.GPU = "1"
	inst, _, err := ResourcesToInstanceType(&ps)
	assert.NoError(t, err)
	fmt.Println(inst)
	assert.Equal(t, "p2.xlarge", inst)
}

func TestHasInstanceType(t *testing.T) {
	_ = Setup("aws", "us-east-1", "t2.nano")
	ps := api.PodSpec{}
	specType := "m4.xlarge"
	ps.InstanceType = specType
	inst, sustainedCPU, err := ResourcesToInstanceType(&ps)
	assert.Nil(t, err)
	assert.Equal(t, specType, inst)
	assert.False(t, sustainedCPU)
	specType = "t2.small"
	ps.InstanceType = specType
	wantSustainedCPU := true
	ps.Resources.SustainedCPU = &wantSustainedCPU
	inst, sustainedCPU, err = ResourcesToInstanceType(&ps)
	assert.Nil(t, err)
	assert.Equal(t, specType, inst)
	assert.True(t, sustainedCPU)
}

func TestIsUnsupportedInstance(t *testing.T) {
	_ = Setup("aws", "us-east-1", "t2.nano")
	selector.unsupportedInstances.Insert("ZZ")
	v := IsUnsupportedInstance("ZZ.top")
	assert.True(t, v)
}

func TestNoMatch(t *testing.T) {
	_ = Setup("aws", "us-east-1", "t2.nano")
	ps := api.PodSpec{}
	ps.Resources.CPU = "1000"
	ps.Resources.Memory = "1"
	_, _, err := ResourcesToInstanceType(&ps)
	assert.NotNil(t, err)
	ps.Resources.CPU = "1"
	ps.Resources.Memory = "100000Gi"
	_, _, err = ResourcesToInstanceType(&ps)
	assert.NotNil(t, err)
}

func TestAWSResourcesToInstanceType(t *testing.T) {
	_ = Setup("aws", "us-east-1", "t2.nano")
	f := false
	testCases := []struct {
		Resources    api.ResourceSpec
		instanceType string
		sustainedCPU bool
	}{
		// {
		// 	Resources: api.ResourceSpec{},
		// 	instanceType: "",
		// 	sustainedCPU: ,
		// },
		{
			Resources:    api.ResourceSpec{Memory: "0.5Gi", CPU: "0.5"},
			instanceType: "t3.nano",
			sustainedCPU: true,
		},
		{
			Resources:    api.ResourceSpec{Memory: "0.5Gi", CPU: "1.0"},
			instanceType: "t3.nano",
			sustainedCPU: true,
		},
		{
			Resources:    api.ResourceSpec{Memory: "2.0Gi", CPU: "1.0"},
			instanceType: "t3.small",
			sustainedCPU: true,
		},
		{
			Resources:    api.ResourceSpec{Memory: "4.0Gi", CPU: "1.0"},
			instanceType: "t3.medium",
			sustainedCPU: true,
		},
		{
			Resources:    api.ResourceSpec{Memory: "1.5Gi", CPU: "1.5"},
			instanceType: "t3.small",
			sustainedCPU: true,
		},
		{
			Resources:    api.ResourceSpec{Memory: "4.0Gi", CPU: "1.0", GPU: "1"},
			instanceType: "p2.xlarge",
			sustainedCPU: false,
		},
		{
			Resources:    api.ResourceSpec{Memory: "180.0Gi", CPU: "48.0"},
			instanceType: "m5.12xlarge",
			sustainedCPU: false,
		},
		{
			Resources:    api.ResourceSpec{Memory: "15.0Gi", CPU: "32.0"},
			instanceType: "c5.9xlarge",
			sustainedCPU: false,
		},
		{
			Resources:    api.ResourceSpec{Memory: "1Gi", CPU: "1.0", SustainedCPU: &f},
			instanceType: "c5.large",
			sustainedCPU: false,
		},
	}
	for _, tc := range testCases {
		it, sus := selector.getInstanceFromResources(tc.Resources)
		assert.Equal(t, tc.instanceType, it)
		assert.Equal(t, tc.sustainedCPU, sus)
	}
}

func TestAzureResourcesToInstanceType(t *testing.T) {
	_ = Setup("azure", "East US", "Standard_B1s")
	testCases := []struct {
		Resources    api.ResourceSpec
		instanceType string
	}{
		{
			Resources:    api.ResourceSpec{Memory: "3Gi", CPU: "1.0"},
			instanceType: "Standard_DS1_v2",
		},
		{
			Resources:    api.ResourceSpec{Memory: "1Gi", CPU: "0.2"},
			instanceType: "Standard_B1ms",
		},
	}

	for _, tc := range testCases {
		it, sus := selector.getInstanceFromResources(tc.Resources)
		assert.Equal(t, tc.instanceType, it)
		assert.Equal(t, false, sus)
	}
}

func TestNoSetup(t *testing.T) {
	selector = nil
	ps := api.PodSpec{}
	_, _, err := ResourcesToInstanceType(&ps)
	assert.NotNil(t, err)
}
