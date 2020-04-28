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
