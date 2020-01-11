package server

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/stretchr/testify/assert"
)

func createServerForCreate() (*InstanceProvider, func()) {
	kv := make(map[string]registry.Registryer)
	podReg, closer := registry.SetupTestPodRegistry()
	kv["Pod"] = podReg
	cm := NewControllerManager(make(map[string]Controller))
	cm.startControllersHelper()
	s := &InstanceProvider{
		KV:                kv,
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
	podReg := s.KV["Pod"].(*registry.PodRegistry)
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
	podReg := s.KV["Pod"].(*registry.PodRegistry)
	p, err := podReg.GetPod("foo")
	assert.NoError(t, err)
	assert.NotNil(t, p)
}
