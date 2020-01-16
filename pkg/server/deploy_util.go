package server

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"path/filepath"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/manager"
	"github.com/elotl/cloud-instance-provider/pkg/nodeclient"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	defaultMode := int32(0644)
	if cmVol.DefaultMode != nil {
		defaultMode = *cmVol.DefaultMode
	}
	optional := cmVol.Optional != nil && *cmVol.Optional
	var items []api.KeyToPath
	if len(cmVol.Items) == 0 {
		items = make([]api.KeyToPath, 0, len(cm.Data)+len(cm.BinaryData))
		for k, _ := range cm.Data {
			items = append(items, api.KeyToPath{Key: k})
		}
		for k, _ := range cm.BinaryData {
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

func makeSecretPackage() error {
	return nil
}

func deployVolumes(pod *api.Pod, node *api.Node, rm *manager.ResourceManager, nodeClientFactory nodeclient.ItzoClientFactoryer) error {
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
			fmt.Println("secret volumes not implemented yet")
		}
	}
	return nil
}
