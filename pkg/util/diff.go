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
	"reflect"
)

func MapDiff(spec, status map[string]interface{}) ([]string, []string, []string) {
	// Go through status, things in status but not in spec get added
	// to delete.  Go through spec, things not in status or in status
	// but not equal, get upserted.
	return MapUserDiff(spec, status, reflect.DeepEqual)
}

func MapUserDiff(spec, status map[string]interface{}, differ func(a, b interface{}) bool) ([]string, []string, []string) {
	add := make([]string, 0)
	update := make([]string, 0)
	delete := make([]string, 0)
	for k := range status {
		_, exists := spec[k]
		if !exists {
			delete = append(delete, k)
		}
	}
	for k, v := range spec {
		status_val, exists := status[k]
		if !exists {
			add = append(add, k)
		} else if !differ(v, status_val) {
			update = append(update, k)
		}
	}
	return add, update, delete
}
