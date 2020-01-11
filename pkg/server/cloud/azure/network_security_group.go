package azure

import (
	"fmt"
	"strconv"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/elotl/cloud-instance-provider/pkg/api"
)

var (
	minPriority = 100
	maxPriority = 4096
	denyAllName = "deny-all"
)

func servicePortToString(port int, portRangeSize int) string {
	strPort := strconv.Itoa(port)
	if portRangeSize > 1 {
		endingPort := port + portRangeSize - 1
		strPort = fmt.Sprintf("%s-%d", strPort, endingPort)
	}
	return strPort
}

//////////////////////////////////////////////////////////////////////
type NSGRule struct {
	ruleName          string
	protocol          api.Protocol
	ports             []string
	sourceRanges      []string
	destinationRanges []string
}

func (rule *NSGRule) toAzureSecurityRule(priority int) network.SecurityRule {
	protocol := network.SecurityRuleProtocolTCP
	if rule.protocol == api.ProtocolUDP {
		protocol = network.SecurityRuleProtocolUDP
	}

	sr := network.SecurityRule{
		Name: to.StringPtr(rule.ruleName),
		SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
			Protocol:                   protocol,
			SourceAddressPrefixes:      to.StringSlicePtr(rule.sourceRanges),
			SourcePortRange:            to.StringPtr("*"),
			DestinationPortRanges:      to.StringSlicePtr(rule.ports),
			Access:                     network.SecurityRuleAccessAllow,
			Direction:                  network.SecurityRuleDirectionInbound,
			Priority:                   to.Int32Ptr(int32(priority)),
			DestinationAddressPrefixes: to.StringSlicePtr(rule.destinationRanges),
		},
	}
	return sr
}

func azureDenyAllRule() network.SecurityRule {
	return network.SecurityRule{
		Name: to.StringPtr(denyAllName),
		SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
			Protocol:                 network.SecurityRuleProtocolAsterisk,
			SourceAddressPrefix:      to.StringPtr("*"),
			SourcePortRange:          to.StringPtr("*"),
			DestinationAddressPrefix: to.StringPtr("*"),
			DestinationPortRange:     to.StringPtr("*"),
			Access:                   network.SecurityRuleAccessDeny,
			Direction:                network.SecurityRuleDirectionInbound,
			Priority:                 to.Int32Ptr(int32(maxPriority)),
		},
	}
}
