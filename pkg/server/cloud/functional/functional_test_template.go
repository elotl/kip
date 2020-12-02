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

package functional

import (
	"fmt"
	"testing"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/util"
	"github.com/stretchr/testify/assert"
)

type TestState struct {
	CloudClient cloud.CloudClient
	Node1       *api.Node
	Node2       *api.Node
}

func SetupCloudFunctionalTest(t *testing.T, c cloud.CloudClient, imageID, rootDevice, instanceType string) (*TestState, error) {
	fmt.Printf("Running cloud functional tests with %+v\n", c)

	ts := &TestState{CloudClient: c}
	err := SetupFirewallRules(t, c)
	if err != nil {
		fmt.Println("could not setup firewall rules", err)
		t.FailNow()
	}
	inst, err := c.ListInstances()
	assert.Nil(t, err)
	assert.Len(t, inst, 0)

	ts.Node1 = api.GetFakeNode()
	ts.Node1.Spec.BootImage = imageID
	ts.Node1.Spec.InstanceType = instanceType

	ts.Node2 = api.GetFakeNode()
	ts.Node2.Spec.BootImage = imageID
	ts.Node2.Spec.InstanceType = instanceType

	img := cloud.Image{
		ID:         imageID,
		Name:       imageID, // GCE implementation gets these reversed
		RootDevice: rootDevice,
	}

	err = ts.startNodes(img)
	if err != nil {
		ts.Cleanup(t)
		return nil, err
	}
	return ts, nil
}

// StartNode for individual nodes takes too long in azure, start the
// nodes in parallel
func (ts *TestState) startNodes(img cloud.Image) error {
	startResults := make(chan error)
	nodes := []*api.Node{ts.Node1, ts.Node2}
	for _, node := range nodes {
		go func(n *api.Node) {
			fmt.Println("starting node", n.Name)
			instanceID, err := ts.CloudClient.StartNode(n, img, "", "")
			if err != nil {
				startResults <- err
				return
			}
			n.Status.InstanceID = instanceID
			startResults <- nil
		}(node)
	}
	for range nodes {
		err := <-startResults
		if err != nil {
			return util.WrapError(err, "Failed to Start Node, failing test")
		}
	}
	fmt.Println("all nodes started, waiting for nodes to be running")
	addresses1, err := ts.CloudClient.WaitForRunning(ts.Node1)
	if err != nil {
		return util.WrapError(err, "Failed to wait for first node, failing test")
	}
	ts.Node1.Status.Addresses = addresses1

	addresses2, err := ts.CloudClient.WaitForRunning(ts.Node2)
	if err != nil {
		return util.WrapError(err, "Failed to wait for second node, failing test")
	}
	ts.Node2.Status.Addresses = addresses2
	return nil
}

func (ts *TestState) Cleanup(t *testing.T) {
	if r := recover(); r != nil {
		msg := fmt.Sprintf("Recovered functional test, cleaning up %v", r)
		assert.Fail(t, msg)
	}
	deleteInstances(t, ts.CloudClient)
}

func SetupFirewallRules(t *testing.T, c cloud.CloudClient) error {
	extraGroups := []string{}
	extraCIDRs := []string{cloud.PublicCIDR}
	fmt.Println("creating security group")
	err := c.EnsureMilpaSecurityGroups(extraCIDRs, extraGroups)
	return err
}

func RunSpotInstanceTest(t *testing.T, c cloud.CloudClient, imageID, rootDevice string) {
	fmt.Printf("Booting spot instance\n")
	spotNode := api.GetFakeNode()
	// For the last 3 months there has always been an AZ in us-east-1
	// that can boot an m3.medium spot instance.
	spotNode.Spec.InstanceType = "t3.micro"
	spotNode.Spec.BootImage = imageID
	img := cloud.Image{
		ID:         imageID,
		RootDevice: rootDevice,
	}
	instanceID, err := c.StartSpotNode(spotNode, img, "", "")
	if err != nil {
		msg := fmt.Sprintf("Failed to Start Spot Node, failing test %s", err)
		assert.Fail(t, msg)
		return
	}
	fmt.Printf("Got spot instance: %s", instanceID)
	spotNode.Status.InstanceID = instanceID
	_, err = c.WaitForRunning(spotNode)
	assert.Nil(t, err)
}

func deleteInstances(t *testing.T, c cloud.CloudClient) {
	instances, err := c.ListInstances()
	// This is shitty but keeps us from screwing up and deleting
	// groups that aren't in this cluster or doing something else
	// wrong.  We need a better way to be safe but for now, we'll rely
	// on manual cleanup.
	assert.NoError(t, err)
	if len(instances) > 6 {
		t.Fatal("Too many instances listed, something went wrong, clean manually")
	}
	// Azure is dog slow to stop VMs.  We can stop them in parallel
	stopResults := make(chan error)
	for _, inst := range instances {
		go func(instid string) {
			fmt.Printf("Stopping instance %s\n", instid)
			err := c.StopInstance(instid)
			stopResults <- err
		}(inst.ID)
	}
	for i := range instances {
		err := <-stopResults
		assert.NoError(t, err)
		fmt.Printf("%d instances stopped\n", i+1)
	}
}
