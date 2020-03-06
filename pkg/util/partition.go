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
