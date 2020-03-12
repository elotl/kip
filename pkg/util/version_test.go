package util

import (
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestGetVersionInfo(t *testing.T) {
	testCases := []struct {
		Version     string
		VersionInfo api.VersionInfo
	}{
		{
			Version: "v0.0.1-205-g58ce23b-dirty",
			VersionInfo: api.VersionInfo{
				Major:        "0",
				Minor:        "0",
				GitVersion:   "v0.0.1-205",
				GitCommit:    "g58ce23b",
				GitTreeState: "dirty",
			},
		},
		{
			Version: "v10.33.1-205-g58ce23b-dirty",
			VersionInfo: api.VersionInfo{
				Major:        "10",
				Minor:        "33",
				GitVersion:   "v10.33.1-205",
				GitCommit:    "g58ce23b",
				GitTreeState: "dirty",
			},
		},
		{
			Version: "v11.22.33-205-g58ce23b",
			VersionInfo: api.VersionInfo{
				Major:        "11",
				Minor:        "22",
				GitVersion:   "v11.22.33-205",
				GitCommit:    "g58ce23b",
				GitTreeState: "clean",
			},
		},
		{
			Version: "v5.11.8",
			VersionInfo: api.VersionInfo{
				Major:        "5",
				Minor:        "11",
				GitVersion:   "v5.11.8",
				GitCommit:    "N/A",
				GitTreeState: "clean",
			},
		},
	}
	for _, tc := range testCases {
		vi := getVersionInfo(tc.Version)
		assert.Equal(t, tc.VersionInfo, vi)
	}
}
