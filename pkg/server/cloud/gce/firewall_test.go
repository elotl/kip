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

package gce

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPortStringToInstancePort(t *testing.T) {
	tests := []struct {
		portString string
		port       int
		prange     int
	}{
		{"80", 80, 1},
		{"8080", 8080, 1},
		{"80-80", 80, 1},
		{"80-81", 80, 2},
		{"500-600", 500, 101},
	}
	for i, tc := range tests {
		p, r := portStringToInstancePort(tc.portString)
		msg := fmt.Sprintf("test case %d fiailed", i)
		assert.Equal(t, tc.port, p, msg)
		assert.Equal(t, tc.prange, r, msg)
	}
}
