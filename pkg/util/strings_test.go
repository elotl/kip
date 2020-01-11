package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterKeysWithPrefix(t *testing.T) {
	m := map[string]string{
		"elotl.co.yes":     "",
		"k8s.io.something": "",
		"foo":              "",
		"elotl.":           "",
		"elotl":            "",
	}

	filter := []string{"elotl.", "k8s.io"}
	res := FilterKeysWithPrefix(m, filter)
	expected := map[string]string{
		"foo":   "",
		"elotl": "",
	}
	assert.Equal(t, expected, res)
}
