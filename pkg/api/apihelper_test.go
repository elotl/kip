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

package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type myStructV1 struct {
	TypeMeta `json:",inline"`
	A        int    `json:"a"`
	B        string `json:"b"`
}

func (myStructV1) IsMilpaObject() {}

type myStructV2 struct {
	TypeMeta `json:",inline"`
	A        int    `json:"a"`
	B        string `json:"b"`
	C        string `json:"c"`
}

func (myStructV2) IsMilpaObject() {}

func createV1() myStructV1 {
	s := myStructV1{
		A: 1,
		B: "foobar",
	}
	s.Create()
	s.Kind = "myStructV1"
	return s
}

func createV2(a int, b string, c string) myStructV2 {
	s := myStructV2{
		A: a,
		B: b,
		C: c,
	}
	s.Create()
	s.Kind = "myStructV2"
	return s
}

func TestEncodeDecode(t *testing.T) {
	addKnownTypes("v1", myStructV1{})
	inp := createV1()
	json, err := Encode(inp)
	assert.Nil(t, err)
	outp, err := Decode(json)
	assert.Nil(t, err)
	assert.Equal(t, &inp, outp)
}

func TestDecodeInto(t *testing.T) {
	addKnownTypes("v1", myStructV1{})
	inp := createV1()
	json, err := Encode(inp)
	assert.Nil(t, err)
	var outp myStructV1
	err = DecodeInto(json, &outp)
	assert.Nil(t, err)
	assert.Equal(t, inp, outp)
}

func TestUpgrade(t *testing.T) {
	addKnownTypes("v1", myStructV1{})
	addKnownTypes("v2", myStructV2{})
	// Upgrade function for v1 -> v2 conversion.
	upgradeFuncs["v1"] = func(apiversion, kind string, obj MilpaObject) (MilpaObject, error) {
		assert.Equal(t, "v1", apiversion)
		assert.Equal(t, "myStructV1", kind)
		o1 := obj.(*myStructV1)
		o2 := createV2(o1.A, o1.B, "conversion succeeded")
		return &o2, nil
	}
	LatestAPIVersion = "v1"
	v1 := createV1()
	json, err := Encode(v1)
	assert.Nil(t, err)
	LatestAPIVersion = "v2"
	var v2 myStructV2
	err = DecodeInto(json, &v2)
	assert.Nil(t, err)
	assert.Equal(t, "v2", v2.APIVersion)
	assert.Equal(t, "myStructV2", v2.Kind)
	assert.Equal(t, 1, v2.A)
	assert.Equal(t, "foobar", v2.B)
	assert.Equal(t, "conversion succeeded", v2.C)
}

func TestUpgradeFailure(t *testing.T) {
	addKnownTypes("v1", myStructV1{})
	addKnownTypes("v2", myStructV2{})
	// Upgrade function for v1 -> v2 conversion.
	upgradeFuncs["v1"] = func(apiversion, kind string, obj MilpaObject) (MilpaObject, error) {
		assert.Equal(t, "v1", apiversion)
		assert.Equal(t, "myStructV1", kind)
		return nil, fmt.Errorf("testing upgrade failure")
	}
	LatestAPIVersion = "v1"
	v1 := createV1()
	json, err := Encode(v1)
	assert.Nil(t, err)
	LatestAPIVersion = "v2"
	var v2 myStructV2
	err = DecodeInto(json, &v2)
	assert.Error(t, err)
}

func TestUpgradeMissing(t *testing.T) {
	addKnownTypes("v1", myStructV1{})
	addKnownTypes("v2", myStructV2{})
	// Upgrade function for v1 -> v2 conversion.
	delete(upgradeFuncs, "v1")
	LatestAPIVersion = "v1"
	v1 := createV1()
	json, err := Encode(v1)
	assert.Nil(t, err)
	LatestAPIVersion = "v2"
	var v2 myStructV2
	err = DecodeInto(json, &v2)
	assert.Error(t, err)
}
