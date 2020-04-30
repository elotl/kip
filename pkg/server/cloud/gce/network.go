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
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/util"
	"google.golang.org/api/compute/v1"
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

func (c *gceClient) getVPCRegionCIDRs(vpcName string) ([]string, error) {
	ctx := context.Background()
	resp, err := c.service.Networks.Get(c.projectID, vpcName).Context(ctx).Do()
	if err != nil {
		return nil, util.WrapError(err, "error querying VPC %s in GCP", vpcName)
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
	vpcCIDRs := make([]string, len(resp.Subnetworks))
	for i := range subnets {
		vpcCIDRs[i] = subnets[i].IpCidrRange
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
	if err := lister.Pages(context.Background(), f); err != nil {
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
			return "", "", util.WrapError(err, "Could not parse CIDR returned from GCP API: %s", subnetCIDR)
		}
		if ipnet.Contains(ip) {
			return subnets[i].Name, subnets[i].IpCidrRange, nil
		}
	}

	return "", "", fmt.Errorf("Could not determine this machine's subnet from local metadata and querying the API. Please specify a subnet name at cloud.gce.subnetName in provider.yaml")
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
	klog.Warningln("AddRoute not implemented for GCE")
	return nil
}

func (c *gceClient) RemoveRoute(destinationCIDR, instanceID string) error {
	klog.Warningln("RemoveRoute not implemented for GCE")
	return nil
}

func (c *gceClient) ModifySourceDestinationCheck(instanceID string, isEnabled bool) error {
	klog.Warningln("ModifySourceDestinationCheck not implemented for GCE")
	return nil
}

func (c *gceClient) GetDNSInfo() ([]string, []string, error) {
	// resolv.conf contents:
	//
	// domain c.milpa-207719.internal
	// search c.milpa-207719.internal. google.internal.
	// nameserver 169.254.169.254

	klog.Warningln("Need to improve GETDNSInfo()")

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
