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
	"k8s.io/klog"
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

func (c *gceClient) getVPCCIDRs(vpcName string) ([]string, error) {
	ctx := context.Background()
	resp, err := c.service.Networks.Get(c.projectID, vpcName).Context(ctx).Do()
	if err != nil {
		return nil, util.WrapError(err, "error querying VPC %s in GCP", vpcName)
	}
	if resp == nil {
		return nil, nilResponseError("Networks.Get")
	}
	//spew.Dump(resp)
	if len(resp.Subnetworks) == 0 {
		return nil, fmt.Errorf("Network error: no subnetworks found in %s network", vpcName)
	}
	// Todo, need to get cidrs from all the subnets
	//resp, err := c.service.Subnetworks.Get(c.projectID, c.region, subnetName).Context(ctx).Do()
	// cidrs := make([]string{}, 0, len(resp.Subnetworks))
	// for _, subnetwork := range resp.Subnetworks {
	// 	c.service.Get()
	// }
	// return resp.IPv4Range, nil

	/*
		       @ Brendan

			   If you have a vpc url you can do something like the following. This however will only retrieve per region
			   Not everything the VPC owns

			   vpcURL := "https://www.googleapis.com/compute/v1/projects/credentials-testing/global/networks/testing-vpc"
			   listCall := c.client.Subnetworks.List(c.projectID, c.region)
			   listCall = listCall.Filter("network eq " + vpcURL)
			   f := func(page *compute.SubnetworkList) error {
			       for _, subnet := range page.Items {
			           fmt.Println(subnet.IpCidrRange)
			       }
			       return nil
			   }
			   if err := listCall.Pages(context.Background(), f); err != nil {
			       return nil , err
			   }
	*/
	return nil, TODO()
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

func (c *gceClient) autodetectSubnet() (string, string, error) {
	return "", "", TODO()
}

func (c *gceClient) getSubnetCIDR(subnetName string) (string, error) {
	ctx := context.Background()
	resp, err := c.service.Subnetworks.Get(c.projectID, c.region, subnetName).Context(ctx).Do()
	if err != nil {
		return "", util.WrapError(err, "Error looking up subnet %s", subnetName)
	}
	if resp == nil {
		return "", nilResponseError("Subnetworks.Get")
	}
	return resp.IpCidrRange, nil
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
	return TODO()
}

func (c *gceClient) RemoveRoute(destinationCIDR, instanceID string) error {
	return TODO()
}

func (c *gceClient) ModifySourceDestinationCheck(instanceID string, isEnabled bool) error {
	return TODO()
}

func (c *gceClient) GetDNSInfo() ([]string, []string, error) {
	// resolv.conf contents:
	//
	// domain c.milpa-207719.internal
	// search c.milpa-207719.internal. google.internal.
	// nameserver 169.254.169.254

	klog.Warningln("Need to implement GETDNSInfo()")

	zoneLetter := c.zone[len(c.zone)-1]
	s := fmt.Sprintf("%s.%s.internal.", string(zoneLetter), c.projectID)
	searches := []string{s, "google.internal."}
	nameservers := []string{"169.254.169.254"}
	return nameservers, searches, nil
}

func (c *gceClient) GetVPCCIDRs() []string {
	return c.vpcCIDRs
}

func (c *gceClient) IsAvailable() (bool, error) {
	klog.Errorln("Need to implement: gce.IsAvailable()")
	return true, nil
}
