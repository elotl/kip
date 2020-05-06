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
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/util"
	"google.golang.org/api/compute/v1"
)

func (c *gceClient) SetBootSecurityGroupIDs(ids []string) {
	c.bootSecurityGroupIDs = ids
}

func (c *gceClient) GetBootSecurityGroupIDs() []string {
	return c.bootSecurityGroupIDs
}

func (c *gceClient) getRuleDescription() string {
	t := time.Now().Format(time.RFC3339)
	return fmt.Sprintf("Created by kip controller %s at %s", c.controllerID, t)
}

func (c *gceClient) EnsureSecurityGroup(sgName string, ports []cloud.InstancePort, sourceRanges []string) (*cloud.SecurityGroup, error) {
	sg, err := c.FindSecurityGroup(sgName)
	if err != nil {
		return nil, util.WrapError(err, "error finding security group")
	}
	if sg == nil {
		return c.CreateSecurityGroup(sgName, ports, sourceRanges)
	}
	err = c.UpdateSecurityGroup(*sg, ports, sourceRanges)
	if err != nil {
		return nil, util.WrapError(
			err, "could not merge new rules into existing security group")
	}
	// We have seen eventual consistency errors here, retry it if we
	// can't find the group
	for i := 0; i < 10; i++ {
		sg, err = c.FindSecurityGroup(sgName)
		if sg != nil || err != nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	if sg == nil && err == nil {
		err = fmt.Errorf("could not find security group %s after creation", sgName)
	}
	return sg, err
}

func formatInstancePort(port, portRangeSize int) string {
	portStr := strconv.Itoa(port)
	if portRangeSize > 1 {
		endPort := strconv.Itoa(port + portRangeSize - 1)
		portStr += "-" + endPort
	}
	return portStr
}

func gceProtocolToKipProtocol(proto string) api.Protocol {
	proto = strings.ToLower(proto)
	switch proto {
	case "tcp":
		return api.ProtocolTCP
	case "udp":
		return api.ProtocolUDP
	case "sctp":
		return api.ProtocolSCTP
	case "ICMP":
		return api.ProtocolICMP
	}
	return api.Protocol("")
}

func portStringToInstancePort(p string) (int, int) {
	parts := strings.Split(p, "-")
	portStr := "0"
	endPortStr := ""
	if len(parts) == 1 {
		portStr = parts[0]
	} else if len(parts) >= 2 {
		portStr = parts[0]
		endPortStr = parts[1]
	}
	size := 1
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Printf("error converting port %s: %v", portStr, err)
	}
	if endPortStr != "" {
		endPort, err := strconv.Atoi(endPortStr)
		if err != nil {
			log.Printf("error converting ending port of range %s: %v", endPortStr, err)
		}
		if endPort > port {
			size = endPort - port + 1
		}
	}
	return port, size
}

func allowedRulesToPorts(fwa []*compute.FirewallAllowed) []cloud.InstancePort {
	instancePorts := make([]cloud.InstancePort, 0, len(fwa))
	for _, fw := range fwa {
		if fw == nil {
			continue
		}
		proto := gceProtocolToKipProtocol(fw.IPProtocol)
		if proto == "" {
			// Todo: output a warning
			continue
		}
		for _, port := range fw.Ports {
			port, size := portStringToInstancePort(port)
			instancePorts = append(instancePorts, cloud.InstancePort{
				Protocol:      proto,
				Port:          port,
				PortRangeSize: size,
			})
		}
	}
	return instancePorts
}

func firewallToSecurityGroup(fw *compute.Firewall) *cloud.SecurityGroup {
	return &cloud.SecurityGroup{
		Name:         fw.Name,
		ID:           fw.Name,
		Ports:        allowedRulesToPorts(fw.Allowed),
		SourceRanges: fw.SourceRanges,
	}
}

func (c *gceClient) toFirewallRule(sgName string, ports []cloud.InstancePort, sourceRanges []string) *compute.Firewall {
	return &compute.Firewall{
		Allowed:      portsToAllowedRules(ports),
		Description:  c.getRuleDescription(),
		Direction:    "INGRESS",
		Name:         sgName,
		Network:      c.getNetworkURL(),
		SourceRanges: sourceRanges,
		TargetTags:   []string{CreateKipCellNetworkTag(c.controllerID)},
	}
}

func (c *gceClient) EnsureMilpaSecurityGroups(extraCIDRs, extraGroupIDs []string) error {
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
	cidrs := make([]string, len(c.vpcCIDRs))
	copy(cidrs, c.vpcCIDRs)
	for _, cidr := range extraCIDRs {
		if cidr != "" {
			cidrs = append(cidrs, cidr)
		}
	}

	apiGroupName := CreateKipCellNetworkTag(c.controllerID)
	apiGroup, err := c.EnsureSecurityGroup(apiGroupName, milpaPorts, cidrs)
	if err != nil {
		return util.WrapError(err, "could not setup Kip API cloud firewall rules")
	}
	ids := append(extraGroupIDs, apiGroup.ID)
	c.SetBootSecurityGroupIDs(ids)
	return nil
}

func (c *gceClient) FindSecurityGroup(sgName string) (*cloud.SecurityGroup, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	resp, err := c.service.Firewalls.Get(c.projectID, sgName).Context(ctx).Do()
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	if resp == nil {
		return nil, nilResponseError("Firewalls.Get")
	}

	sg := firewallToSecurityGroup(resp)
	return sg, nil
}

func portsToAllowedRules(ports []cloud.InstancePort) []*compute.FirewallAllowed {
	allowed := make([]*compute.FirewallAllowed, len(ports))
	for i := range ports {
		proto := strings.ToLower(string(ports[i].Protocol))
		portString := formatInstancePort(ports[i].Port, ports[i].PortRangeSize)
		allowed[i] = &compute.FirewallAllowed{
			IPProtocol: proto,
		}
		if proto == "tcp" || proto == "udp" || proto == "sctp" {
			allowed[i].Ports = []string{portString}
		}
	}
	return allowed
}

func (c *gceClient) UpdateSecurityGroup(cloudSG cloud.SecurityGroup, ports []cloud.InstancePort, sourceRanges []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	rule := c.toFirewallRule(cloudSG.Name, ports, sourceRanges)
	resp, err := c.service.Firewalls.Patch(c.projectID, cloudSG.Name, rule).Context(ctx).Do()
	if err != nil {
		return err
	}
	if resp == nil {
		return nilResponseError("Firewalls.Patch")
	}
	if err := c.waitOnOperation(resp.Name, c.getGlobalOperation); err != nil {
		return err
	}
	return nil
}

func (c *gceClient) CreateSecurityGroup(sgName string, ports []cloud.InstancePort, sourceRanges []string) (*cloud.SecurityGroup, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	rule := c.toFirewallRule(sgName, ports, sourceRanges)
	op, err := c.service.Firewalls.Insert(c.projectID, rule).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	if err := c.waitOnOperation(op.Name, c.getGlobalOperation); err != nil {
		return nil, err
	}
	sg := cloud.NewSecurityGroup(sgName, sgName, ports, sourceRanges)
	return &sg, nil
}

func (c *gceClient) AttachSecurityGroups(node *api.Node, groups []string) error {
	// Todo, pull this from short term instance cache
	inst, err := c.getInstanceSpec(node.Status.InstanceID)
	if err != nil {
		return util.WrapError(err, "error retrieving instance's network tags fingerprint from GKE")
	}
	fp := ""
	if inst.Tags != nil {
		fp = inst.Tags.Fingerprint
	}
	allTags := append(groups, c.bootSecurityGroupIDs...)
	rb := &compute.Tags{
		Fingerprint: fp,
		Items:       allTags,
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	op, err := c.service.Instances.SetTags(c.projectID, c.zone, node.Status.InstanceID, rb).Context(ctx).Do()
	if err != nil {
		return util.WrapError(err, "Error attaching instance tags %s", err)
	}
	if err := c.waitOnOperation(op.Name, c.getGlobalOperation); err != nil {
		return err
	}
	return nil
}
