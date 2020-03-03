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

package aws

import (
	"io/ioutil"
	"net/http"
	"time"
)

const (
	metadataTimeout = 2 * time.Second
	metadataURL     = "http://169.254.169.254/latest/meta-data/"
)

// This function grabs the ec2 metadata for the local machine that
// milpa is running on.  However, if milpa is not running within AWS
// the standard AWS metadata query hangs for about 15s. I tried
// modifiying the AWS HTTP client timeout but that didn't work so
// we'll just use our own client.
func GetMetadata(p string) (string, error) {
	if len(p) > 0 && p[0] == '/' {
		p = p[1:]
	}
	url := metadataURL + p
	timeout := time.Duration(metadataTimeout)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil || resp.StatusCode/200 != 1 {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}
