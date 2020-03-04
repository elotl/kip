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
	"bytes"
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/nodeclient"
	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/stretchr/testify/assert"
)

func setupDeployTestServer() (InstanceProvider, func()) {
	nodeReg, closer1 := registry.SetupTestNodeRegistry()
	podReg, closer2 := registry.SetupTestPodRegistry()
	s := InstanceProvider{
		Registries: map[string]registry.Registryer{
			"Node": nodeReg,
			"Pod":  podReg,
		},
		ItzoClientFactory: nodeclient.NewMockItzoClientFactory(),
	}
	return s, func() { closer1(); closer2() }
}

func TestDeploy(t *testing.T) {
	s, closer := setupDeployTestServer()
	defer closer()
	nodeReg := s.Registries["Node"].(*registry.NodeRegistry)
	node := api.GetFakeNode()
	node.Status.Addresses = api.NewNetworkAddresses("1.2.3.4", "")
	_, err := nodeReg.CreateNode(node)
	assert.NoError(t, err)
	podReg := s.Registries["Pod"].(*registry.PodRegistry)
	pod := api.GetFakePod()
	pod.Status.BoundNodeName = node.Name
	pod.Status.Phase = api.PodRunning
	_, err = podReg.CreatePod(pod)
	assert.NoError(t, err)
	var buf bytes.Buffer
	buf.WriteString("foobar")
	err = s.deploy(pod.Name, "mypkg", &buf) //make([]byte, 1))
	assert.NoError(t, err)
}

func TestDeployNoPod(t *testing.T) {
	s, closer := setupDeployTestServer()
	defer closer()
	var buf bytes.Buffer
	buf.WriteString("foobar")
	err := s.deploy("mypod", "mypkg", &buf)
	assert.Error(t, err)
}

func TestDeployNoNode(t *testing.T) {
	s, closer := setupDeployTestServer()
	defer closer()
	podReg := s.Registries["Pod"].(*registry.PodRegistry)
	pod := api.GetFakePod()
	pod.Status.BoundNodeName = ""
	pod.Status.Phase = api.PodRunning
	_, err := podReg.CreatePod(pod)
	assert.NoError(t, err)
	var buf bytes.Buffer
	buf.WriteString("foobar")
	err = s.deploy(pod.Name, "mypkg", &buf)
	assert.Error(t, err)
}
