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
