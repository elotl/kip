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

	"github.com/docker/libkv/store"
	"github.com/elotl/kip/pkg/server/registry"
	"github.com/elotl/kip/pkg/util"
	"github.com/elotl/kip/pkg/util/hash"
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
	assert.NoError(t, err)
	assert.Equal(t, u, u2)
	uu, err := uuid.FromString(u)
	assert.NoError(t, err)
	encoded := hash.Base32EncodeNoPad(uu.Bytes())
	controllerID, err := getControllerID(kvstore)
	assert.NoError(t, err)
	assert.Equal(t, encoded, controllerID)
}
