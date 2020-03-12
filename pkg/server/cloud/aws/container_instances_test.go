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
			arn:      "arn:aws:ecs:us-east-1:689494258501:task-definition/kip-yav6wbd6vfehrkm2gcmqoirsk4_default_hellofargate:4",
			name:     "default_hellofargate",
			revision: 4,
		},
		{
			arn:      "arn:aws:ecs:us-east-1:689494258501:task-definition/kip-differentcluster_default_hellofargate:4",
			name:     "",
			revision: 0,
		},
		{
			arn:      "arn:aws:ecs:us-east-1:689494258501:task-definition/kip-yav6wbd6vfehrkm2gcmqoirsk4-default_hellofargate:4",
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
