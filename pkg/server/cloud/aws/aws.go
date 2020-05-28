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
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/util"
	"github.com/elotl/kip/pkg/util/errors"
	"k8s.io/klog"
)

const (
	awsTimeout = 10 * time.Second
)

var (
	maxAWSUserTags = 45
)

type AwsEC2 struct {
	client               *ec2.EC2
	elb                  *elb.ELB
	ecs                  *ecs.ECS
	ecsClusterName       string
	controllerID         string
	nametag              string
	vpcID                string
	vpcCIDR              string
	subnetID             string
	availabilityZone     string
	usePublicIPs         bool
	region               string
	bootSecurityGroupIDs []string
	cloudStatus          *cloud.LinkedAZSubnetStatus
}

func getEC2Client() (*ec2.EC2, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, util.WrapError(err, "Error creating EC2 client session")
	}

	config := aws.NewConfig()
	config = config.WithHTTPClient(&http.Client{
		Timeout: awsTimeout,
	})
	ec2Client := ec2.New(sess, config)
	return ec2Client, nil
}

// Todo, see if we can either share clients or cut down on the
// copypaste.
func getELBClient() (*elb.ELB, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, util.WrapError(err, "Error creating ELB client session")
	}

	config := aws.NewConfig()
	config = config.WithHTTPClient(&http.Client{
		Timeout: awsTimeout,
	})
	client := elb.New(sess, config)
	return client, nil

}

func getECSClient() (*ecs.ECS, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, util.WrapError(err, "Error creating ECS client session")
	}
	config := aws.NewConfig()
	config = config.WithHTTPClient(&http.Client{
		Timeout: awsTimeout,
	})
	client := ecs.New(sess, config)
	return client, nil
}

func CheckConnection() error {
	client, err := getEC2Client()
	if err != nil {
		return util.WrapError(err, "Check connection failed setting up an ec2 client")
	}
	klog.V(2).Infof("Checking for credential errors")
	val, err := client.Config.Credentials.Get()
	if err != nil {
		return util.WrapError(err, "Error validating AWS credentials")
	}
	klog.V(2).Infof("Using credentials from %s", val.ProviderName)
	// Validate that region is set. I'm pretty sure that all our
	// authentication methods need this to be set.
	if client.Config.Region == nil || *client.Config.Region == "" {
		return fmt.Errorf("Empty region in AWS configuraiton, please specify a region in the config file or environment")
	}
	klog.V(2).Infof("Validating read access")
	_, err = client.DescribeInstances(nil)
	return err
}

type EC2ClientConfig struct {
	ControllerID   string
	Nametag        string
	VPCID          string
	SubnetID       string
	ECSClusterName string
	PrivateIPOnly  bool
}

// Parsing our server.json configuration should have put all confg
// into environment variables, load necessary config from there.
func NewEC2Client(config EC2ClientConfig) (*AwsEC2, error) {
	if config.ControllerID == "" {
		return nil, fmt.Errorf("ControllerID is a required parameter")
	}
	if config.Nametag == "" {
		return nil, fmt.Errorf("Nametag is a required parameter")
	}
	ec2Client, err := getEC2Client()
	if err != nil {
		return nil, util.WrapError(err, "Error creating EC2 client")
	}
	elbClient, err := getELBClient()
	if err != nil {
		return nil, util.WrapError(err, "Error creating ELB client")
	}
	var ecsClient *ecs.ECS
	if config.ECSClusterName != "" {
		ecsClient, err = getECSClient()
		if err != nil {
			return nil, util.WrapError(err, "Error creating ECS client")
		}
	}
	client := &AwsEC2{
		client:         ec2Client,
		elb:            elbClient,
		ecs:            ecsClient,
		ecsClusterName: config.ECSClusterName,
		controllerID:   config.ControllerID,
		nametag:        config.Nametag,
	}
	client.vpcID, client.vpcCIDR, err = client.assertVPCExists(config.VPCID)
	if err != nil {
		return nil, err
	}
	client.subnetID = config.SubnetID
	if client.subnetID == "" {
		client.subnetID, err = detectCurrentSubnet()
		if err != nil {
			return nil, util.WrapError(err, "Could not detect current subnet from metadata service. Please supply an AWS subnet id in provider.yaml")
		}
	}
	client.region = os.Getenv("AWS_REGION")
	client.cloudStatus, err = cloud.NewLinkedAZSubnetStatus(client)
	if err != nil {
		return nil, util.WrapError(
			err, "Error setting up cloud status keeper")
	}

	subnetAttrs, err := client.getSubnetAttributes()
	if err != nil {
		return nil, util.WrapError(err, "Error getting subnet attributes")
	}
	client.availabilityZone = subnetAttrs.AZ
	client.usePublicIPs = !config.PrivateIPOnly
	if subnetAttrs.AddressAffinity == cloud.PrivateAddress {
		klog.V(2).Infoln("cells will run in a private subnet (no route to internet gateway)")
		client.usePublicIPs = false
	}
	return client, nil
}

func (c *AwsEC2) getSubnetAttributes() (cloud.SubnetAttributes, error) {
	var sn cloud.SubnetAttributes
	subnets, err := c.GetSubnets()
	if err != nil {
		return sn, err
	}
	if len(subnets) == 0 {
		return sn, fmt.Errorf("no subnets found")
	}
	for _, sn := range subnets {
		if sn.ID == c.subnetID {
			return sn, nil
		}
	}
	return sn, fmt.Errorf("could not match the provided subnetID %s to any subnet in the VPC", c.subnetID)
}

func (c *AwsEC2) CloudStatusKeeper() cloud.StatusKeeper {
	return c.cloudStatus
}

func (c *AwsEC2) GetVPCCIDRs() []string {
	return []string{c.vpcCIDR}
}

func (m *AwsEC2) GetAttributes() cloud.CloudAttributes {
	return cloud.CloudAttributes{
		DiskProductName: api.StorageGP2,
		FixedSizeVolume: false,
		Provider:        cloud.ProviderAWS,
		Region:          m.region,
	}
}

func filterLabelsForTags(resource string, labels map[string]string) (map[string]string, error) {
	illegalKeys := []string{"Node", cloud.ControllerTagKey}
	ignoredPrefixes := []string{"aws:"}
	ignoredPrefixes = append(ignoredPrefixes, util.InternalLabelPrefixes...)
	i := 0
	allErrs := []error{}
	filteredLabels := make(map[string]string)
	for k, v := range labels {
		i++
		// constraints:
		// instances <= 50 tags (reserve 5 for milpa)
		// key - 127 chars name cannot begin with "aws:"
		// value - 255 chars
		// Key can't be one of our internal Milpa tag keys
		if i > maxAWSUserTags {
			e := fmt.Errorf("error tagging resource %s: Users are limited to %d tags", resource, maxAWSUserTags)
			allErrs = append(allErrs, e)
			break
		}
		ignoreKey := false
		for _, ignoredPrefix := range ignoredPrefixes {
			if strings.HasPrefix(k, ignoredPrefix) {
				ignoreKey = true
				break
			}
		}
		if ignoreKey {
			continue
		}
		if util.StringInSlice(k, illegalKeys) {
			allErrs = append(allErrs,
				fmt.Errorf("error tagging instance %s, illegal keys: %v",
					resource, illegalKeys))
			continue
		}
		if len(k) > 127 {
			allErrs = append(allErrs,
				fmt.Errorf("error tagging instance %s, keys are limited to 127 chars", resource))
			continue
		}
		if len(v) > 255 {
			allErrs = append(allErrs,
				fmt.Errorf("error tagging instance %s, values are limited to 255 chars", resource))
			continue
		}
		filteredLabels[k] = v
	}
	podName := labels[cloud.PodNameTagKey]
	if podName != "" {
		if labels["Name"] == "" {
			filteredLabels["Name"] = podName
		}
		delete(filteredLabels, cloud.PodNameTagKey)
	}
	var err error
	if len(allErrs) > 0 {
		err = errors.NewAggregate(allErrs)
	}
	return filteredLabels, err
}

func ec2TagsFromLabels(resource string, labels map[string]string) ([]*ec2.Tag, error) {
	filteredLabels, err := filterLabelsForTags(resource, labels)
	awsTags := make([]*ec2.Tag, 0, len(labels))
	for k, v := range filteredLabels {
		awsTags = append(awsTags, &ec2.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		})
	}
	return awsTags, err
}
