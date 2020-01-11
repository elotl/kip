package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitTaskDef(t *testing.T) {
	tests := []struct {
		arn      string
		name     string
		revision int
	}{
		{
			arn:      "arn:aws:ecs:us-east-1:689494258501:task-definition/milpa-yav6wbd6vfehrkm2gcmqoirsk4_default_hellofargate:4",
			name:     "default_hellofargate",
			revision: 4,
		},
		{
			arn:      "arn:aws:ecs:us-east-1:689494258501:task-definition/milpa-differentcluster_default_hellofargate:4",
			name:     "",
			revision: 0,
		},
		{
			arn:      "arn:aws:ecs:us-east-1:689494258501:task-definition/milpa-yav6wbd6vfehrkm2gcmqoirsk4-default_hellofargate:4",
			name:     "",
			revision: 0,
		},
	}
	controllerID := "yav6wbd6vfehrkm2gcmqoirsk4"
	for _, tc := range tests {
		n, r := SplitTaskDef(tc.arn, controllerID)
		assert.Equal(t, tc.name, n)
		assert.Equal(t, tc.revision, r)
	}
}
