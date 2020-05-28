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

package gce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		s     string
		zone  string
		isErr bool
	}{
		{
			s:    "projects/832569367454/zones/us-west1-a",
			zone: "us-west1-a",
		},
		{
			s:    "projects/832569367454/zones/us-west1-a/",
			zone: "us-west1-a",
		},
		{
			s:     "",
			isErr: true,
		},
	}
	for _, tc := range tests {
		resp, err := extractZoneFromResponse(tc.s)
		if tc.isErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.zone, resp)
		}
	}
}
