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
