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
	"os"

	"cloud.google.com/go/compute/metadata"
	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/cloud"
	"google.golang.org/api/compute/v1"
)

func NI() error {
	return fmt.Errorf("Not implemented yet!")
}

type gceClient struct {
	service              *compute.Service
	controllerID         string
	nametag              string
	projectID            string
	region               string
	zone                 string
	vpcName              string
	subnetName           string
	vpcCIDR              string // Unsure if this is needed?
	subnetCIDR           string // Unsure if this is needed?
	usePublicIPs         bool
	bootSecurityGroupIDs []string
	cloudStatus          *cloud.AZSubnetStatus
}

func NewGCEClient(controllerID, nametag, projectID string, opts ...ClientOption) (*gceClient, error) {
	client := &gceClient{
		service:      nil,
		controllerID: controllerID,
		nametag:      nametag,
		projectID:    projectID,
		usePublicIPs: true,
	}
	var err error
	if client.projectID == "" {
		client.projectID, err = client.autodetectProject()
		if err != nil {
			return nil, err
		}
	}
	for _, opt := range opts {
		err = opt.Apply(client)
		if err != nil {
			return nil, err
		}
	}
	if client.service == nil {
		client.service, err = serviceFromEnvironment()
		if err != nil {
			return nil, err
		}
	}
	if client.zone == "" {
		client.region, client.zone, err = client.autodetectRegionAndZone()
		if err != nil {
			return nil, err
		}
	}
	err = client.setupVPC(client.vpcName)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Try to get credentials from environment variables or from
// the environment the machine is running in
func serviceFromEnvironment() (*compute.Service, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (c *gceClient) autodetectProject() (string, error) {
	envVal := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if envVal != "" {
		return envVal, nil
	} else if metadata.OnGCE() {
		md := newMetadataClient()
		projectID, _ := md.ProjectID()
		if projectID != "" {
			return projectID, nil
		}
	}
	return "", fmt.Errorf("Could not get GCE project ID from environment or from instance metadata service")
}

func (c *gceClient) GetAttributes() cloud.CloudAttributes {
	return cloud.CloudAttributes{
		DiskProductName: api.StandardPersistentDisk,
		FixedSizeVolume: false,
		Provider:        cloud.ProviderGCE,
		Region:          c.region,
	}
}

func (c *gceClient) CloudStatusKeeper() cloud.StatusKeeper {
	return c.cloudStatus
}

func (c *gceClient) GetRegistryAuth() (string, string, error) {
	return "", "", NI()
}
