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

package kubeconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type Kubeconfig struct {
	Config    *clientcmdapi.Config
	UserName  string
	tokenFile string
	mutex     *sync.Mutex
}

// Create a new Kubeconfig object based on a token and CA cert.
func NewFromToken(userName, clusterName, serverURL, tokenFile, rootCAFile string) (*Kubeconfig, error) {
	if clusterName == "" {
		clusterName = "default"
	}
	if serverURL == "" {
		host, port := os.Getenv("KUBERNETES_SERVICE_HOST"), os.Getenv("KUBERNETES_SERVICE_PORT")
		if len(host) == 0 || len(port) == 0 {
			return nil, fmt.Errorf("no Kubernetes master host or port set")
		}
		serverURL = fmt.Sprintf("https://%s:%s", host, port)
	}
	token, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		return nil, err
	}
	caCert, err := ioutil.ReadFile(rootCAFile)
	if err != nil {
		return nil, err
	}
	cfg := CreateFromToken(serverURL, clusterName, userName, caCert, token)
	return &Kubeconfig{
		mutex:     &sync.Mutex{},
		Config:    cfg,
		UserName:  userName,
		tokenFile: tokenFile,
	}, nil
}

// Create a new Kubeconfig object based on a kubeconfig file.
func LoadFromFile(filename string) (*Kubeconfig, error) {
	kc, err := clientcmd.LoadFromFile(filename)
	if err != nil {
		return nil, err
	}
	return &Kubeconfig{
		mutex:  &sync.Mutex{},
		Config: kc,
	}, nil
}

// Update our kubeconfig in case the token has changed. Note: Kubernetes will
// probably transition to serviceaccount tokens with limited lifetime. This
// function can be called periodically to update the config.
func (kc *Kubeconfig) Refresh() error {
	if kc.tokenFile == "" {
		// This was created from a kubeconfig file, not a token. No need to
		// refresh.
		return nil
	}
	token, err := ioutil.ReadFile(kc.tokenFile)
	if err != nil {
		return err
	}
	kc.mutex.Lock()
	defer kc.mutex.Unlock()
	kc.Config.AuthInfos[kc.UserName] = &clientcmdapi.AuthInfo{
		Token: string(token),
	}
	return nil
}

func CreateFromToken(serverURL, clusterName, userName string, caCert, token []byte) *clientcmdapi.Config {
	contextName := fmt.Sprintf("%s@%s", userName, clusterName)
	config := &clientcmdapi.Config{
		Clusters: map[string]*clientcmdapi.Cluster{
			clusterName: {
				Server:                   serverURL,
				CertificateAuthorityData: caCert,
			},
		},
		Contexts: map[string]*clientcmdapi.Context{
			contextName: {
				Cluster:  clusterName,
				AuthInfo: userName,
			},
		},
		AuthInfos:      map[string]*clientcmdapi.AuthInfo{},
		CurrentContext: contextName,
	}
	config.AuthInfos[userName] = &clientcmdapi.AuthInfo{
		Token: string(token),
	}
	return config
}

// Save kubeconfig to a file.
func (kc *Kubeconfig) WriteToFile(filename string) error {
	err := clientcmd.WriteToFile(*kc.Config, filename)
	if err != nil {
		return err
	}
	return nil
}

// Serialize kubeconfig.
func (kc *Kubeconfig) toJSON() ([]byte, error) {
	data, err := json.Marshal(kc.Config)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Create kubeconfig from JSON data. This is intended for testing, since we
// don't know the path of the token file the config has been created from.
func fromJSON(data []byte) (*Kubeconfig, error) {
	config := clientcmdapi.Config{}
	err := json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	userName := ""
	for un, _ := range config.AuthInfos {
		if un != "" {
			userName = un
		}
	}
	kc := Kubeconfig{
		Config:   &config,
		UserName: userName,
	}
	return &kc, nil
}
