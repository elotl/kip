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

package aws

func breakIntoChunks(vals []*string, chunkSize int) [][]*string {
	var chunked [][]*string
	for i := 0; i < len(vals); i += chunkSize {
		end := i + chunkSize
		if end > len(vals) {
			end = len(vals)
		}
		chunked = append(chunked, vals[i:end])
	}
	return chunked
}
