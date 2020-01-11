package util

import (
	"strings"

	"github.com/elotl/cloud-instance-provider/pkg/api"
)

var (
	// VERSION is set during the build from the file at milpa/version
	VERSION string = ""
	// GIT_REVISION and GIT_DIRTY are set during build.
	GIT_REVISION string = ""
	GIT_DIRTY    string = ""
)

func Version() string {
	version := VERSION
	if GIT_REVISION != "" {
		version = version + "-" + GIT_REVISION
	}
	if GIT_DIRTY != "" {
		version = version + "-" + GIT_DIRTY
	}
	return version
}

func GetVersionInfo() api.VersionInfo {
	parts := strings.SplitN(VERSION, ".", 3)
	dirtyString := GIT_DIRTY
	if dirtyString == "" {
		dirtyString = "clean"
	}
	v := api.VersionInfo{
		GitVersion:   VERSION,
		GitCommit:    GIT_REVISION,
		GitTreeState: dirtyString,
	}
	if len(parts) >= 2 {
		v.Major = parts[0]
		v.Minor = parts[1]
	}
	return v
}
