package main

import (
	"fmt"

	"github.com/elotl/kip/pkg/util/k8s"
	"github.com/elotl/node-cli/manager"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func maybeLoadKubeConfig(path string) (*clientcmdapi.Config, error) {
	if path == "" {
		return nil, fmt.Errorf("no kubeconfig provided; trying to continue")
	}
	kubeConfig, err := clientcmd.LoadFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("loading kubeconfig %s: %v; trying to continue", path, err)
	}
	return kubeConfig, nil
}

func loadOrCreateKubeConfig(path, mainKubeConfig, saSecret string, rm *manager.ResourceManager) (*clientcmdapi.Config, error) {
	if path == "" && saSecret == "" {
		return nil, fmt.Errorf("no kubeconfig or service account secret provided")
	}
	if path != "" {
		if saSecret != "" {
			return nil, fmt.Errorf("kubeconfig and service account secret are mutually exclusive")
		}
		kubeConfig, err := clientcmd.LoadFromFile(path)
		if err != nil {
			return nil, fmt.Errorf("unable to load kubeconfig %s: %v", path, err)
		}
		return kubeConfig, nil
	}
	kubeConfig, err := k8s.CreateKubeConfigFromSecret(rm, mainKubeConfig, saSecret)
	if err != nil {
		return nil, fmt.Errorf("creating kubeconfig from service account secret: %v", err)
	}
	return kubeConfig, nil
}
