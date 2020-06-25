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
	"fmt"
	"regexp"
	"testing"

	"github.com/elotl/kip/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestSetupInstanceSelector(t *testing.T) {
	defaultInstanceType := "t2.nano"
	err := Setup("aws", "us-east-1", "", defaultInstanceType)
	assert.NoError(t, err)
}

func TestHappy(t *testing.T) {
	defaultInstanceType := "t2.nano"
	_ = Setup("aws", "us-east-1", "", defaultInstanceType)
	ps := api.PodSpec{}
	ps.Resources.CPU = "1"
	ps.Resources.Memory = "1Gi"
	ps.Resources.DedicatedCPU = true
	inst, sustainedCPU, err := ResourcesToInstanceType(&ps)
	assert.NoError(t, err)
	assert.Equal(t, "c5.large", inst)
	assert.False(t, *sustainedCPU)
	ps.Resources = api.ResourceSpec{}
	inst, sustainedCPU, err = ResourcesToInstanceType(&ps)
	assert.NoError(t, err)
	assert.Equal(t, inst, defaultInstanceType)
	assert.Nil(t, sustainedCPU)
}

func TestAWSGPUInstance(t *testing.T) {
	defaultInstanceType := "t2.nano"
	_ = Setup("aws", "us-east-1", "", defaultInstanceType)
	ps := api.PodSpec{}
	ps.Resources.GPU = "1"
	inst, _, err := ResourcesToInstanceType(&ps)
	assert.NoError(t, err)
	fmt.Println(inst)
	assert.Equal(t, "p2.xlarge", inst)
}

func TestGCEDefaultGPUInstance(t *testing.T) {
	err := Setup("gce", "us-west-1", "us-west1-a", "f1-micro")
	assert.NoError(t, err)
	ps := api.PodSpec{}
	ps.Resources.GPU = "1"
	inst, _, err := ResourcesToInstanceType(&ps)
	assert.NoError(t, err)
	assert.Equal(t, "n1-standard-1", inst)
}

func TestGCESpecificGPUInstance(t *testing.T) {
	err := Setup("gce", "us-west-1", "us-west1-a", "f1-micro")
	assert.NoError(t, err)
	ps := api.PodSpec{}
	ps.Resources.GPU = "1 nvidia-tesla-p100"
	inst, _, err := ResourcesToInstanceType(&ps)
	assert.NoError(t, err)
	assert.Equal(t, "n1-standard-1", inst)
}

func TestHasInstanceType(t *testing.T) {
	_ = Setup("aws", "us-east-1", "", "t2.nano")
	ps := api.PodSpec{}
	specType := "m4.xlarge"
	ps.InstanceType = specType
	inst, sustainedCPU, err := ResourcesToInstanceType(&ps)
	assert.Nil(t, err)
	assert.Equal(t, specType, inst)
	assert.Nil(t, sustainedCPU)
	specType = "t2.small"
	ps.InstanceType = specType
	wantSustainedCPU := true
	ps.Resources.SustainedCPU = &wantSustainedCPU
	inst, sustainedCPU, err = ResourcesToInstanceType(&ps)
	assert.Nil(t, err)
	assert.Equal(t, specType, inst)
	if sustainedCPU == nil {
		t.Error("sustainedCPU should be true")
	} else {
		assert.True(t, *sustainedCPU)
	}
}

func TestIsUnsupportedInstance(t *testing.T) {
	_ = Setup("aws", "us-east-1", "", "t2.nano")
	selector.unsupportedInstances.Insert("ZZ")
	v := IsUnsupportedInstance("ZZ.top")
	assert.True(t, v)
}

func TestNoMatch(t *testing.T) {
	_ = Setup("aws", "us-east-1", "", "t2.nano")
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

type instanceTypeSpec struct {
	Resources        api.ResourceSpec
	instanceTypeGlob string
	instanceType     string
	sustainedCPU     bool
}

func runInstanceTypeTests(t *testing.T, testCases []instanceTypeSpec) {
	for _, tc := range testCases {
		var (
			re  *regexp.Regexp
			err error
		)
		msg := fmt.Sprintf("Test instanceSpec: %#v, glob: %s", tc.Resources, tc.instanceTypeGlob)
		if tc.instanceTypeGlob != "" {
			re, err = globToRegexp(tc.instanceTypeGlob)
			fmt.Println(tc.instanceTypeGlob)
			assert.NoError(t, err, msg)
		}
		it, sus := selector.getInstanceFromResources(tc.Resources, re)
		assert.Equal(t, tc.instanceType, it, msg)
		assert.Equal(t, tc.sustainedCPU, sus, msg)
	}
}

func TestAWSResourcesToInstanceType(t *testing.T) {
	_ = Setup("aws", "us-east-1", "", "t2.nano")
	f := false
	testCases := []instanceTypeSpec{
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
		{
			Resources:        api.ResourceSpec{Memory: "0.5Gi", CPU: "0.5"},
			instanceTypeGlob: "c5*",
			instanceType:     "c5.large",
			sustainedCPU:     false,
		},
		{
			Resources:        api.ResourceSpec{Memory: "15Gi", CPU: "32.0"},
			instanceTypeGlob: "m5.*",
			instanceType:     "m5.12xlarge",
			sustainedCPU:     false,
		},
	}
	runInstanceTypeTests(t, testCases)
}

func TestGCEResourcesToInstanceType(t *testing.T) {
	err := Setup("gce", "us-west-1", "us-west1-a", "f1-micro")
	assert.NoError(t, err)
	f := false
	testCases := []instanceTypeSpec{
		{
			Resources:    api.ResourceSpec{Memory: "0.5Gi", CPU: "0.5"},
			instanceType: "g1-small",
			sustainedCPU: false,
		},
		{
			Resources:    api.ResourceSpec{Memory: "0.5Gi", CPU: "1.0"},
			instanceType: "e2-micro",
			sustainedCPU: false,
		},
		{
			Resources:    api.ResourceSpec{Memory: "2.0Gi", CPU: "1.0"},
			instanceType: "n1-standard-1",
			sustainedCPU: false,
		},
		{
			Resources:    api.ResourceSpec{Memory: "4.0Gi", CPU: "1.0"},
			instanceType: "e2-medium",
			sustainedCPU: false,
		},
		{
			Resources:    api.ResourceSpec{Memory: "1.5Gi", CPU: "1.5"},
			instanceType: "e2-highcpu-2",
			sustainedCPU: false,
		},
		{
			Resources:    api.ResourceSpec{Memory: "4.0Gi", CPU: "1.0", GPU: "1"},
			instanceType: "n1-standard-2",
			sustainedCPU: false,
		},
		{
			Resources:    api.ResourceSpec{Memory: "180.0Gi", CPU: "48.0"},
			instanceType: "n2-standard-48",
			sustainedCPU: false,
		},
		{
			Resources:    api.ResourceSpec{Memory: "15.0Gi", CPU: "32.0"},
			instanceType: "n1-highcpu-32",
			sustainedCPU: false,
		},
		{
			Resources:    api.ResourceSpec{Memory: "1Gi", CPU: "1.0", SustainedCPU: &f},
			instanceType: "e2-micro",
			sustainedCPU: false,
		},
	}
	runInstanceTypeTests(t, testCases)
}

func TestAzureResourcesToInstanceType(t *testing.T) {
	_ = Setup("azure", "East US", "", "Standard_B1s")
	testCases := []instanceTypeSpec{
		{
			Resources:    api.ResourceSpec{Memory: "3Gi", CPU: "1.0"},
			instanceType: "Standard_DS1_v2",
		},
		{
			Resources:    api.ResourceSpec{Memory: "1Gi", CPU: "0.2"},
			instanceType: "Standard_B1ms",
		},
	}
	runInstanceTypeTests(t, testCases)
}

func TestNoSetup(t *testing.T) {
	selector = nil
	ps := api.PodSpec{}
	_, _, err := ResourcesToInstanceType(&ps)
	assert.NotNil(t, err)
}

func TestGlobToRegexp(t *testing.T) {
	tests := []struct {
		glob      string
		regexpstr string
		iserr     bool
	}{
		{
			glob:      "",
			regexpstr: "",
			iserr:     false,
		},
		{
			glob:      "foo",
			regexpstr: "^foo$",
			iserr:     false,
		},
		{
			glob:      "n1-*",
			regexpstr: "^n1-.*$",
			iserr:     false,
		},
		{
			glob:      "c5.*",
			regexpstr: "^c5\\..*$",
			iserr:     false,
		},
	}
	for _, tc := range tests {
		re, err := globToRegexp(tc.glob)
		if tc.iserr {
			assert.Error(t, err)
			continue
		}
		assert.NoError(t, err)
		if tc.regexpstr == "" {
			assert.Nil(t, re)
		} else {
			assert.Equal(t, tc.regexpstr, re.String())
		}
	}
}
