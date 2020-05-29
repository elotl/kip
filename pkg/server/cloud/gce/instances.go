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
			k = replaceReservedLabelChars(k)
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
		klog.Errorf("error retrieving instance specification: %v", err)
		return nil, err
	}
	return instance, nil
}

func (c *gceClient) getInstanceStatus(instanceID string) (string, error) {
	instance, err := c.getInstanceSpec(instanceID)
	if err != nil {
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

func (c *gceClient) createInstanceSpec(node *api.Node, image cloud.Image, metadata string) (*compute.Instance, error) {
	decodedUserData, err := base64.StdEncoding.DecodeString(metadata)
	if err != nil {
		return nil, util.WrapError(err, "Could not decode metadata string")
	}

	name := makeInstanceID(c.controllerID, node.Name)
	diskType := c.getDiskTypeURL()
	volSizeGiB := cloud.ToSaneVolumeSize(node.Spec.Resources.VolumeSize)
	disks := c.getAttachedDiskSpec(true, int64(volSizeGiB), name, diskType, image.Name)
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
			Items: c.createInstanceMetadata(string(decodedUserData)),
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

func (c *gceClient) createInstanceMetadata(decodedUserData string) []*compute.MetadataItems {
	// disableIPAliases disables IP alias issues from the google network daemon
	disableIPAliases := `#!/bin/bash
    readonly inst_cfg_file="/etc/default/instance_configs.cfg.template"

    if [[ ! -f "$inst_cfg_file" ]]; then
        touch "$inst_cfg_file"
    fi

cat << EOF >> "$inst_cfg_file"
[IpForwarding]
ip_aliases = false
EOF

    $(/usr/bin/google_instance_setup)
    `
	items := []*compute.MetadataItems{
		{
			Key:   "user-data",
			Value: &decodedUserData,
		},
		{
			Key:   "startup-script",
			Value: &disableIPAliases,
		},
	}
	for k, v := range c.gkeMetadata {
		val := v
		items = append(items, &compute.MetadataItems{
			Key:   k,
			Value: &val,
		})
	}
	return items
}

// this function handles the starting of both regular and spot type instances
// it is called in the exported StartNode and StartSpotNode functions
func (c *gceClient) startNode(node *api.Node, image cloud.Image, metadata string) (*cloud.StartNodeResult, error) {
	klog.V(2).Infof("Starting instance for node: %v", node)
	spec, err := c.createInstanceSpec(node, image, metadata)
	if err != nil {
		return nil, err
	}
	klog.V(2).Infof("Starting node with security groups: %v subnet: '%s'",
		c.bootSecurityGroupIDs, c.subnetName)
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	op, err := c.service.Instances.Insert(c.projectID, c.zone, spec).Context(ctx).Do()
	if err != nil {
		return nil, util.WrapError(err, "startup error")
	}
	// Todo: catch and convert errors to notify us of
	// out of capacity errors or invalid machine types
	// see pkg/server/cloud/aws/instances.StartNode()
	if err := waitOnOperation(op.Name, c.getZoneOperation); err != nil {
		return nil, err
	}
	startResult := &cloud.StartNodeResult{
		InstanceID:       spec.Name,
		AvailabilityZone: c.zone,
	}
	return startResult, nil
}

func (c *gceClient) StartNode(node *api.Node, image cloud.Image, metadata string) (*cloud.StartNodeResult, error) {
	return c.startNode(node, image, metadata)
}

// In we dictate whether the node is a spot based on the node passed in
// this is decided in createInstanceSpec which is called in the unexported
// startNode function. StartSpotNode is necessary to fullfil the interface.
func (c *gceClient) StartSpotNode(node *api.Node, image cloud.Image, metadata string) (*cloud.StartNodeResult, error) {
	return c.startNode(node, image, metadata)
}

func (c *gceClient) WaitForRunning(node *api.Node) ([]api.NetworkAddress, error) {
	start := time.Now()
	eventualConsistencyTimeout := 30 * time.Second
	for {
		status, err := c.getInstanceStatus(node.Status.InstanceID)
		if err != nil {
			// We likely called this right after calling StartNode
			// Allow for eventual consistency hickups
			if isNotFoundError(err) &&
				time.Since(start) < eventualConsistencyTimeout {
				time.Sleep(5 * time.Second)
				continue
			}
			klog.Errorf("Error waiting for instance to start: %v", err)
			return nil, err
		}
		if time.Since(start) > waitForRunningTimeout {
			return nil, fmt.Errorf("WaitForRunning timeout for instance %s after %s", node.Status.InstanceID, waitForRunningTimeout.String())
		}
		klog.V(5).Infof("status: %s", status)
		if status == statusInstanceRunning {
			break
		}
		time.Sleep(5 * time.Second)
	}
	instance, err := c.getInstanceSpec(node.Status.InstanceID)
	if err != nil {
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
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	op, err := c.service.Instances.Delete(c.projectID, c.zone, instanceID).Context(ctx).Do()
	if err != nil {
		klog.Errorf("Error terminating instance: %v", err)
		// todo, check on status of instance, set status of instance
		// based on that, prepare to come back and clean this
		// inconsistency up
		return err
	}
	if err := waitOnOperation(op.Name, c.getZoneOperation); err != nil {
		return err
	}
	return nil
}

func (c *gceClient) getFirstVolume(instanceID string) *compute.AttachedDisk {
	instance, err := c.getInstanceSpec(instanceID)
	if err != nil {
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
	if vol == nil {
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
	if err := waitOnOperation(op.Name, c.getZoneOperation); err != nil {
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
	if err := waitOnOperation(op.Name, c.getZoneOperation); err != nil {
		return err
	}
	return nil
}

func (c *gceClient) GetImage(spec cloud.BootImageSpec) (cloud.Image, error) {
	project, ok := spec["project"]
	if !ok {
		return cloud.Image{}, fmt.Errorf(
			"project is a required boot image value. Please specify cells.bootImageSpec.project in provider.yaml")
	}
	image, ok := spec["image"]
	if !ok {
		return cloud.Image{}, fmt.Errorf(
			"image is a required boot image value. Please specify cells.bootImageSpec.image in provider.yaml")
	}
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	resp, err := c.service.Images.Get(project, image).Context(ctx).Do()
	if err != nil {
		return cloud.Image{}, util.WrapError(err, "Error looking up boot image %s/%s", project, image)
	}
	if resp == nil {
		return cloud.Image{}, nilResponseError("Images.Get")
	}
	var creationTime *time.Time
	if resp.CreationTimestamp != "" {
		ts, err := time.Parse(time.RFC3339, resp.CreationTimestamp)
		if err != nil {
			klog.Warningf(
				"invalid image creation date %s", resp.CreationTimestamp)
		} else {
			creationTime = &ts
		}
	}
	return cloud.Image{
		ID:           resp.Name,
		Name:         resp.SelfLink,
		RootDevice:   "",
		CreationTime: creationTime,
	}, nil
}

func (c *gceClient) AssignInstanceProfile(node *api.Node, instanceProfile string) error {
	klog.Errorf("In GCE Instances must be stopped to assign service account")
	return nil
}
