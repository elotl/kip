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

package aws

import (
	"flag"
	"fmt"
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud/functional"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/stretchr/testify/assert"
)

// should probably load this from some static config
const (
	testControllerID = "AwsFunctionalTest"
	vpcID            = "vpc-841834e2"
	defaultSubnetID  = "subnet-12a8a13f"
	imageAmi         = "ami-e2dea19d"
	instanceType     = "t2.nano"
)

func getEC2(t *testing.T, controllerID string) *AwsEC2 {
	if !util.AWSEnvVarsSet() {
		t.Fatal("Neet to setup AWS env vars for tests")
	}
	e, err := NewEC2Client(controllerID, controllerID, vpcID, defaultSubnetID, "")
	if err != nil {
		msg := "Error getting EC2 Client: " + err.Error()
		assert.FailNow(t, msg)
	}

	return e
}

// If the user is running functional tests on their laptop, choose a
// subnetID for them, otherwise we'll leave it blank and pull out the
// subnet from the metadata service.
func maybeSetSubnet() string {
	_, err := detectCurrentVPC()
	if err != nil {
		return defaultSubnetID
	}
	return ""
}

var runFunctional = flag.Bool("functional", false, "run functional system tests")

func TestAWSCloud(t *testing.T) {
	if !(*runFunctional) {
		t.Skip("skipping cloud functional tests")
	}
	fmt.Printf("Running Functional Tests\n")

	if !util.AWSEnvVarsSet() {
		t.Fatal("Neet to setup AWS env vars for tests")
	}
	controllerID := api.SimpleNameGenerator.GenerateName(testControllerID)
	//subnetID := maybeSetSubnet()
	// While we're running 2 builds in different vpcs, just use the default
	// subnet
	subnetID := defaultSubnetID
	c, err := NewEC2Client(controllerID, controllerID, vpcID, subnetID, "")
	assert.Nil(t, err)
	ts, err := functional.SetupCloudFunctionalTest(t, c, imageAmi, instanceType)
	if err != nil {
		assert.FailNow(t, "Failed to setup functional test: %s", err.Error())
	}
	defer ts.Cleanup(t)
	t.Run("GetRegistryAuthTest", func(t *testing.T) {
		functional.ContainerAuthTest(t, ts.CloudClient)
	})
	t.Run("BootSpotInstanceTest", func(t *testing.T) {
		functional.RunSpotInstanceTest(t, ts.CloudClient, imageAmi)
	})
}
