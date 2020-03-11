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
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	uuid "github.com/satori/go.uuid"
	"k8s.io/klog"
)

const (
	awsInstanceProduct  = "Linux/UNIX"
	resizeTimeout       = 60 * time.Second
	maxUserInstanceTags = 45
)

func (e *AwsEC2) StopInstance(instanceID string) error {
	awsInstanceIDs := []*string{aws.String(instanceID)}
	_, err := e.client.TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: awsInstanceIDs,
	})
	if err != nil {
		klog.Errorf("Error terminating instance: %v", err)
		// todo, check on status of instance, set status of instance
		// based on that, prepare to come back and clean this
		// inconsistency up
		return err
	}
	return nil
}

func (e *AwsEC2) getNodeTags(node *api.Node) []*ec2.Tag {
	nametag := util.CreateUnboundNodeNameTag(e.nametag)
	tags := []*ec2.Tag{
		&ec2.Tag{
			Key:   aws.String("Name"),
			Value: aws.String(nametag),
		},
		&ec2.Tag{
			Key:   aws.String("Node"),
			Value: aws.String(node.Name),
		},
		&ec2.Tag{
			Key:   aws.String(cloud.ControllerTagKey),
			Value: aws.String(e.controllerID),
		},
		&ec2.Tag{
			Key:   aws.String(cloud.NametagTagKey),
			Value: aws.String(e.nametag),
		},
	}
	return tags
}

func (e *AwsEC2) getBlockDeviceMapping(volSizeGiB int32) []*ec2.BlockDeviceMapping {
	awsVolSize := aws.Int64(int64(volSizeGiB))
	devices := []*ec2.BlockDeviceMapping{
		&ec2.BlockDeviceMapping{
			DeviceName: aws.String("xvda"),
			Ebs: &ec2.EbsBlockDevice{
				VolumeType:          aws.String("gp2"),
				DeleteOnTermination: aws.Bool(true),
				VolumeSize:          awsVolSize,
			}},
	}
	return devices
}

func (e *AwsEC2) getInstanceNetworkSpec(privateIPOnly bool) []*ec2.InstanceNetworkInterfaceSpecification {
	networkSpec := []*ec2.InstanceNetworkInterfaceSpecification{
		&ec2.InstanceNetworkInterfaceSpecification{
			AssociatePublicIpAddress: aws.Bool(!privateIPOnly),
			DeviceIndex:              aws.Int64(0), // seems to work
			Groups:                   aws.StringSlice(e.bootSecurityGroupIDs),
			SecondaryPrivateIpAddressCount: aws.Int64(1),
		},
	}
	// Let AWS figure out the subnet/AZ if we didn't specify a subnet
	networkSpec[0].SubnetId = aws.String(e.subnetID)
	return networkSpec
}

func subnetMatchNeeds(privateIP bool, subnet cloud.SubnetAttributes) bool {
	return (subnet.AddressAffinity == cloud.AnyAddress) ||
		(privateIP && subnet.AddressAffinity == cloud.PrivateAddress) ||
		(!privateIP && subnet.AddressAffinity == cloud.PublicAddress)
}

func (e *AwsEC2) getFirstVolume(instanceId string) *ec2.Volume {
	input := &ec2.DescribeVolumesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("attachment.instance-id"),
				Values: []*string{
					aws.String(instanceId),
				},
			},
		},
	}
	result, err := e.client.DescribeVolumes(input)
	if err != nil {
		klog.Errorf("Error retrieving list of volumes attached to %s: %v",
			instanceId, err)
		return nil
	}
	return result.Volumes[0]
}

func (e *AwsEC2) ResizeVolume(node *api.Node, size int64) (error, bool) {
	// Note: we assume that there's only one volume attached to an instance.
	vol := e.getFirstVolume(node.Status.InstanceID)
	if vol == nil || vol.Size == nil || vol.VolumeId == nil {
		return fmt.Errorf("Error retrieving volume info for node %s: %v",
			node.Name, vol), false
	}
	if *vol.Size >= size {
		klog.V(2).Infof("Volume on node %s is %dGiB >= %dGiB",
			node.Name, *vol.Size, size)
		return nil, false
	}
	klog.V(2).Infof("Resizing volume to %dGiB for node: %v", size, node)
	result, err := e.client.ModifyVolume(&ec2.ModifyVolumeInput{
		Size:     aws.Int64(size),
		VolumeId: aws.String(*vol.VolumeId),
	})
	if err != nil {
		return util.WrapError(err, "Failed to resize volume"), false
	}
	// These fields are pointers, check if any of them is nil.
	targetsize := int64(0)
	if result.VolumeModification.TargetSize != nil {
		targetsize = *result.VolumeModification.TargetSize
	}
	state := "N/A"
	if result.VolumeModification.ModificationState != nil {
		state = *result.VolumeModification.ModificationState
	}
	statusmsg := "N/A"
	if result.VolumeModification.StatusMessage != nil {
		statusmsg = *result.VolumeModification.StatusMessage
	}
	if targetsize != size {
		klog.Errorf("Error resizing volume for %v to %dGiB: state %s status %s",
			node, size, state, statusmsg)
		return util.WrapError(err, "Failed to resize volume"), false
	}
	timeout := time.Now().Add(resizeTimeout)
	for time.Now().Before(timeout) {
		time.Sleep(1 * time.Second)
		vol = e.getFirstVolume(node.Status.InstanceID)
		if vol == nil || vol.Size == nil || vol.VolumeId == nil {
			return fmt.Errorf("Error retrieving volume info for node %s: %v",
				node.Name, vol), false
		}
		if *vol.Size >= size {
			klog.V(2).Infof("Volume on node %s is %dGiB >= %dGiB",
				node.Name, *vol.Size, size)
			return nil, true
		} else {
			klog.V(2).Infof("Resizing volume on %s: currently %dGiB, requested %dGiB",
				node.Name, *vol.Size, size)
		}
	}
	return fmt.Errorf(
		"Volume resize request timeout on node %s", node.Name), false
}

func (e *AwsEC2) GetImageId(tags cloud.BootImageTags) (string, error) {
	// Let's set some reasonable defaults if BootImageTags is not present in
	// the server config.
	if tags.Company == "" {
		tags.Company = "elotl"
	}
	if tags.Product == "" {
		tags.Product = "kip"
	}
	input := &ec2.DescribeImagesInput{
		Owners: []*string{
			aws.String(e.imageOwnerID),
		},
		Filters: []*ec2.Filter{{
			Name:   aws.String("state"),
			Values: aws.StringSlice([]string{"available"}),
		}},
	}
	resp, err := e.client.DescribeImages(input)
	if err != nil {
		klog.Errorf("Error getting image list for tags %v: %v", tags, err)
		return "", err
	}
	if len(resp.Images) == 0 {
		msg := fmt.Sprintf("No images found for owner %v", e.imageOwnerID)
		klog.Errorf("%s", msg)
		return "", fmt.Errorf("%s", msg)
	}
	var images []cloud.Image
	for _, img := range resp.Images {
		image := cloud.Image{
			Id:   *img.ImageId,
			Name: *img.Name,
		}
		images = append(images, image)
	}
	return cloud.GetBestImage(images, tags)
}

func (e *AwsEC2) StartNode(node *api.Node, metadata string) (*cloud.StartNodeResult, error) {
	klog.V(2).Infof("Starting instance for node: %v", node)
	tags := e.getNodeTags(node)
	tagSpec := ec2.TagSpecification{
		ResourceType: aws.String("instance"),
		Tags:         tags,
	}
	volSizeGiB := cloud.ToSaneVolumeSize(node.Spec.Resources.VolumeSize)
	devices := e.getBlockDeviceMapping(volSizeGiB)
	networkSpec := e.getInstanceNetworkSpec(node.Spec.Resources.PrivateIPOnly)
	klog.V(2).Infof("Starting node with security groups: %v subnet: '%s'",
		e.bootSecurityGroupIDs, e.subnetID)
	result, err := e.client.RunInstances(&ec2.RunInstancesInput{
		ImageId:             aws.String(node.Spec.BootImage),
		InstanceType:        aws.String(node.Spec.InstanceType),
		MinCount:            aws.Int64(1),
		MaxCount:            aws.Int64(1),
		TagSpecifications:   []*ec2.TagSpecification{&tagSpec},
		NetworkInterfaces:   networkSpec,
		BlockDeviceMappings: devices,
		UserData:            aws.String(metadata),
	})
	if err != nil {
		if isSubnetConstrainedError(err) {
			return nil, &cloud.NoCapacityError{
				OriginalError: err.Error(),
				SubnetID:      e.subnetID,
			}
		} else if isAZConstrainedError(err) || isInstanceConstrainedError(err) {
			return nil, &cloud.NoCapacityError{
				OriginalError: err.Error(),
			}
		}
		return nil, util.WrapError(err, "Could not run instance")
	}
	if len(result.Instances) == 0 {
		return nil, fmt.Errorf("Could not get instance info at result.Instances")
	}
	cloudID := aws.StringValue(result.Instances[0].InstanceId)
	klog.V(2).Infof("Started instance: %s", cloudID)
	startResult := &cloud.StartNodeResult{
		InstanceID:       cloudID,
		AvailabilityZone: e.availabilityZone,
	}
	return startResult, nil
}

// This isn't terribly different from Start node but there are
// some minor differences.  We'll capture errors correctly here and there
func (e *AwsEC2) StartSpotNode(node *api.Node, metadata string) (*cloud.StartNodeResult, error) {
	klog.V(2).Infof("Starting instance for node: %v", node)
	tags := e.getNodeTags(node)
	tagSpec := ec2.TagSpecification{
		ResourceType: aws.String("instance"),
		Tags:         tags,
	}
	var err error
	//var subnet *cloud.SubnetAttributes
	klog.V(2).Infof("Starting spot node in: %s", e.subnetID)
	volSizeGiB := cloud.ToSaneVolumeSize(node.Spec.Resources.VolumeSize)
	devices := e.getBlockDeviceMapping(volSizeGiB)
	networkSpec := e.getInstanceNetworkSpec(node.Spec.Resources.PrivateIPOnly)
	klog.V(2).Infof("Starting node with security groups: %v subnet: '%s'",
		e.bootSecurityGroupIDs, e.subnetID)
	result, err := e.client.RunInstances(&ec2.RunInstancesInput{
		ImageId:             aws.String(node.Spec.BootImage),
		InstanceType:        aws.String(node.Spec.InstanceType),
		MinCount:            aws.Int64(1),
		MaxCount:            aws.Int64(1),
		TagSpecifications:   []*ec2.TagSpecification{&tagSpec},
		NetworkInterfaces:   networkSpec,
		BlockDeviceMappings: devices,
		UserData:            aws.String(metadata),
		InstanceMarketOptions: &ec2.InstanceMarketOptionsRequest{
			MarketType: aws.String("spot"),
			SpotOptions: &ec2.SpotMarketOptions{
				InstanceInterruptionBehavior: aws.String("terminate"),
				SpotInstanceType:             aws.String("one-time"),
			},
		},
	})

	if err != nil {
		if isSubnetConstrainedError(err) {
			return nil, &cloud.NoCapacityError{
				OriginalError: err.Error(),
				SubnetID:      e.subnetID,
			}
		} else if isAZConstrainedError(err) || isInstanceConstrainedError(err) {
			return nil, &cloud.NoCapacityError{
				OriginalError: err.Error(),
			}
		} else if isUnsupportedInstanceError(err) {
			return nil, &cloud.UnsupportedInstanceError{err.Error()}
		}
		return nil, util.WrapError(err, "Could not run instance")
	}
	if len(result.Instances) == 0 {
		return nil, fmt.Errorf("Could not get instance info at result.Instances")
	}
	cloudID := aws.StringValue(result.Instances[0].InstanceId)
	klog.V(2).Infof("Started instance: %s", cloudID)
	startResult := &cloud.StartNodeResult{
		InstanceID:       cloudID,
		AvailabilityZone: e.availabilityZone,
	}
	return startResult, nil
}

func (e *AwsEC2) WaitForRunning(node *api.Node) ([]api.NetworkAddress, error) {
	awsInstanceIDs := []*string{&node.Status.InstanceID}
	dii := &ec2.DescribeInstancesInput{InstanceIds: awsInstanceIDs}
	// Due to eventual consistency, after we create an instance and
	// get its instanceID back from RunInstances, the rest of AWS
	// might not know about that instanceID yet.
	err := util.Retry(
		30*time.Second,
		func() error {
			waitErr := e.client.WaitUntilInstanceRunning(dii)
			return waitErr
		},
		func(err error) bool {
			return strings.HasPrefix(err.Error(), "InvalidInstanceID.NotFound")
		})
	if err != nil {
		return nil, fmt.Errorf("waiting for instance to start: %v", err)
	}
	reply, err := e.client.DescribeInstances(dii)
	if err != nil {
		return nil, util.WrapError(err, "describing instances failed")
	}
	if len(reply.Reservations) == 0 || len(reply.Reservations[0].Instances) == 0 {
		return nil, fmt.Errorf("no instances found when waiting for running")
	}
	instance := reply.Reservations[0].Instances[0]
	dnii := &ec2.DescribeNetworkInterfacesInput{
		NetworkInterfaceIds: []*string{
			instance.NetworkInterfaces[0].NetworkInterfaceId,
		},
	}
	ifreply, err := e.client.DescribeNetworkInterfaces(dnii)
	if err != nil {
		return nil, util.WrapError(err, "describing network interface failed")
	}
	if len(ifreply.NetworkInterfaces) == 0 || len(ifreply.NetworkInterfaces[0].PrivateIpAddresses) <= 1 {
		return nil, fmt.Errorf("missing private IP address(es)")
	}
	addresses := api.NewNetworkAddresses(
		aws.StringValue(instance.PrivateIpAddress),
		aws.StringValue(instance.PrivateDnsName),
	)
	if !node.Spec.Resources.PrivateIPOnly {
		addresses = api.SetPublicAddresses(
			aws.StringValue(instance.PublicIpAddress),
			aws.StringValue(instance.PublicDnsName),
			addresses)
	}
	nodeIPAddress := api.GetPrivateIP(addresses)
	for _, addr := range ifreply.NetworkInterfaces[0].PrivateIpAddresses {
		ip := aws.StringValue(addr.PrivateIpAddress)
		if ip != nodeIPAddress {
			addresses = api.SetPodIP(ip, addresses)
			break
		}
	}
	return addresses, nil
}

func (e *AwsEC2) SetSustainedCPU(node *api.Node, enabled bool) error {
	creditSpec := "standard"
	if enabled {
		creditSpec = "unlimited"
	}
	req := []*ec2.InstanceCreditSpecificationRequest{{
		CpuCredits: aws.String(creditSpec),
		InstanceId: aws.String(node.Status.InstanceID),
	}}
	output, err := e.client.ModifyInstanceCreditSpecification(
		&ec2.ModifyInstanceCreditSpecificationInput{
			ClientToken:                  aws.String(uuid.NewV4().String()),
			InstanceCreditSpecifications: req,
		})
	if err != nil {
		return util.WrapError(err, "Error setting sustained CPU for node %s", node.Name)
	}
	if len(output.UnsuccessfulInstanceCreditSpecifications) > 0 {
		msg := aws.StringValue(output.UnsuccessfulInstanceCreditSpecifications[0].Error.Message)
		return fmt.Errorf("Error setting sustained CPU: %s", msg)
	}
	return nil
}

// return the ec2 tag from a list of tags, emptystring if it doesn't exist
func getTag(tags []*ec2.Tag, target string) string {
	for _, t := range tags {
		if *t.Key == target {
			return *t.Value
		}
	}
	return ""
}

func (e *AwsEC2) ListInstancesFilterID(ids []string) ([]cloud.CloudInstance, error) {
	filters := []*ec2.Filter{
		{
			Name:   aws.String("instance-id"),
			Values: aws.StringSlice(ids),
		},
		{
			Name:   aws.String("instance-state-name"),
			Values: []*string{aws.String("running"), aws.String("pending")},
		},
	}
	return e.listInstancesHelper(filters)
}

func (e *AwsEC2) ListInstances() ([]cloud.CloudInstance, error) {
	filters := []*ec2.Filter{
		{
			Name:   aws.String(fmt.Sprintf("tag:%s", cloud.ControllerTagKey)),
			Values: []*string{aws.String(e.controllerID)},
		},
		{
			Name:   aws.String("vpc-id"),
			Values: []*string{aws.String(e.vpcID)},
		},
		{
			Name:   aws.String("instance-state-name"),
			Values: []*string{aws.String("running"), aws.String("pending")},
		},
	}
	return e.listInstancesHelper(filters)
}

func (e *AwsEC2) listInstancesHelper(filters []*ec2.Filter) ([]cloud.CloudInstance, error) {
	// Todo: if filters is nil we need to page through results since AWS
	// will only return 1000 results per page
	params := &ec2.DescribeInstancesInput{
		Filters: filters,
	}
	instances := make([]cloud.CloudInstance, 0, 10)
	var nextToken *string
	for {
		params.NextToken = nextToken
		resp, err := e.client.DescribeInstances(params)
		if err != nil {
			err = util.WrapError(err, "error listing instances")
			return nil, err
		}

		for _, res := range resp.Reservations {
			for _, inst := range res.Instances {
				NodeName := getTag(inst.Tags, "Node")
				instances = append(instances,
					cloud.CloudInstance{
						ID:       *inst.InstanceId,
						NodeName: NodeName,
					})
			}
		}
		nextToken = resp.NextToken
		if nextToken == nil {
			break
		}
	}
	return instances, nil
}

// Tagging with user lables is a best effort, in other words, we allow
// this to generate errors but will try to continue with tagging if
// the user breaks some tag constraints.
func (e *AwsEC2) AddInstanceTags(iid string, labels map[string]string) error {
	awsTags, err := ec2TagsFromLabels(iid, labels)
	if err != nil {
		klog.Warning(err)
	}
	if len(awsTags) > 0 {
		_, err = e.client.CreateTags(&ec2.CreateTagsInput{
			Resources: aws.StringSlice([]string{iid}),
			Tags:      awsTags,
		})
	}
	return err
}

// For a list of AWS Errors:
// https://docs.aws.amazon.com/AWSEC2/latest/APIReference/errors-overview.html
func isSubnetConstrainedError(err error) bool {
	if awsErr, ok := err.(awserr.Error); ok {
		switch awsErr.Code() {
		case "Unsupported":
			// Kinda guessing at this from reading sourcecode of juju
			return strings.Contains(awsErr.Message(), "Availability Zone")
		case "InsufficientFreeAddressesInSubnet", "InsufficientAddressCapacity":
			return true
		}
	}
	return false
}

func isAZConstrainedError(err error) bool {
	if awsErr, ok := err.(awserr.Error); ok {
		switch awsErr.Code() {
		case "InsufficientInstanceCapacity", "InsufficientCapacity":
			// Note according to the docs, "InsufficientCapacity"
			// pertains only to instance imports. Older forum posts
			// show InsufficientCapacity errors when there's no
			// instance capacity.  Going to keep it in this case here.
			return true
		}
	}
	return false
}

func isInstanceConstrainedError(err error) bool {
	if awsErr, ok := err.(awserr.Error); ok {
		switch awsErr.Code() {
		case "InstanceLimitExceeded", "MaxSpotInstanceCountExceeded":
			return true
		}
	}
	return false
}

func isUnsupportedInstanceError(err error) bool {
	if awsErr, ok := err.(awserr.Error); ok {
		if strings.Contains(awsErr.Error(), "unsupported instance type") {
			return true
		}
	}
	return false
}

// Other AWS Errors to be aware of:
// invalid parameters:
// InvalidParameter, InvalidParameterCombination, InvalidParameterValue
// UnsupportedInstanceAttribute, UnsupportedOperation
// InvalidAvailabilityZone

func (e *AwsEC2) AssignInstanceProfile(node *api.Node, instanceProfile string) error {
	_, err := e.client.AssociateIamInstanceProfile(
		&ec2.AssociateIamInstanceProfileInput{
			IamInstanceProfile: &ec2.IamInstanceProfileSpecification{
				Arn: aws.String(instanceProfile),
			},
			InstanceId: aws.String(node.Status.InstanceID),
		})
	return err
}
