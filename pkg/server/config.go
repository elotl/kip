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

package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/api/validation"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/server/cloud/aws"
	"github.com/elotl/kip/pkg/server/cloud/azure"
	"github.com/elotl/kip/pkg/server/cloud/gce"
	"github.com/elotl/kip/pkg/server/nodemanager"
	"github.com/elotl/kip/pkg/util"
	vutil "github.com/elotl/kip/pkg/util/validation"
	"github.com/elotl/kip/pkg/util/validation/field"
	"github.com/elotl/kip/pkg/util/yaml"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog"
)

const (
	blankTemplateValue = "FILL_IN"
)

var (
	defaultStatusInterval              = 5
	defaultCloudAPIHealthCheckInterval = 60
	defaultCloudAPIHealthCheckTimeout  = 180
	defaultStatusHealthCheckTimeout    = 90

	defaultCPUCapacity    = resource.MustParse("20")
	defaultMemoryCapacity = resource.MustParse("100Gi")
	defaultPodCapacity    = resource.MustParse("20")
)

// ServerConfigFile stores the parsed json from provider.yaml
type ServerConfigFile struct {
	api.TypeMeta `json:",inline"`
	Cloud        MultiCloudConfig `json:"cloud"`
	Etcd         EtcdConfig       `json:"etcd"`
	Cells        CellsConfig      `json:"cells"`
	Testing      TestingConfig    `json:"testing"`
	Kubelet      KubeletConfig    `json:"kubelet"`
}

// Kubelet stores kubelet-specific configuration such as capacity and labels.
type KubeletConfig struct {
	// Deprecated: CPU, Memory and Pods are copied into Capacity, and are only
	// present for backward compatibility.
	CPU      *resource.Quantity `json:"cpu"`
	Memory   *resource.Quantity `json:"memory"`
	Pods     *resource.Quantity `json:"pods"`
	Capacity v1.ResourceList    `json:"capacity"`
	Labels   map[string]string  `json:"labels"`
}

type MultiCloudConfig struct {
	AWS   *AWSConfig   `json:"aws,omitempty"`
	GCE   *GCEConfig   `json:"gce,omitempty"`
	Azure *AzureConfig `json:"azure,omitempty"`
}

type AWSConfig struct {
	Region                string `json:"region"`
	AccessKeyID           string `json:"accessKeyID"`
	SecretAccessKey       string `json:"secretAccessKey"`
	VPCID                 string `json:"vpcID,omitempty"`
	SubnetID              string `json:"subnetID,omitempty"`
	EcsClusterName        string `json:"ecsClusterName"`
	EndpointURL           string `json:"endpointURL"`
	InsecureTLSSkipVerify bool   `json:"insecureTLSSkipVerify"`
}

// See https://github.com/Azure/azure-sdk-for-go/blob/master/README.md
// for more info on SDK login credentials.  We might want to support
// CertificatePath and CertificatePassword.
type AzureConfig struct {
	SubscriptionID string `json:"subscriptionID"`
	Location       string `json:"location"`
	VNetName       string `json:"virtualNetworkName"`
	ResourceGroup  string `json:"resourceGroup"`
	SubnetName     string `json:"subnetName"`
	TenantID       string `json:"tenantID"`
	ClientID       string `json:"clientID"`
	ClientSecret   string `json:"clientSecret"`
}

type GCEConfig struct {
	ProjectID       string          `json:"projectID"`
	CredentialsFile string          `json:"credentialsFile,omitempty"`
	Credentials     *GCECredentials `json:"credentials,omitempty"`
	Zone            string          `json:"zone,omitempty"`
	VPCName         string          `json:"vpcName,omitempty"`
	SubnetName      string          `json:"subnetName,omitempty"`
}

type GCECredentials struct {
	ClientEmail string `json:"clientEmail"`
	PrivateKey  string `json:"privateKey"`
}

type EtcdConfig struct {
	Client   EtcdClientConfig   `json:"client"`
	Internal InternalEtcdConfig `json:"internal"`
}

type EtcdClientConfig struct {
	Endpoints []string `json:"endpoints"`
	CertFile  string   `json:"certFile"`
	KeyFile   string   `json:"keyFile"`
	CAFile    string   `json:"caFile"`
}

type InternalEtcdConfig struct {
	DataDir    string `json:"dataDir"`
	ConfigFile string `json:"configFile"`
}

type CellsConfig struct {
	BootImageSpec          cloud.BootImageSpec           `json:"bootImageSpec"`
	DefaultInstanceType    string                        `json:"defaultInstanceType"`
	DefaultVolumeSize      string                        `json:"defaultVolumeSize"`
	StandbyCells           []nodemanager.StandbyNodeSpec `json:"standbyCells"`
	CloudInitFile          string                        `json:"cloudInitFile"`
	Itzo                   ItzoConfig                    `json:"itzo"`
	ExtraCIDRs             []string                      `json:"extraCIDRs"`
	ExtraSecurityGroups    []string                      `json:"extraSecurityGroups"`
	Nametag                string                        `json:"nametag"`
	StatusInterval         int                           `json:"statusInterval"`
	HealthCheck            HealthCheckConfig             `json:"healthcheck"`
	PrivateIPOnly          *bool                         `json:"privateIPOnly"`
	CellConfig             map[string]string             `json:"cellConfig"`
	DefaultIAMPermissions  string                        `json:"defaultIAMPermissions"`
	UseCloudParameterStore bool                          `json:"useCloudParameterStore"`
}

type HealthCheckConfig struct {
	Status   *StatusHealthCheck   `json:"status"`
	CloudAPI *CloudAPIHealthCheck `json:"cloudAPI"`
}

type StatusHealthCheck struct {
	HealthyTimeout int `json:"healthyTimeout"`
}

type CloudAPIHealthCheck struct {
	Interval       int `json:"interval"`
	HealthyTimeout int `json:"healthyTimeout"`
}

type ItzoConfig struct {
	Version string `json:"version"`
	URL     string `json:"url"`
}

type TestingConfig struct {
	ControllerID string `json:"controllerID"`
}

func serverConfigFileWithDefaults() *ServerConfigFile {
	sc := ServerConfigFile{
		TypeMeta: api.TypeMeta{
			Kind:       "serverConfig",
			APIVersion: "v1",
		},
		Etcd: EtcdConfig{
			Internal: InternalEtcdConfig{
				DataDir: "/opt/milpa/data",
			},
		},
		Cells: CellsConfig{
			BootImageSpec:     cloud.BootImageSpec{},
			StandbyCells:      []nodemanager.StandbyNodeSpec{},
			DefaultVolumeSize: "5Gi",
			StatusInterval:    defaultStatusInterval,
		},
		Kubelet: KubeletConfig{
			Capacity: v1.ResourceList{
				"cpu":    defaultCPUCapacity,
				"memory": defaultMemoryCapacity,
				"pods":   defaultPodCapacity,
			},
			Labels: map[string]string{},
		},
	}
	return &sc
}

func setEnvIfNotSet(envVar, value string) error {
	if os.Getenv(envVar) == "" && value != "" {
		err := os.Setenv(envVar, value)
		if err != nil {
			msg := fmt.Sprintf("Error setting %s env var", envVar)
			return util.WrapError(err, msg)
		}
	}
	return nil
}

func setupAWSRegion(configRegion string) error {
	winningRegionVal := configRegion
	if os.Getenv("AWS_REGION") != "" {
		winningRegionVal = os.Getenv("AWS_REGION")
	}
	if os.Getenv("AWS_DEFAULT_REGION") != "" {
		winningRegionVal = os.Getenv("AWS_DEFAULT_REGION")
	}
	klog.V(2).Infof("using AWS region %q", winningRegionVal)
	return os.Setenv("AWS_REGION", winningRegionVal)
}

func setupAwsEnvVars(c *AWSConfig) error {
	if err := setEnvIfNotSet("AWS_ACCESS_KEY_ID", c.AccessKeyID); err != nil {
		return err
	}
	if err := setEnvIfNotSet("AWS_SECRET_ACCESS_KEY", c.SecretAccessKey); err != nil {
		return err
	}
	region := c.Region
	if region == "" {
		region = aws.AutoDetectRegion()
	}
	if region != "" {
		if err := setupAWSRegion(region); err != nil {
			return err
		}
	}
	klog.V(2).Infof("Validating connection to AWS")
	if err := aws.CheckConnection(c.EndpointURL, c.InsecureTLSSkipVerify); err != nil {
		return util.WrapError(err, "Error validating connection to AWS")
	}
	klog.V(2).Infof("Validated access to AWS")
	return nil
}

func setupAzureEnvVars(c *AzureConfig) error {
	if err := setEnvIfNotSet("AZURE_TENANT_ID", c.TenantID); err != nil {
		return err
	}
	if err := setEnvIfNotSet("AZURE_CLIENT_ID", c.ClientID); err != nil {
		return err
	}
	if err := setEnvIfNotSet("AZURE_CLIENT_SECRET", c.ClientSecret); err != nil {
		return err
	}
	if err := os.Setenv("SUBSCRIPTION_ID", c.SubscriptionID); err != nil {
		return err
	}
	klog.V(2).Infof("Validating connection to Azure")
	if err := azure.CheckConnection(c.SubscriptionID); err != nil {
		return util.WrapError(err, "Error validating connection to Azure")
	}
	klog.V(2).Infof("Validated access to Azure")
	return nil
}

func configureCloudProvider(cf *ServerConfigFile, controllerID, nametag string) (cloud.CloudClient, error) {
	// see which cloud is non-null, take first
	mockCloudAPI := os.Getenv("MOCK_CLOUD_API")
	if val, err := strconv.ParseBool(mockCloudAPI); err == nil && val {
		klog.Warningf("Running with a mocked cloud API client. This kip installation can not be changed to run a real k8s cluster")
		return cloud.NewMockClient(), nil
	}
	cc := cf.Cloud
	numClouds := 0
	if cc.AWS != nil {
		numClouds++
	}
	if cc.Azure != nil {
		numClouds++
	}
	if cc.GCE != nil {
		numClouds++
	}
	if numClouds > 1 {
		return nil, fmt.Errorf("Multiple clouds configured in cloud section of provider.yaml")
	}
	privateIPOnly := false
	if cf.Cells.PrivateIPOnly != nil && *cf.Cells.PrivateIPOnly {
		privateIPOnly = true
	}
	if cc.AWS != nil {
		errs := validateAWSConfig(cc.AWS)
		if len(errs) > 0 {
			err := fmt.Errorf("Invalid AWS Cloud Config: %v", errs.ToAggregate())
			return nil, err
		}

		err := setupAwsEnvVars(cc.AWS)
		if err != nil {
			return nil, util.WrapError(err, "Could not configure AWS cloud client authorization")
		}

		// Gross: if vpc is "default", the NewEC2Client will
		// attempt to figure out the VPCID and the actual ID
		// will be available from there

		client, err := aws.NewEC2Client(aws.EC2ClientConfig{
			ControllerID:          controllerID,
			Nametag:               nametag,
			VPCID:                 cc.AWS.VPCID,
			SubnetID:              cc.AWS.SubnetID,
			ECSClusterName:        cc.AWS.EcsClusterName,
			PrivateIPOnly:         privateIPOnly,
			EndpointURL:           cc.AWS.EndpointURL,
			InsecureTLSSkipVerify: cc.AWS.InsecureTLSSkipVerify,
		})

		if err != nil {
			return nil, util.WrapError(err, "Error creating AWS cloud client")
		}
		return client, nil
	} else if cc.Azure != nil {
		errs := validateAzureConfig(cc.Azure)
		if len(errs) > 0 {
			err := fmt.Errorf("Invalid Azure Cloud Config: %v", errs.ToAggregate())
			return nil, err
		}
		err := setupAzureEnvVars(cc.Azure)
		if err != nil {
			return nil, util.WrapError(err, "Could not configure Azure cloud client")
		}
		client, err := azure.NewAzureClient(
			controllerID,
			nametag,
			cc.Azure.SubscriptionID,
			cc.Azure.Location,
			cc.Azure.VNetName,
			cc.Azure.SubnetName,
			cc.Azure.ResourceGroup,
		)
		if err != nil {
			return nil, util.WrapError(err, "Error creating Azure cloud client")
		}
		return client, nil
	} else if cc.GCE != nil {
		errs := validateGCEConfig(cc.GCE)
		if len(errs) > 0 {
			err := fmt.Errorf("Invalid GCE Cloud Config: %v", errs.ToAggregate())
			return nil, err
		}
		options := make([]gce.ClientOption, 0, 4)
		options = append(options, gce.WithZone(cc.GCE.Zone))
		options = append(options, gce.WithVPCName(cc.GCE.VPCName))
		options = append(options, gce.WithSubnetName(cc.GCE.SubnetName))
		options = append(options, gce.WithPrivateIPOnly(privateIPOnly))
		if cc.GCE.CredentialsFile != "" {
			options = append(options, gce.WithCredentialsFile(cc.GCE.CredentialsFile))
		}
		if cc.GCE.Credentials != nil {
			options = append(options, gce.WithCredentials(cc.GCE.Credentials.ClientEmail, cc.GCE.Credentials.PrivateKey))
		}
		client, err := gce.NewGCEClient(controllerID, nametag, cc.GCE.ProjectID, options...)
		if err != nil {
			return nil, util.WrapError(err, "Error creating GCE cloud client")
		}
		return client, nil
	} else {
		return nil, fmt.Errorf("You must specify a cloud configuration in provider.yaml")
	}
}

func ParseConfig(path string) (*ServerConfigFile, error) {
	var err error

	// unmarshal into ServerConfigFile
	if _, err = os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("Missing config file %s", path)
	}
	configData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, util.WrapError(err, "Could not read provider.yaml")
	}
	configFile := serverConfigFileWithDefaults()
	decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(configData), bufferSize)
	err = decoder.Decode(&configFile)
	if err != nil {
		return nil, util.WrapError(err, "Error parsing provider.yaml")
	}

	setConfigDefaults(configFile)

	updateCapacityFromDeprecatedFields(configFile)

	return configFile, nil
}

func updateCapacityFromDeprecatedFields(config *ServerConfigFile) {
	if config.Kubelet.CPU != nil {
		config.Kubelet.Capacity[v1.ResourceCPU] = *config.Kubelet.CPU
	}
	if config.Kubelet.Memory != nil {
		config.Kubelet.Capacity[v1.ResourceMemory] = *config.Kubelet.Memory
	}
	if config.Kubelet.Pods != nil {
		config.Kubelet.Capacity[v1.ResourcePods] = *config.Kubelet.Pods
	}
}

// Sets default values for parameters that can only be set once the
// ServerConfigFile has been parsed
func setConfigDefaults(config *ServerConfigFile) {
	if config.Cells.HealthCheck.Status == nil && config.Cells.HealthCheck.CloudAPI == nil {
		config.Cells.HealthCheck = HealthCheckConfig{
			Status: &StatusHealthCheck{
				HealthyTimeout: defaultStatusHealthCheckTimeout,
			},
		}
	}
	if config.Cells.HealthCheck.Status != nil {
		if config.Cells.HealthCheck.Status.HealthyTimeout == 0 {
			config.Cells.HealthCheck.Status.HealthyTimeout = defaultStatusHealthCheckTimeout
		}
	} else if config.Cells.HealthCheck.CloudAPI != nil {
		if config.Cells.HealthCheck.CloudAPI.HealthyTimeout == 0 {
			config.Cells.HealthCheck.CloudAPI.HealthyTimeout = defaultCloudAPIHealthCheckTimeout
		}
		if config.Cells.HealthCheck.CloudAPI.Interval <= 0 {
			config.Cells.HealthCheck.CloudAPI.Interval = defaultCloudAPIHealthCheckInterval
		}
	}
}

func ConfigureCloud(configFile *ServerConfigFile, controllerID, nametag string) (cloud.CloudClient, error) {
	cloudClient, err := configureCloudProvider(configFile, controllerID, nametag)
	if err != nil {
		return nil, fmt.Errorf("Error setting up cloud client: %v", err)
	}

	connectWithPublicIPs := cloudClient.ConnectWithPublicIPs()
	if connectWithPublicIPs {
		klog.V(2).Infof("controller will connect to nodes via public IPs")
	} else {
		klog.V(2).Infof("controller will connect to nodes via private IPs")
	}
	err = cloudClient.EnsureMilpaSecurityGroups(
		configFile.Cells.ExtraCIDRs,
		configFile.Cells.ExtraSecurityGroups,
	)
	if err != nil {
		return nil, util.WrapError(err, "Error setting up cloud client security groups")
	}
	return cloudClient, err
}

const awsRegionFormat string = "[a-z]{2}-[a-z]+-[0-9]"

var awsRegionRegexp = regexp.MustCompile("^" + awsRegionFormat + "$")

func validateAWSConfig(cf *AWSConfig) field.ErrorList {
	allErrs := field.ErrorList{}

	fldPath := field.NewPath("cloud.aws")
	if cf.Region == blankTemplateValue {
		allErrs = append(allErrs, field.Required(fldPath.Child("region"), "aws region must be set or pulled from the environment"))
	} else if cf.Region != "" && !awsRegionRegexp.MatchString(cf.Region) {
		regexError := vutil.RegexError(awsRegionFormat, "us-east-1")
		allErrs = append(allErrs, field.Invalid(fldPath.Child("region"), cf.Region, regexError))
	}

	// validation of these parameters can be problematic credentials
	// can come from the config file or a ~HOME/.aws/ folder, or from
	// env vars or from an IAM role attached to an instance.

	if cf.AccessKeyID == blankTemplateValue {
		allErrs = append(allErrs, field.Required(fldPath.Child("accessKeyID"), "accessKeyID must be set or pulled from the environment"))
	}
	if cf.SecretAccessKey == blankTemplateValue {
		allErrs = append(allErrs, field.Required(fldPath.Child("secretAccessKey"), "secretAccessKey must be set or pulled from the environment"))
	}

	return allErrs
}

func validateAzureConfig(cf *AzureConfig) field.ErrorList {
	allErrs := field.ErrorList{}
	fldPath := field.NewPath("cloud.azure")

	// Items that must be set in provider.yaml
	if cf.SubscriptionID == blankTemplateValue || cf.SubscriptionID == "" {
		allErrs = append(allErrs, field.Required(fldPath.Child("subscriptionID"), "azure subscriptionID must be set in provider.yaml"))
	}
	if cf.Location == blankTemplateValue || cf.Location == "" {
		allErrs = append(allErrs, field.Required(fldPath.Child("location"), "azure Location must be set in provider.yaml"))
	}
	// Items that can be set in the Environment
	if cf.TenantID == blankTemplateValue {
		allErrs = append(allErrs, field.Required(fldPath.Child("tenantID"), "tenantID must be set in provider.yaml or pulled from the environment"))
	}
	if cf.ClientID == blankTemplateValue {
		allErrs = append(allErrs, field.Required(fldPath.Child("clientID"), "clientID must be set in provider.yaml or pulled from the environment"))
	}
	if cf.ClientSecret == blankTemplateValue {
		allErrs = append(allErrs, field.Required(fldPath.Child("clientSecret"), "clientSecret must be set in provider.yaml or pulled from the environment"))
	}
	if cf.ResourceGroup == blankTemplateValue || cf.ResourceGroup == "" {
		allErrs = append(allErrs, field.Required(fldPath.Child("resourceGroup"), "resourceGroup must be set in provider.yaml"))
	}

	return allErrs
}

func validateGCEConfig(cf *GCEConfig) field.ErrorList {
	allErrs := field.ErrorList{}
	fldPath := field.NewPath("cloud.gce")
	if cf.CredentialsFile != "" && cf.Credentials != nil {
		allErrs = append(allErrs, field.Invalid(fldPath, "Multiple credential sources specified", "please specify ONE of credentialsFile or credentials structure"))
	}
	if cf.CredentialsFile != "" {
		allErrs = append(allErrs, validation.ValidateFileExists(cf.CredentialsFile, fldPath.Child("credentialsFile"))...)
	}
	return allErrs
}

func validateServerConfigFile(cf *ServerConfigFile) field.ErrorList {
	allErrs := field.ErrorList{}

	cells := cf.Cells
	fldPath := field.NewPath("cells")

	if cells.CloudInitFile != "" {
		allErrs = append(allErrs, validation.ValidateFileExists(cells.CloudInitFile, fldPath.Child("cloudInitFile"))...)
	}

	allErrs = append(allErrs, validation.ValidateResourceParses(cells.DefaultVolumeSize, fldPath.Child("defaultVolumeSize"))...)

	if cells.DefaultInstanceType == "" {
		allErrs = append(allErrs, field.Required(fldPath.Child("defaultInstanceType"), ""))
	}

	if cells.StatusInterval < 0 {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("statusInterval"), cells.StatusInterval, "cells.statusInterval must be >= 1"))
	}

	if cells.HealthCheck.Status != nil && cells.HealthCheck.CloudAPI != nil {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("healthcheck"), "multiple healthchecks configured", "cannot set both status and cloudAPI healthchecks"))
	}
	if cells.HealthCheck.Status != nil && cells.HealthCheck.Status.HealthyTimeout <= 0 {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("healthcheck.status.healthyTimeout"), cells.HealthCheck.Status.HealthyTimeout, "healthyTimeout must be positive"))
	}
	if cells.HealthCheck.CloudAPI != nil {
		if cells.HealthCheck.CloudAPI.HealthyTimeout <= 0 {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("healthcheck.cloudAPI.healthyTimeout"), cells.HealthCheck.CloudAPI.HealthyTimeout, "healthyTimeout must be positive"))
		}
		if cells.HealthCheck.CloudAPI.Interval < 10 {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("healthcheck.cloudAPI.interval"), cells.HealthCheck.CloudAPI.Interval, "cloudAPI interval must be greater than or equal to 10 (seconds)"))
		}
	}

	// Sadly we can't validate the default instance type until
	// after we initialize the instanceselector and the instance
	// selector needs the cloud config in order to be initialized
	// allErrs = append(allErrs, validation.ValidateInstanceType(nodes.DefaultInstanceType, fldPath.Child("defaultInstanceType"))...)

	snPath := fldPath.Child("standbyCells")
	for i, n := range cells.StandbyCells {
		snPath = snPath.Index(i)
		allErrs = append(allErrs, validation.ValidateNonnegativeField(int64(n.Count), snPath.Child("count"))...)
		if n.InstanceType == "" {
			allErrs = append(allErrs, field.Required(snPath.Child("instanceType"), ""))
		}
	}

	if len(cells.Nametag) > 0 {
		for _, msg := range validation.NameIsDNS952Label(cells.Nametag, false) {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("nametag"), cells.Nametag, msg))
		}
	}

	return allErrs
}

func validateBootImageSpec(spec cloud.BootImageSpec, cloudClient cloud.CloudClient) error {
	img, err := cloudClient.GetImage(spec)
	if err != nil {
		return util.WrapError(err, "could not get machine image for %v", spec)
	}
	if img.ID == "" {
		return fmt.Errorf("could not find machine image for %v", spec)
	}
	return nil
}
