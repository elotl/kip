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
	"math"

	"k8s.io/apimachinery/pkg/api/resource"
)

func CPUCoresFraction(q *resource.Quantity) float32 {
	i := q.ScaledValue(resource.Nano)
	f := float64(i) * math.Pow10(int(resource.Nano))
	return float32(f)
}

func ToMiBFloat32(q *resource.Quantity) float32 {
	f := float64(q.Value()) / float64(1024*1024)
	return float32(f)
}

func ToGiBFloat32(q *resource.Quantity) float32 {
	f := float64(q.Value()) / float64(1024*1024*1024)
	return float32(f)
}

func ToMiBRoundUp(q *resource.Quantity) int32 {
	f := ToMiBFloat32(q)
	return int32(math.Ceil(float64(f)))
}
func ToGiBRoundUp(q *resource.Quantity) int32 {
	f := ToGiBFloat32(q)
	return int32(math.Ceil(float64(f)))
}

func ToGBFloat32(q *resource.Quantity) float32 {
	f := float64(q.Value()) / float64(1000*1000*1000)
	return float32(f)
}

func ToGBRoundUp(q *resource.Quantity) int32 {
	f := ToGBFloat32(q)
	return int32(math.Ceil(float64(f)))
}
