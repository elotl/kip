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
	"math"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/elotl/cloud-instance-provider/pkg/util/sets"
	"k8s.io/klog"
)

func (c *AwsEC2) SetBootSecurityGroupIDs(ids []string) {
	c.bootSecurityGroupIDs = ids
}

func (c *AwsEC2) GetBootSecurityGroupIDs() []string {
	return c.bootSecurityGroupIDs
}

func (c *AwsEC2) EnsureMilpaSecurityGroups(extraCIDRs, extraGroupIDs []string) error {
	milpaPorts := []cloud.InstancePort{
		{
			Protocol:      api.ProtocolTCP,
			Port:          cloud.RestAPIPort,
			PortRangeSize: 1,
		},
		{
			Protocol:      api.ProtocolTCP,
			Port:          1,
			PortRangeSize: math.MaxUint16,
		},
		{
			Protocol:      api.ProtocolUDP,
			Port:          1,
			PortRangeSize: math.MaxUint16,
		},
		{
			Protocol:      api.ProtocolICMP,
			Port:          -1,
			PortRangeSize: 1,
		},
	}
	vpcCIDRs := []string{c.vpcCIDR}
	cidrs := append(vpcCIDRs, extraCIDRs...)
	apiGroupName := util.CreateSecurityGroupName(c.controllerID, cloud.MilpaAPISGName)
	apiGroup, err := c.EnsureSecurityGroup(apiGroupName, milpaPorts, cidrs)
	if err != nil {
		return util.WrapError(err, "Could not setup Milpa API cloud firewall rules")
	}
	ids := append(extraGroupIDs, apiGroup.ID)
	klog.V(2).Infoln("security group name", apiGroupName, ids)
	c.SetBootSecurityGroupIDs(ids)
	return nil
}

func (e *AwsEC2) FindSecurityGroup(sgName string) (*cloud.SecurityGroup, error) {
	filters := []*ec2.Filter{
		&ec2.Filter{
			Name:   aws.String("group-name"),
			Values: aws.StringSlice([]string{sgName}),
		},
		&ec2.Filter{
			Name:   aws.String("tag-key"),
			Values: aws.StringSlice([]string{cloud.ControllerTagKey}),
		},
		&ec2.Filter{
			Name:   aws.String("tag-value"),
			Values: aws.StringSlice([]string{e.controllerID}),
		},
		&ec2.Filter{
			Name:   aws.String("vpc-id"),
			Values: aws.StringSlice([]string{e.vpcID}),
		},
	}
	output, err := e.client.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{
		Filters:    filters,
		MaxResults: aws.Int64(1000),
	})
	if err != nil {
		return nil, util.WrapError(err, "Could not list Security Groups")
	}
	if len(output.SecurityGroups) == 0 {
		return nil, nil
	}
	sg := awsSGToMilpa(output.SecurityGroups[0])
	return &sg, nil
}

// Notice this calls findSecurityGroup twice, not the most
// efficient...  Currently this is only used to set up the milpa
// security group.  If it's used more, we'll need to do something else
// (possibly return the changes from UpdateSecurityGroup and see if we
// need to re-fetch the SG.
func (e *AwsEC2) EnsureSecurityGroup(sgName string, ports []cloud.InstancePort, sourceRanges []string) (*cloud.SecurityGroup, error) {
	sg, err := e.FindSecurityGroup(sgName)
	if err != nil {
		return nil, util.WrapError(err, "Error finding security group")
	}
	if sg == nil {
		return e.CreateSecurityGroup(sgName, ports, sourceRanges)
	}
	err = e.UpdateSecurityGroup(*sg, ports, sourceRanges)
	if err != nil {
		return nil, util.WrapError(
			err, "Could not merge new rules into existing security group")
	}
	// We have seen eventual consistency errors here, retry it if we
	// can't find the group
	for i := 0; i < 10; i++ {
		sg, err = e.FindSecurityGroup(sgName)
		if sg != nil || err != nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	if sg == nil && err == nil {
		err = fmt.Errorf("Could not find security group %s after creation", sgName)
	}
	return sg, err
}

func (e *AwsEC2) tagSecurityGroup(groupName, groupID string) error {
	_, err := e.client.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{&groupID},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String(groupName),
			},
			{
				Key:   aws.String(cloud.ControllerTagKey),
				Value: aws.String(e.controllerID),
			},
			{
				Key:   aws.String(cloud.NamespaceTagKey),
				Value: aws.String("default"),
			},
			{
				Key:   aws.String(cloud.NametagTagKey),
				Value: aws.String(e.nametag),
			},
		},
	})
	return err
}

func (e *AwsEC2) CreateSecurityGroup(sgName string, ports []cloud.InstancePort, sourceRanges []string) (*cloud.SecurityGroup, error) {
	createRes, err := e.client.CreateSecurityGroup(&ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(sgName),
		Description: aws.String(fmt.Sprintf("MilpaSG %s %s", e.nametag, sgName)),
		VpcId:       aws.String(e.vpcID),
	})
	if err != nil {
		return nil, util.WrapError(err, "Could not create security group")
	}
	groupID := aws.StringValue(createRes.GroupId)

	err = util.Retry(
		90*time.Second,
		func() error { return e.tagSecurityGroup(sgName, groupID) },
		func(err error) bool {
			// todo, look at casting AWS errors, we can get codes, etc.
			return strings.HasPrefix(err.Error(), "InvalidGroup.NotFound")
		})
	if err != nil {
		return nil, util.WrapError(err, "Could not tag security group")
	}

	rules := cloud.MakeIngressRules(ports, sourceRanges)
	ipPermissions := makeIPPermissions(rules)
	err = util.Retry(
		45*time.Second,
		func() error {
			_, err = e.client.AuthorizeSecurityGroupIngress(
				&ec2.AuthorizeSecurityGroupIngressInput{
					GroupId:       aws.String(groupID),
					IpPermissions: ipPermissions,
				})
			return err
		},
		func(err error) bool {
			return strings.HasPrefix(err.Error(), "InvalidGroup.NotFound")
		})

	if err != nil {
		_ = e.DeleteSecurityGroup(groupID)
		return nil, util.WrapError(
			err, "Unable to set security group %s ingress", sgName)
	}

	sg := cloud.NewSecurityGroup(groupID, sgName, ports, sourceRanges)
	return &sg, nil
}

// In non-default VPC, docs say you can only delete Groups
// by ID.
func (e *AwsEC2) DeleteSecurityGroup(groupID string) error {
	_, err := e.client.DeleteSecurityGroup(&ec2.DeleteSecurityGroupInput{
		GroupId: aws.String(groupID),
	})
	if err != nil {
		return util.WrapError(err, "Could not delete security group")
	}
	klog.V(2).Infof("Deleted security group %s", groupID)
	return nil
}

func awsSGToMilpa(sg *ec2.SecurityGroup) cloud.SecurityGroup {
	// We protect against ports with extra CIDRs here.  We
	sourceRangesSet := sets.NewString()
	ports := make([]cloud.InstancePort, len(sg.IpPermissions))

	for i := 0; i < len(sg.IpPermissions); i++ {
		port := int(*(sg.IpPermissions[i].FromPort))
		portRangeSize := int(*(sg.IpPermissions[i].ToPort)) - port + 1
		if portRangeSize <= 0 {
			portRangeSize = 1
		}
		ports[i] = cloud.InstancePort{
			Port:          port,
			PortRangeSize: portRangeSize,
			Protocol:      api.MakeProtocol(*sg.IpPermissions[i].IpProtocol),
		}
		for i := 0; i < len(sg.IpPermissions[0].IpRanges); i++ {
			sourceRangesSet.Insert(*sg.IpPermissions[0].IpRanges[i].CidrIp)
		}
	}
	return cloud.NewSecurityGroup(
		*(sg.GroupId), *(sg.GroupName), ports, sourceRangesSet.List())
}

// go through and figure out what rules need to be deleted and what
// rules need to be added in order to make our security group match
// the spec the user has asked for.  We do the merge instead of
// deleting everything and re-adding because we don't want to delete
// existing rules that aren't changing sincec services might depend on
// those rules.  We have to be careful because AWS doesn't allow
// duplicate rules to exist (but does allow overlapping rules)
func (e *AwsEC2) UpdateSecurityGroup(cloudSG cloud.SecurityGroup, specPorts []cloud.InstancePort, sourceRanges []string) error {
	addIngress, deleteIngress := cloud.MergeSecurityGroups(cloudSG, specPorts, sourceRanges)
	// performance: these two calls could be in parallel
	var errMsg string
	if len(addIngress) > 0 {
		newPerms := makeIPPermissions(addIngress)
		_, err := e.client.AuthorizeSecurityGroupIngress(
			&ec2.AuthorizeSecurityGroupIngressInput{
				GroupId:       aws.String(cloudSG.ID),
				IpPermissions: newPerms,
			})
		if err != nil {
			errMsg = fmt.Sprintf(
				"Error authoizing ingress ports for group: %s", err.Error())
		}
	}
	// Note: In AWS we can ask to revoke permissions that don't
	// actually exist and we won't get an error.  In the Input,
	// anything that exists will be revoked and things that don't
	// exist will be ignored.  This makes diffing easier.
	if len(deleteIngress) > 0 {
		deletePerms := makeIPPermissions(deleteIngress)
		_, err := e.client.RevokeSecurityGroupIngress(
			&ec2.RevokeSecurityGroupIngressInput{
				GroupId:       aws.String(cloudSG.ID),
				IpPermissions: deletePerms,
			})
		if err != nil {
			errMsg += fmt.Sprintf(
				"Error revoking ingress ports for group: %s", err.Error())
		}
	}
	if errMsg != "" {
		return fmt.Errorf(errMsg)
	}
	return nil
}

func makeIPPermissions(rules []cloud.IngressRule) []*ec2.IpPermission {
	ipPermissions := make([]*ec2.IpPermission, len(rules))
	for i := 0; i < len(rules); i++ {
		fromPort := rules[i].Port
		// AWS ranges are inclusive, our range size, by definition, excludes
		// the ending port
		toPort := rules[i].Port + rules[i].PortRangeSize - 1
		if rules[i].PortRangeSize <= 0 {
			toPort = fromPort
		}
		ipPermissions[i] = &ec2.IpPermission{
			IpProtocol: aws.String(string(rules[i].Protocol)),
			FromPort:   aws.Int64(int64(fromPort)),
			ToPort:     aws.Int64(int64(toPort)),
			IpRanges:   []*ec2.IpRange{&ec2.IpRange{CidrIp: &rules[i].Source}},
		}
	}
	return ipPermissions
}

func (e *AwsEC2) AttachSecurityGroups(node *api.Node, groups []string) error {
	allGroups := append(e.bootSecurityGroupIDs, groups...)
	for i := range allGroups {
		allGroups[i] = strings.TrimSpace(allGroups[i])
	}
	_, err := e.client.ModifyInstanceAttribute(
		&ec2.ModifyInstanceAttributeInput{
			InstanceId: aws.String(node.Status.InstanceID),
			Groups:     aws.StringSlice(allGroups),
		})
	return err
}
