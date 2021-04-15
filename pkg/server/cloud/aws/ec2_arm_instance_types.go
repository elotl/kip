package aws

import "strings"

var ARMInstanceTypes = []string{
	"t4g.",
	"c6g.",
	"c6gd.",
	"c6gn.",
	"m6g.",
	"m6gd.",
	"r6g.",
	"r6gd.",
	"x2gd.",
}

func checkInstanceTypeArch(instanceType string) string {
	for _, instanceFamilyPrefix := range ARMInstanceTypes {
		if strings.HasPrefix(instanceType, instanceFamilyPrefix) {
			return arm64BootImageArch
		}
	}
	return defaultBootImageArch
}
