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
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/util"
	"k8s.io/klog"

	"google.golang.org/api/compute/v1"
)

func convertLabelKeys(labels map[string]string) map[string]string {
	convertedLabels := make(map[string]string)
	for k, v := range labels {
		switch k {
		case cloud.ControllerTagKey:
			k = "kip-controller-id"
		case cloud.NametagTagKey:
			k = "kip-nametag"
		case cloud.NamespaceTagKey:
			k = "kip-namespace"
		case cloud.PodNameTagKey:
			k = "kip-pod-name"
		default:
			k = strings.ToLower(k)
		}
		v = strings.ToLower(v)
		v = replaceReservedLabelChars(v)
		convertedLabels[k] = v
	}
	return convertedLabels
}

func (c *gceClient) getInstanceSpec(instanceID string) (*compute.Instance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	instance, err := c.service.Instances.Get(c.projectID, c.zone, instanceID).Context(ctx).Do()
	if err != nil {
		// TODO error handling for googleapi errors
		klog.Errorf("error retrieving instance specification: %v", err)
		return nil, err
	}
	return instance, nil
}

func (c *gceClient) getInstanceStatus(instanceID string) (string, error) {
	instance, err := c.getInstanceSpec(instanceID)
	if err != nil {
		// TODO error handling for googleapi errors
		klog.Errorf("error retrieving instance status: %v", err)
		return "", err
	}
	return instance.Status, nil
}

func (c *gceClient) getInstanceLabels(nodeName string) map[string]string {
	nametag := c.createUnboundNodeNameTag()
	return map[string]string{
		"name":             nametag,
		"node":             nodeName,
		controllerLabelKey: c.controllerID,
		nametagLabelKey:    c.nametag,
	}
}

func (c *gceClient) getAttachedDiskSpec(isBoot bool, size int64, name, typeURL, imageURL string) []*compute.AttachedDisk {
	if size < minimumDiskSize {
		klog.V(2).Infof("GCE does not allow disk smaller than 10GiB. requested size: %dGiB, using default: 10GiB", size)
		size = minimumDiskSize
	}
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

func (c *gceClient) createInstanceSpec(node *api.Node, metadata string) (*compute.Instance, error) {
	md, err := base64.StdEncoding.DecodeString(metadata)
	if err != nil {
		return nil, util.WrapError(err, "Could not decode metadata string")
	}
	mds := string(md)
	name := makeInstanceID(c.controllerID, node.Name)
	diskType := c.getDiskTypeURL()
	volSizeGiB := cloud.ToSaneVolumeSize(node.Spec.Resources.VolumeSize)
	disks := c.getAttachedDiskSpec(true, int64(volSizeGiB), name, diskType, node.Spec.BootImage)
	labels := c.getInstanceLabels(node.Name)
	networkInterfaces := c.getInstanceNetworkSpec(node.Spec.Resources.PrivateIPOnly)
	instanceType := c.getInstanceTypeURL(node.Spec.InstanceType)
	spec := &compute.Instance{
		Disks:             disks,
		Labels:            labels,
		MachineType:       instanceType,
		Name:              name,
		NetworkInterfaces: networkInterfaces,
		Tags: &compute.Tags{
			Items: c.bootSecurityGroupIDs,
		},
		Metadata: &compute.Metadata{
			Items: []*compute.MetadataItems{
				{
					Key:   "user-data",
					Value: &mds,
				},
			},
		},
	}
	if node.Spec.Spot {
		ar := false
		spec.Scheduling = &compute.Scheduling{
			AutomaticRestart:  &ar,
			OnHostMaintenance: "TERMINATE",
			Preemptible:       true,
		}
	}
	return spec, nil
}

func (c *gceClient) StartNode(node *api.Node, metadata string) (*cloud.StartNodeResult, error) {
	klog.V(2).Infof("Starting instance for node: %v", node)
	spec, err := c.createInstanceSpec(node, metadata)
	if err != nil {
		return nil, err
	}
	klog.V(2).Infof("Starting node with security groups: %v subnet: '%s'",
		c.bootSecurityGroupIDs, c.subnetName)
	op, err := c.service.Instances.Insert(c.projectID, c.zone, spec).Do()
	if err != nil {
		// TODO add error checking for googleapi using helpers in util
		return nil, util.WrapError(err, "startup error")
	}
	if err := c.waitOnOperation(op.Name, c.getZoneOperation); err != nil {
		return nil, err
	}
	startResult := &cloud.StartNodeResult{
		InstanceID:       spec.Name,
		AvailabilityZone: c.zone,
	}
	return startResult, nil
}

func (c *gceClient) StartSpotNode(node *api.Node, metadata string) (*cloud.StartNodeResult, error) {
	klog.V(2).Infof("Starting instance for node: %v", node)
	spec, err := c.createInstanceSpec(node, metadata)
	if err != nil {
		return nil, err
	}
	klog.V(2).Infof("Starting node with security groups: %v subnet: '%s'",
		c.bootSecurityGroupIDs, c.subnetName)
	op, err := c.service.Instances.Insert(c.projectID, c.zone, spec).Do()
	if err != nil {
		// TODO add error checking for googleapi using helpers in util
		return nil, util.WrapError(err, "startup error")
	}
	if err := c.waitOnOperation(op.Name, c.getZoneOperation); err != nil {
		return nil, err
	}
	startResult := &cloud.StartNodeResult{
		InstanceID:       spec.Name,
		AvailabilityZone: c.zone,
	}
	return startResult, nil
}

func (c *gceClient) WaitForRunning(node *api.Node) ([]api.NetworkAddress, error) {
	for {
		status, err := c.getInstanceStatus(node.Status.InstanceID)
		if err != nil {
			klog.Errorf("Error waiting for instance to start: %v", err)
			// TODO add error checking for googleapi using helpers in util
			return nil, err
		}

		klog.V(2).Infof("status: %s", status)
		if status == statusInstanceRunning {
			break
		}
		time.Sleep(10 * time.Second)
	}
	instance, err := c.getInstanceSpec(node.Status.InstanceID)
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
	op, err := c.service.Instances.Delete(c.projectID, c.zone, instanceID).Do()
	if err != nil {
		klog.Errorf("Error terminating instance: %v", err)
		// todo, check on status of instance, set status of instance
		// based on that, prepare to come back and clean this
		// inconsistency up
		return err
	}
	if err := c.waitOnOperation(op.Name, c.getZoneOperation); err != nil {
		return err
	}
	return nil
}

func (c *gceClient) getFirstVolume(instanceID string) *compute.AttachedDisk {
	instance, err := c.getInstanceSpec(instanceID)
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

func (c *gceClient) ResizeVolume(node *api.Node, size int64) (error, bool) {
	vol := c.getFirstVolume(node.Status.InstanceID)
	// in GCE zonal standard persistent disks cannot be smaller than 10GiB
	if vol == nil || vol.InitializeParams.DiskSizeGb < 10 {
		return fmt.Errorf("Error retrieving volume info for node %s: %v",
			node.Name, vol), false
	}

	if vol.InitializeParams.DiskSizeGb > size {
		klog.V(2).Infof("Volume on node %s is %dGiB >= %dGiB",
			node.Name, vol.InitializeParams.DiskSizeGb, size)
		return nil, false
	}

	klog.V(2).Infof("Resizing volume to %dGiB for node: %v", size, node)
	// TODO proper way to retrieve diskname
	diskName := c.nametag + "-boot-volume"
	resizeRequest := compute.DisksResizeRequest{SizeGb: size}
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	op, err := c.service.Disks.Resize(c.projectID, c.zone, diskName, &resizeRequest).Context(ctx).Do()
	if err != nil {
		return util.WrapError(err, "Failed to resize volume"), false
	}
	if err := c.waitOnOperation(op.Name, c.getZoneOperation); err != nil {
		return err, false
	}
	return nil, true
}

func (c *gceClient) SetSustainedCPU(node *api.Node, enabled bool) error {
	// Not supported in GCE return nil
	return nil
}

func (c *gceClient) ListInstancesFilterID(ids []string) ([]cloud.CloudInstance, error) {
	instances, err := c.ListInstances()
	if err != nil {
		// TODO add error checking for googleapi using helpers in util
		return nil, err
	}
	var filteredList []cloud.CloudInstance
	find := func(id string) {
		for _, ci := range instances {
			if ci.ID == id {
				filteredList = append(filteredList, ci)
			}
		}
	}
	for _, id := range ids {
		find(id)
	}
	return filteredList, nil
}

func (c *gceClient) ListInstances() ([]cloud.CloudInstance, error) {
	listCall := c.service.Instances.List(c.projectID, c.zone)
	filter := c.getLabelCompareFilter(controllerLabelKey, c.controllerID)
	listCall = listCall.Filter(filter)
	var instances []cloud.CloudInstance
	f := func(page *compute.InstanceList) error {
		for _, instance := range page.Items {
			instances = append(instances, cloud.CloudInstance{
				ID:       instance.Name,
				NodeName: instance.Labels["node"],
			})
		}
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	if err := listCall.Pages(ctx, f); err != nil {
		// TODO add error checking for googleapi using helpers in util
		return nil, err
	}
	return instances, nil
}

func (c *gceClient) AddInstanceTags(iid string, labels map[string]string) error {
	// in GCE "labels" are what AWS considers tags
	labels = convertLabelKeys(labels)

	// Todo, pull this from short term instance cache
	inst, err := c.getInstanceSpec(iid)
	if err != nil {
		return util.WrapError(err, "error retrieving instance's label fingerprint from GKE")
	}
	mergedLabels := inst.Labels
	for k, v := range labels {
		mergedLabels[k] = v
	}

	labelRequest := compute.InstancesSetLabelsRequest{
		LabelFingerprint: inst.LabelFingerprint,
		Labels:           mergedLabels,
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	op, err := c.service.Instances.SetLabels(c.projectID, c.zone, iid, &labelRequest).Context(ctx).Do()
	if err != nil {
		return util.WrapError(err, "Error attaching instance labels %s", err)
	}
	if err := c.waitOnOperation(op.Name, c.getZoneOperation); err != nil {
		return err
	}
	return nil
}

func (c *gceClient) GetImageID(spec cloud.BootImageSpec) (string, error) {
	project, ok := spec["project"]
	if !ok {
		return "", fmt.Errorf("project is a required boot image value. Please specify cells.bootImageSpec.project in provider.yaml")
	}
	image, ok := spec["image"]
	if !ok {
		return "", fmt.Errorf("image is a required boot image value. Please specify cells.bootImageSpec.image in provider.yaml")
	}
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	resp, err := c.service.Images.Get(project, image).Context(ctx).Do()
	if err != nil {
		return "", util.WrapError(err, "Error looking up boot image %s/%s", project, image)
	}
	if resp == nil {
		return "", nilResponseError("Images.Get")
	}
	return resp.SelfLink, nil
}

func (c *gceClient) AssignInstanceProfile(node *api.Node, instanceProfile string) error {
	klog.Errorf("In GCE Instances must be stopped to assign service account")
	return nil
}
