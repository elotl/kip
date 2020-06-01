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
	"fmt"
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/compute/metadata"
	"k8s.io/klog"
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

func extractZoneFromResponse(zone string) (string, error) {
	if len(zone) == 0 {
		return "", fmt.Errorf("Got empty zone response from metadata service")
	}
	if zone[len(zone)-1] == '/' {
		zone = zone[:len(zone)-1]
	}
	return zone[strings.LastIndex(zone, "/")+1:], nil
}

// Querying the zone from some images adds a trailing slash to the
// zone response.  e.g.
// projects/832569367454/zones/us-west1-a/
// vs
// projects/832569367454/zones/us-west1-a
// The latter breaks the metadata library that we use so lets handle
// it here.
func getZoneFromMetadata(c *metadata.Client) (string, error) {
	s, err := getMetadataTrimmed(c, "instance/zone")
	if err != nil {
		return "", err
	}
	return extractZoneFromResponse(s)
}

func getGKEMetadata() map[string]string {
	if !metadata.OnGCE() {
		return nil
	}
	keys := []string{
		"cluster-location",
		"cluster-name",
		"cluster-uid",
	}
	c := newMetadataClient()
	gkeMD := make(map[string]string)
	for _, k := range keys {
		url := "instance/attributes/" + k
		val, err := getMetadataTrimmed(c, url)
		if err != nil {
			// Pretty sure we can continue in this case. The cluster
			// will still function, if super important, daemonSets
			// might not run on our virtual node but we are running
			// them mostly for apprarances sake.  They should be
			// patched to not run on our kip nodes.
			klog.Warningln("unable to retrieve gke metadata key", k)
			continue
		}
		gkeMD[k] = val
	}
	return gkeMD
}
