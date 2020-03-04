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

package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNamespaceFromString(t *testing.T) {
	vals := [][]string{
		{"a", ""},
		{"a_b", "a"},
		{"aaa_b", "aaa"},
		{"_bbb", ""},
	}
	for i, val := range vals {
		ns := GetNamespaceFromString(val[0])
		assert.Equal(t, val[1], ns, "Test %d failed", i+1)
	}
}

func TestGetNameFromString(t *testing.T) {
	vals := [][]string{
		{"aaa", "aaa"},
		{"_a", "a"},
		{"_aaa", "aaa"},
		{"a_b", "b"},
		{"a_bbb", "bbb"},
		{"a_", ""},
	}
	for i, val := range vals {
		ns := GetNameFromString(val[0])
		assert.Equal(t, val[1], ns, "Test %d failed", i+1)
	}
}
