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
	"os"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

func GetServerURLFromInCluster() string {
	config, err := rest.InClusterConfig()
	if err != nil {
		klog.Errorf("trying to determine API server URL: %v", err)
		return ""
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Errorf("trying to determine API server URL: %v", err)
		return ""
	}
	endpoints := clientset.CoreV1().Endpoints(v1.NamespaceDefault)
	ep, err := endpoints.Get("kubernetes", v1.GetOptions{})
	if err != nil {
		klog.Errorf("trying to determine API server URL: %v", err)
		return ""
	}
	for _, subset := range ep.Subsets {
		if len(subset.Addresses) > 0 && len(subset.Ports) > 0 {
			host := subset.Addresses[0].IP
			if host == "" {
				host = subset.Addresses[0].Hostname
			}
			if host == "" {
				klog.Errorf("empty host in kubernetes svc address")
				return ""
			}
			port := subset.Ports[0].Port
			return fmt.Sprintf("https://%s:%d", host, port)
		}
	}
	klog.Errorf("no endpoint found for kubernetes svc")
	return ""
}

func GetServerURL(kubeconfigPath string) string {
	if kubeconfigPath != "" {
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err == nil {
			klog.V(4).Infof("server URL from kubeconfig: %q", config.Host)
			return config.Host
		}
		klog.Errorf("building config from kubeconfig: %v", err)
	}
	if masterURI := os.Getenv("MASTER_URI"); masterURI != "" {
		klog.V(4).Infof("server URL from MASTER_URI: %q", masterURI)
		return masterURI
	}
	if serverURL := GetServerURLFromInCluster(); serverURL != "" {
		klog.V(4).Infof("server URL from kubernetes svc EP: %q", serverURL)
		return serverURL
	}
	return ""
}
