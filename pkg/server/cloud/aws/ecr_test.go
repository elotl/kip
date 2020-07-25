package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegionFromImage(t *testing.T) {
	tests := []struct {
		image  string
		region string
		isErr  bool
	}{
		{
			image:  "12345929.dkr.ecr.us-west-1.amazonaws.com/helloserver:latest",
			region: "us-west-1",
		},
	}
	for _, tc := range tests {
		r, err := regionFromImage(tc.image)
		if tc.isErr {
			assert.Error(t, err)
			continue
		}
		assert.NoError(t, err)
		assert.Equal(t, tc.region, r)
	}
}
