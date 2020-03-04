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
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
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
				ImageOwnerID:    "1000",
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
				// No boot image tags specified.
				cf.Cells.BootImageTags = cloud.BootImageTags{}
			},
			errors: 0,
		},
	}
	for i, test := range tests {
		cf := serverConfigFileWithDefaults()
		cf.Cells.BootImageTags = cloud.BootImageTags{
			Company: "elotl-test",
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
