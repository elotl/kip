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
	"net"
)

func CIDRInsideCIDR(cidrA string, cidrB string) bool {
	return CIDRInsideCIDRs(cidrA, []string{cidrB})
}

func CIDRInsideCIDRs(cidrA string, cidrs []string) bool {
	ipA, _, err := net.ParseCIDR(cidrA)
	if err != nil {
		return false
	}
	for _, cidrB := range cidrs {
		_, networkB, err := net.ParseCIDR(cidrB)
		if err != nil {
			continue
		}
		if networkB.Contains(ipA) {
			return true
		}
	}
	return false
}

func NextIP(ip net.IP, inc uint) net.IP {
	i := ip.To4()
	v := uint(i[0])<<24 + uint(i[1])<<16 + uint(i[2])<<8 + uint(i[3])
	v += inc
	v3 := byte(v & 0xFF)
	v2 := byte((v >> 8) & 0xFF)
	v1 := byte((v >> 16) & 0xFF)
	v0 := byte((v >> 24) & 0xFF)
	return net.IPv4(v0, v1, v2, v3)
}
