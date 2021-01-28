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
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/util"
	uuid "github.com/satori/go.uuid"
	"k8s.io/klog"
)

const (
	awsInstanceProduct                  = "Linux/UNIX"
	resizeTimeout                       = 60 * time.Second
	maxUserInstanceTags                 = 45
	awsCreationDateFormat               = "2006-01-02T15:04:05.000Z"
	elotlOwnerID                        = "689494258501"
	elotlImageNameFilter                = "elotl-kip-*"
)

var (
	BootTimeout time.Duration = 15 * time.Minute
	AwsInstanceAvailableState = "available"
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

func (e *AwsEC2) getBlockDeviceMapping(image cloud.Image, volSizeGiB int32) []*ec2.BlockDeviceMapping {
	awsVolSize := aws.Int64(int64(volSizeGiB))
	devices := []*ec2.BlockDeviceMapping{
		&ec2.BlockDeviceMapping{
			DeviceName: aws.String(image.RootDevice),
			Ebs: &ec2.EbsBlockDevice{
				VolumeType:          aws.String("gp2"),
				DeleteOnTermination: aws.Bool(true),
				VolumeSize:          awsVolSize,
			}},
	}
	return devices
}

func (e *AwsEC2) getInstanceNetworkSpec(privateIPOnly bool) []*ec2.InstanceNetworkInterfaceSpecification {
	associatePublicIPAddress := true
	if privateIPOnly || !e.usePublicIPs {
		associatePublicIPAddress = false
	}
	networkSpec := []*ec2.InstanceNetworkInterfaceSpecification{
		&ec2.InstanceNetworkInterfaceSpecification{
			AssociatePublicIpAddress:       aws.Bool(associatePublicIPAddress),
			DeviceIndex:                    aws.Int64(0), // seems to work
			Groups:                         aws.StringSlice(e.bootSecurityGroupIDs),
			SecondaryPrivateIpAddressCount: aws.Int64(1),
		},
	}
	networkSpec[0].SubnetId = aws.String(e.subnetID)
	return networkSpec
}

func (e *AwsEC2) getFirstVolume(instanceID string) *ec2.Volume {
	input := &ec2.DescribeVolumesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("attachment.instance-id"),
				Values: []*string{
					aws.String(instanceID),
				},
			},
		},
	}
	result, err := e.client.DescribeVolumes(input)
	if err != nil {
		klog.Errorf("Error retrieving list of volumes attached to %s: %v",
			instanceID, err)
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
	if aws.Int64Value(vol.Size) >= size {
		klog.V(2).Infof("Volume on node %s is %dGiB >= %dGiB",
			node.Name, aws.Int64Value(vol.Size), size)
		return nil, false
	}
	klog.V(2).Infof("Resizing volume to %dGiB for node: %v", size, node)
	result, err := e.client.ModifyVolume(&ec2.ModifyVolumeInput{
		Size:     aws.Int64(size),
		VolumeId: aws.String(aws.StringValue(vol.VolumeId)),
	})
	if err != nil {
		return util.WrapError(err, "Failed to resize volume"), false
	}
	// These fields are pointers, check if any of them is nil.
	targetsize := int64(0)
	if result.VolumeModification != nil &&
		result.VolumeModification.TargetSize != nil {
		targetsize = aws.Int64Value(result.VolumeModification.TargetSize)
	}
	state := "N/A"
	if result.VolumeModification != nil &&
		result.VolumeModification.ModificationState != nil {
		state = aws.StringValue(result.VolumeModification.ModificationState)
	}
	statusmsg := "N/A"
	if result.VolumeModification != nil &&
		result.VolumeModification.StatusMessage != nil {
		statusmsg = aws.StringValue(result.VolumeModification.StatusMessage)
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
		state, reason, err := e.getVolumeModificationState(aws.StringValue(vol.VolumeId))
		if err != nil {
			return fmt.Errorf("Retrieving volume modification state failed for node %s: %v", node.Name, err), false
		}
		if state == ec2.VolumeModificationStateFailed {
			return fmt.Errorf("Volume modification failed for node %s: %v. Reason: %s",
				node.Name, vol, reason), false
		}
		if aws.Int64Value(vol.Size) >= size &&
			(state == ec2.VolumeModificationStateOptimizing ||
				state == ec2.VolumeModificationStateCompleted) {
			klog.V(2).Infof("Volume on node %s is %dGiB >= %dGiB",
				node.Name, aws.Int64Value(vol.Size), size)
			return nil, true
		} else {
			klog.V(2).Infof("Resizing volume on %s: currently %dGiB, state: %s, requested %dGiB",
				node.Name, aws.Int64Value(vol.Size), state, size)
		}
	}
	return fmt.Errorf(
		"Volume resize request timeout on node %s", node.Name), false
}

func (e *AwsEC2) getVolumeModificationState(volumeID string) (string, string, error) {
	result, err := e.client.DescribeVolumesModifications(&ec2.DescribeVolumesModificationsInput{
		VolumeIds: []*string{aws.String(volumeID)},
	})
	if err != nil {
		return "", "", err
	}
	for i := range result.VolumesModifications {
		if result.VolumesModifications[i] != nil &&
			aws.StringValue(result.VolumesModifications[i].VolumeId) == volumeID {
			mod := aws.StringValue(result.VolumesModifications[i].ModificationState)
			reason := aws.StringValue(result.VolumesModifications[i].StatusMessage)
			return mod, reason, err
		}
	}
	klog.Warningf("No volume modifications returned from api for volume %s. Will retry.", volumeID)
	return "", "", nil
}

func bootImageSpecToDescribeImagesInput(spec cloud.BootImageSpec) *ec2.DescribeImagesInput {
	input := &ec2.DescribeImagesInput{}
	if len(spec) < 1 {
		input.Owners = aws.StringSlice([]string{elotlOwnerID})
		input.Filters = []*ec2.Filter{
			{
				Name:   aws.String("name"),
				Values: aws.StringSlice([]string{elotlImageNameFilter}),
			},
		}
		return input
	}
	for key, value := range spec {
		switch key {
		case "executableUsers":
			users := strings.Fields(value)
			input.ExecutableUsers = aws.StringSlice(users)
		case "owners":
			owners := strings.Fields(value)
			input.Owners = aws.StringSlice(owners)
		case "imageIDs":
			imageIDs := strings.Fields(value)
			input.ImageIds = aws.StringSlice(imageIDs)
		case "filters":
			filters := strings.Fields(value)
			ec2Filters := make([]*ec2.Filter, len(filters))
			for i, filter := range filters {
				parts := strings.SplitN(filter, "=", 2)
				filterName := parts[0]
				filterValues := strings.Split(parts[1], ",")
				ec2Filters[i] = &ec2.Filter{
					Name:   aws.String(filterName),
					Values: aws.StringSlice(filterValues),
				}
			}
			input.Filters = ec2Filters
		default:
			klog.Warningf("invalid boot image spec key: %q (=%q)", key, value)
		}
	}
	return input
}

func getRootDeviceVolumeSize(blockDevices []*ec2.BlockDeviceMapping, rootDeviceName string) int32 {
	var rootDiskSize int32
	for _, blockDevice := range blockDevices {
		if aws.StringValue(blockDevice.DeviceName) == rootDeviceName && blockDevice.Ebs != nil {
			rootDiskSize = int32(aws.Int64Value(blockDevice.Ebs.VolumeSize))
			break
		}
	}
	return rootDiskSize
}

func (e *AwsEC2) GetImage(spec cloud.BootImageSpec) (cloud.Image, error) {
	input := bootImageSpecToDescribeImagesInput(spec)
	resp, err := e.client.DescribeImages(input)
	if err != nil {
		klog.Errorf("getting image list for spec %+v: %v", spec, err)
		return cloud.Image{}, err
	}
	if len(resp.Images) < 1 {
		msg := fmt.Sprintf("no images found for spec %+v", spec)
		klog.Errorf("%s", msg)
		return cloud.Image{}, fmt.Errorf("%s", msg)
	}
	images := make([]cloud.Image, len(resp.Images))
	for i, img := range resp.Images {
		var creationTime *time.Time
		if img.CreationDate != nil {
			ts, err := time.Parse(awsCreationDateFormat, *img.CreationDate)
			if err != nil {
				klog.Warningf(
					"invalid image creation date %s", *img.CreationDate)
			} else {
				creationTime = &ts
			}
		}
		rootDeviceName := aws.StringValue(img.RootDeviceName)
		if rootDeviceName == "" {
			klog.Warningf("cannot get root device name from image: %v", img.Name)
		}
		rootDiskSize := getRootDeviceVolumeSize(img.BlockDeviceMappings, rootDeviceName)
		images[i] = cloud.Image{
			Name:           aws.StringValue(img.Name),
			RootDevice:     aws.StringValue(img.RootDeviceName),
			ID:             aws.StringValue(img.ImageId),
			CreationTime:   creationTime,
			VolumeDiskSize: rootDiskSize,
		}
	}
	cloud.SortImagesByCreationTime(images)
	return images[len(images)-1], nil
}

func (e *AwsEC2) StartNode(node *api.Node, image cloud.Image, metadata, iamPermissions string) (string, error) {
	klog.V(2).Infof("Starting instance for node: %v", node)
	tags := e.getNodeTags(node)
	tagSpec := ec2.TagSpecification{
		ResourceType: aws.String("instance"),
		Tags:         tags,
	}
	volSizeGiB := cloud.ToSaneVolumeSize(node.Spec.Resources.VolumeSize, image)
	devices := e.getBlockDeviceMapping(image, volSizeGiB)
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
		IamInstanceProfile:  getIAMInstanceProfileSpecification(iamPermissions),
	})
	if err != nil {
		if isSubnetConstrainedError(err) {
			return "", &cloud.NoCapacityError{
				OriginalError: err.Error(),
				SubnetID:      e.subnetID,
			}
		} else if isAZConstrainedError(err) || isInstanceConstrainedError(err) {
			return "", &cloud.NoCapacityError{
				OriginalError: err.Error(),
			}
		}
		return "", util.WrapError(err, "Could not run instance")
	}
	if len(result.Instances) == 0 {
		return "", fmt.Errorf("Could not get instance info at result.Instances")
	}
	cloudID := aws.StringValue(result.Instances[0].InstanceId)
	klog.V(2).Infof("Started instance: %s", cloudID)
	return cloudID, nil
}

// We need to ensure the dedicated host fulfills two constraints
// 1) in a state of "available"
// 2) no tenant is currently occupying the host
func (e *AwsEC2) ReleaseDedicatedHosts() error {
	hosts, err := e.listAvailableDedicatedHosts()
	if err != nil {
		return err
	}
	var hostIdsForRelease []*string
	for _, host := range hosts {
		if len(host.Instances) > 0 {
			continue
		}
		hostIdsForRelease = append(hostIdsForRelease, host.HostId)
	}
	resp, err := e.client.ReleaseHosts(&ec2.ReleaseHostsInput{
		HostIds: hostIdsForRelease,
	})
	if err != nil {
		return err
	}
	// We do not want to return these as actual errors since in many cases such
	// as mac1.metal hosts there is a 24 hour limit before you are allow to release
	// the associated host.
	for _, host := range resp.Unsuccessful {
		klog.Warningf("unable to release host: %s, error: %v",
			aws.StringValue(host.ResourceId), aws.StringValue(host.Error.Message))
	}
	return nil
}

func (e *AwsEC2) listAvailableDedicatedHosts() ([]*ec2.Host, error) {
	describeOutput, err := e.client.DescribeHosts(&ec2.DescribeHostsInput{
		Filter: []*ec2.Filter{
			{
				Name:   aws.String("state"),
				Values: []*string{aws.String(AwsInstanceAvailableState)},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return describeOutput.Hosts, nil
}

func (e *AwsEC2) retrieveOrAllocateHost(node *api.Node) (*string, error) {
	describeOutput, err := e.client.DescribeHosts(&ec2.DescribeHostsInput{
		Filter: []*ec2.Filter{
			{
				Name:   aws.String("state"),
				Values: aws.StringSlice([]string{AwsInstanceAvailableState}),
			}, {
				Name:   aws.String("instance-type"),
				Values: aws.StringSlice([]string{node.Spec.InstanceType}),
			}, {
				Name:   aws.String("availability-zone"),
				Values: aws.StringSlice([]string{e.availabilityZone}),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// if there is a host available check if it has capacity
	if len(describeOutput.Hosts) > 0 {
		// check if host has enough capacity to run instance
		for _, host := range describeOutput.Hosts {
			if aws.Int64Value(host.AvailableCapacity.AvailableVCpus) > 0 {
				// should we check resource requirements here?
				return host.HostId, nil
			}
		}
	}
	// if no host has availability we will allocate a new host to the account
	// and return the id of the newly allocated host
	allocateOutput, err := e.client.AllocateHosts(&ec2.AllocateHostsInput{
		AutoPlacement:    aws.String("on"),
		AvailabilityZone: aws.String(e.availabilityZone),
		InstanceType:     aws.String(node.Spec.InstanceType),
		Quantity:         aws.Int64(1),
	})
	if err != nil {
		return nil, err
	}
	return allocateOutput.HostIds[0], nil
}

func (e *AwsEC2) isHostAvailable(hostId *string) bool {
	describeOutput, err := e.client.DescribeHosts(&ec2.DescribeHostsInput{
		HostIds: aws.StringSlice([]string{aws.StringValue(hostId)}),
	})
	if err != nil {
		klog.Errorf("cannot describe host %s: %v", *hostId, err)
		return false
	}
	if len(describeOutput.Hosts) != 1 {
		return false
	}
	host := describeOutput.Hosts[0]
	return aws.StringValue(host.State) == AwsInstanceAvailableState
}

func (e *AwsEC2) waitForHostAvailable(ctx context.Context, hostId *string) bool {
	hostAvailable := false
	ticker := time.NewTicker(1*time.Second)
	for {
		select {
		case <-ctx.Done():
			return hostAvailable
		case <-ticker.C:
			klog.V(2).Infof("checking host %s availability...", *hostId)
			hostAvailable = e.isHostAvailable(hostId)
			if hostAvailable {
				return hostAvailable
			}
			klog.V(2).Infof("host not available yet")
		default:
			//
		}
	}
}

func (e *AwsEC2) StartDedicatedNode(node *api.Node, image cloud.Image, metadata, iamPermissions string) (string, error) {
	klog.V(2).Infof("Starting instance for node: %v", node)
	hostId, err := e.retrieveOrAllocateHost(node)
	if err != nil {
		return "", fmt.Errorf("Could not retrieve or allocate dedicated host: %v", err)
	}
	tags := e.getNodeTags(node)
	tagSpec := ec2.TagSpecification{
		ResourceType: aws.String("instance"),
		Tags:         tags,
	}
	volSizeGiB := cloud.ToSaneVolumeSize(node.Spec.Resources.VolumeSize, image)
	klog.V(2).Infof("calculated volume size for node: %v", volSizeGiB)
	devices := e.getBlockDeviceMapping(image, volSizeGiB)
	networkSpec := e.getInstanceNetworkSpec(node.Spec.Resources.PrivateIPOnly)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	hostAvailable := e.waitForHostAvailable(ctx, hostId)
	if !hostAvailable {
		return "", util.WrapError(err, "Could not run instance: host %s not available", *hostId)
	}

	klog.V(2).Infof("Starting node with security groups: %v subnet: '%s'",
		e.bootSecurityGroupIDs, e.subnetID)
	klog.V(2).Info("Block devices for a node")
	for _, device := range devices {
		klog.V(2).Infof("Device: %s volume size: %d", *device.DeviceName, *device.Ebs.VolumeSize)
	}
	result, err := e.client.RunInstances(&ec2.RunInstancesInput{
		ImageId:             aws.String(node.Spec.BootImage),
		InstanceType:        aws.String(node.Spec.InstanceType),
		MinCount:            aws.Int64(1),
		MaxCount:            aws.Int64(1),
		TagSpecifications:   []*ec2.TagSpecification{&tagSpec},
		NetworkInterfaces:   networkSpec,
		BlockDeviceMappings: devices,
		UserData:            aws.String(metadata),
		IamInstanceProfile:  getIAMInstanceProfileSpecification(iamPermissions),
		Placement: &ec2.Placement{
			HostId:  hostId,
			Tenancy: aws.String("host"),
		},
	})
	if err != nil {
		if isSubnetConstrainedError(err) {
			return "", &cloud.NoCapacityError{
				OriginalError: err.Error(),
				SubnetID:      e.subnetID,
			}
		} else if isAZConstrainedError(err) || isInstanceConstrainedError(err) {
			return "", &cloud.NoCapacityError{
				OriginalError: err.Error(),
			}
		}
		return "", util.WrapError(err, "Could not run instance")
	}
	if len(result.Instances) == 0 {
		return "", fmt.Errorf("Could not get instance info at result.Instances")
	}
	cloudID := aws.StringValue(result.Instances[0].InstanceId)
	klog.V(2).Infof("Started instance: %s", cloudID)
	return cloudID, nil
}

func getIAMInstanceProfileSpecification(iamPermissions string) *ec2.IamInstanceProfileSpecification {
	if iamPermissions == "" {
		return nil
	}
	profile := &ec2.IamInstanceProfileSpecification{}
	if strings.Contains(iamPermissions, ":") {
		profile.Arn = aws.String(iamPermissions)
	} else {
		profile.Name = aws.String(iamPermissions)
	}
	return profile
}

// This isn't terribly different from Start node but there are
// some minor differences.  We'll capture errors correctly here and there
func (e *AwsEC2) StartSpotNode(node *api.Node, image cloud.Image, metadata, iamPermissions string) (string, error) {
	klog.V(2).Infof("Starting instance for node: %v", node)
	tags := e.getNodeTags(node)
	tagSpec := ec2.TagSpecification{
		ResourceType: aws.String("instance"),
		Tags:         tags,
	}
	var err error
	//var subnet *cloud.SubnetAttributes
	klog.V(2).Infof("Starting spot node in: %s", e.subnetID)
	volSizeGiB := cloud.ToSaneVolumeSize(node.Spec.Resources.VolumeSize, image)
	devices := e.getBlockDeviceMapping(image, volSizeGiB)
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
		IamInstanceProfile: getIAMInstanceProfileSpecification(iamPermissions),
	})

	if err != nil {
		if isSubnetConstrainedError(err) {
			return "", &cloud.NoCapacityError{
				OriginalError: err.Error(),
				SubnetID:      e.subnetID,
			}
		} else if isAZConstrainedError(err) || isInstanceConstrainedError(err) {
			return "", &cloud.NoCapacityError{
				OriginalError: err.Error(),
			}
		} else if isUnsupportedInstanceError(err) {
			return "", &cloud.UnsupportedInstanceError{err.Error()}
		}
		return "", util.WrapError(err, "Could not run instance")
	}
	if len(result.Instances) == 0 {
		return "", fmt.Errorf("Could not get instance info at result.Instances")
	}
	cloudID := aws.StringValue(result.Instances[0].InstanceId)
	klog.V(2).Infof("Started instance: %s", cloudID)
	return cloudID, nil
}

func (e *AwsEC2) WaitForRunning(node *api.Node) ([]api.NetworkAddress, error) {
	awsInstanceIDs := []*string{&node.Status.InstanceID}
	dii := &ec2.DescribeInstancesInput{InstanceIds: awsInstanceIDs}
	// Due to eventual consistency, after we create an instance and
	// get its instanceID back from RunInstances, the rest of AWS
	// might not know about that instanceID yet.
	err := util.Retry(
		BootTimeout,
		func() error {
			waitErr := e.client.WaitUntilInstanceRunning(dii)
			return waitErr
		},
		func(err error) bool {
			// TODO: remove this log line
			klog.V(2).Infof("retrying err: %s", err.Error())
			return strings.HasPrefix(err.Error(), "InvalidInstanceID.NotFound") || strings.Contains(err.Error(), "ResourceNotReady")
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
	if !node.Spec.Resources.PrivateIPOnly && e.usePublicIPs {
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

func (e *AwsEC2) AddIAMPermissions(node *api.Node, instanceProfile string) error {
	instanceID := node.Status.InstanceID
	out, err := e.client.DescribeIamInstanceProfileAssociations(
		&ec2.DescribeIamInstanceProfileAssociationsInput{
			Filters: []*ec2.Filter{
				{
					Name:   aws.String("instance-id"),
					Values: aws.StringSlice([]string{instanceID}),
				},
			},
		})
	if err != nil {
		return util.WrapError(err, "DescribeIamInstanceProfileAssociations() "+instanceID)
	}

	if out == nil || len(out.IamInstanceProfileAssociations) == 0 {
		klog.V(5).Infof("adding IAM instance profile %s to instance %s", instanceProfile, instanceID)
		_, err = e.client.AssociateIamInstanceProfile(
			&ec2.AssociateIamInstanceProfileInput{
				IamInstanceProfile: getIAMInstanceProfileSpecification(instanceProfile),
				InstanceId:         aws.String(instanceID),
			})
		if err != nil {
			return util.WrapError(err, "AssociateIamInstanceProfile() "+instanceID)
		}
		klog.V(5).Infof("added IAM instance profile %s to %s", instanceProfile, instanceID)
		return nil
	}

	n := len(out.IamInstanceProfileAssociations)
	if n > 1 {
		klog.Warningf("DescribeIamInstanceProfileAssociations() %s: got %d results", instanceID, n)
	}

	for _, association := range out.IamInstanceProfileAssociations {
		if association == nil {
			return fmt.Errorf("DescribeIamInstanceProfileAssociations() %s: nil association", instanceID)
		}
		klog.V(5).Infof("replacing current IAM instance profile %s on %s association ID %s state %s",
			aws.StringValue(association.IamInstanceProfile.Arn),
			aws.StringValue(association.InstanceId),
			aws.StringValue(association.AssociationId),
			aws.StringValue(association.State))
		_, err = e.client.ReplaceIamInstanceProfileAssociation(
			&ec2.ReplaceIamInstanceProfileAssociationInput{
				AssociationId:      association.AssociationId,
				IamInstanceProfile: getIAMInstanceProfileSpecification(instanceProfile),
			})
		if err != nil {
			return util.WrapError(err, "ReplaceIamInstanceProfileAssociation() "+instanceID)
		}
		klog.V(5).Infof("replaced IAM instance profile association %s to %s for %s",
			aws.StringValue(association.AssociationId), instanceProfile, instanceID)
		break
	}

	return nil
}
