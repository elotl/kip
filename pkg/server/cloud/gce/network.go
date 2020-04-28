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
	"strings"

	"cloud.google.com/go/compute/metadata"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/util"
)

func zoneToRegion(zone string) (string, error) {
	parts := strings.Split(zone, "-")
	if len(parts) != 3 {
		return "", fmt.Errorf("unknown zone format, expecting geo-region-zone format")
	}
	return fmt.Sprintf("%s-%s", parts[0], parts[1]), nil
}

func (c *gceClient) ConnectWithPublicIPs() bool {
	if !c.usePublicIPs {
		return false
	} else {
		return !metadata.OnGCE()
	}
}

func (c *gceClient) setupVPC(vpcName string) error {
	if vpcName == "" {
		detectedVPC, err := c.detectCurrentVPC()
		if err != nil {
			vpcError := fmt.Errorf("could not autodetect VPC, please set the cloud vpc name in provider.yaml. The error was: %s", err)
			return vpcError
		} else {
			c.vpcName = detectedVPC
		}
	}
	// make sure the VPC exists
	ctx := context.Background()
	resp, err := c.service.Networks.Get(c.projectID, vpcName).Context(ctx).Do()
	if err != nil {
		return util.WrapError(err, "error querying VPC %s in GCP", vpcName)
	}
	c.vpcName = resp.Name
	c.vpcCIDR = resp.IPv4Range

	return nil
}

func (c *gceClient) detectCurrentVPC() (string, error) {
	// if we're inside GCE, get the instance's VPC name
	if !metadata.OnGCE() {
		return "", fmt.Errorf("instance is not running inside GCE, could not determine VPC name automatically")
	}
	md := newMetadataClient()
	interfaces, err := getMetadataLines(md, interfacesPath)
	if err != nil {
		return "", util.WrapError(err, "error querying instance interfaces from GCE metadata service")
	}
	if len(interfaces) == 0 {
		return "", fmt.Errorf("could not find host instance's interface in GCE metadata service")
	} else if len(interfaces) > 1 {
		return "", fmt.Errorf("multiple interfaces connected to the host instance")
	}
	iface := interfaces[0]
	path := fmt.Sprintf("%s%snetwork", interfacesPath, iface)
	fullNetwork, err := getMetadataTrimmed(md, path)
	if err != nil {
		return "", util.WrapError(err, "error querying instance network from GCE metadata service")
	}
	parts := strings.Split(fullNetwork, "/")
	if len(parts) <= 0 {
		return "", fmt.Errorf("improperly formatted VPC network name returned from GCE metadata service: %s", fullNetwork)
	}
	return parts[len(parts)-1], nil
}

func (c *gceClient) autodetectRegionAndZone() (string, string, error) {
	if !metadata.OnGCE() {
		return "", "", fmt.Errorf("instance is not running inside GCE, could not determine zone automatically. Please specify the zone in cloud.gce.zone in provider.yaml")
	}

	md := newMetadataClient()
	zone, err := md.Zone()
	if err != nil {
		return "", "", util.WrapError(err, "error getting zone from GCE metadata service")
	}
	region, err := zoneToRegion(zone)
	if err != nil {
		return "", "", util.WrapError(err, "error parsing zone from GCE metadata service")
	}
	return region, zone, nil
}

func (c *gceClient) GetSubnets() ([]cloud.SubnetAttributes, error) {
	sns := []cloud.SubnetAttributes{{
		Name:            c.subnetName,
		ID:              c.subnetName,
		CIDR:            c.subnetCIDR,
		AZ:              c.zone,
		AddressAffinity: cloud.AnyAddress,
	}}
	return sns, nil
}

func (c *gceClient) GetAvailabilityZones() ([]string, error) {
	return []string{c.zone}, nil
}

func (c *gceClient) AddRoute(destinationCIDR, instanceID string) error {
	return NI()
}

func (c *gceClient) RemoveRoute(destinationCIDR string) error {
	return NI()
}

func (c *gceClient) ModifySourceDestinationCheck(instanceID string, isEnabled bool) error {
	return NI()
}

func (c *gceClient) GetDNSInfo() ([]string, []string, error) {
	return nil, nil, NI()
}

func (c *gceClient) GetVPCCIDRs() []string {
	return []string{c.vpcCIDR}
}

func (c *gceClient) IsAvailable() (bool, error) {
	return false, NI()
}
