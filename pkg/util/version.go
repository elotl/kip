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

package util

import (
	"strings"

	"github.com/elotl/cloud-instance-provider/pkg/api"
)

var (
	// VERSION is set during the build from the file at kip/version
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
