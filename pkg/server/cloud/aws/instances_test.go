package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/stretchr/testify/assert"
)

// func bootImageSpecToDescribeImagesInput(spec cloud.BootImageSpec) *ec2.DescribeImagesInput
func TestBootImageSpecToDescribeImagesInput(t *testing.T) {
	testCases := []struct {
		Spec  cloud.BootImageSpec
		Input ec2.DescribeImagesInput
	}{
		{
			Spec: cloud.BootImageSpec{},
			Input: ec2.DescribeImagesInput{
				Owners: aws.StringSlice([]string{elotlOwnerID}),
				Filters: []*ec2.Filter{
					{
						Name:   aws.String("name"),
						Values: aws.StringSlice([]string{elotlImageNameFilter}),
					},
				},
			},
		},
		{
			Spec: cloud.BootImageSpec{
				"filters": "name=elotl-kip-*",
			},
			Input: ec2.DescribeImagesInput{
				Filters: []*ec2.Filter{
					{
						Name:   aws.String("name"),
						Values: aws.StringSlice([]string{"elotl-kip-*"}),
					},
				},
			},
		},
		{
			Spec: cloud.BootImageSpec{
				"imageIDs": "ami-123 ami-456 ami-789",
				"owners":   "12345 9999999",
				"filters":  "name=elotl-kip-* tag-key=elotl-image-tag",
			},
			Input: ec2.DescribeImagesInput{
				ImageIds: aws.StringSlice([]string{"ami-123", "ami-456", "ami-789"}),
				Owners:   aws.StringSlice([]string{"12345", "9999999"}),
				Filters: []*ec2.Filter{
					{
						Name:   aws.String("name"),
						Values: aws.StringSlice([]string{"elotl-kip-*"}),
					},
					{
						Name:   aws.String("tag-key"),
						Values: aws.StringSlice([]string{"elotl-image-tag"}),
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		input := bootImageSpecToDescribeImagesInput(tc.Spec)
		assert.Equal(t, tc.Input, *input)
	}
}
