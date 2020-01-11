package server

import (
	"testing"

	"github.com/docker/libkv/store"
	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/elotl/cloud-instance-provider/pkg/util/hash"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestClusterID(t *testing.T) {
	name, closer := util.MakeTempFileName("milpa_cid")
	defer closer()
	kvstore := registry.CreateKVStore(name)
	_, err := kvstore.Get(etcdClusterUUIDPath)
	assert.Equal(t, store.ErrKeyNotFound, err)
	u, err := ensureClusterUUID(kvstore)
	assert.NoError(t, err)
	u2, err := ensureClusterUUID(kvstore)
	assert.Equal(t, u, u2)
	uu, err := uuid.FromString(u)
	assert.NoError(t, err)
	encoded := hash.Base32EncodeNoPad(uu.Bytes())
	controllerID, err := getControllerID(kvstore)
	assert.NoError(t, err)
	assert.Equal(t, encoded, controllerID)
}
