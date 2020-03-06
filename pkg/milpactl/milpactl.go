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

package milpactl

import "strings"

// Expand nicknames and properly capitalize resource type names
func CleanupResourceName(resource string) string {
	resourceLower := strings.ToLower(resource)
	forms := map[string]string{
		// lower to upper
		"event":   "Event",
		"metric":  "Metric",
		"node":    "Node",
		"pod":     "Pod",
		"service": "Service",
		"secret":  "Secret",

		// Plurals
		"events":   "Event",
		"metrics":  "Metric",
		"nodes":    "Node",
		"pods":     "Pod",
		"services": "Service",
		"secrets":  "Secret",

		// Nicknames
		"ev":  "Event",
		"no":  "Node",
		"po":  "Pod",
		"svc": "Service",
	}
	if cleanedUp, ok := forms[resourceLower]; ok {
		return cleanedUp
	}
	return resource
}
