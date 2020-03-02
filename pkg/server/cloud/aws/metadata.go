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
	"bytes"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/fullsailor/pkcs7"
)

const (
	metadataTimeout   = 2 * time.Second
	metadataURL       = "http://169.254.169.254/latest/meta-data/"
	AWSCertificatePem = `-----BEGIN CERTIFICATE-----
MIIC7TCCAq0CCQCWukjZ5V4aZzAJBgcqhkjOOAQDMFwxCzAJBgNVBAYTAlVTMRkw
FwYDVQQIExBXYXNoaW5ndG9uIFN0YXRlMRAwDgYDVQQHEwdTZWF0dGxlMSAwHgYD
VQQKExdBbWF6b24gV2ViIFNlcnZpY2VzIExMQzAeFw0xMjAxMDUxMjU2MTJaFw0z
ODAxMDUxMjU2MTJaMFwxCzAJBgNVBAYTAlVTMRkwFwYDVQQIExBXYXNoaW5ndG9u
IFN0YXRlMRAwDgYDVQQHEwdTZWF0dGxlMSAwHgYDVQQKExdBbWF6b24gV2ViIFNl
cnZpY2VzIExMQzCCAbcwggEsBgcqhkjOOAQBMIIBHwKBgQCjkvcS2bb1VQ4yt/5e
ih5OO6kK/n1Lzllr7D8ZwtQP8fOEpp5E2ng+D6Ud1Z1gYipr58Kj3nssSNpI6bX3
VyIQzK7wLclnd/YozqNNmgIyZecN7EglK9ITHJLP+x8FtUpt3QbyYXJdmVMegN6P
hviYt5JH/nYl4hh3Pa1HJdskgQIVALVJ3ER11+Ko4tP6nwvHwh6+ERYRAoGBAI1j
k+tkqMVHuAFcvAGKocTgsjJem6/5qomzJuKDmbJNu9Qxw3rAotXau8Qe+MBcJl/U
hhy1KHVpCGl9fueQ2s6IL0CaO/buycU1CiYQk40KNHCcHfNiZbdlx1E9rpUp7bnF
lRa2v1ntMX3caRVDdbtPEWmdxSCYsYFDk4mZrOLBA4GEAAKBgEbmeve5f8LIE/Gf
MNmP9CM5eovQOGx5ho8WqD+aTebs+k2tn92BBPqeZqpWRa5P/+jrdKml1qx4llHW
MXrs3IgIb6+hUIB+S8dz8/mmO0bpr76RoZVCXYab2CZedFut7qc3WUH9+EUAH5mw
vSeDCOUMYQR7R9LINYwouHIziqQYMAkGByqGSM44BAMDLwAwLAIUWXBlk40xTwSw
7HX32MxXYruse9ACFBNGmdX2ZBrVNGrN9N2f6ROk0k9K
-----END CERTIFICATE-----`
)

// This function grabs the ec2 metadata for the local machine that
// milpa is running on.  However, if milpa is not running within AWS
// the standard AWS metadata query hangs for about 15s. I tried
// modifiying the AWS HTTP client timeout but that didn't work so
// we'll just use our own client.
func GetMetadata(p string) (string, error) {
	if len(p) > 0 && p[0] == '/' {
		p = p[1:]
	}
	url := metadataURL + p
	timeout := time.Duration(metadataTimeout)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil || resp.StatusCode/200 != 1 {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

// The AWS Go SDK is missing the MarketplaceProductCodes field
type UpdatedEC2InstanceIdentityDocument struct {
	ec2metadata.EC2InstanceIdentityDocument
	MarketplaceProductCodes []string `json:"marketplaceProductCodes"`
}

func parseCertificate(certPem string) (*x509.Certificate, error) {
	dec, _ := pem.Decode([]byte(certPem))
	if dec == nil {
		return nil, fmt.Errorf("failed to find any PEM-encoded data block")
	}
	return x509.ParseCertificate(dec.Bytes)
}

func getInstanceIdentityDocument(metadataClient *ec2metadata.EC2Metadata, certPem string) (UpdatedEC2InstanceIdentityDocument, error) {
	doc := UpdatedEC2InstanceIdentityDocument{}
	cert, err := parseCertificate(certPem)
	if err != nil {
		return doc,
			awserr.New("CertificateParsingError",
				"failed to parse the PEM-encoded certificate", err)
	}

	resp, err := metadataClient.GetDynamicData("instance-identity/pkcs7")
	if err != nil {
		return doc,
			awserr.New("EC2MetadataRequestError",
				"failed to get EC2 instance PKCS7 signature", err)
	}

	dec, err := base64.StdEncoding.DecodeString(resp)
	if err != nil {
		return doc,
			awserr.New("Base64DecodeError",
				"failed to base64-decode PKCS7 resopnse", err)
	}

	parsed, err := pkcs7.Parse(dec)
	if err != nil {
		return doc,
			awserr.New("EC2MetadataRequestError",
				"failed to parse PKCS7 response", err)
	}

	parsed.Certificates = []*x509.Certificate{cert}
	if err := parsed.Verify(); err != nil {
		return doc,
			awserr.New("EC2MetadataVerificationError",
				"failed to verify PKCS7 signature with certificate", err)
	}

	if err := json.NewDecoder(bytes.NewReader(parsed.Content)).Decode(&doc); err != nil {
		return doc,
			awserr.New("SerializationError",
				"failed to decode EC2 instance identity document", err)
	}
	return doc, nil
}

func (c *AwsEC2) validateInstanceDocumentInfo(instDoc UpdatedEC2InstanceIdentityDocument) error {
	reply, err := c.client.DescribeInstances(&ec2.DescribeInstancesInput{
		InstanceIds: []*string{&instDoc.InstanceID},
	})
	if err != nil {
		return util.WrapError(err, "Could not get instance info from EC2 API")
	}
	if len(reply.Reservations) == 0 || len(reply.Reservations[0].Instances) == 0 {
		return fmt.Errorf("EC2 API says the reported Marketplace Instance %s is not running", instDoc.InstanceID)
	}
	instance := reply.Reservations[0].Instances[0]
	if instance.LaunchTime == nil {
		return fmt.Errorf("Could not get AWS Marketplace instance launch time from EC2 API")
	}
	launchTime := *instance.LaunchTime
	if launchTime.Sub(instDoc.PendingTime) > 2*time.Second {
		return fmt.Errorf("The reported launch time of the AWS Marketplace instance %s is different from the time returned by the metadata service %s", instDoc.PendingTime.String(), launchTime.String())
	}
	return nil
}
