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
	"fmt"
	"time"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/util"
	"k8s.io/klog"

	"google.golang.org/api/compute/v1"
)

func (c *gceClient) getNodeSpec(instanceID string) (*compute.Instance, error) {
	instance, err := c.service.Instances.Get(c.projectID, c.zone, instanceID).Do()
	if err != nil {
		// TODO error handling for googleapi errors
		klog.Errorf("error retrieving instance specification: %v", err)
		return nil, err
	}
	return instance, nil
}

func (c *gceClient) getNodeStatus(instanceID string) (string, error) {
	instance, err := c.getNodeSpec(instanceID)
	if err != nil {
		// TODO error handling for googleapi errors
		klog.Errorf("error retrieving instance status: %v", err)
		return "", err
	}
	return instance.Status, nil
}

func (c *gceClient) getFirstVolume(instanceID string) *compute.AttachedDisk {
	instance, err := c.getNodeSpec(instanceID)
	if err != nil {
		// TODO error handling for googleapi errors
		klog.Errorf("error retrieving instance volume: %v", err)
		return nil
	}

	volumes := instance.Disks
	if len(volumes) == 0 {
		return nil
	}

	return volumes[0]
}

func (c *gceClient) getNodeLabels() map[string]string {
	// TODO this is different from the one in utils in that it
	// uses unix timestamps to accommodate gcp naming convention
	nametag := c.createUnboundNodeNameTag()
	return map[string]string{
		"name": nametag,
		"node": c.nametag,
	}
}

func (c *gceClient) getAttachedDisk(isBoot bool, size int64, name, typeURL, imageURL string) []*compute.AttachedDisk {
	diskSpec := []*compute.AttachedDisk{
		{
			Boot:       isBoot,
			AutoDelete: true,
			// RW or RO
			Mode: "READ_WRITE",
			// SCSI or NVMe
			Interface: "SCSI",
			InitializeParams: &compute.AttachedDiskInitializeParams{
				DiskName:    name,
				DiskSizeGb:  size,
				DiskType:    typeURL,
				SourceImage: imageURL,
			},
		},
	}
	return diskSpec
}

func (c *gceClient) getInstanceNetworkSpec(privateIPOnly bool) []*compute.NetworkInterface {
	networkURL := c.getNetworkURL()
	subNetworkURL := c.getSubNetworkURL()

	accessConfig := &compute.AccessConfig{
		Name:        "External NAT",
		NetworkTier: "STANDARD",
		Type:        "ONE_TO_ONE_NAT",
	}
	// If we are only using private IP set the access config to nil
	if privateIPOnly || !c.usePublicIPs {
		accessConfig = nil
	}
	networkSpec := []*compute.NetworkInterface{
		{
			AccessConfigs: []*compute.AccessConfig{accessConfig},
			AliasIpRanges: []*compute.AliasIpRange{
				{
					IpCidrRange: "/32",
				},
			},
			Network:    networkURL,
			Subnetwork: subNetworkURL,
		},
	}
	return networkSpec
}

func (c *gceClient) StartNode(node *api.Node, metadata string) (*cloud.StartNodeResult, error) {
	klog.V(2).Infof("Starting instance for node: %v", node)
	bootVolName := c.nametag + "-boot-volume"
	diskType := c.getDiskTypeURL()
	volSizeGiB := cloud.ToSaneVolumeSize(node.Spec.Resources.VolumeSize)
	disks := c.getAttachedDisk(true, int64(volSizeGiB), bootVolName, diskType, node.Spec.BootImage)
	labels := c.getNodeLabels()
	networkInterfaces := c.getInstanceNetworkSpec(node.Spec.Resources.PrivateIPOnly)
	kipNetworkTag := CreateKipCellNetworkTag(c.controllerID)
	instanceType := c.getInstanceTypeURL(node.Spec.InstanceType)
	spec := &compute.Instance{
		Disks:             disks,
		Labels:            labels,
		MachineType:       instanceType,
		Name:              c.nametag,
		NetworkInterfaces: networkInterfaces,
		Tags: &compute.Tags{
			Items: []string{kipNetworkTag},
		},
	}
	klog.V(2).Infof("Starting node with security groups: %v subnet: '%s'",
		c.bootSecurityGroupIDs, c.subnetName)
	operation, err := c.service.Instances.Insert(c.projectID, c.zone, spec).Do()
	if err != nil {
		// TODO add error checking for googleapi using helpers in util
		return nil, util.WrapError(err, "startup error")
	}
	cloudID := c.nametag
	startResult := &cloud.StartNodeResult{
		InstanceID:       cloudID,
		AvailabilityZone: c.zone,
	}
	return startResult, nil
}

func (c *gceClient) StartSpotNode(node *api.Node, metadata string) (*cloud.StartNodeResult, error) {
	klog.V(2).Infof("Starting instance for node: %v", node)
	volName := c.nametag + "-boot-volume"
	diskType := c.getDiskTypeURL()
	volSizeGiB := cloud.ToSaneVolumeSize(node.Spec.Resources.VolumeSize)
	disks := c.getAttachedDisk(true, int64(volSizeGiB), volName, diskType, node.Spec.BootImage)
	networkInterfaces := c.getInstanceNetworkSpec(node.Spec.Resources.PrivateIPOnly)
	labels := c.getNodeLabels()
	kipNetworkTag := CreateKipCellNetworkTag(c.controllerID)
	instanceType := c.getInstanceTypeURL(node.Spec.InstanceType)
	autoRestart := false
	spec := &compute.Instance{
		Disks:             disks,
		Labels:            labels,
		MachineType:       instanceType,
		Name:              c.nametag,
		NetworkInterfaces: networkInterfaces,
		Scheduling: &compute.Scheduling{
			AutomaticRestart:  &autoRestart,
			OnHostMaintenance: "TERMINATE",
			Preemptible:       true,
		},
		Tags: &compute.Tags{
			Items: []string{kipNetworkTag},
		},
	}
	klog.V(2).Infof("Starting node with security groups: %v subnet: '%s'",
		c.bootSecurityGroupIDs, c.subnetName)
	operation, err := c.service.Instances.Insert(c.projectID, c.zone, spec).Do()
	if err != nil {
		// TODO add error checking for googleapi using helpers in util
		return nil, util.WrapError(err, "startup error")
	}
	cloudID := c.nametag
	startResult := &cloud.StartNodeResult{
		InstanceID:       cloudID,
		AvailabilityZone: c.zone,
	}
	return startResult, nil
}

func (c *gceClient) WaitForRunning(node *api.Node) ([]api.NetworkAddress, error) {
	for {
		status, err := c.getNodeStatus(node.Status.InstanceID)
		if err != nil {
			klog.Errorf("Error waiting for instance to start: %v", err)
			// TODO add error checking for googleapi using helpers in util
			return nil, err
		}

		klog.V(2).Infof("status: %s", status)
		if status == "RUNNING" {
			break
		}
		time.Sleep(10 * time.Second)
	}
	instance, err := c.getNodeSpec(node.Status.InstanceID)
	if err != nil {
		// TODO add error checking for googleapi using helpers in util
		return nil, err
	}

	if len(instance.NetworkInterfaces) == 0 {
		return nil, fmt.Errorf("missing private IP address(es)")
	}

	addresses := api.NewNetworkAddresses(
		instance.NetworkInterfaces[0].NetworkIP,
		instance.Hostname,
	)

	if !node.Spec.Resources.PrivateIPOnly && c.usePublicIPs {
		if len(instance.NetworkInterfaces[0].AccessConfigs) == 0 {
			return nil, fmt.Errorf("missing Public IP address")
		}

		addresses = api.SetPublicAddresses(
			instance.NetworkInterfaces[0].AccessConfigs[0].NatIP,
			instance.NetworkInterfaces[0].AccessConfigs[0].PublicPtrDomainName,
			addresses,
		)
	}

	if len(instance.NetworkInterfaces[0].AliasIpRanges) == 0 {
		return nil, fmt.Errorf("missing Pod IP address")
	}

	podIP, err := getPodIpFromCIDR(instance.NetworkInterfaces[0].AliasIpRanges[0].IpCidrRange)
	if err != nil {
		klog.Errorf("Error retrieving Pod IP: %v", err)
		return nil, err
	}
	addresses = api.SetPodIP(podIP, addresses)

	return addresses, nil
}

func (c *gceClient) StopInstance(instanceID string) error {
	_, err := c.service.Instances.Delete(c.projectID, c.zone, instanceID).Do()
	if err != nil {
		klog.Errorf("Error terminating instance: %v", err)
		// todo, check on status of instance, set status of instance
		// based on that, prepare to come back and clean this
		// inconsistency up
		return err
	}
	return nil
}

func (c *gceClient) ResizeVolume(node *api.Node, size int64) (error, bool) {
	vol := c.getFirstVolume(node.Status.InstanceID)
	if vol == nil || vol.DiskSizeGb == nil {
		return fmt.Errorf("Error retrieving volume info for node %s: %v",
			node.Name, vol), false
	}

	if *vol.DiskSizeGb > size {
		klog.V(2).Infof("Volume on node %s is %dGiB >= %dGiB",
			node.Name, *vol.Size, size)
		return nil, false
	}

	klog.V(2).Infof("Resizing volume to %dGiB for node: %v", size, node)
	resizeRequest := compute.DisksResizeRequest{SizeGb: size}
	_, err := c.client.Disks.Resize(c.projectID, c.zone, diskName, &resizeRequest).Do()
	if err != nil {
		return util.WrapError(err, "Failed to resize volume"), false
	}

	return nil, true
}

func (c *gceClient) SetSustainedCPU(node *api.Node, enabled bool) error {
	return nil
}

func (c *gceClient) ListInstancesFilterID(ids []string) ([]cloud.CloudInstance, error) {
	return nil, TODO()
}

func (c *gceClient) ListInstances() ([]cloud.CloudInstance, error) {
	return nil, TODO()
}

func (c *gceClient) AddInstanceTags(iid string, labels map[string]string) error {
	return TODO()
}

func (c *gceClient) GetImageID(spec cloud.BootImageSpec) (string, error) {
	klog.Errorln("Need to get boot image from spec")
	bootDiskImageURL := c.getProjectURL() + "/ubuntu-os-cloud/global/images/ubuntu-1804-bionic-v20200414"
	return bootDiskImageURL, nil
}

func (c *gceClient) AssignInstanceProfile(node *api.Node, instanceProfile string) error {
	return TODO()
}
