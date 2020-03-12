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
	"fmt"
	"strconv"
	"strings"

	"github.com/elotl/cloud-instance-provider/pkg/api"
)

var (
	// VERSION is set during the build from "git describe --dirty".
	VERSION string = ""
)

func Version() string {
	version := VERSION
	if version == "" {
		// Something's broken, maybe manual "go build".
		version = "N/A"
	}
	return version
}

func getVersionInfo(version string) api.VersionInfo {
	// The output from "git describe --dirty" is something like
	// "v0.0.1-205-g58ce23b-dirty", where the last "-dirty" part is optional.
	semver := "N/A"
	nCommits := int64(0)
	gitCommit := "N/A"
	dirtyString := "N/A"
	parts := strings.Split(version, "-")
	if len(parts) == 1 {
		semver = parts[0]
		dirtyString = "clean"
	}
	if len(parts) >= 3 {
		semver = parts[0]
		nCommits, _ = strconv.ParseInt(parts[1], 10, 32)
		if nCommits > 0 {
			semver = fmt.Sprintf("%s-%d", semver, nCommits)
		}
		gitCommit = parts[2]
		dirtyString = "clean"
	}
	if len(parts) == 4 {
		dirtyString = parts[3]
	}
	v := api.VersionInfo{
		GitVersion:   semver,
		GitCommit:    gitCommit,
		GitTreeState: dirtyString,
	}
	parts = strings.SplitN(semver, ".", 3)
	if len(parts) >= 2 {
		major := parts[0]
		if len(major) > 1 && major[0] == 'v' {
			major = major[1:]
		}
		minor := parts[1]
		v.Major = major
		v.Minor = minor
	}
	return v
}

func GetVersionInfo() api.VersionInfo {
	return getVersionInfo(VERSION)
}
