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
	"strings"
	"time"

	"cloud.google.com/go/compute/metadata"
	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/util"
	"google.golang.org/api/compute/v1"
	"k8s.io/klog"
)

const (
	defaultTimeout              = 20 * time.Second
	waitForRunningTimeout       = 3 * time.Minute
	controllerLabelKey          = "kip-controller-id"
	nameLabelKey                = "name"
	namespaceLabelKey           = "kip-namespace"
	minimumDiskSize       int64 = 10
	nametagLabelKey             = "kip-nametag"
	podNameLabelKey             = "kip-pod-name"
	statusOperationDone         = "DONE"
	statusInstanceRunning       = "RUNNING"
	apiRetries                  = 10
)

type gceClient struct {
	service              *compute.Service
	clientEmail          string
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
	gkeMetadata          map[string]string
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
	client.gkeMetadata = getGKEMetadata()

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
	return "", "", fmt.Errorf("Not implemented in gce")
}

func nilResponseError(call string) error {
	return fmt.Errorf("Nil response from GCE API %s RPC", call)
}

func (c *gceClient) getGlobalOperation(opName string) (*compute.Operation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	resp, err := c.service.GlobalOperations.Get(c.projectID, opName).Context(ctx).Do()
	if err != nil {
		return nil, util.WrapError(err, "Error could not retrieve global operation")
	}
	if resp == nil {
		return nil, nilResponseError("GlobalOperations.Get")
	}
	return resp, nil
}

func (c *gceClient) getRegionOperation(opName string) (*compute.Operation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	resp, err := c.service.RegionOperations.Get(c.projectID, c.region, opName).Context(ctx).Do()
	if err != nil {
		return nil, util.WrapError(err, "Error could not retrieve region operation")
	}
	if resp == nil {
		return nil, nilResponseError("RegionOperations.Get")
	}
	return resp, nil
}

func (c *gceClient) getZoneOperation(opName string) (*compute.Operation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	resp, err := c.service.ZoneOperations.Get(c.projectID, c.zone, opName).Context(ctx).Do()
	if err != nil {
		return nil, util.WrapError(err, "Error could not retrieve zone operation")
	}
	if resp == nil {
		return nil, nilResponseError("ZoneOperations.Get")
	}
	return resp, nil
}

func waitBackoff(i int) time.Duration {
	waitTimes := []time.Duration{1, 1, 2, 3, 5}
	if i < len(waitTimes) {
		return waitTimes[i] * time.Second
	}
	return waitTimes[len(waitTimes)-1] * time.Second
}

// In GCE operations will immediately succeed from a call, however that does not
// mean they have completed execution errorless. Here we wait for an operation
// to finish so we can check handle errors as we find necessary
func waitOnOperation(opName string, getOperation func(string) (*compute.Operation, error)) error {
	i := -1
	for {
		i += 1
		op, err := getOperation(opName)
		if err != nil {
			return err
		}

		if op.Status != statusOperationDone {
			time.Sleep(waitBackoff(i))
			continue
		}
		// check if the operation is not nil
		// if not nil it will have a *compute.OperationError
		if op.Error != nil {
			klog.Errorf("Operation %s was not successful", opName)
			var errors []string
			for _, e := range op.Error.Errors {
				errors = append(errors, e.Message)
			}
			return fmt.Errorf("Operation failed with error(s): %s", strings.Join(errors, ", "))
		}
		break
	}
	return nil
}
