package aws

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/elotl/cloud-instance-provider/pkg/util/sets"
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
		return "", fmt.Errorf("Could not get subnet info from metadata service, no subnets found. Please manually specify a subnet ID in server.yml")
	} else if len(subnetIDs) > 1 {
		return "", fmt.Errorf("Instance is connected to multiple subnets. Please specify a subnet ID in server.yml where instances will be launched into.")
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
			vpcError := fmt.Errorf("Could not autodetect VPC ID, please set the cloud vpcID in /etc/mlpa/server.yml. The error was: %s", err)
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

func (e *AwsEC2) GetSubnets() ([]cloud.SubnetAttributes, error) {
	klog.V(2).Infof("Getting subnets and availability zones for VPC %s", e.vpcID)
	vpcFilters := []*ec2.Filter{
		{
			Name: aws.String("vpc-id"),
			Values: []*string{
				aws.String(e.vpcID),
			},
		},
	}
	snResp, err := e.client.DescribeSubnets(&ec2.DescribeSubnetsInput{
		Filters: vpcFilters,
	})
	if err != nil {
		return nil, util.WrapError(err, "Error getting VPC subnets from AWS")
	}
	if len(snResp.Subnets) == 0 {
		return nil, fmt.Errorf("No subnets found in VPC")
	}
	rtResp, err := e.client.DescribeRouteTables(&ec2.DescribeRouteTablesInput{
		Filters: vpcFilters,
	})
	if err != nil {
		return nil, util.WrapError(err, "Error getting VPC route tables from AWS")
	}
	subnets, err := makeMilpaSubnets(snResp.Subnets, rtResp.RouteTables)
	return subnets, err
}

func (az *AwsEC2) GetAvailabilityZones() ([]string, error) {
	sns, err := az.GetSubnets()
	if err != nil {
		return nil, err
	}
	azs := sets.NewString()
	for _, sn := range sns {
		azs.Insert(sn.AZ)
	}
	return azs.List(), nil
}

func makeMilpaSubnets(awsSubnets []*ec2.Subnet, rts []*ec2.RouteTable) ([]cloud.SubnetAttributes, error) {
	// Todo, return an error if we have a subnet length of 0
	subnets := make([]cloud.SubnetAttributes, len(awsSubnets))
	for i, subnet := range awsSubnets {
		subnetID := aws.StringValue(subnet.SubnetId)
		addressType := cloud.PrivateAddress
		isPublic, err := isSubnetPublic(rts, subnetID)
		if err != nil {
			klog.Errorf("could not compute if %s is public subnet: %v", subnetID, err)
			continue
		}
		if isPublic {
			addressType = cloud.PublicAddress
		}
		subnetInfo := cloud.SubnetAttributes{
			ID:                 subnetID,
			CIDR:               aws.StringValue(subnet.CidrBlock),
			AZ:                 aws.StringValue(subnet.AvailabilityZone),
			AddressAffinity:    addressType,
			AvailableAddresses: int(aws.Int64Value(subnet.AvailableIpAddressCount)),
		}
		subnets[i] = subnetInfo
	}
	// Now, if we have public AND private address subnets then we are
	// good to go, that means the user has setup their network with 2
	// different zones (public and private addressing:
	// https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Scenario2.html).
	// However, if we have only public subnets (pretty likely, esp if
	// the user created their VPC for Milpa), then we want to mark the
	// public subnets as available for any address type.  That way,
	// the user can launch instances with both public and private IP
	// addresses using milpa (which is what they want!).
	onlyPublic := true
	for i := range subnets {
		if subnets[i].AddressAffinity == cloud.PrivateAddress {
			onlyPublic = false
			break
		}
	}
	if onlyPublic {
		for i := range subnets {
			subnets[i].AddressAffinity = cloud.AnyAddress
		}
	}
	return subnets, nil
}

// Taken from k8s.  Works as advertised (finds the route table associated
// with the subnet and figures out if there's an internet gateway route
// in that table
func isSubnetPublic(rt []*ec2.RouteTable, subnetID string) (bool, error) {
	var subnetTable *ec2.RouteTable
	for _, table := range rt {
		for _, assoc := range table.Associations {
			if aws.StringValue(assoc.SubnetId) == subnetID {
				subnetTable = table
				break
			}
		}
	}
	if subnetTable == nil {
		// If there is no explicit association, the subnet will be implicitly
		// associated with the VPC's main routing table.
		for _, table := range rt {
			for _, assoc := range table.Associations {
				if aws.BoolValue(assoc.Main) == true {
					klog.V(4).Infof("Assuming implicit use of main routing table %s for %s",
						aws.StringValue(table.RouteTableId), subnetID)
					subnetTable = table
					break
				}
			}
		}
	}

	if subnetTable == nil {
		return false, fmt.Errorf("Could not locate routing table for subnet %s", subnetID)
	}

	for _, route := range subnetTable.Routes {
		// There is no direct way in the AWS API to determine if a subnet is public or private.
		// A public subnet is one which has an internet gateway route
		// we look for the gatewayId and make sure it has the prefix of igw to differentiate
		// from the default in-subnet route which is called "local"
		// or other virtual gateway (starting with vgv)
		// or vpc peering connections (starting with pcx).
		if strings.HasPrefix(aws.StringValue(route.GatewayId), "igw") {
			return true, nil
		}
	}

	return false, nil
}

func (e *AwsEC2) ControllerInsideVPC() bool {
	vpcID, err := detectCurrentVPC()
	if err != nil {
		return false
	} else if vpcID != "" && vpcID == e.vpcID {
		return true
	}
	return false
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

func (e *AwsEC2) RemoveRoute(destinationCIDR string) error {
	out, err := e.client.DescribeRouteTables(&ec2.DescribeRouteTablesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("vpc-id"),
				Values: []*string{
					aws.String(e.vpcID),
				},
			},
		},
	})
	if err != nil {
		return util.WrapError(err, "Could not list vpc route tables")
	}
	for _, table := range out.RouteTables {
		if len(table.Associations) == 0 {
			continue
		}
		for _, route := range table.Routes {
			if destinationCIDR == aws.StringValue(route.DestinationCidrBlock) {
				_, err = e.client.DeleteRoute(&ec2.DeleteRouteInput{
					DestinationCidrBlock: route.DestinationCidrBlock,
					RouteTableId:         table.RouteTableId,
				})
				if err != nil {
					return util.WrapError(
						err,
						"Error deleting old route table entry for table %s",
						*table.RouteTableId,
					)
				}
				break
			}
		}
	}
	return nil
}

func (e *AwsEC2) AddRoute(destinationCIDR, instanceID string) error {
	out, err := e.client.DescribeRouteTables(&ec2.DescribeRouteTablesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("vpc-id"),
				Values: []*string{
					aws.String(e.vpcID),
				},
			},
		},
	})
	if err != nil {
		return util.WrapError(err, "Could not list vpc route tables")
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
				err, "Error creating route table entry in table %s",
				*table.RouteTableId)

		}
	}
	return nil
}
