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
