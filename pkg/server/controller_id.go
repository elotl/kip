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
	"github.com/docker/libkv/store"
	"github.com/elotl/kip/pkg/etcd"
	"github.com/elotl/kip/pkg/util"
	"github.com/elotl/kip/pkg/util/hash"
	uuid "github.com/satori/go.uuid"
)

// Note: we might want to eventually create a registry with various
// piece of info for the server. Right now this is the only thing that
// would go in it so I'm not creating the registry just for storing
// the UUID of this server...
const (
	etcdClusterInfoPath string = "milpa/cluster"
	etcdClusterUUIDPath string = "milpa/cluster/uuid"
)

// Internally we store a UUID4 for the controller that is a string
// with 36 characters.  However, we want to use that string in various
// identifiers in the cloud so we convert it back to a UUID, take the
// bits, base32 encode them, strip padding and end up with a 26
// character string which works out better for example, prepending
// "kiyot-" to the encoded string gives us a string of length 32.
// That string can be used in most cloud resource names and tags.
func getControllerID(etcdClient etcd.Storer) (string, error) {
	controllerUUIDString, err := ensureClusterUUID(etcdClient)
	if err != nil {
		return "", util.WrapError(err, "")
	}
	controllerUUID, err := uuid.FromString(controllerUUIDString)
	if err != nil {
		return "", util.WrapError(err, "")
	}
	controllerID := hash.Base32EncodeNoPad(controllerUUID.Bytes())
	return controllerID, nil
}

func ensureClusterUUID(etcdClient etcd.Storer) (string, error) {
	pair, err := etcdClient.Get(etcdClusterUUIDPath)
	if err == nil {
		clusterUUID := string(pair.Value)
		return clusterUUID, nil
	} else if err == store.ErrKeyNotFound {
		clusterUUID := uuid.NewV4().String()
		_, _, err = etcdClient.AtomicPut(etcdClusterUUIDPath, []byte(clusterUUID), nil, nil)
		if err != nil {
			return clusterUUID, util.WrapError(err, "Error storing new controller UUID")
		}
		return clusterUUID, nil
	} else {
		return "", util.WrapError(err, "Error pulling controller UUID from storage")
	}
}
