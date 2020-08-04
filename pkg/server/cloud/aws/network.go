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

package aws

import (
	"fmt"
	"net"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/util"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog"
)

// Errors to be aware of in the future:
// ElasticIP: AddressLimitExceeded, IncorrectInstanceState (must not be pending for association)
//
// Todo: limit the number of rules we can have in a security group
// also watch out for having too many security groups in a VPC

func (e *AwsEC2) CreateSGName(svcName string) string {
	return util.CreateSecurityGroupName(e.controllerID, svcName)
}

func detectCurrentVPC() (string, error) {
	macString, err := GetMetadata("network/interfaces/macs/")
	if err != nil {
		return "", util.WrapError(err, "Error getting instance mac addresses")
	}

	vpcs := sets.NewString()
	for _, mac := range strings.Split(macString, "\n") {
		mac = strings.Trim(mac, "/")
		vpcPath := fmt.Sprintf("network/interfaces/macs/%s/vpc-id", mac)
		vpcResponse, err := GetMetadata(vpcPath)
		if err != nil {
			klog.Errorf("Could not get vpc for mac address: %s\n", mac)
			continue
		}
		vpcs.Insert(vpcResponse)
	}
	if len(vpcs) == 0 {
		return "", fmt.Errorf("Could not get network info from metadata service, no vpcs found")
	} else if len(vpcs) > 1 {
		return "", fmt.Errorf("Instance is connected to multiple VPCs")
	}
	vpcID, _ := vpcs.PopAny()
	return vpcID, nil
}

func detectCurrentSubnet() (string, error) {
	macString, err := GetMetadata("network/interfaces/macs/")
	if err != nil {
		return "", util.WrapError(err, "Error getting instance mac addresses")
	}

	subnetIDs := sets.NewString()
	for _, mac := range strings.Split(macString, "\n") {
		mac = strings.Trim(mac, "/")
		idPath := fmt.Sprintf("network/interfaces/macs/%s/subnet-id", mac)
		response, err := GetMetadata(idPath)
		if err != nil {
			klog.Errorf("Could not get subnet ID for mac address %s, %v\n", mac, err)
			continue
		}
		subnetIDs.Insert(response)
	}
	if len(subnetIDs) == 0 {
		return "", fmt.Errorf("Could not get subnet info from metadata service, no subnets found. Please manually specify a subnetID in provider.yaml")
	} else if len(subnetIDs) > 1 {
		return "", fmt.Errorf("Instance is connected to multiple subnets. Please specify a subnet ID in provider.yaml where instances will be launched into.")
	}
	subnetID, _ := subnetIDs.PopAny()
	return subnetID, nil
}

func (e *AwsEC2) assertVPCExists(vpcID string) (string, string, error) {
	// If the user didn't supply a VPC, detect the VPC that milpa is
	// connected to
	if vpcID == "" {
		detectedVPC, err := detectCurrentVPC()
		if err != nil {
			vpcError := fmt.Errorf("Could not autodetect VPC ID, please set the cloud vpcID in provider.yaml. The error was: %s", err)
			return "", "", vpcError
		} else {
			vpcID = detectedVPC
		}
	}

	var filters []*ec2.Filter
	// VPC can either be the default or the user can specify the name
	// of a VPC
	if strings.ToLower(vpcID) == "default" {
		filters = []*ec2.Filter{
			&ec2.Filter{
				Name:   aws.String("isDefault"),
				Values: aws.StringSlice([]string{"true"}),
			},
		}
	} else {
		filters = []*ec2.Filter{
			&ec2.Filter{
				Name:   aws.String("vpc-id"),
				Values: aws.StringSlice([]string{vpcID}),
			},
		}
	}

	//output, err := e.client.DescribeVpcs(&ec2.DescribeVpcsInput{Filters: filters})
	vpcs := make([]ec2.Vpc, 0, 1)
	err := e.client.DescribeVpcsPages(&ec2.DescribeVpcsInput{Filters: filters},
		func(page *ec2.DescribeVpcsOutput, lastPage bool) bool {
			if page == nil {
				return false
			}
			for i := range page.Vpcs {
				vpcs = append(vpcs, *(page.Vpcs[i]))
			}
			return page.NextToken != nil
		})
	if err != nil {
		return "", "", util.WrapError(err, "Could not get VPC")
	} else if len(vpcs) < 1 {
		err := fmt.Errorf(
			"Could not find VPC %s, you must create the specified VPC before running Milpa",
			vpcID)
		return "", "", err
	}
	// // if we found a VPC, grab the data out of it
	klog.V(2).Infoln("Current vpc: ",
		aws.StringValue(vpcs[0].VpcId),
		aws.StringValue(vpcs[0].CidrBlock))
	return aws.StringValue(vpcs[0].VpcId), aws.StringValue(vpcs[0].CidrBlock), nil
}

func (e *AwsEC2) getSubnetAttributes(subnetID string) (snAttrs cloud.SubnetAttributes, err error) {
	klog.V(2).Infof("Getting subnets and availability zones for VPC %s", e.vpcID)
	describeSubnetsFilter := []*ec2.Filter{
		{
			Name: aws.String("vpc-id"),
			Values: []*string{
				aws.String(e.vpcID),
			},
		},
		{
			Name: aws.String("subnet-id"),
			Values: []*string{
				aws.String(subnetID),
			},
		},
	}
	snResp, err := e.client.DescribeSubnets(&ec2.DescribeSubnetsInput{
		Filters: describeSubnetsFilter,
	})
	if err != nil {
		return snAttrs, util.WrapError(err, "Error getting subnet %s from AWS", subnetID)
	}
	if len(snResp.Subnets) == 0 || snResp.Subnets[0] == nil {
		return snAttrs, fmt.Errorf("Could not find subnet %s in in VPC", subnetID)
	}
	rt, err := e.getSubnetRouteTable(subnetID)
	snAttrs, err = makeSubnetAttrs(snResp.Subnets[0], rt)
	return snAttrs, err
}

func (e *AwsEC2) getSubnetRouteTable(subnetID string) (*ec2.RouteTable, error) {
	describeVPCRouteTablesFilter := []*ec2.Filter{
		{
			Name: aws.String("vpc-id"),
			Values: []*string{
				aws.String(e.vpcID),
			},
		},
	}
	rtResp, err := e.client.DescribeRouteTables(&ec2.DescribeRouteTablesInput{
		Filters: describeVPCRouteTablesFilter,
	})
	if err != nil {
		return nil, util.WrapError(err, "Error getting VPC route tables from AWS")
	}
	var defaultRT *ec2.RouteTable
	for i, rt := range rtResp.RouteTables {
		for _, assoc := range rt.Associations {
			if aws.BoolValue(assoc.Main) {
				defaultRT = rtResp.RouteTables[i]
			}
			if aws.StringValue(assoc.SubnetId) == subnetID {
				return rt, nil
			}
		}
	}
	if defaultRT == nil {
		return nil, fmt.Errorf("Could not get route table associated with subnet %s in in VPC", subnetID)
	}
	return defaultRT, nil
}

func (az *AwsEC2) IsAvailable() (bool, error) {
	out, err := az.client.DescribeAvailabilityZones(&ec2.DescribeAvailabilityZonesInput{
		ZoneNames: aws.StringSlice([]string{az.availabilityZone}),
	})
	if err != nil {
		return false, err
	}
	if len(out.AvailabilityZones) != 1 {
		return false, fmt.Errorf("invalid AZ reply")
	}
	state := strings.ToLower(aws.StringValue(out.AvailabilityZones[0].State))
	return state == "available", nil
}

func makeSubnetAttrs(awsSubnet *ec2.Subnet, rt *ec2.RouteTable) (cloud.SubnetAttributes, error) {
	snAttrs := cloud.SubnetAttributes{}
	subnetID := aws.StringValue(awsSubnet.SubnetId)
	addressType := cloud.PrivateAddress
	isPublic, err := isSubnetPublic(rt, subnetID)
	if err != nil {
		return snAttrs, util.WrapError(err, "could not compute if %s is public subnet", subnetID)

	}
	if isPublic {
		addressType = cloud.AnyAddress
	}
	snAttrs = cloud.SubnetAttributes{
		ID:                 subnetID,
		CIDR:               aws.StringValue(awsSubnet.CidrBlock),
		AZ:                 aws.StringValue(awsSubnet.AvailabilityZone),
		AddressAffinity:    addressType,
		AvailableAddresses: int(aws.Int64Value(awsSubnet.AvailableIpAddressCount)),
	}
	return snAttrs, nil
}

// Taken from k8s.  Works as advertised (giventhe route table associated
// with the subnet, figure out if there's an internet gateway route
// in that table. If so then the subnet supports public addresses.
func isSubnetPublic(subnetTable *ec2.RouteTable, subnetID string) (bool, error) {
	if subnetTable == nil {
		return false, fmt.Errorf("Could not locate routing table for subnet %s", subnetID)
	}

	for _, route := range subnetTable.Routes {
		// There is no direct way in the AWS API to determine if a
		// subnet is public or private.  A public subnet is one which
		// has an internet gateway route we look for the gatewayId and
		// make sure it has the prefix of igw to differentiate from
		// the default in-subnet route which is called "local" or
		// other virtual gateway (starting with vgv) or vpc peering
		// connections (starting with pcx).
		if strings.HasPrefix(aws.StringValue(route.GatewayId), "igw") {
			return true, nil
		}
	}

	return false, nil
}

func (e *AwsEC2) ConnectWithPublicIPs() bool {
	if !e.usePublicIPs {
		return false
	} else {
		return !e.ControllerInsideVPC()
	}
}

func (e *AwsEC2) ControllerInsideVPC() bool {
	vpcID, err := detectCurrentVPC()
	inside := false
	if err == nil && vpcID != "" && vpcID == e.vpcID {
		inside = true
		klog.V(2).Infoln("controller is inside the VPC")
	} else {
		klog.V(2).Infoln("controller is outside the VPC")
	}
	return inside
}

func (e *AwsEC2) ModifySourceDestinationCheck(instanceID string, isEnabled bool) error {
	_, err := e.client.ModifyInstanceAttribute(
		&ec2.ModifyInstanceAttributeInput{
			InstanceId: aws.String(instanceID),
			SourceDestCheck: &ec2.AttributeBooleanValue{
				Value: aws.Bool(isEnabled),
			},
		})
	return err
}

func (e *AwsEC2) RemoveRoute(destinationCIDR, instanceID string) error {
	out, err := e.client.DescribeRouteTables(&ec2.DescribeRouteTablesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: aws.StringSlice([]string{e.vpcID}),
			},
			{
				Name:   aws.String("association.main"),
				Values: aws.StringSlice([]string{"false"}),
			},
			{
				Name:   aws.String("association.subnet-id"),
				Values: aws.StringSlice([]string{e.subnetID}),
			},
		},
	})
	if err != nil {
		return util.WrapError(err, "listing vpc route tables")
	}
	for _, table := range out.RouteTables {
		if len(table.Associations) == 0 {
			continue
		}
		for _, route := range table.Routes {
			// Ignore any routes that are not instance routes. If the instance
			// has been terminated, the route will be a blackhole and EC2 will
			// set route.NetworkInterfaceId while clearing up route.InstanceId.
			routeIID := aws.StringValue(route.InstanceId)
			routeNetIID := aws.StringValue(route.NetworkInterfaceId)
			if routeIID == "" && routeNetIID == "" {
				continue
			}
			// If instanceID is provided, match on it.
			if instanceID != "" && instanceID != routeIID {
				continue
			}
			// If destinationCIDR is provided, match on it.
			if destinationCIDR != "" &&
				destinationCIDR != aws.StringValue(route.DestinationCidrBlock) {
				continue
			}
			// If both instanceID and destinationCIDR are empty, remove
			// blackhole instance routes.
			if destinationCIDR == "" && instanceID == "" &&
				aws.StringValue(route.State) != ec2.RouteStateBlackhole {
				continue
			}
			_, err = e.client.DeleteRoute(&ec2.DeleteRouteInput{
				DestinationCidrBlock: route.DestinationCidrBlock,
				RouteTableId:         table.RouteTableId,
			})
			if err != nil {
				return util.WrapError(
					err,
					"deleting route %s in table %s",
					aws.StringValue(route.DestinationCidrBlock),
					aws.StringValue(table.RouteTableId))
			}
			klog.V(5).Infof("removed route %s in table %s",
				aws.StringValue(route.DestinationCidrBlock),
				aws.StringValue(table.RouteTableId))
		}
	}
	return nil
}

func (e *AwsEC2) AddRoute(destinationCIDR, instanceID string) error {
	if destinationCIDR == "" || instanceID == "" {
		return fmt.Errorf(
			"invalid input: empty value (got %q %q)",
			destinationCIDR, instanceID)
	}
	out, err := e.client.DescribeRouteTables(&ec2.DescribeRouteTablesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: aws.StringSlice([]string{e.vpcID}),
			},
			{
				Name:   aws.String("association.main"),
				Values: aws.StringSlice([]string{"false"}),
			},
			{
				Name:   aws.String("association.subnet-id"),
				Values: aws.StringSlice([]string{e.subnetID}),
			},
		},
	})
	if err != nil {
		return util.WrapError(err, "listing vpc route tables")
	}
	for _, table := range out.RouteTables {
		if len(table.Associations) == 0 {
			continue
		}
		_, err = e.client.CreateRoute(&ec2.CreateRouteInput{
			DestinationCidrBlock: aws.String(destinationCIDR),
			InstanceId:           aws.String(instanceID),
			RouteTableId:         table.RouteTableId,
		})
		if err != nil {
			return util.WrapError(
				err,
				"creating route %s via %s in table %s",
				destinationCIDR, instanceID, aws.StringValue(table.RouteTableId))
		}
		klog.V(5).Infof("added route %s in table %s",
			destinationCIDR, aws.StringValue(table.RouteTableId))
	}
	return nil
}

func (e *AwsEC2) GetDNSInfo() ([]string, []string, error) {
	_, cidr, err := net.ParseCIDR(e.vpcCIDR)
	if err != nil {
		return nil, nil, util.WrapError(
			err, "failed to parse VPC CIDR %q", e.vpcCIDR)
	}
	out, err := e.client.DescribeVpcs(&ec2.DescribeVpcsInput{
		VpcIds: aws.StringSlice([]string{e.vpcID}),
	})
	if err != nil {
		return nil, nil, util.WrapError(
			err, "DescribeVpcs %q", e.vpcID)
	}
	if len(out.Vpcs) != 1 || out.Vpcs[0] == nil {
		return nil, nil, util.WrapError(
			err, "DescribeVpcs %q malformed reply %+v", e.vpcID, *out)
	}
	dhcpOptionsID := aws.StringValue(out.Vpcs[0].DhcpOptionsId)
	dhcpOut, err := e.client.DescribeDhcpOptions(&ec2.DescribeDhcpOptionsInput{
		DhcpOptionsIds: aws.StringSlice([]string{dhcpOptionsID}),
	})
	if err != nil {
		return nil, nil, util.WrapError(
			err, "DescribeDhcpOptions %q", dhcpOptionsID)
	}
	if len(dhcpOut.DhcpOptions) != 1 || dhcpOut.DhcpOptions[0] == nil {
		return nil, nil, util.WrapError(
			err,
			"DescribeDhcpOptions %q malformed reply %+v",
			dhcpOptionsID,
			*dhcpOut)
	}
	var nameservers []string
	var searches []string
	for _, cfg := range dhcpOut.DhcpOptions[0].DhcpConfigurations {
		if aws.StringValue(cfg.Key) == "domain-name-servers" {
			if len(cfg.Values) == 1 &&
				aws.StringValue(cfg.Values[0].Value) == "AmazonProvidedDNS" {
				// The DNS server in a VPC is always the base address + 2.
				ip := util.NextIP(cidr.IP, 2)
				nameservers = []string{ip.String()}
			} else {
				for _, value := range cfg.Values {
					nameservers = append(
						nameservers, aws.StringValue(value.Value))
				}
			}
		} else if aws.StringValue(cfg.Key) == "domain-name" {
			for _, value := range cfg.Values {
				searches = append(searches, aws.StringValue(value.Value))
			}
		}
	}
	return nameservers, searches, nil
}
