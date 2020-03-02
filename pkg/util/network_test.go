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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCIDRInsideCIDRs(t *testing.T) {
	cases := []struct {
		A      string
		B      []string
		inside bool
	}{
		{
			A:      "10.0.1.0/24",
			B:      []string{"192.168.5.0/24"},
			inside: false,
		},
		{
			A:      "10.1.1.0/24",
			B:      []string{"10.1.0.0/16"},
			inside: true,
		},
		{
			A:      "10.1.1.0/24",
			B:      []string{"192.168.5.0/24", "10.1.0.0/16"},
			inside: true,
		},
		{
			A:      "99.10.1.1.0/24",
			B:      []string{"10.1.0.0/16"},
			inside: false,
		},
	}
	for i, tc := range cases {
		result := CIDRInsideCIDRs(tc.A, tc.B)
		msg := fmt.Sprintf("testcase %d (zero offset) failed", i)
		assert.Equal(t, tc.inside, result, msg)
	}
}
