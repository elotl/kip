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
	"net"
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

func TestNextIP(t *testing.T) {
	testCases := []struct {
		BaseIP net.IP
		Inc    uint
		NextIP net.IP
	}{
		{
			BaseIP: net.ParseIP("10.10.0.0"),
			Inc:    1,
			NextIP: net.ParseIP("10.10.0.1"),
		},
		{
			BaseIP: net.ParseIP("10.10.0.255"),
			Inc:    2,
			NextIP: net.ParseIP("10.10.1.1"),
		},
		{
			BaseIP: net.ParseIP("192.168.100.255"),
			Inc:    1,
			NextIP: net.ParseIP("192.168.101.0"),
		},
		{
			BaseIP: net.ParseIP("255.255.255.255"),
			Inc:    2,
			NextIP: net.ParseIP("0.0.0.1"),
		},
	}
	for i, tc := range testCases {
		nextIP := NextIP(tc.BaseIP, tc.Inc)
		msg := fmt.Sprintf("testcase %d failed", i+1)
		assert.Equal(t, tc.NextIP, nextIP, msg)
	}
}
