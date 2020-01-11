package aws

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/elotl/cloud-instance-provider/pkg/util"
)

var (
	ecsAuthUntil    = time.Now().UTC()
	ecsAuthUsername = ""
	ecsAuthPassword = ""
)

func (e *AwsEC2) GetRegistryAuth() (string, string, error) {
	// We don't want the token to expire while we are deploying a pod
	// so pad our
	aFewMinutesFromNow := time.Now().UTC().Add(15 * time.Minute)
	if ecsAuthUntil.Before(aFewMinutesFromNow) ||
		ecsAuthUsername == "" ||
		ecsAuthPassword == "" {

		svc := ecr.New(session.New())
		input := &ecr.GetAuthorizationTokenInput{}
		result, err := svc.GetAuthorizationToken(input)

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
			return "", "", fmt.Errorf(
				"Could not decode authorization token from AWS")
		}
		parts := strings.Split(decoded, ":")
		if len(parts) != 2 {
			return "", "", fmt.Errorf(
				"Invalid format of ECS authorization from AWS")
		}
		ecsAuthUntil = validUntil
		ecsAuthUsername = parts[0]
		ecsAuthPassword = parts[1]
	}
	return ecsAuthUsername, ecsAuthPassword, nil
}
