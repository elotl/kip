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
