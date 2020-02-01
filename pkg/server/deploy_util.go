package server

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"path/filepath"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/nodeclient"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/virtual-kubelet/node-cli/manager"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
			archivePath = filepath.Join(item.Path, item.Key)
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
			//data, err = base64.StdEncoding.DecodeString(string(binaryData))
			// if err != nil {
			// 	msg := fmt.Sprintf("volume %s items %s/%s references improperly formatted key %s: %v", secVol.SecretName, sec.Namespace, sec.Name, item.Key, err)
			// 	if optional {
			// 		glog.Warning(msg)
			// 		continue
			// 	}
			// 	return nil, fmt.Errorf(msg)
			// }
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
			archivePath = filepath.Join(item.Path, item.Key)
		}
		packageItems[archivePath] = packageFile{
			data: data,
			mode: mode,
		}
	}
	return packageItems, nil
}

func deployPodVolumes(pod *api.Pod, node *api.Node, rm *manager.ResourceManager, nodeClientFactory nodeclient.ItzoClientFactoryer) error {
	client := nodeClientFactory.GetClient(node.Status.Addresses)
	for _, vol := range pod.Spec.Volumes {
		if vol.ConfigMap != nil {
			optional := vol.ConfigMap.Optional != nil && *vol.ConfigMap.Optional
			// get the configmap
			configMap, err := rm.GetConfigMap(vol.ConfigMap.Name, pod.Namespace)
			if err != nil {
				if !(errors.IsNotFound(err) && optional) {
					return util.WrapError(err, "Couldn't get configMap %v/%v", pod.Namespace, vol.ConfigMap.Name)
				}
				configMap = &v1.ConfigMap{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: pod.Namespace,
						Name:      vol.ConfigMap.Name,
					},
				}
			}
			packageFiles, err := getConfigMapFiles(vol.ConfigMap, configMap)
			if err != nil {
				return util.WrapError(err, "couldn't get configMap payload %v/%v", pod.Namespace, vol.ConfigMap.Name)
			}
			payload, err := makeDeployPackage(packageFiles)
			if err != nil {
				return util.WrapError(err, "error creating tar.gz package %s for %s", vol.Name, pod.Name)
			}
			client.Deploy(pod.Name, vol.Name, bufio.NewReader(payload))
			if err != nil {
				return util.WrapError(err, "error deploying package %s to %s", vol.Name, pod.Name)
			}
		} else if vol.Secret != nil {
			optional := vol.Secret.Optional != nil && *vol.Secret.Optional
			secret, err := rm.GetSecret(vol.Secret.SecretName, pod.Namespace)
			if err != nil {
				if !(errors.IsNotFound(err) && optional) {
					return util.WrapError(err, "Couldn't get secret %v/%v", pod.Namespace, vol.Secret.SecretName)
				}
				secret = &v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: pod.Namespace,
						Name:      vol.Secret.SecretName,
					},
				}
			}
			packageFiles, err := getSecretFiles(vol.Secret, secret)
			if err != nil {
				return util.WrapError(err, "couldn't get secret payload %v/%v", pod.Namespace, vol.Secret.SecretName)
			}
			payload, err := makeDeployPackage(packageFiles)
			if err != nil {
				return util.WrapError(err, "error creating tar.gz package %s for %s", vol.Name, pod.Name)
			}
			client.Deploy(pod.Name, vol.Name, bufio.NewReader(payload))
			if err != nil {
				return util.WrapError(err, "error deploying package %s to %s", vol.Name, pod.Name)
			}
		}
	}
	return nil
}
