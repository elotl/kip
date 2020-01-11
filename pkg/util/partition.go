package util

import "reflect"

// take a slice and move all elements that satisfy isFront to the
// front of the slice.  isFront is a function that, given an index
// of an element, looks up that elemenet in the slice and returns
// true if the element should be at the front of the slice.
func PartitionSlice(slice interface{}, isFront func(i int) bool) {
	rv := reflect.ValueOf(slice)
	swapper := reflect.Swapper(slice)
	a := 0
	b := rv.Len() - 1
	for a < b {
		if isFront(a) {
			a++
		} else {
			swapper(a, b)
			b--
		}
	}
}
