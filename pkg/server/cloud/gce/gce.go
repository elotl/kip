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
	"context"
	"fmt"
	"os"
	"runtime"
	"time"

	"cloud.google.com/go/compute/metadata"
	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/util"
	"google.golang.org/api/compute/v1"
)

const (
	defaultTimeout              = 20 * time.Second
	controllerLabelKey          = "kip-controller-id"
	nameLabelKey                = "name"
	namespaceLabelKey           = "kip-namespace"
	minimumDiskSize       int64 = 10
	nametagLabelKey             = "kip-nametag"
	podNameLabelKey             = "kip-pod-name"
	statusOperationDone         = "DONE"
	statusInstanceRunning       = "RUNNING"
)

func TODO() error {
	msg := "TODO: Not implemented yet!"
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		fName := "unknown"
		details := runtime.FuncForPC(pc)
		if details != nil {
			fName = details.Name()
		}
		msg = fmt.Sprintf("TODO: [%s.%d] %s not implemented yet", file, line, fName)
	}
	return fmt.Errorf(msg)
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
	vpcCIDRs             []string // Unsure if this is needed?
	subnetCIDR           string   // Unsure if this is needed?
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

	// Setup VPC Parameters
	if client.vpcName == "" {
		client.vpcName, err = client.detectCurrentVPC()
		if err != nil {
			return nil, err
		}
	}
	client.vpcCIDRs, err = client.getVPCRegionCIDRs(client.vpcName)
	if err != nil {
		return nil, err
	}

	// Setup subnet parameters
	if client.subnetName == "" {
		client.subnetName, client.subnetCIDR, err = client.autodetectSubnet()
		if err != nil {
			return nil, err
		}
	} else {
		client.subnetCIDR, err = client.getSubnetCIDR(client.subnetName)
		if err != nil {
			return nil, err
		}
	}

	client.cloudStatus, err = cloud.NewAZSubnetStatus(client)
	if err != nil {
		return nil, util.WrapError(err, "Error creating gce cloud status keeper")
	}

	return client, nil
}

// Try to get credentials from environment variables or from
// the environment the machine is running in
func serviceFromEnvironment() (*compute.Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	service, err := compute.NewService(ctx)
	return service, err
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
		Zone:            c.zone,
	}
}

func (c *gceClient) CloudStatusKeeper() cloud.StatusKeeper {
	return c.cloudStatus
}

func (c *gceClient) GetRegistryAuth() (string, string, error) {
	return "", "", TODO()
}

func nilResponseError(call string) error {
	return fmt.Errorf("Nil response from GCE API %s RPC", call)
}
