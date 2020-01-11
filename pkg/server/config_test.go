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
				cf.Nodes.BootImageTags = cloud.BootImageTags{}
			},
			errors: 0,
		},
		// Now that we have added marketplace licensing, these
		// are no longer errors
		{
			mod: func(cf *ServerConfigFile) {
				cf.License.Username = ""
				cf.License.Password = ""
				cf.License.Id = ""
				cf.License.Key = ""
			},
			errors: 0,
		},
		{
			mod: func(cf *ServerConfigFile) {
				cf.License.Username = blankTemplateValue
				cf.License.Password = blankTemplateValue
				cf.License.Id = blankTemplateValue
				cf.License.Key = blankTemplateValue
			},
			errors: 0,
		},
	}
	for i, test := range tests {
		cf := serverConfigFileWithDefaults()
		cf.Nodes.BootImageTags = cloud.BootImageTags{
			Company: "elotl-test",
		}
		cf.Nodes.DefaultInstanceType = "t2.nano"
		cf.License.Username = "tester"
		cf.License.Password = "tester_password"
		cf.License.Id = "tester_id"
		cf.License.Key = "some long key"
		test.mod(cf)
		errs := validateServerConfigFile(cf)
		if len(errs) != test.errors {
			t.Errorf("Expected %d errors from test %d, got %d: %v",
				test.errors, i+1, len(errs), errs)
		}
	}
}
