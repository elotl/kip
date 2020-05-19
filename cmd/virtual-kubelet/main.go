// Copyright © 2017 The virtual-kubelet authors
// Copyright © 2020 Elotl Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"net"
	"os"
	"strings"

	"github.com/elotl/kip/pkg/klog"
	"github.com/elotl/kip/pkg/server"
	"github.com/elotl/kip/pkg/util/habitat"
	"github.com/elotl/kip/pkg/util/k8s"
	cli "github.com/virtual-kubelet/node-cli"
	opencensuscli "github.com/virtual-kubelet/node-cli/opencensus"
	"github.com/virtual-kubelet/node-cli/opts"
	"github.com/virtual-kubelet/node-cli/provider"
	"github.com/virtual-kubelet/virtual-kubelet/log"
	"github.com/virtual-kubelet/virtual-kubelet/trace"
	"github.com/virtual-kubelet/virtual-kubelet/trace/opencensus"
)

var (
	buildVersion = "N/A"
	buildTime    = "N/A"
	k8sVersion   = "v1.17.0" // This should follow the version of k8s.io/kubernetes we are importing
)

func getInternalIP() string {
	internalIP := os.Getenv("VKUBELET_POD_IP")
	if internalIP != "" {
		return internalIP
	}
	internalIP = habitat.GetMyIP()
	if internalIP != "" {
		return internalIP
	}
	ips := habitat.GetIPAddresses()
	if len(ips) > 0 {
		return ips[0]
	}
	return ""
}

func main() {
	ctx := cli.ContextWithCancelOnSignal(context.Background())

	log.L = klog.NewKlogAdapter()

	trace.T = opencensus.Adapter{}
	traceConfig := opencensuscli.Config{
		AvailableExporters: map[string]opencensuscli.ExporterInitFunc{
			"ocagent": initOCAgent,
		},
	}

	serverConfig := &ServerConfig{}

	o, err := opts.FromEnv()
	if err != nil {
		log.G(ctx).Fatal(err)
	}
	o.Provider = "kip"
	o.Version = strings.Join([]string{k8sVersion, "vk", buildVersion}, "-")
	o.PodSyncWorkers = 10

	internalIP := getInternalIP()
	if internalIP == "" {
		log.G(ctx).Fatal("unable to determine internal IP address")
	}

	cert := os.Getenv("APISERVER_CERT_LOCATION")
	key := os.Getenv("APISERVER_KEY_LOCATION")
	ips := []net.IP{net.ParseIP(internalIP)}
	if err := ensureCert(o.NodeName, cert, key, ips); err != nil {
		log.G(ctx).Fatal(err)
	}

	node, err := cli.New(ctx,
		cli.WithBaseOpts(o),
		cli.WithCLIVersion(buildVersion, buildTime),
		cli.WithProvider("kip",
			func(cfg provider.InitConfig) (provider.Provider, error) {
				serverURL := k8s.GetServerURL(o.KubeConfigPath)
				if serverURL == "" {
					log.G(ctx).Fatal("can't determine API server URL, " +
						"please set --kubeconfig or MASTER_URI")
				}
				return server.NewInstanceProvider(
					cfg.ConfigPath,
					cfg.NodeName,
					internalIP,
					serverURL,
					serverConfig.NetworkAgentSecret,
					serverConfig.ClusterDNS,
					cfg.KubeClusterDomain,
					cfg.DaemonPort,
					serverConfig.DebugServer,
					cfg.ResourceManager,
					ctx.Done(),
				)
			}),
		cli.WithPersistentFlags(traceConfig.FlagSet()),
		cli.WithPersistentFlags(serverConfig.FlagSet()),
		cli.WithPersistentPreRunCallback(func() error {
			return opencensuscli.Configure(ctx, &traceConfig, o)
		}),
	)

	if err != nil {
		log.G(ctx).Fatal(err)
	}

	if err := node.Run(ctx); err != nil {
		log.G(ctx).Fatal(err)
	}
}
