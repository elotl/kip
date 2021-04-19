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
	"time"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/server/cloud/functional"
	"github.com/elotl/kip/pkg/util"
	"github.com/stretchr/testify/assert"
	"k8s.io/klog"
)

// should probably load this from some static config
const (
	testControllerID = "AwsFunctionalTest"
	vpcID            = "vpc-841834e2"
	defaultSubnetID  = "subnet-12a8a13f"
	imageAmi         = "ami-e2dea19d"
	imageARMAmi      = "ami-02316a849a44166f1"
	rootDevice       = "xvda" // Update if imageAmi is changed.
	instanceType     = "t2.nano"
)

var runFunctional = flag.Bool("functional", false, "run functional system tests")

func getEC2(t *testing.T, controllerID string) *AwsEC2 {
	if !util.AWSEnvVarsSet() {
		t.Fatal("Neet to setup AWS env vars for tests")
	}
	e, err := NewEC2Client(EC2ClientConfig{
		ControllerID: controllerID,
		Nametag:      controllerID,
		VPCID:        vpcID,
		SubnetID:     defaultSubnetID,
	})
	if err != nil {
		msg := "Error getting EC2 Client: " + err.Error()
		assert.FailNow(t, msg)
	}

	return e
}

// Each functional test run creates a unique securty group that we need to delete
func ensureSecurityGroupDeleted(e *AwsEC2) error {
	apiGroupName := util.CreateSecurityGroupName(e.controllerID, cloud.MilpaAPISGName)
	sg, err := e.FindSecurityGroup(apiGroupName)
	if err != nil {
		// Todo, return if security group not found
		return util.WrapError(err, "Error finding security group")
	}
	if sg == nil {
		return nil
	}

	// Try to delete groups for up to 3 minutes
	// this has flaked a few times and we've had dependency violations for
	// up to 2 mintues...  It's getting worse for AWS.
	// Azure takes for freakin' ever
	for i := 1; i < 60; i++ {
		fmt.Printf("Deleting security group: %s - %s\n", sg.Name, sg.ID)
		err = e.DeleteSecurityGroup(sg.ID)
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return util.WrapError(err, "Could not delete security group")
	}
	klog.Infof("Deleted security group %s", sg.ID)
	return nil
}

func TestAWSCloud(t *testing.T) {
	if !(*runFunctional) {
		t.Skip("skipping aws cloud functional tests")
	}
	fmt.Printf("Running Functional Tests\n")

	if !util.AWSEnvVarsSet() {
		t.Fatal("Neet to setup AWS env vars for tests")
	}
	controllerID := api.SimpleNameGenerator.GenerateName(testControllerID)
	c, err := NewEC2Client(EC2ClientConfig{
		ControllerID: controllerID,
		Nametag:      controllerID,
		VPCID:        vpcID,
		SubnetID:     defaultSubnetID,
	})
	assert.Nil(t, err)
	defer func() {
		err := ensureSecurityGroupDeleted(c)
		if err != nil {
			assert.FailNow(t, "Failed to delete cell security group")
		}
	}()

	ts, err := functional.SetupCloudFunctionalTest(t, c, imageAmi, rootDevice, instanceType)
	if err != nil {
		assert.FailNow(t, "Failed to setup functional test: %s", err.Error())
	}
	defer ts.Cleanup(t)
	t.Run("GetRegistryAuthTest", func(t *testing.T) {
		AWSContainerAuthTest(t, ts.CloudClient)
	})
	t.Run("BootSpotInstanceTest", func(t *testing.T) {
		functional.RunSpotInstanceTest(t, ts.CloudClient, imageAmi, "t3.micro", rootDevice)
	})
	t.Run("BootSpotARM64InstanceTest", func(t *testing.T) {
		functional.RunSpotInstanceTest(t, ts.CloudClient, imageARMAmi, "t4g.micro", rootDevice)
	})
}

func AWSContainerAuthTest(t *testing.T, c cloud.CloudClient) {
	username1, password1, err := c.GetRegistryAuth("689494258501.dkr.ecr.us-east-1.amazonaws.com/helloserver:latest")
	assert.NoError(t, err, "Error getting container authorization")
	assert.Equal(t, "AWS", username1)

	// Make sure we cache passwords
	username2, password2, err := c.GetRegistryAuth("689494258501.dkr.ecr.us-east-1.amazonaws.com/helloserver:latest")
	assert.NoError(t, err, "Error getting container authorization second time")
	assert.Equal(t, username1, username2)
	assert.Equal(t, password1, password2)

	// // Get auth for different region, make sure we get a new password
	_, password3, err := c.GetRegistryAuth("689494258501.dkr.ecr.us-west-1.amazonaws.com/helloserver:latest")
	assert.NoError(t, err, "Error getting container authorization in other region")
	assert.NotEqual(t, password1, password3)
}
