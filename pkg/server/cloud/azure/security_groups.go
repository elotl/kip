package azure

import (
	"context"
	"fmt"
	"math"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"k8s.io/klog"
)

func (az *AzureClient) EnsureMilpaNSG() error {
	// to keep things simple, the cluster NSG Name is the same as the
	// resourceGroupName
	nsgName := cloud.MilpaAPISGName
	resourceGroupName := controllerResourceGroupName(az.controllerID)
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	nsg, err := az.nsgs.Get(timeoutCtx, resourceGroupName, nsgName, "")
	if isNotFoundError(err) {
		parameters := network.SecurityGroup{
			Location: to.StringPtr(az.region),
		}
		timeoutCtx, cancel = context.WithTimeout(ctx, azureDefaultTimeout)
		defer cancel()
		future, err := az.nsgs.CreateOrUpdate(timeoutCtx, resourceGroupName, nsgName, parameters)
		if err != nil {
			return util.WrapError(err, "Error creating cluster network security group %s", nsgName)
		}
		klog.V(2).Infof("Creating cluster network security group %s", nsgName)
		timeoutCtx, cancel = context.WithTimeout(ctx, azureWaitTimeout)
		defer cancel()
		err = future.WaitForCompletionRef(timeoutCtx, az.nsgs.Client)
		if err != nil {
			return util.WrapError(err, "Error waiting for cluster network security group %s to be created", nsgName)
		}
		nsg, err = future.Result(az.nsgs)
		if err != nil {
			return util.WrapError(err, "Error getting result of network security gruop %s creation", nsgName)
		}
	} else if err != nil {
		return util.WrapError(err, "Error checking for existence of cluster network security group %s", nsgName)
	}
	az.nsgID = to.String(nsg.ID)
	az.nsgName = to.String(nsg.Name)
	return nil
}

func (az *AzureClient) SetBootSecurityGroupIDs(ids []string) {
	az.bootASGNames = ids
}

func (az *AzureClient) GetBootSecurityGroupIDs() []string {
	return az.bootASGNames
}

func (az *AzureClient) EnsureMilpaSecurityGroups(extraCIDRs, extraGroupIDs []string) error {
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
			Port:          1,
			PortRangeSize: 255,
		},
	}

	if len(extraGroupIDs) > 0 {
		az.SetBootSecurityGroupIDs(extraGroupIDs)
		return nil
	}
	err := az.EnsureMilpaNSG()
	if err != nil {
		return util.WrapError(err, "Error ensuring milpa network security group exists")
	}
	cidrs := append(az.virtualNetwork.CIDRs, extraCIDRs...)
	_, err = az.EnsureSecurityGroup(cloud.MilpaAPISGName, milpaPorts, cidrs)
	return err
}

func (az *AzureClient) setClusterNSG(rules []network.SecurityRule) error {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	future, err := az.nsgs.CreateOrUpdate(
		timeoutCtx,
		controllerResourceGroupName(az.controllerID),
		az.nsgName,
		network.SecurityGroup{
			Location: to.StringPtr(az.region),
			SecurityGroupPropertiesFormat: &network.SecurityGroupPropertiesFormat{
				SecurityRules: &rules,
			},
		},
	)
	if err != nil {
		return fmt.Errorf("cannot set azure cluster nsg: %v", err)
	}

	timeoutCtx, cancel = context.WithTimeout(ctx, azureWaitTimeout)
	defer cancel()
	err = future.WaitForCompletionRef(timeoutCtx, az.nsgs.Client)
	if err != nil {
		return fmt.Errorf("cannot get future response for updating azure NSG: %v", err)
	}
	_, err = future.Result(az.nsgs)
	return err
}

func (az *AzureClient) EnsureSecurityGroup(sgName string, ports []cloud.InstancePort, sourceRanges []string) (*cloud.SecurityGroup, error) {
	grpID := 0
	rules := []NSGRule{}
	for _, port := range ports {
		if port.Protocol == api.ProtocolICMP {
			continue
		}
		for _, cidr := range sourceRanges {
			ruleName := fmt.Sprintf("%s-%d", sgName, grpID)
			strPort := servicePortToString(port.Port, port.PortRangeSize)
			rule := NSGRule{
				ruleName:          ruleName,
				protocol:          port.Protocol,
				ports:             []string{strPort},
				sourceRanges:      []string{cidr},
				destinationRanges: az.virtualNetwork.CIDRs,
			}
			rules = append(rules, rule)
			grpID++
		}
	}

	azureRules := make([]network.SecurityRule, len(rules))
	for i, rule := range rules {
		priority := i + minPriority
		azureRule := rule.toAzureSecurityRule(priority)
		azureRules[i] = azureRule
	}
	err := az.setClusterNSG(azureRules)
	if err != nil {
		return nil, util.WrapError(err, "Error setting cluster network security rules")
	}

	// This return value isn't used anywhere but we'll create it
	sg := &cloud.SecurityGroup{
		ID:           az.nsgID,
		Name:         az.nsgName,
		Ports:        ports,
		SourceRanges: sourceRanges,
	}

	return sg, nil
}

func (az *AzureClient) CreateSGName(svcName string) string {
	return util.CreateSecurityGroupName(az.controllerID, "default", svcName)
}

func (az *AzureClient) AttachSecurityGroups(node *api.Node, groups []string) error {
	return nil
}
