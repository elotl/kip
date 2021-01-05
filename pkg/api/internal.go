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

type PodParameters struct {
	Secrets     map[string]map[string][]byte   `json:"secrets"`
	Credentials map[string]RegistryCredentials `json:"credentials"`
	Spec        PodSpec                        `json:"spec"`
	Annotations map[string]string
	PodName     string
	NodeName    string
	PodIP       string
	PodHostname string
}

type RegistryCredentials struct {
	Server   string `json:"server"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type PodStatusReply struct {
	UnitStatuses     []UnitStatus    `json:"unitStatus"`
	InitUnitStatuses []UnitStatus    `json:"initUnitStatus"`
	ResourceUsage    ResourceMetrics `json:"resourceUsage,omitempty"`
	PodIP            string          `json:"podIP"`
}

type PortForwardParams struct {
	PodName string
	Port    string
}

type ExecParams struct {
	PodName     string
	UnitName    string
	Command     []string
	Interactive bool
	TTY         bool
	SkipNSEnter bool
}

type AttachParams struct {
	PodName     string
	UnitName    string
	Interactive bool
	TTY         bool
}

type RunCmdParams struct {
	Command []string
}
