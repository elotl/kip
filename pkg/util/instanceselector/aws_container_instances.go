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

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"k8s.io/apimachinery/pkg/api/resource"
)

// This was taken from virtual kubelet's fargate implementation
// and made to work with our resource spec:
//https://github.com/virtual-kubelet/virtual-kubelet/blob/master/providers/aws/fargate/fargate.go

const (
	// EC2 compute resource units.

	// VCPU is one virtual CPU core in EC2.
	VCPU int64 = 1024
	// MiB is 2^20 bytes.
	MiB int64 = 1024 * 1024
	// GiB is 2^30 bytes.
	GiB int64 = 1024 * MiB
)

// TaskSize represents a Fargate task size.
type taskSize struct {
	cpu    int64
	memory memorySizeRange
}

// MemorySizeRange represents a range of Fargate task memory sizes.
type memorySizeRange struct {
	min int64
	max int64
	inc int64
}

var (
	// Fargate task size table.
	// https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#task_size
	//
	// VCPU     Memory (in MiBs, available in 1GiB increments)
	// ====     ===================
	//  256     512, 1024 ...  2048
	//  512          1024 ...  4096
	// 1024          2048 ...  8192
	// 2048          4096 ... 16384
	// 4096          8192 ... 30720
	//
	taskSizeTable = []taskSize{
		{VCPU / 4, memorySizeRange{512 * MiB, 512 * MiB, 1}},
		{VCPU / 4, memorySizeRange{1 * GiB, 2 * GiB, 1 * GiB}},
		{VCPU / 2, memorySizeRange{1 * GiB, 4 * GiB, 1 * GiB}},
		{1 * VCPU, memorySizeRange{2 * GiB, 8 * GiB, 1 * GiB}},
		{2 * VCPU, memorySizeRange{4 * GiB, 16 * GiB, 1 * GiB}},
		{4 * VCPU, memorySizeRange{8 * GiB, 30 * GiB, 1 * GiB}},
	}
)

func FargateInstanceSelector(rs *api.ResourceSpec) (int64, int64, error) {
	// Fargate tasks have explicit CPU and memory limits. Both are
	// required and specify the maximum amount of resources for the
	// task. The limits must match a task size on taskSizeTable.
	//
	var cpu int64
	var memory int64
	memoryRequest := int64(0)
	cpuRequest := int64(0)
	if rs.Memory == "" {
		memoryRequest = 0.0
	} else {
		memQuantity, err := resource.ParseQuantity(rs.Memory)
		if err != nil {
			return 0, 0, err
		}
		memoryRequest = int64(util.ToMiBFloat32(&memQuantity))
	}

	if rs.CPU == "" {
		cpuRequest = 0.0
	} else {
		cpuQuantity, err := resource.ParseQuantity(rs.CPU)
		if err != nil {
			return 0, 0, err
		}
		cpuRequest = int64(util.CPUCoresFraction(&cpuQuantity) * float32(VCPU))
	}

	// Find the smallest Fargate task size that can satisfy the total resource request.
	for _, row := range taskSizeTable {
		if cpuRequest <= row.cpu {
			for mem := row.memory.min; mem <= row.memory.max; mem += row.memory.inc {
				if memoryRequest <= mem/MiB {
					cpu = row.cpu
					memory = mem / MiB
					break
				}
			}

			if cpu != 0 {
				break
			}
		}
	}

	// Fail if the resource requirements cannot be satisfied by any Fargate task size.
	if cpu == 0 {
		return 0, 0, fmt.Errorf("resource requirements (cpu:%v, memory:%v) are too high", rs.CPU, rs.Memory)
	}

	// Fargate task CPU size is specified in vCPU/1024s and memory size is specified in MiBs.
	return cpu, memory, nil
}
