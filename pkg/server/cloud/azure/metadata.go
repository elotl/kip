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

package azure

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/elotl/kip/pkg/util"
	"k8s.io/klog"
)

const (
	metadataTimeout = 4 * time.Second
	metadataURL     = "http://169.254.169.254/metadata/" //instance?api-version=2017-08-01"
)

// This function grabs the azuremetadata for the local machine that
// kip is running on.  Times out after a couple of seconds
func GetMetadata(p string) (string, error) {
	if len(p) > 0 && p[0] == '/' {
		p = p[1:]
	}
	url := metadataURL + p + "?api-version=2017-08-01"
	timeout := time.Duration(metadataTimeout)
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", util.WrapError(err, "Invalid metadata request")
	}
	req.Header.Set("Metadata", "true")
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode/200 != 1 {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func getMetadataInstanceName() (string, string) {
	p := "instance/compute"
	vm := struct {
		ResourceGroupName string `json:"resourceGroupName"`
		Name              string `json:"name"`
	}{}
	data, err := GetMetadata(p)

	if err != nil {
		return "", ""
	}
	err = json.Unmarshal([]byte(data), &vm)
	if err != nil {
		klog.Errorln("Could not unmarshal azure instance metadata", err.Error())
		return "", ""
	}
	return vm.ResourceGroupName, vm.Name
}
