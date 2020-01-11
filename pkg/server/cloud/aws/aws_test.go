package aws

import (
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/stretchr/testify/assert"
)

func TestAWSTagsFromLabels(t *testing.T) {
	tests := []struct {
		k       string
		v       string
		fail    bool
		numTags int
	}{
		{"foo", "bar", false, 1},
		{"aws:foo", "bar", false, 0},
		{cloud.ControllerTagKey, "bar", true, 0},
		{"foooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo", "bar", true, 0},
		{"foo", "baaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaar", true, 0},
	}
	for i, tc := range tests {
		maxAWSUserTags = 5
		tags, err := filterLabelsForTags("test", map[string]string{tc.k: tc.v})
		if tc.fail {
			assert.Error(t, err, "test %d should have returned an error", i+1)
		} else {
			assert.NoError(t, err, "test %d should have not raised an error", i+1)
		}
		assert.Equal(t, tc.numTags, len(tags))
	}
	mytags := map[string]string{"key1": "val1", "key2": "val2"}
	maxAWSUserTags = 2
	_, err := filterLabelsForTags("test", mytags)
	assert.NoError(t, err)
	maxAWSUserTags = 1
	_, err = filterLabelsForTags("test", mytags)
	assert.Error(t, err)
}

// eventually test for getting VPC resolver address
