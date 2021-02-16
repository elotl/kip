package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/elotl/kip/pkg/server/cloud"
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

func TestGetRootDeviceVolumeSpecs(t *testing.T) {
	notRootDeviceName := "not-root"
	rootDeviceName := "root-device"
	var volumeSize int64 = 100
	var expectedVolumeSize int32 = 100
	testCases := []struct {
		caseName             string
		blockDevices         []*ec2.BlockDeviceMapping
		rootDeviceName       string
		expectedRootDiskSize int32
		expectedVolumeType   string
		expectedIops         *int64
		expectedThroughput   *int64
	}{
		{
			caseName: "root-device-found",
			blockDevices: []*ec2.BlockDeviceMapping{
				{
					DeviceName: &notRootDeviceName,
				},
				{
					DeviceName: &rootDeviceName,
					Ebs: &ec2.EbsBlockDevice{
						VolumeSize: &volumeSize,
						VolumeType: aws.String("gp3"),
						Iops:       &volumeSize,
						Throughput: &volumeSize,
					},
				},
			},
			rootDeviceName:       rootDeviceName,
			expectedRootDiskSize: expectedVolumeSize,
			expectedVolumeType:   "gp3",
			expectedIops:         aws.Int64(volumeSize),
			expectedThroughput:   aws.Int64(volumeSize),

		},
		{
			caseName: "gp3-no-throughput-set",
			blockDevices: []*ec2.BlockDeviceMapping{
				{
					DeviceName: &notRootDeviceName,
				},
				{
					DeviceName: &rootDeviceName,
					Ebs: &ec2.EbsBlockDevice{
						VolumeSize: &volumeSize,
						VolumeType: aws.String("gp3"),
						Iops:       &volumeSize,
					},
				},
			},
			rootDeviceName:       rootDeviceName,
			expectedRootDiskSize: expectedVolumeSize,
			expectedVolumeType:   "gp3",
			expectedIops:         aws.Int64(volumeSize),
			expectedThroughput:   nil,

		},
		{
			caseName: "gp3-no-iops-set",
			blockDevices: []*ec2.BlockDeviceMapping{
				{
					DeviceName: &notRootDeviceName,
				},
				{
					DeviceName: &rootDeviceName,
					Ebs: &ec2.EbsBlockDevice{
						VolumeSize: &volumeSize,
						VolumeType: aws.String("gp3"),
						Throughput: &volumeSize,
					},
				},
			},
			rootDeviceName:       rootDeviceName,
			expectedRootDiskSize: expectedVolumeSize,
			expectedVolumeType:   "gp3",
			expectedIops:         nil,
			expectedThroughput:   aws.Int64(volumeSize),

		},
		{
			caseName:       "empty-volume-list",
			blockDevices:   []*ec2.BlockDeviceMapping{},
			rootDeviceName: rootDeviceName,
			expectedVolumeType: "gp2",
		},
		{
			caseName: "root-device-not-found",
			blockDevices: []*ec2.BlockDeviceMapping{
				{
					DeviceName: &notRootDeviceName,
					Ebs: &ec2.EbsBlockDevice{
						VolumeSize: &volumeSize,
					},
				},
			},
			rootDeviceName: rootDeviceName,
			expectedVolumeType: "gp2",
		},
		{
			caseName: "root-device-io1",
			blockDevices: []*ec2.BlockDeviceMapping{
				{
					DeviceName: &rootDeviceName,
					Ebs: &ec2.EbsBlockDevice{
						VolumeSize: &volumeSize,
						VolumeType: aws.String("io1"),
						Iops: aws.Int64(25000),
					},
				},
			},
			rootDeviceName: rootDeviceName,
			expectedRootDiskSize: expectedVolumeSize,
			expectedVolumeType: "io1",
			expectedIops: aws.Int64(25000),
		},
		{
			caseName: "root-device-io2",
			blockDevices: []*ec2.BlockDeviceMapping{
				{
					DeviceName: &rootDeviceName,
					Ebs: &ec2.EbsBlockDevice{
						VolumeSize: &volumeSize,
						VolumeType: aws.String("io2"),
					},
				},
			},
			rootDeviceName: rootDeviceName,
			expectedRootDiskSize: expectedVolumeSize,
			expectedVolumeType: "io2",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			rootDisk := getRootDeviceVolumeSpecs(testCase.blockDevices, testCase.rootDeviceName)
			assert.Equal(t, testCase.expectedRootDiskSize, rootDisk.VolumeSize)
			assert.Equal(t, testCase.expectedVolumeType, rootDisk.VolumeType)
			assert.Equal(t, aws.Int64Value(testCase.expectedThroughput), aws.Int64Value(rootDisk.Throughput))
			assert.Equal(t, aws.Int64Value(testCase.expectedIops), aws.Int64Value(rootDisk.Iops))
		})
	}
}
