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
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestCPUCoresFraction(t *testing.T) {
	cpuSize := resource.MustParse("250m")
	assert.Equal(t, float32(0.25), CPUCoresFraction(&cpuSize))
	cpuSize = resource.MustParse("500m")
	assert.Equal(t, float32(0.5), CPUCoresFraction(&cpuSize))
	cpuSize = resource.MustParse("1")
	assert.Equal(t, float32(1), CPUCoresFraction(&cpuSize))
	cpuSize = resource.MustParse("5")
	assert.Equal(t, float32(5), CPUCoresFraction(&cpuSize))
	cpuSize = resource.MustParse("1.5")
	assert.Equal(t, float32(1.5), CPUCoresFraction(&cpuSize))
}

func TestToGiBRoundUp(t *testing.T) {
	memorySize := resource.MustParse("500Mi")
	assert.Equal(t, int32(1), ToGiBRoundUp(&memorySize))
	memorySize = resource.MustParse("1.5Gi")
	assert.Equal(t, int32(2), ToGiBRoundUp(&memorySize))
	memorySize = resource.MustParse("1500Mi")
	assert.Equal(t, int32(2), ToGiBRoundUp(&memorySize))
	memorySize = resource.MustParse("100Mi")
	assert.Equal(t, int32(1), ToGiBRoundUp(&memorySize))
	memorySize = resource.MustParse("1000Mi")
	assert.Equal(t, int32(1), ToGiBRoundUp(&memorySize))
}

func TestToGBRoundUp(t *testing.T) {
	memorySize := resource.MustParse("500M")
	assert.Equal(t, int32(1), ToGBRoundUp(&memorySize))
	memorySize = resource.MustParse("1.5G")
	assert.Equal(t, int32(2), ToGBRoundUp(&memorySize))
	memorySize = resource.MustParse("1500M")
	assert.Equal(t, int32(2), ToGBRoundUp(&memorySize))
	memorySize = resource.MustParse("100M")
	assert.Equal(t, int32(1), ToGBRoundUp(&memorySize))
	memorySize = resource.MustParse("1000M")
	assert.Equal(t, int32(1), ToGBRoundUp(&memorySize))
	memorySize, _ = resource.ParseQuantity("")
	assert.Equal(t, int32(0), ToGBRoundUp(&memorySize))
}

func TestToGiBFloat32(t *testing.T) {
	memorySize := resource.MustParse("512Mi")
	assert.Equal(t, float32(0.5), ToGiBFloat32(&memorySize))
	memorySize = resource.MustParse("1.5Gi")
	assert.Equal(t, float32(1.5), ToGiBFloat32(&memorySize))
	memorySize = resource.MustParse("200Gi")
	assert.Equal(t, float32(200), ToGiBFloat32(&memorySize))
}

func TestToGBFloat32(t *testing.T) {
	memorySize := resource.MustParse("500M")
	assert.Equal(t, float32(0.5), ToGBFloat32(&memorySize))
	memorySize = resource.MustParse("1.5G")
	assert.Equal(t, float32(1.5), ToGBFloat32(&memorySize))
	memorySize = resource.MustParse("200G")
	assert.Equal(t, float32(200), ToGBFloat32(&memorySize))
	memorySize = resource.MustParse("1000000000")
	assert.Equal(t, float32(1), ToGBFloat32(&memorySize))
}
