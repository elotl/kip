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

package k8s

import (
	"fmt"
	"strings"

	"github.com/elotl/kip/pkg/util"
	"github.com/elotl/kip/pkg/util/kubeconfig"
	"github.com/elotl/node-cli/manager"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func CreateNetworkAgentKubeconfig(rm *manager.ResourceManager, serverURL, networkAgentSecret string) (*clientcmdapi.Config, error) {
	parts := strings.SplitN(networkAgentSecret, "/", 2)
	name := parts[len(parts)-1]
	namespace := "default"
	if len(parts) > 1 {
		namespace = parts[0]
	}
	secret, err := rm.GetSecret(name, namespace)
	if err != nil {
		return nil, util.WrapError(err, "retrieving %q/%q", namespace, name)
	}
	token, ok := secret.Data["token"]
	if !ok {
		return nil, fmt.Errorf("missing token in network agent secret")
	}
	cacrt, ok := secret.Data["ca.crt"]
	if !ok {
		return nil, fmt.Errorf("missing CA cert in network agent secret")
	}
	cfg := kubeconfig.CreateFromToken(serverURL, "", name, cacrt, token)
	return cfg, nil
}

func ValidateKubeconfig(config *clientcmdapi.Config) error {
	cc, err := clientcmd.NewDefaultClientConfig(*config, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		return util.WrapError(err, "validating network agent access")
	}
	clientset, err := kubernetes.NewForConfig(cc)
	if err != nil {
		return util.WrapError(err, "validating network agent access")
	}
	_, err = clientset.ServerVersion()
	if err != nil {
		return util.WrapError(err, "validating network agent access")
	}
	return nil
}
