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
