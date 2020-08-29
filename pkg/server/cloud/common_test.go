package cloud

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToSaneVolumeSize(t *testing.T) {
	testCases := []struct {
		Name            string
		volSpecSize         string
		imageVolumeDiskSize int32
		expectedSize        int32
	}{
		{
			"spec size bigger",
			"5Gi",
			2,
			5,
		},
		{
			"image size bigger",
			"5Gi",
			10,
			10,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got := ToSaneVolumeSize(testCase.volSpecSize, Image{
				ID:             "dummy-id",
				Name:           "dummy-name",
				RootDevice:     "dummy-root-device",
				CreationTime:   nil,
				VolumeDiskSize: &testCase.imageVolumeDiskSize,
			})
			assert.Equal(t, got, testCase.expectedSize)
		})
	}

}
