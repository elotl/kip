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

package azure

import (
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/stretchr/testify/assert"
)

func TestFilterTagsFromLabels(t *testing.T) {
	tests := []struct {
		k    string
		v    string
		fail bool
	}{
		{"foo", "bar", false},
		{cloud.ControllerTagKey, "bar", true},
		{"foooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo", "bar", true},
		{"foo", "baaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaar", true},
	}
	for i, tc := range tests {
		maxAzureUserTags = 5
		_, err := filterLabelsForTags("test", map[string]string{tc.k: tc.v})
		if tc.fail {
			assert.Error(t, err, "test %d should have returned an error", i+1)
		} else {
			assert.NoError(t, err, "test %d should have not raised an error", i+1)
		}
	}
	mytags := map[string]string{"key1": "val1", "key2": "val2"}
	maxAzureUserTags = 2
	_, err := filterLabelsForTags("test", mytags)
	assert.NoError(t, err)
	maxAzureUserTags = 1
	_, err = filterLabelsForTags("test", mytags)
	assert.Error(t, err)
}
