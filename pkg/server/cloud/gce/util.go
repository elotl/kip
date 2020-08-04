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
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/elotl/kip/pkg/util/hash"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/api/googleapi"
)

const (
	gceAuthAPIEndpoint    = "https://www.googleapis.com/auth/"
	gceComputeAPIEndpoint = "https://www.googleapis.com/compute/v1/"
)

var labelValueRegex = regexp.MustCompile("[^a-z0-9-_]+")

func isNotFoundError(err error) bool {
	return isHTTPErrorCode(err, http.StatusNotFound)
}

func isHTTPErrorCode(err error, code int) bool {
	apiErr, ok := err.(*googleapi.Error)
	return ok && apiErr.Code == code
}

func CreateKipCellNetworkTag(controllerID string) string {
	return fmt.Sprintf("kip-%s", controllerID)
}

func getPodIpFromCIDR(cidr string) (string, error) {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", err
	}
	_, bits := ipNet.Mask.Size()
	if bits != 32 {
		return "", fmt.Errorf("cannot get pod ip over an ip range")
	}
	addr := ipNet.IP.String()
	return addr, nil
}

func (c *gceClient) createUnboundNodeNameTag() string {
	return fmt.Sprintf(
		"kip-node-%s-%d", c.nametag, time.Now().Unix())
}

func (c *gceClient) getLabelCompareFilter(mapKey, compare string) string {
	return fmt.Sprintf("labels.%s = %s", mapKey, compare)
}

func (c *gceClient) getNetworkURL() string {
	return gceComputeAPIEndpoint + strings.Join([]string{"projects", c.projectID, "global", "networks", c.vpcName},
		"/",
	)
}

func (c *gceClient) getSubnetworkURL() string {
	return gceComputeAPIEndpoint + strings.Join(
		[]string{"projects", c.projectID, "regions", c.region, "subnetworks", c.subnetName},
		"/",
	)
}

func (c *gceClient) getProjectURL() string {
	return gceComputeAPIEndpoint + strings.Join([]string{"projects", c.projectID},
		"/",
	)
}

// TOOD figure out better way to pull disk type
func (c *gceClient) getDiskTypeURL() string {
	return gceComputeAPIEndpoint + strings.Join(
		[]string{"projects", c.projectID, "zones", c.zone, "diskTypes", "pd-standard"},
		"/",
	)
}

func (c *gceClient) getDiskImageURL(project, image string) string {
	return gceComputeAPIEndpoint + strings.Join(
		[]string{"projects", project, "global", "images", image},
		"/",
	)
}

func (c *gceClient) getVPCURL() string {
	return gceComputeAPIEndpoint + strings.Join(
		[]string{"projects", c.projectID, "global", "networks", c.vpcName},
		"/",
	)
}

func (c *gceClient) getInstanceTypeURL(instanceType string) string {
	return gceComputeAPIEndpoint + strings.Join(
		[]string{"projects", c.projectID, "zones", c.zone, "machineTypes", instanceType},
		"/",
	)
}

func (c *gceClient) getAcceleratorTypeURL(acceleratorType string) string {
	return gceComputeAPIEndpoint + strings.Join(
		[]string{"projects", c.projectID, "zones", c.zone, "acceleratorTypes", acceleratorType},
		"/",
	)
}

func getServiceAccountScopes(scopes []string) []string {
	scopeLen := len(scopes)
	if scopeLen == 1 {
		return []string{gceAuthAPIEndpoint + scopes[0]}
	}
	scopeURLs := make([]string, scopeLen)
	for _, scope := range scopes {
		s := gceAuthAPIEndpoint + scope
		scopeURLs = append(scopeURLs, s)
	}
	return scopeURLs
}

func makeInstanceID(controllerID, nodeName string) string {
	nodeUUID, err := uuid.FromString(nodeName)
	var compressedID string
	if err != nil {
		compressedID = strings.Replace(nodeName, "-", "", -1)
		if len(compressedID) > 26 {
			compressedID = compressedID[:26]
		}
	} else {
		compressedID = hash.Base32EncodeNoPad(nodeUUID.Bytes())
	}

	return fmt.Sprintf("kip-%s-%s", controllerID, compressedID)
}

func replaceReservedLabelChars(label string) string {
	return labelValueRegex.ReplaceAllString(label, "-")
}
