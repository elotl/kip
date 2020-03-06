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

package hash

import (
	"encoding/base32"
	"encoding/hex"
	"hash"
	"hash/fnv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

// DeepHashObject writes specified object to hash using the spew library
// which follows pointers and prints actual values of the nested objects
// ensuring the hash does not change when a pointer changes.
func DeepHashObject(hasher hash.Hash, objectToWrite interface{}) {
	hasher.Reset()
	printer := spew.ConfigState{
		Indent:         " ",
		SortKeys:       true,
		DisableMethods: true,
		SpewKeys:       true,
	}
	printer.Fprintf(hasher, "%#v", objectToWrite)
}

func computeHashHelper(obj interface{}) []byte {
	hasher := fnv.New128a()
	DeepHashObject(hasher, obj)
	in := make([]byte, 0, hasher.Size())
	return hasher.Sum(in)
}

func ComputeHash128Hex(obj interface{}) string {
	return hex.EncodeToString(computeHashHelper(obj))
}

var noPadEncoder = base32.StdEncoding.WithPadding(base32.NoPadding)

// This hash ends up 26 characters long
func ComputeHash128b32(obj interface{}) string {
	return Base32EncodeNoPad(computeHashHelper(obj))
}

func Base32EncodeNoPad(b []byte) string {
	return strings.ToLower(noPadEncoder.EncodeToString(b))
}
