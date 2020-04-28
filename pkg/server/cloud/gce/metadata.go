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

package gce

import (
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/compute/metadata"
)

const (
	metadataTimeout = 4 * time.Second
	interfacesPath  = "instance/network-interfaces/"
)

func newMetadataClient() *metadata.Client {
	timeout := time.Duration(metadataTimeout)
	client := &http.Client{
		Timeout: timeout,
	}
	return metadata.NewClient(client)
}

func getMetadataTrimmed(c *metadata.Client, suffix string) (s string, err error) {
	s, err = c.Get(suffix)
	s = strings.TrimSpace(s)
	return
}

func getMetadataLines(c *metadata.Client, suffix string) ([]string, error) {
	j, err := c.Get(suffix)
	if err != nil {
		return nil, err
	}
	s := strings.Split(strings.TrimSpace(j), "\n")
	for i := range s {
		s[i] = strings.TrimSpace(s[i])
	}
	return s, nil
}
