package aws

import (
	"fmt"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/util"
	"k8s.io/klog"
)

const (
	ssmBasePath           = "/kip/cells"
	ssmMaxResults         = 10
	ssmParameterChunkSize = 4096
)

// We use SSM to store parameters such as the internal Kip CA certificate,
// instance private key and certificate, etc. The parameter key is
// /kip/cells/<instance-ID>/<parameter-name>, so instances can retrieve all
// their configuration parameters in one step using their instance ID.

func createPath(names ...string) string {
	e := append([]string{ssmBasePath}, names...)
	return filepath.Join(e...)
}

func (e *AwsEC2) getInstanceRoleID(instanceID string) (string, error) {
	associations, err := e.client.DescribeIamInstanceProfileAssociations(
		&ec2.DescribeIamInstanceProfileAssociationsInput{
			Filters: []*ec2.Filter{
				{
					Name:   aws.String("instance-id"),
					Values: aws.StringSlice([]string{instanceID}),
				},
			},
		})
	if err != nil {
		return "", util.WrapError(err, "DescribeIamInstanceProfileAssociations() "+instanceID)
	}
	if len(associations.IamInstanceProfileAssociations) <= 0 {
		return "", util.WrapError(err, "DescribeIamInstanceProfileAssociations() not found "+instanceID)
	}
	if associations.IamInstanceProfileAssociations[0].IamInstanceProfile == nil {
		return "", util.WrapError(err, "DescribeIamInstanceProfileAssociations() no instance profile "+instanceID)
	}
	profileARN := aws.StringValue(associations.IamInstanceProfileAssociations[0].IamInstanceProfile.Arn)
	parts := strings.SplitN(profileARN, "/", 2)
	if len(parts) < 2 {
		return "", fmt.Errorf("%s invalid instance profile ARN %q", instanceID, profileARN)
	}
	profileName := parts[1]

	out, err := e.iam.GetInstanceProfile(&iam.GetInstanceProfileInput{
		InstanceProfileName: aws.String(profileName),
	})
	if err != nil {
		return "", util.WrapError(err, "getting instance profile for "+instanceID)
	}
	if out.InstanceProfile == nil || len(out.InstanceProfile.Roles) < 1 {
		return "", fmt.Errorf("got invalid instance profile for %s", instanceID)
	}
	if len(out.InstanceProfile.Roles) > 1 {
		klog.Warningf("instance %s has multiple roles in its profile", instanceID)
	}

	roleID := aws.StringValue(out.InstanceProfile.Roles[0].RoleId)

	return roleID, nil
}

func (e *AwsEC2) getParameterTags(instanceID string) []*ssm.Tag {
	// Tag the parameter with the AWS "user ID" (which is the IAM role ID in
	// its instance profile and the instance ID separated by a colon), so an
	// IAM policy with a condition "aws:resourceTag/AWSUserID" ==
	// "${aws:userid}" can be used to restrict access to parameters on a per
	// instance basis.
	userID := "unknown"
	roleID, err := e.getInstanceRoleID(instanceID)
	if err != nil {
		klog.Warningf("failed to get instance role ID for %s: %v", instanceID, err)
	} else {
		userID = fmt.Sprintf("%s:%s", roleID, instanceID)
	}

	return []*ssm.Tag{
		&ssm.Tag{
			Key:   aws.String(cloud.AWSUserIDTagKey),
			Value: aws.String(userID),
		},
		&ssm.Tag{
			Key:   aws.String(cloud.ControllerTagKey),
			Value: aws.String(e.controllerID),
		},
		&ssm.Tag{
			Key:   aws.String(cloud.NametagTagKey),
			Value: aws.String(e.nametag),
		},
	}
}

func processInChunks(name, value string, maxSize int, fn func(chunkName, chunkValue string) error) error {
	i := 0
	for len(value) > maxSize {
		n := maxSize
		for n >= maxSize-utf8.UTFMax && !utf8.RuneStart(value[n]) {
			n--
		}
		chunkName := fmt.Sprintf("%s_%d", name, i)
		chunkValue := value[:n]
		err := fn(chunkName, chunkValue)
		if err != nil {
			return util.WrapError(err, "adding parameter "+chunkName)
		}
		i++
		value = value[n:]
	}
	if len(value) > 0 {
		chunkName := fmt.Sprintf("%s_%d", name, i)
		chunkValue := value
		err := fn(chunkName, chunkValue)
		if err != nil {
			return util.WrapError(err, "adding parameter "+chunkName)
		}
	}
	return nil
}

func (e *AwsEC2) AddInstanceParameter(instanceID, name, value string, isSecret bool) error {
	// Try to add the parameter as is; if it is larger than the maximum size
	// allowed on AWS, split it into chunks. The instance side will have to
	// re-assemble it after retrieving all its parameters.
	if len(value) <= ssmParameterChunkSize {
		return e.addInstanceParameter(instanceID, name, value, isSecret)
	}

	klog.V(2).Infof("%s trying to add parameter %s > %d bytes, splitting it into chunks",
		instanceID, name, ssmParameterChunkSize)
	err := processInChunks(name, value, ssmParameterChunkSize, func(chunkName, chunkValue string) error {
		return e.addInstanceParameter(instanceID, chunkName, chunkValue, isSecret)
	})
	if err != nil {
		return util.WrapError(err, "processing chunked parameters")
	}

	return nil
}

func (e *AwsEC2) addInstanceParameter(instanceID, name, value string, isSecret bool) error {
	pType := aws.String(ssm.ParameterTypeString)
	if isSecret {
		pType = aws.String(ssm.ParameterTypeSecureString)
	}
	in := &ssm.PutParameterInput{
		Name:        aws.String(createPath(instanceID, name)),
		Description: aws.String(fmt.Sprintf("Kip cell %s parameter %s", instanceID, name)),
		Overwrite:   aws.Bool(false),
		Value:       aws.String(value),
		Tier:        aws.String(ssm.ParameterTierIntelligentTiering),
		Type:        pType,
		Tags:        e.getParameterTags(instanceID),
	}
	out, err := e.ssm.PutParameter(in)
	if err != nil {
		return err
	}
	klog.V(5).Infof("created new parameter %s for %s, tier %s version %d",
		instanceID, name, aws.StringValue(out.Tier), aws.Int64Value(out.Version))
	return nil
}

func (e *AwsEC2) getAllParameterNames(base string) ([]string, error) {
	in := &ssm.GetParametersByPathInput{
		MaxResults:     aws.Int64(ssmMaxResults),
		Path:           aws.String(base),
		Recursive:      aws.Bool(true),
		WithDecryption: aws.Bool(false),
	}
	names := make([]string, 0)
	err := e.ssm.GetParametersByPathPages(in,
		func(page *ssm.GetParametersByPathOutput, lastPage bool) bool {
			for _, param := range page.Parameters {
				if param == nil {
					klog.Warningf("got nil SSM parameter for base %s", base)
					continue
				}
				names = append(names, aws.StringValue(param.Name))
			}
			return true
		})
	if err != nil {
		return nil, util.WrapError(err, fmt.Sprintf("ssm.GetParametersByPathPages() %s", base))
	}
	return names, nil
}

func (e *AwsEC2) DeleteInstanceParameter(instanceID, name string) error {
	names := []string{createPath(instanceID, name)}
	if name == "" {
		var err error
		parameterNames, err := e.getAllParameterNames(createPath(instanceID))
		if err != nil {
			return util.WrapError(err, fmt.Sprintf("getAllParameternames() %s %s", instanceID, name))
		}
		names = parameterNames
		klog.V(2).Infof("removing all instance parameters for %s: %v", instanceID, names)
	}

	if len(names) <= 0 {
		klog.Warningf("no instance parameters found for %s", instanceID)
		return nil
	}

	out, err := e.ssm.DeleteParameters(&ssm.DeleteParametersInput{
		Names: aws.StringSlice(names),
	})
	if err != nil {
		return util.WrapError(err, fmt.Sprintf("DeleteParameters() %s %s", instanceID, name))
	}

	msg := fmt.Sprintf("deleted parameters %v invalid parameters %v",
		aws.StringValueSlice(out.DeletedParameters), aws.StringValueSlice(out.InvalidParameters))
	if len(out.InvalidParameters) > 0 {
		klog.Warning(msg)
	} else {
		klog.V(5).Info(msg)
	}

	return nil
}
