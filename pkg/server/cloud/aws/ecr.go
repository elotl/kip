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
	"encoding/base64"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/elotl/kip/pkg/util"
)

type ecsAuth struct {
	username   string
	password   string
	validUntil time.Time
}

var (
	ecsRegionAuth sync.Map
)

func regionFromImage(image string) (string, error) {
	server, _, err := util.ParseImageSpec(image)
	if err != nil {
		return "", err
	}
	// server should look like: ACCOUNT.dkr.ecr.REGION.amazonaws.com
	parts := strings.Split(server, ".")
	if len(parts) != 6 {
		return "", fmt.Errorf("Unknown ECR server format: %s", server)
	}
	return parts[3], nil
}

func (e *AwsEC2) GetRegistryAuth(image string) (string, string, error) {
	region, err := regionFromImage(image)
	if err != nil {
		return "", "", util.WrapError(err, "could not parse region from ECR repository %s", image)
	}

	var regionAuth ecsAuth
	regionAuthIface, ok := ecsRegionAuth.Load(region)
	if !ok {
		// This should never happen. Just in case, don't panic so we
		// don't bring the entire system down repeatedly.
		return "", "", fmt.Errorf("wrong type stored for ECS registry credentials")
	}
	regionAuth = regionAuthIface.(ecsAuth)
	// We don't want auth to expire when deploying a pod so we pad
	// our expire time by a bit
	aFewMinutesFromNow := time.Now().UTC().Add(15 * time.Minute)
	if regionAuth.username == "" ||
		regionAuth.password == "" ||
		regionAuth.validUntil.Before(aFewMinutesFromNow) {
		cfg := aws.NewConfig()
		ses, err := session.NewSession(cfg.WithRegion(region))
		if err != nil {
			return "", "", err
		}
		svc := ecr.New(ses)
		result, err := svc.GetAuthorizationToken(&ecr.GetAuthorizationTokenInput{})
		if err != nil {
			return "", "", util.WrapError(err, "Error requesting ECS credentials")
		}
		if result.AuthorizationData == nil && len(result.AuthorizationData) == 0 {
			return "", "", util.WrapError(
				err, "Error in ECS credentials response: No credentails returned")
		}
		creds := *(result.AuthorizationData)[0]
		// AWS credentials are valid for 12 hours
		validUntil := time.Now().UTC().Add(12 * time.Hour)
		if creds.ExpiresAt != nil {
			validUntil = creds.ExpiresAt.UTC()
		}
		// Auth is base64 encoded HTTP Auth format (username:password)
		decodedByte, err := base64.StdEncoding.DecodeString(*creds.AuthorizationToken)
		decoded := string(decodedByte)
		if err != nil {
			return "", "", fmt.Errorf("Could not decode authorization token from AWS")
		}
		parts := strings.Split(decoded, ":")
		if len(parts) != 2 {
			return "", "", fmt.Errorf("Invalid format of ECS authorization from AWS")
		}
		regionAuth = ecsAuth{
			validUntil: validUntil,
			username:   parts[0],
			password:   parts[1],
		}
		ecsRegionAuth.Store(region, regionAuth)
	}
	return regionAuth.username, regionAuth.password, nil
}
