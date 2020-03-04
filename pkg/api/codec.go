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

// Allows us to change out our storage format (json, yaml, msgpack, etc.)
package api

import (
	"github.com/json-iterator/go"

	"k8s.io/klog"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type TypeVersioner interface {
	GetAPIVersion() string
}

// Storage codec is the codec interface that's used to serialize and
// unserialize objects to and from a format that can be used by our
// storage system (etcd).  Examples of possible codecs would be
// Json/msgpack/protobufs, etc.
type StorageCodec interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
}

// +k8s:deepcopy-gen=false
type JsonCodec struct{}

// +k8s:deepcopy-gen=false
type IndentingJsonCodec struct {
	JsonCodec
}

func (c JsonCodec) Marshal(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}

func warnIfUnversioned(t TypeVersioner) {
	version := t.GetAPIVersion()
	if version == "" {
		klog.Warningf("Found empty API version in registry for %v", t)
	}
}

func (c JsonCodec) Unmarshal(data []byte, v interface{}) error {
	// Everything in the registry must have an API version string.
	// Warn if we ever unmarshall anything with an empty version
	err := json.Unmarshal(data, v)
	switch t := v.(type) {
	case TypeVersioner:
		warnIfUnversioned(t)
	default:
	}
	return err
}

func (c IndentingJsonCodec) Marshal(i interface{}) ([]byte, error) {
	return json.MarshalIndent(i, "", "    ")
}

// MilpaCodec is the codec used by the registry.  It knows about types
// and versions and should know how to update versions.  Internally
// this codec should use a StorageCodec for doing the actual
// serialization.
type MilpaCodec interface {
	Marshal(MilpaObject) ([]byte, error)
	Unmarshal([]byte, MilpaObject) error
}

// +k8s:deepcopy-gen=false
type VersioningCodec struct{}

func (c VersioningCodec) Marshal(i MilpaObject) ([]byte, error) {
	return Encode(i)
}

func (c VersioningCodec) Unmarshal(data []byte, v MilpaObject) error {
	return DecodeInto(data, v)
}
