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

package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseImageSpec(t *testing.T) {
	tests := []struct {
		image     string
		server    string
		repoImage string
		err       bool
	}{
		{
			image:     "user/repo:tag",
			server:    "",
			repoImage: "user/repo:tag",
		},
		{
			image:     "someurl.com/user/repo",
			server:    "someurl.com",
			repoImage: "user/repo",
		},
		{
			image:     "ACCOUNT.dkr.ecr.REGION.amazonaws.com/imagename:tag",
			server:    "ACCOUNT.dkr.ecr.REGION.amazonaws.com",
			repoImage: "imagename:tag",
		},
		{
			image:     "ACCOUNT.dkr.ecr.REGION.amazonaws.com/multiple/level/repo/imagename:tag",
			server:    "ACCOUNT.dkr.ecr.REGION.amazonaws.com",
			repoImage: "multiple/level/repo/imagename:tag",
		},
		{
			image:     "602401143452.dkr.ecr.us-east-1.amazonaws.com/eks/kube-dns/sidecar:1.14.10",
			server:    "602401143452.dkr.ecr.us-east-1.amazonaws.com",
			repoImage: "eks/kube-dns/sidecar:1.14.10",
		},
	}
	for _, test := range tests {
		server, repoImage, err := ParseImageSpec(test.image)
		if test.err {
			assert.Error(t, err)
			continue
		}
		assert.NoError(t, err)
		assert.Equal(t, server, test.server)
		assert.Equal(t, repoImage, test.repoImage)
	}
}
