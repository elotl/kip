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
	"context"
	"encoding/json"
	"testing"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/clientapi"
	"github.com/elotl/kip/pkg/server/registry"
	"github.com/stretchr/testify/assert"
)

func createServerForCreate() (*InstanceProvider, func()) {
	podReg, closer := registry.SetupTestPodRegistry()
	regs := map[string]registry.Registryer{
		"Pod": podReg,
	}
	cm := NewControllerManager(make(map[string]Controller))
	cm.startControllersHelper()
	s := &InstanceProvider{
		Registries:        regs,
		Encoder:           api.VersioningCodec{},
		controllerManager: cm,
	}
	return s, closer
}

func TestServerCreateReqOK(t *testing.T) {
	s, closer := createServerForCreate()
	defer closer()
	pod := api.GetFakePod()
	podBits, err := json.Marshal(pod)
	assert.NoError(t, err)
	req := &clientapi.CreateRequest{
		Manifest: podBits,
	}
	reply, err := s.Create(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, int32(201), reply.Status)
	assert.Len(t, reply.Warning, 0)
	podReg := s.Registries["Pod"].(*registry.PodRegistry)
	p, err := podReg.GetPod(pod.Name)
	assert.NoError(t, err)
	assert.NotNil(t, p)
}

func TestServerCreateReqExtraFields(t *testing.T) {
	s, closer := createServerForCreate()
	defer closer()
	podBits := []byte(`{
	"kind": "Pod",
	"apiVersion": "v1",
	"metadata": {
		"name": "foo",
		"labels": {
			"Label1": "Value1"
		},
		"creationTimestamp": "2018-11-20T22:21:29.351788931-08:00",
		"uid": "45c58ebd-d8a3-471e-84a5-f960ac078b3e",
		"namespace": "default"
	},
	"spec": {
		"phase": "Running",
		"containers": []
	},
	"status": {}
}`)
	req := &clientapi.CreateRequest{
		Manifest: podBits,
	}
	reply, err := s.Create(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, int32(201), reply.Status)
	assert.True(t, len(reply.Warning) > 0)
	assert.Contains(t, string(reply.Warning), "containers")
	podReg := s.Registries["Pod"].(*registry.PodRegistry)
	p, err := podReg.GetPod("foo")
	assert.NoError(t, err)
	assert.NotNil(t, p)
}
