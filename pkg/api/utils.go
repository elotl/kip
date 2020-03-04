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

package api

func AllPodUnits(pod *Pod) []Unit {
	units := make([]Unit, 0, len(pod.Spec.Units)+len(pod.Spec.InitUnits))
	for i := 0; i < len(pod.Spec.InitUnits); i++ {
		units = append(units, pod.Spec.InitUnits[i])
	}
	for i := 0; i < len(pod.Spec.Units); i++ {
		units = append(units, pod.Spec.Units[i])
	}
	return units
}
