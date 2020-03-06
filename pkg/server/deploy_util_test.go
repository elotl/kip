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
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/nodeclient"
	"github.com/elotl/cloud-instance-provider/pkg/util/k8s/eventrecorder"
	"github.com/kubernetes/kubernetes/pkg/kubelet/network/dns"
	"github.com/stretchr/testify/assert"
	"github.com/virtual-kubelet/node-cli/manager"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
)

func tarPkgToPackageFile(tarfile io.Reader) (map[string]packageFile, error) {
	gzr, err := gzip.NewReader(tarfile)
	if err != nil {
		return nil, err
	}
	defer gzr.Close()
	tr := tar.NewReader(gzr)

	tfContents := make(map[string]packageFile)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if header.Typeflag == tar.TypeReg {
			data := make([]byte, header.Size)
			read_so_far := int64(0)
			for read_so_far < header.Size {
				n, err := tr.Read(data[read_so_far:])
				if err == io.EOF {
					break
				}
				if err != nil {
					return nil, err
				}
				read_so_far += int64(n)
			}
			tfContents[header.Name[7:]] = packageFile{
				data: data,
				mode: int32(header.Mode),
			}
		}
	}
	return tfContents, nil
}

func TestMakeDeployPackage(t *testing.T) {
	contents := map[string]packageFile{
		"file1":         packageFile{data: []byte("file1"), mode: 0777},
		"path/to/file2": {data: []byte("file2"), mode: 0400},
	}
	buf, err := makeDeployPackage(contents)
	assert.NoError(t, err)
	tfContents, err := tarPkgToPackageFile(bufio.NewReader(buf))
	assert.NoError(t, err)
	assert.Equal(t, contents, tfContents)
}

func TestGetConfigMapFiles(t *testing.T) {
	trueVal := true
	readonlyVal := int32(0444)
	allPermsVal := int32(0777)
	simpleConfigMap := v1.ConfigMap{
		Data: map[string]string{
			"foo": "foocontent",
			"bar": "barcontent",
		},
		BinaryData: map[string][]byte{
			"zed": []byte("zedstuff"),
		},
	}

	tests := []struct {
		name          string
		vol           api.ConfigMapVolumeSource
		cm            v1.ConfigMap
		isErr         bool
		expectedFiles map[string]packageFile
	}{
		{
			name: "optional is skipped",
			vol: api.ConfigMapVolumeSource{
				Optional: &trueVal,
			},
			cm:            v1.ConfigMap{},
			isErr:         false,
			expectedFiles: map[string]packageFile{},
		},
		{
			name: "no items gets all items, default mode",
			vol: api.ConfigMapVolumeSource{
				Optional: &trueVal,
			},
			cm:    simpleConfigMap,
			isErr: false,
			expectedFiles: map[string]packageFile{
				"foo": packageFile{
					data: []byte("foocontent"),
					mode: defaultVolumeFileMode,
				},
				"bar": packageFile{
					data: []byte("barcontent"),
					mode: defaultVolumeFileMode,
				},
				"zed": packageFile{
					data: []byte("zedstuff"),
					mode: defaultVolumeFileMode,
				},
			},
		},
		{
			name: "only get some items, different default modes",
			vol: api.ConfigMapVolumeSource{
				Optional: &trueVal,
				Items: []api.KeyToPath{
					{
						Key:  "bar",
						Path: "path/to",
						Mode: &allPermsVal,
					},
					{
						Key: "zed",
					},
				},
				DefaultMode: &readonlyVal,
			},
			cm:    simpleConfigMap,
			isErr: false,
			expectedFiles: map[string]packageFile{
				"path/to": packageFile{
					data: []byte("barcontent"),
					mode: allPermsVal,
				},
				"zed": packageFile{
					data: []byte("zedstuff"),
					mode: readonlyVal,
				},
			},
		},
	}
	for _, tc := range tests {
		files, err := getConfigMapFiles(&tc.vol, &tc.cm)
		if tc.isErr {
			assert.Error(t, err, tc.name)
		} else {
			assert.NoError(t, err, tc.name)
			assert.Equal(t, tc.expectedFiles, files, tc.name)
		}
	}
}

func TestDeployVolumes(t *testing.T) {
	trueVal := true
	pod := api.GetFakePod()
	pod.Namespace = "default"
	testNode := api.GetFakeNode()
	tests := []struct {
		name          string
		volumes       []api.Volume
		configMap     *v1.ConfigMap
		secret        *v1.Secret
		expectedFiles map[string]packageFile
		isErr         bool
	}{
		{
			name: "optional packages are skipped",
			volumes: []api.Volume{
				{
					Name: "optional",
					VolumeSource: api.VolumeSource{
						ConfigMap: &api.ConfigMapVolumeSource{
							LocalObjectReference: api.LocalObjectReference{
								Name: "not-present",
							},
							Optional: &trueVal,
						},
					},
				},
			},
			expectedFiles: map[string]packageFile{},
			isErr:         false,
		},
		{
			name: "get configmap, single item",
			volumes: []api.Volume{
				{
					Name: "mytest",
					VolumeSource: api.VolumeSource{
						ConfigMap: &api.ConfigMapVolumeSource{
							LocalObjectReference: api.LocalObjectReference{
								Name: "test-config-map",
							},
							Items: []api.KeyToPath{
								{Key: "bar"},
							},
						},
					},
				},
			},
			configMap: &v1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-config-map",
					Namespace: "default",
				},
				Data: map[string]string{
					"foo": "abc",
					"bar": "123",
				},
			},
			expectedFiles: map[string]packageFile{
				"bar": packageFile{data: []byte("123"), mode: defaultVolumeFileMode},
			},
			isErr: false,
		},
		{
			name: "get secret, single item",
			volumes: []api.Volume{
				{
					Name: "mytest",
					VolumeSource: api.VolumeSource{
						Secret: &api.SecretVolumeSource{
							SecretName: "test-secret",
							Items: []api.KeyToPath{
								{Key: "bar"},
							},
						},
					},
				},
			},
			secret: &v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-secret",
					Namespace: "default",
				},
				Data: map[string][]byte{
					"foo": []byte("abc"), // abc -> YWJj
					"bar": []byte("123"), // 123 -> MTIz
				},
			},
			expectedFiles: map[string]packageFile{
				"bar": packageFile{data: []byte("123"), mode: defaultVolumeFileMode},
			},
			isErr: false,
		},
	}
	for _, tc := range tests {
		indexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
		if tc.configMap != nil {
			assert.Nil(t, indexer.Add(tc.configMap))
		}
		configMapLister := corev1listers.NewConfigMapLister(indexer)
		if tc.secret != nil {
			assert.Nil(t, indexer.Add(tc.secret))
		}
		secretLister := corev1listers.NewSecretLister(indexer)
		rm, err := manager.NewResourceManager(nil, secretLister, configMapLister, nil)
		if err != nil {
			t.Fatal(err)
		}

		// create the nodeClientFactory
		nc := nodeclient.NewMockItzoClientFactory()
		nc.DeployPackage = func(pod, name string, data io.Reader) error {
			tfContents, err := tarPkgToPackageFile(data)
			assert.NoError(t, err, tc.name)
			assert.Equal(t, tc.expectedFiles, tfContents, tc.name)
			return nil
		}
		pod.Spec.Volumes = tc.volumes
		err = deployPodVolumes(pod, testNode, rm, nc)
		if tc.isErr {
			assert.Error(t, err, tc.name)
		} else {
			assert.NoError(t, err, tc.name)
		}
	}
}

func createFakeDNSConfigurer(dnsIP, resolvconfPath, clusterDomain string) *dns.Configurer {
	loggingEventRecorder := eventrecorder.NewLoggingEventRecorder(4)
	nodeRef := &v1.ObjectReference{
		Kind:       "Node",
		APIVersion: "v1",
		Name:       "FakeNode",
	}
	clusterDNS := net.ParseIP(dnsIP)
	return dns.NewConfigurer(
		loggingEventRecorder,
		nodeRef,
		nil,
		[]net.IP{clusterDNS},
		clusterDomain,
		resolvconfPath,
	)
}

func stringPtr(s string) *string {
	return &s
}

func TestCreateResolvconf(t *testing.T) {
	resolverConfigF, err := ioutil.TempFile("", "resolv-conf-test")
	assert.NoError(t, err)
	resolvconfPath := resolverConfigF.Name()
	resolverConfigF.Close()
	defer os.Remove(resolvconfPath)
	defaultDNS := "9.8.7.6"
	defaultSearch := "foo.bar"
	defaultOptions := "ndots:3"
	defaultResolver := fmt.Sprintf(
		"nameserver %s\nsearch %s\noptions %s\n",
		defaultDNS,
		defaultSearch,
		defaultOptions,
	)
	ioutil.WriteFile(
		resolvconfPath,
		[]byte(defaultResolver),
		0644,
	)
	clusterDNS := "1.2.3.4"
	clusterDomain := "cluster.local"
	clusterOptions := "ndots:5"
	clusterResolver := fmt.Sprintf(
		"nameserver %s\nsearch .svc.%s svc.%s %s foo.bar\noptions %s\n",
		clusterDNS,
		clusterDomain,
		clusterDomain,
		clusterDomain,
		clusterOptions,
	)
	dnsConfigurer := createFakeDNSConfigurer(
		clusterDNS, resolvconfPath, clusterDomain)
	testCases := []struct {
		DNSPolicy   v1.DNSPolicy
		DNSConfig   *v1.PodDNSConfig
		HostNetwork bool
		Resolvconf  string
	}{
		{
			DNSPolicy:  v1.DNSClusterFirstWithHostNet,
			Resolvconf: clusterResolver,
		},
		{
			DNSPolicy:  v1.DNSClusterFirst,
			Resolvconf: clusterResolver,
		},
		{
			DNSPolicy:  v1.DNSDefault,
			Resolvconf: defaultResolver,
		},
		{
			DNSPolicy:  v1.DNSNone,
			Resolvconf: "",
		},
		{
			DNSPolicy:   v1.DNSClusterFirstWithHostNet,
			HostNetwork: true,
			Resolvconf:  clusterResolver,
		},
		{
			DNSPolicy:   v1.DNSClusterFirst,
			HostNetwork: true,
			Resolvconf:  defaultResolver,
		},
		{
			DNSPolicy: v1.DNSClusterFirst,
			DNSConfig: &v1.PodDNSConfig{
				Nameservers: []string{
					"11.11.11.11",
					"22.22.22.22",
				},
				Searches: []string{
					"a.b.c.d",
					"e.f.g.h",
				},
				Options: []v1.PodDNSConfigOption{
					{
						Name:  "ndots",
						Value: stringPtr("4"),
					},
				},
			},
			Resolvconf: "nameserver 1.2.3.4\nnameserver 11.11.11.11\nnameserver 22.22.22.22\nsearch .svc.cluster.local svc.cluster.local cluster.local foo.bar a.b.c.d e.f.g.h\noptions ndots:4\n",
		},
		{
			DNSPolicy: v1.DNSDefault,
			DNSConfig: &v1.PodDNSConfig{
				Nameservers: []string{
					"33.33.33.33",
				},
				Searches: []string{
					"i.j.k.l",
				},
				Options: []v1.PodDNSConfigOption{
					{
						Name:  "timeout",
						Value: stringPtr("10"),
					},
				},
			},
			Resolvconf: "nameserver 9.8.7.6\nnameserver 33.33.33.33\nsearch foo.bar i.j.k.l\noptions ndots:3 timeout:10\n",
		},
		{
			DNSPolicy: v1.DNSNone,
			DNSConfig: &v1.PodDNSConfig{
				Nameservers: []string{
					"44.44.44.44",
				},
				Searches: []string{
					"m.n.o.p",
				},
				Options: []v1.PodDNSConfigOption{
					{
						Name: "debug",
					},
					{
						Name:  "attempts",
						Value: stringPtr("5"),
					},
				},
			},
			Resolvconf: "nameserver 44.44.44.44\nsearch m.n.o.p\noptions debug attempts:5\n",
		},
	}
	for i, tc := range testCases {
		pod := &v1.Pod{}
		pod.Name = fmt.Sprintf("dnstest%d", i)
		pod.Spec.DNSPolicy = tc.DNSPolicy
		pod.Spec.DNSConfig = tc.DNSConfig
		pod.Spec.HostNetwork = tc.HostNetwork
		dnsconf, err := dnsConfigurer.GetPodDNS(pod)
		assert.NoError(t, err)
		resolvconf, err := createResolvconf(pod.Name, dnsconf)
		assert.NoError(t, err)
		msg := fmt.Sprintf("Test case %d: %+v", i+1, tc)
		assert.Equal(
			t,
			resolvconfToMap(tc.Resolvconf),
			resolvconfToMap(string(resolvconf)),
			msg,
		)
	}
}

func resolvconfToMap(conf string) map[string][]string {
	lines := strings.Split(conf, "\n")
	output := make(map[string][]string)
	for _, line := range lines {
		line = strings.Trim(strings.Replace(line, "	", " ", -1), " ")
		if line == "" {
			continue
		}
		words := strings.Split(line, " ")
		if len(words) < 1 {
			continue
		}
		k := words[0]
		v := words[1:]
		sort.Strings(v)
		output[k] = v
	}
	return output
}
