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
	"net"
	"strings"

	"cloud.google.com/go/compute/metadata"
	"github.com/elotl/kip/pkg/util"
	"google.golang.org/api/compute/v1"
)

func zoneToRegion(zone string) (string, error) {
	parts := strings.Split(zone, "-")
	if len(parts) != 3 {
		return "", fmt.Errorf("unknown zone format, expecting geo-region-zone format got %q", zone)
	}
	return fmt.Sprintf("%s-%s", parts[0], parts[1]), nil
}

func (c *gceClient) ConnectWithPublicIPs() bool {
	if !c.usePublicIPs {
		return false
	} else if metadata.OnGCE() {
		mdVPC, err := c.detectCurrentVPC()
		if err != nil {
			return true
		}
		if mdVPC == c.vpcName {
			return false
		}
	}
	return true
}

func (c *gceClient) getVPCRegionCIDRs(vpcName string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	resp, err := c.service.Networks.Get(c.projectID, vpcName).Context(ctx).Do()
	if err != nil {
		return nil, util.WrapError(err, "error querying VPC %s in GCE", vpcName)
	}
	if resp == nil {
		return nil, nilResponseError("Networks.Get")
	}
	if len(resp.Subnetworks) == 0 {
		return nil, fmt.Errorf("Network error: no subnetworks found in %s network", vpcName)
	}

	// Clusters shouldn't span regions, only open the firewall rules
	// to all subnets in the controller's region
	subnets, err := c.getRegionSubnets()
	if err != nil {
		return nil, util.WrapError(err, "Error listing network subnets in region")
	}
	vpcCIDRs := make([]string, 0, len(subnets))
	for _, subnet := range subnets {
		vpcCIDRs = append(vpcCIDRs, subnet.IpCidrRange)
		for _, secondary := range subnet.SecondaryIpRanges {
			if secondary != nil {
				vpcCIDRs = append(vpcCIDRs, secondary.IpCidrRange)
			}
		}
	}
	if len(vpcCIDRs) == 0 {
		return nil, fmt.Errorf("Could not list any subnets in %s - %s", vpcName, c.region)
	}
	return vpcCIDRs, nil
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
	zone, err := getZoneFromMetadata(md)
	if err != nil {
		return "", "", util.WrapError(err, "error getting zone from GCE metadata service")
	}
	region, err := zoneToRegion(zone)
	if err != nil {
		return "", "", util.WrapError(err, "error parsing zone from GCE metadata service")
	}
	return region, zone, nil
}

func (c *gceClient) getRegionSubnets() ([]*compute.Subnetwork, error) {
	vpcURL := c.getVPCURL()
	lister := c.service.Subnetworks.List(c.projectID, c.region)
	lister = lister.Filter("network eq " + vpcURL)
	subnets := []*compute.Subnetwork{}
	f := func(page *compute.SubnetworkList) error {
		for _, subnet := range page.Items {
			if subnet != nil {
				subnets = append(subnets, subnet)
			}
		}
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	if err := lister.Pages(ctx, f); err != nil {
		return nil, err
	}
	return subnets, nil
}

func (c *gceClient) autodetectSubnet() (string, string, error) {
	if !metadata.OnGCE() {
		return "", "", fmt.Errorf("instance is not running inside GCE, could not determine the instance's subnet automatically. Please specify the subnet in cloud.gce.subnetName in provider.yaml")
	}

	// get the current subnet's CIDR
	md := newMetadataClient()
	ipStr, err := md.InternalIP()
	if err != nil {
		return "", "", util.WrapError(err, "error getting IP address from GCE metadata service")
	}
	ip := net.ParseIP(ipStr)
	subnets, err := c.getRegionSubnets()
	if err != nil {
		return "", "", util.WrapError(err, "Error listing network subnets in region")
	}
	for i := range subnets {
		subnetCIDR := subnets[i].IpCidrRange
		_, ipnet, err := net.ParseCIDR(subnetCIDR)
		if err != nil {
			return "", "", util.WrapError(err, "Could not parse CIDR returned from GCE API: %s", subnetCIDR)
		}
		if ipnet.Contains(ip) {
			return subnets[i].Name, subnets[i].IpCidrRange, nil
		}
	}

	return "", "", fmt.Errorf("Could not determine this machine's subnet from local metadata and querying the API. Please specify a subnet name at cloud.gce.subnetName in provider.yaml")
}

// As of 7/28/20 this is only used to ensure that the subnet exists
func (c *gceClient) getSubnetCIDR(subnetName string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	resp, err := c.service.Subnetworks.Get(c.projectID, c.region, subnetName).Context(ctx).Do()
	if err != nil {
		return "", util.WrapError(err, "Error looking up subnet %s", subnetName)
	}
	if resp == nil {
		return "", nilResponseError("Subnetworks.Get")
	}
	return resp.IpCidrRange, nil
}

func (c *gceClient) AddRoute(destinationCIDR, instanceID string) error {
	// TODO
	return nil
}

func (c *gceClient) RemoveRoute(destinationCIDR, instanceID string) error {
	// TODO
	return nil
}

func (c *gceClient) ModifySourceDestinationCheck(instanceID string, isEnabled bool) error {
	// TODO
	return nil
}

func (c *gceClient) GetDNSInfo() ([]string, []string, error) {
	// instance resolv.conf contents:
	//
	// domain c.milpa-207719.internal
	// search c.milpa-207719.internal. google.internal.
	// nameserver 169.254.169.254
	//
	// This is fairly symplistic, might need to look into determining
	// if the users environment has other settings
	// Note that the c. is not the zone letter. It does not change from
	// zone to zone.
	s := fmt.Sprintf("c.%s.internal.", c.projectID)
	searches := []string{s, "google.internal."}
	nameservers := []string{"169.254.169.254"}
	return nameservers, searches, nil
}

func (c *gceClient) GetVPCCIDRs() []string {
	return c.vpcCIDRs
}

func (c *gceClient) IsAvailable() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	resp, err := c.service.Zones.Get(c.projectID, c.zone).Context(ctx).Do()
	if err != nil {
		return true, err
	}
	if resp == nil {
		return true, nilResponseError("Zones.Get")
	}
	isAvailable := resp.Status != "DOWN"
	return isAvailable, nil
}
