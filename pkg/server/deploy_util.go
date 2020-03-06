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

package server

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/nodeclient"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/kubernetes/kubernetes/pkg/kubelet/network/dns"
	"github.com/virtual-kubelet/node-cli/manager"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
	"k8s.io/klog"
)

const (
	defaultVolumeFileMode = int32(0644)
)

type packageFile struct {
	data []byte
	mode int32
}

// Creates a tar.gz buffer filled with the package files
func makeDeployPackage(contents map[string]packageFile) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	for path, file := range contents {
		tarFilepath := filepath.Join(".", "ROOTFS", path)
		hdr := &tar.Header{
			Name:     tarFilepath,
			Mode:     int64(file.mode),
			Size:     int64(len(file.data)),
			Typeflag: byte(tar.TypeReg),
			Uid:      0,
			Gid:      0,
		}
		err := tw.WriteHeader(hdr)
		if err != nil {
			return nil, err
		}
		_, err = tw.Write(file.data)
		if err != nil {
			return nil, err
		}
	}
	if err := tw.Close(); err != nil {
		return nil, err
	}
	if err := gw.Close(); err != nil {
		return nil, err
	}
	return &buf, nil
}

func getConfigMapFiles(cmVol *api.ConfigMapVolumeSource, cm *v1.ConfigMap) (map[string]packageFile, error) {
	packageItems := make(map[string]packageFile)
	defaultMode := defaultVolumeFileMode
	if cmVol.DefaultMode != nil {
		defaultMode = *cmVol.DefaultMode
	}
	optional := cmVol.Optional != nil && *cmVol.Optional
	var items []api.KeyToPath
	if len(cmVol.Items) == 0 {
		items = make([]api.KeyToPath, 0, len(cm.Data)+len(cm.BinaryData))
		for k := range cm.Data {
			items = append(items, api.KeyToPath{Key: k})
		}
		for k := range cm.BinaryData {
			items = append(items, api.KeyToPath{Key: k})
		}
	} else {
		items = cmVol.Items
	}

	for _, item := range items {
		var data []byte
		if stringData, ok := cm.Data[item.Key]; ok {
			data = []byte(stringData)
		} else if binaryData, ok := cm.BinaryData[item.Key]; ok {
			data = binaryData
		} else {
			if optional {
				continue
			}
			return nil, fmt.Errorf("volume %s items %s/%s references non-existent config key: %s", cmVol.Name, cm.Namespace, cm.Name, item.Key)
		}
		mode := defaultMode
		if item.Mode != nil {
			mode = *item.Mode
		}
		archivePath := item.Key
		if item.Path != "" {
			archivePath = item.Path
		}
		packageItems[archivePath] = packageFile{
			data: data,
			mode: mode,
		}
	}
	return packageItems, nil
}

func getSecretFiles(secVol *api.SecretVolumeSource, sec *v1.Secret) (map[string]packageFile, error) {
	packageItems := make(map[string]packageFile)
	defaultMode := defaultVolumeFileMode
	if secVol.DefaultMode != nil {
		defaultMode = *secVol.DefaultMode
	}
	optional := secVol.Optional != nil && *secVol.Optional
	var items []api.KeyToPath
	if len(secVol.Items) == 0 {
		items = make([]api.KeyToPath, 0, len(sec.Data))
		for k := range sec.Data {
			items = append(items, api.KeyToPath{Key: k})
		}
	} else {
		items = secVol.Items
	}

	for _, item := range items {
		var data []byte
		if binaryData, ok := sec.Data[item.Key]; ok {
			data = binaryData
		} else {
			if optional {
				continue
			}
			return nil, fmt.Errorf("volume %s items %s/%s references non-existent config key: %s", secVol.SecretName, sec.Namespace, sec.Name, item.Key)
		}
		mode := defaultMode
		if item.Mode != nil {
			mode = *item.Mode
		}
		archivePath := item.Key
		if item.Path != "" {
			archivePath = item.Path
		}
		packageItems[archivePath] = packageFile{
			data: data,
			mode: mode,
		}
	}
	return packageItems, nil
}

func getConfigMapVolumeFiles(namespace string, cmVol *api.ConfigMapVolumeSource, rm *manager.ResourceManager) (map[string]packageFile, error) {
	optional := cmVol.Optional != nil && *cmVol.Optional
	// get the configmap
	configMap, err := rm.GetConfigMap(cmVol.Name, namespace)
	if err != nil {
		if !(errors.IsNotFound(err) && optional) {
			return nil, util.WrapError(err, "Couldn't get configMap %v/%v", namespace, cmVol.Name)
		}
		configMap = &v1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace,
				Name:      cmVol.Name,
			},
		}
	}
	return getConfigMapFiles(cmVol, configMap)
}

func getSecretVolumeFiles(namespace string, secVol *api.SecretVolumeSource, rm *manager.ResourceManager) (map[string]packageFile, error) {
	optional := secVol.Optional != nil && *secVol.Optional
	secret, err := rm.GetSecret(secVol.SecretName, namespace)
	if err != nil {
		if !(errors.IsNotFound(err) && optional) {
			return nil, util.WrapError(err, "Couldn't get secret %v/%v", namespace, secVol.SecretName)
		}
		secret = &v1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace,
				Name:      secVol.SecretName,
			},
		}
	}
	return getSecretFiles(secVol, secret)
}

func getProjectedVolumeFiles(namespace string, vol *api.ProjectedVolumeSource, rm *manager.ResourceManager) (map[string]packageFile, error) {
	defaultMode := api.ProjectedVolumeSourceDefaultMode
	if vol.DefaultMode != nil {
		defaultMode = *vol.DefaultMode
	}
	allPackageFiles := make(map[string]packageFile)
	for _, src := range vol.Sources {
		var (
			packageFiles map[string]packageFile
			err          error
		)
		if src.ConfigMap != nil {
			vol := &api.ConfigMapVolumeSource{
				LocalObjectReference: src.ConfigMap.LocalObjectReference,
				Items:                src.ConfigMap.Items,
				DefaultMode:          &defaultMode,
				Optional:             src.ConfigMap.Optional,
			}
			packageFiles, err = getConfigMapVolumeFiles(namespace, vol, rm)
			if err != nil {
				return nil, util.WrapError(err, "couldn't get projected configMap payload %v/%v", namespace, src.ConfigMap.Name)
			}
		} else if src.Secret != nil {
			vol := &api.SecretVolumeSource{
				SecretName:  src.Secret.LocalObjectReference.Name,
				Items:       src.Secret.Items,
				DefaultMode: &defaultMode,
				Optional:    src.Secret.Optional,
			}
			packageFiles, err = getSecretVolumeFiles(namespace, vol, rm)
			if err != nil {
				return nil, util.WrapError(err, "couldn't get projected secret payload %v/%v", namespace, src.Secret.Name)
			}
		}
		for k, v := range packageFiles {
			allPackageFiles[k] = v
		}
	}
	return allPackageFiles, nil
}

func deployPodVolumes(pod *api.Pod, node *api.Node, rm *manager.ResourceManager, nodeClientFactory nodeclient.ItzoClientFactoryer) error {
	client := nodeClientFactory.GetClient(node.Status.Addresses)
	for _, vol := range pod.Spec.Volumes {
		var (
			packageFiles map[string]packageFile
			err          error
		)
		if vol.ConfigMap != nil {
			packageFiles, err = getConfigMapVolumeFiles(pod.Namespace, vol.ConfigMap, rm)
			if err != nil {
				return util.WrapError(err, "couldn't get configMap payload %v/%v", pod.Namespace, vol.ConfigMap.Name)
			}
		} else if vol.Secret != nil {
			packageFiles, err = getSecretVolumeFiles(pod.Namespace, vol.Secret, rm)
			if err != nil {
				return util.WrapError(err, "couldn't get secret payload %v/%v", pod.Namespace, vol.Secret.SecretName)
			}
		} else if vol.Projected != nil {
			packageFiles, err = getProjectedVolumeFiles(pod.Namespace, vol.Projected, rm)
			if err != nil {
				return err
			}
		}
		// Deploy empty packages as well since they might be
		// referenced in a container (but will have no data)
		payload, err := makeDeployPackage(packageFiles)
		if err != nil {
			return util.WrapError(err, "error creating tar.gz package %s for %s", vol.Name, pod.Name)
		}
		err = client.Deploy(pod.Name, vol.Name, bufio.NewReader(payload))
		if err != nil {
			return util.WrapError(err, "error deploying package %s to %s", vol.Name, pod.Name)
		}
	}
	return nil
}

func deployNetworkAgentToken(cfg *clientcmdapi.Config, pod *api.Pod, node *api.Node, nodeClientFactory nodeclient.ItzoClientFactoryer) error {
	if cfg == nil {
		klog.V(4).Infof("no network agent kubeconfig provided for %s", pod.Name)
		return nil
	}
	data, err := clientcmd.Write(*cfg)
	if err != nil {
		return util.WrapError(err,
			"error serializing network agent kubeconfig for %s", pod.Name)
	}
	packageFiles := map[string]packageFile{
		"kubeconfig/kubeconfig": {
			data: data,
			mode: 0600,
		},
	}
	payload, err := makeDeployPackage(packageFiles)
	if err != nil {
		return util.WrapError(err,
			"error creating kubeconfig package for %s", pod.Name)
	}
	client := nodeClientFactory.GetClient(node.Status.Addresses)
	err = client.Deploy(pod.Name, "kubeconfig", bufio.NewReader(payload))
	if err != nil {
		return util.WrapError(err,
			"error deploying kubeconfig package for %s", pod.Name)
	}
	return nil
}

func deployResolvconf(pod *api.Pod, node *api.Node, dnsConfigurer *dns.Configurer, nodeClientFactory nodeclient.ItzoClientFactoryer) error {
	if dnsConfigurer == nil {
		return fmt.Errorf("no DNS configurer")
	}
	client := nodeClientFactory.GetClient(node.Status.Addresses)
	k8spod, err := milpaToK8sPod("", "", pod)
	if err != nil {
		return util.WrapError(err, "converting pod to generate DNS config")
	}
	dnsconf, err := dnsConfigurer.GetPodDNS(k8spod)
	if err != nil {
		return util.WrapError(err, "creating pod DNS config")
	}
	data, err := createResolvconf(pod.Name, dnsconf)
	if err != nil {
		return util.WrapError(err, "creating pod resolv.conf")
	}
	payload, err := makeDeployPackage(map[string]packageFile{
		"/etc/resolv.conf": packageFile{
			data: data,
			mode: 0644,
		},
	})
	if err != nil {
		return util.WrapError(err, "creating pod resolv.conf package")
	}
	err = client.Deploy(pod.Name, resolvconfVolumeName, bufio.NewReader(payload))
	if err != nil {
		return util.WrapError(
			err, "error deploying resolv.conf package to %s", pod.Name)
	}
	return nil
}

func createResolvconf(podName string, dnsconf *runtimeapi.DNSConfig) ([]byte, error) {
	buf := bytes.Buffer{}
	for _, srv := range dnsconf.Servers {
		_, err := buf.WriteString(fmt.Sprintf("nameserver %s\n", srv))
		if err != nil {
			return nil, util.WrapError(
				err, "creating DNS config for pod %q", podName)
		}
	}
	search := strings.Join(dnsconf.Searches, " ")
	if len(dnsconf.Searches) > 0 {
		_, err := buf.WriteString(fmt.Sprintf("search %s\n", search))
		if err != nil {
			return nil, util.WrapError(
				err, "creating DNS config for pod %q", podName)
		}
	}
	options := strings.Join(dnsconf.Options, " ")
	if len(dnsconf.Options) > 0 {
		_, err := buf.WriteString(fmt.Sprintf("options %s\n", options))
		if err != nil {
			return nil, util.WrapError(
				err, "creating DNS config for pod %q", podName)
		}
	}
	return buf.Bytes(), nil
}

func deployEtcHosts(pod *api.Pod, node *api.Node, dnsConfigurer *dns.Configurer, nodeClientFactory nodeclient.ItzoClientFactoryer) error {
	if dnsConfigurer == nil {
		return fmt.Errorf("no DNS configurer")
	}
	client := nodeClientFactory.GetClient(node.Status.Addresses)
	namespace, podName := util.SplitNamespaceAndName(pod.Name)
	podIPs := []string{api.GetPodIP(node.Status.Addresses)}
	useHostNetwork := api.IsHostNetwork(pod.Spec.SecurityContext)
	data, err := util.CreateEtcHosts(
		dnsConfigurer,
		podName,
		namespace,
		pod.Spec.Hostname,
		pod.Spec.Subdomain,
		podIPs,
		pod.Spec.HostAliases,
		useHostNetwork)
	if err != nil {
		return util.WrapError(err, "creating pod /etc/hosts")
	}
	payload, err := makeDeployPackage(map[string]packageFile{
		"/etc/hosts": packageFile{
			data: data,
			mode: 0644,
		},
	})
	if err != nil {
		return util.WrapError(err, "creating pod /etc/hosts package")
	}
	err = client.Deploy(pod.Name, etchostsVolumeName, bufio.NewReader(payload))
	if err != nil {
		return util.WrapError(
			err, "error deploying /etc/hosts package to %s", pod.Name)
	}
	return nil
}
