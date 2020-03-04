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
	"fmt"
	"strings"
	"time"
)

const (
	NamespaceSeparator = '_'
)

var (
	InternalLabelPrefixes = []string{
		"kiyot.",
		"io.kubernetes",
		"kubernetes.io",
	}
)

func GetNamespaceFromString(n string) string {
	if i := strings.IndexByte(n, NamespaceSeparator); i > 0 {
		return n[:i]
	}
	return ""
}

func GetNameFromString(n string) string {
	i := strings.IndexByte(n, NamespaceSeparator)
	if i >= 0 && i < len(n)-1 {
		return n[i+1:]
	} else if i == len(n)-1 {
		return ""
	}
	return n
}

func WithNamespace(ns, name string) string {
	return ns + string(NamespaceSeparator) + name
}

func SplitNamespaceAndName(n string) (string, string) {
	parts := strings.SplitN(n, string(NamespaceSeparator), 2)
	if len(parts) == 0 {
		return "", ""
	} else if len(parts) == 1 {
		return "", parts[0]
	} else {
		return parts[0], parts[1]
	}
}

//GCP requires names to follow the regex: [a-z]([-a-z0-9]*[a-z0-9])?
func CreateSecurityGroupName(controllerID, svcName string) string {
	return strings.ToLower(fmt.Sprintf("kip-%s-%s", controllerID, svcName))
}

func CreateUnboundNodeNameTag(nametag string) string {
	return fmt.Sprintf(
		"Kip Node %s %s", nametag, time.Now().UTC().Format(time.Stamp))
}

func CreateBoundNodeNameTag(nametag, podName string) string {
	return fmt.Sprintf("Kip Node %s %s", nametag, podName)
}

func CreateResourceGroupName(region string) string {
	return fmt.Sprintf("kip-%s", region)
}

func CreateClusterResourceGroupName(controllerID string) string {
	return fmt.Sprintf("kip-%s", controllerID)
}

func CreateClusterResourcePrefix(controllerID string) string {
	return fmt.Sprintf("kip-%s-", controllerID)
}
