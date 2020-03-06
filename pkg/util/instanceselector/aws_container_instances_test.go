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

package instanceselector

import (
	"fmt"
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestFargateInstanceSelector(t *testing.T) {
	cases := []struct {
		memSpec string
		cpuSpec string
		mem     int64
		cpu     int64
		error   bool
	}{
		{
			memSpec: "0",
			cpuSpec: "0",
			mem:     512,
			cpu:     256,
			error:   false,
		},
		{
			memSpec: "512Mi",
			cpuSpec: ".25",
			mem:     512,
			cpu:     256,
			error:   false,
		},
		{
			memSpec: "513Mi",
			cpuSpec: ".25",
			mem:     1024,
			cpu:     256,
			error:   false,
		},
		{
			memSpec: "8Gi",
			cpuSpec: ".25",
			mem:     8 * 1024,
			cpu:     VCPU,
			error:   false,
		},
		{
			memSpec: "8Gi",
			cpuSpec: ".25",
			mem:     8 * 1024,
			cpu:     VCPU,
			error:   false,
		},
		{
			memSpec: "700Mi",
			cpuSpec: "2.5",
			mem:     8 * 1024,
			cpu:     4 * VCPU,
			error:   false,
		},
		{
			memSpec: "700Mi",
			cpuSpec: "4.1",
			mem:     0,
			cpu:     0,
			error:   true,
		},
		{
			memSpec: "35Gi",
			cpuSpec: "5",
			mem:     0,
			cpu:     0,
			error:   true,
		},
	}

	err := Setup("aws", "us-east-1", "t3.nano")
	if err != nil {
		assert.Fail(t, "Failed to setup instanceselector")
		return
	}
	for i, tc := range cases {
		rs := api.ResourceSpec{
			Memory: tc.memSpec,
			CPU:    tc.cpuSpec,
		}
		cpu, mem, err := ResourcesToContainerInstance(&rs)
		if (tc.error && err == nil) ||
			(!tc.error && err != nil) {
			msg := fmt.Sprintf("test case %d said error should be %v but it was %v", i+1, tc.error, err)
			assert.Fail(t, msg)
			continue
		}
		assert.Equal(t, tc.cpu, cpu, "test case %d: cpu failed", i+1)
		assert.Equal(t, tc.mem, mem, "test case %d: memory failed", i+1)
	}
}
