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

package main

import "github.com/spf13/pflag"

type ServerConfig struct {
	DebugServer            bool
	NetworkAgentSecret     string
	NetworkAgentKubeConfig string
	ClusterDNS             string
	OverrideInstanceData   bool
	InstanceDataPath       string
}

func (c *ServerConfig) FlagSet() *pflag.FlagSet {
	flags := pflag.NewFlagSet("serverconfig", pflag.ContinueOnError)
	flags.BoolVar(&c.DebugServer, "debug-server", c.DebugServer, "Enable a listener in the server for inspecting internal kip structures.")
	flags.StringVar(&c.NetworkAgentSecret, "network-agent-secret", c.NetworkAgentSecret, "Service account secret for the cell network agent, in the form of <namespace>/<name>")
	flags.StringVar(&c.NetworkAgentKubeConfig, "network-agent-kubeconfig", c.NetworkAgentKubeConfig, "Network agent kubeconfig file, mutually exclusive with --network-agent-secret")
	flags.StringVar(&c.ClusterDNS, "cluster-dns", c.ClusterDNS, "Default cluster DNS server to use; if not specified, the kube-system/kube-dns service IP will be used")
	flags.BoolVar(&c.OverrideInstanceData, "override-instance-data", c.OverrideInstanceData, "set this to true to specify custom instance-data path")
	flags.StringVar(&c.InstanceDataPath, "instance-data-path", c.InstanceDataPath, "instance-data path; ignored if override-instance-data isn't set to true")
	return flags
}
