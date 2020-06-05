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

package server

import (
	"fmt"
	"testing"

	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestValidateAws(t *testing.T) {
	// all == FILL_IN
	// others OK but region has wrong regex
	tests := []struct {
		c      AWSConfig
		errors int
	}{
		{
			// We now allow empty configs to support IAM roles and env vars
			c:      AWSConfig{},
			errors: 0,
		},
		{
			c: AWSConfig{
				Region:          "FILL_IN",
				AccessKeyID:     "FILL_IN",
				SecretAccessKey: "FILL_IN",
			},
			errors: 3,
		},
		{
			// if Region is non-empty and doesn't match the expected
			// format of an AWS region, throw an error
			c: AWSConfig{
				Region:          "None",
				AccessKeyID:     "1234",
				SecretAccessKey: "abcd",
			},
			errors: 1,
		},
		{
			c: AWSConfig{
				Region:          "us-east-1",
				AccessKeyID:     "1234",
				SecretAccessKey: "abcd",
				VPCID:           "vpc-12345",
			},
			errors: 0,
		},
	}
	for i, test := range tests {
		errs := validateAWSConfig(&test.c)
		if len(errs) != test.errors {
			t.Errorf("Expected %d errors from test %d, got %d: %v",
				test.errors, i+1, len(errs), errs)
		}
	}
}

func TestConfigValidation(t *testing.T) {
	tests := []struct {
		mod    func(cf *ServerConfigFile)
		errors int
	}{
		{
			mod:    func(cf *ServerConfigFile) {},
			errors: 0,
		},
		{
			mod: func(cf *ServerConfigFile) {
				// No boot image spec specified.
				cf.Cells.BootImageSpec = cloud.BootImageSpec{}
			},
			errors: 0,
		},
	}
	for i, test := range tests {
		cf := serverConfigFileWithDefaults()
		cf.Cells.BootImageSpec = cloud.BootImageSpec{
			"name": "elotl-test",
		}
		cf.Cells.DefaultInstanceType = "t2.nano"
		test.mod(cf)
		errs := validateServerConfigFile(cf)
		if len(errs) != test.errors {
			t.Errorf("Expected %d errors from test %d, got %d: %v",
				test.errors, i+1, len(errs), errs)
		}
	}
}

func mustParseQuantity(str string) *resource.Quantity {
	q := resource.MustParse(str)
	return &q
}

//func updateCapacityFromDeprecatedFields(config *ServerConfigFile)
func TestUpdateCapacityFromDeprecatedFields(t *testing.T) {
	testCases := []struct {
		CPU      *resource.Quantity
		Memory   *resource.Quantity
		Pods     *resource.Quantity
		Capacity v1.ResourceList
		Result   v1.ResourceList
	}{
		{
			CPU:      mustParseQuantity("100m"),
			Memory:   mustParseQuantity("1Gi"),
			Pods:     mustParseQuantity("100"),
			Capacity: v1.ResourceList{},
			Result: v1.ResourceList{
				v1.ResourceCPU:    resource.MustParse("100m"),
				v1.ResourceMemory: resource.MustParse("1Gi"),
				v1.ResourcePods:   resource.MustParse("100"),
			},
		},
		{
			Capacity: v1.ResourceList{},
			Result:   v1.ResourceList{},
		},
		{
			CPU: mustParseQuantity("100m"),
			Capacity: v1.ResourceList{
				v1.ResourceCPU:    resource.MustParse("200m"),
				v1.ResourceMemory: resource.MustParse("1Gi"),
				v1.ResourcePods:   resource.MustParse("100"),
			},
			Result: v1.ResourceList{
				v1.ResourceCPU:    resource.MustParse("100m"),
				v1.ResourceMemory: resource.MustParse("1Gi"),
				v1.ResourcePods:   resource.MustParse("100"),
			},
		},
		{
			CPU:      mustParseQuantity("100m"),
			Capacity: v1.ResourceList{},
			Result: v1.ResourceList{
				v1.ResourceCPU: resource.MustParse("100m"),
			},
		},
		{
			Memory: mustParseQuantity("10Mi"),
			Capacity: v1.ResourceList{
				v1.ResourceMemory: resource.MustParse("1Gi"),
			},
			Result: v1.ResourceList{
				v1.ResourceMemory: resource.MustParse("10Mi"),
			},
		},
		{
			CPU: mustParseQuantity("1.5"),
			Capacity: v1.ResourceList{
				v1.ResourceMemory: resource.MustParse("1Gi"),
			},
			Result: v1.ResourceList{
				v1.ResourceCPU:    resource.MustParse("1.5"),
				v1.ResourceMemory: resource.MustParse("1Gi"),
			},
		},
	}
	for i, tc := range testCases {
		config := ServerConfigFile{
			Kubelet: KubeletConfig{
				CPU:      tc.CPU,
				Memory:   tc.Memory,
				Pods:     tc.Pods,
				Capacity: tc.Capacity,
			},
		}
		updateCapacityFromDeprecatedFields(&config)
		msg := fmt.Sprintf(
			"test case %d failed: input %+v result %+v", i+1, tc, config.Kubelet.Capacity)
		assert.Equal(t, tc.Result, config.Kubelet.Capacity, msg)
	}
}
