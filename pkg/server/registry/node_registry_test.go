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

package registry

import (
	"reflect"
	"testing"

	"github.com/elotl/kip/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestNodeCreateGet(t *testing.T) {
	nodeRegistry, closer := SetupTestNodeRegistry()
	defer closer()

	n1 := api.GetFakeNode()
	_, err := nodeRegistry.CreateNode(n1)
	if err != nil {
		t.Error(err)
	}
	n2, err := nodeRegistry.GetNode(n1.Name)
	if err != nil {
		t.Error(err)
	}
	n3, err := nodeRegistry.GetNode(n2.Name)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(n3, n2) {
		t.Errorf("did not get the same node back:\n%#v\n%#v", n3, n2)
	}
}

func TestNodeDoubleCreate(t *testing.T) {
	nodeRegistry, closer := SetupTestNodeRegistry()
	defer closer()

	n := api.GetFakeNode()
	_, err := nodeRegistry.CreateNode(n)
	assert.Nil(t, err)
	_, err = nodeRegistry.CreateNode(n)
	assert.NotNil(t, err)
}

// func TestNodeDelete(t *testing.T) {
// 	nodeRegistry, closer := SetupTestNodeRegistry()
// 	defer closer()
// 	n1 := api.GetFakeNode()
// 	_, err := nodeRegistry.CreateNode(n1)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	n2, err := nodeRegistry.DeleteNode(n1)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if n2.Spec.Phase != api.NodeTerminated ||
// 		n2.Status.Phase != n1.Status.Phase {
// 		t.Errorf("Delete should set spec to terminated but not set status")
// 	}
// 	if n2.DeletionTimestamp == nil {
// 		t.Errorf("Delete should set Deletion timestamp")
// 	}
// }

func TestListNodes(t *testing.T) {
	nodeRegistry, closer := SetupTestNodeRegistry()
	defer closer()
	n1 := api.GetFakeNode()
	_, err := nodeRegistry.CreateNode(n1)
	if err != nil {
		t.Error(err)
	}
	n2 := api.GetFakeNode()
	_, err = nodeRegistry.CreateNode(n2)
	if err != nil {
		t.Error(err)
	}

	apilist, err := nodeRegistry.List()
	nodes := apilist.(*api.NodeList)
	if err != nil {
		t.Error(err)
	} else if len(nodes.Items) != 2 {
		t.Errorf("Expected to get 2 nodes back, got %d", len(nodes.Items))
	}

	nodes, err = nodeRegistry.ListNodes(func(n *api.Node) bool {
		return n.Name == n1.Name
	})
	if err != nil {
		t.Error(err)
	} else if len(nodes.Items) != 1 {
		t.Errorf("Error listing nodes with predicate, should have filtered down to 1 node, got %d", len(nodes.Items))
	}
}

func TestUpdateNode(t *testing.T) {
	// GIVEN
	nodeRegistry, closer := SetupTestNodeRegistry()
	defer closer()
	n1 := api.GetFakeNode()
	n1.Spec.InstanceType = "t3.nano"
	_, err := nodeRegistry.CreateNode(n1)
	assert.NoError(t, err)
	n, err := nodeRegistry.GetNode(n1.Name)
	assert.NoError(t, err)
	assert.Equal(t, "t3.nano", n.Spec.InstanceType)

	// WHEN
	n.Spec.InstanceType = "t3.micro"
	n2, err := nodeRegistry.UpdateNode(n)

	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, n2)
	assert.Equal(t, "t3.micro", n2.Spec.InstanceType)
}

// func TestPurgeNode(t *testing.T) {
// 	nodeRegistry, closer := SetupTestNodeRegistry()
// 	defer closer()
// 	n1 := api.GetFakeNode()
// 	_, err := nodeRegistry.CreateNode(n1)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	nodeRegistry.PurgeNode(n1.Name)
// 	_, err = nodeRegistry.GetNode(n1.Name)
// 	if err == nil {
// 		t.Errorf("Expected error when requesting node that has been purged")
// 	}
// }
