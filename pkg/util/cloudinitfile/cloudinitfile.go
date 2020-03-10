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

// Useful for 1 allowing the user to update the cloud-init file
// without restarting the serer.  Also validates a users cloud-init data and has helpers for managing milpa data
package cloudinitfile

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/elotl/cloud-instance-provider/pkg/util/filewatcher"
	"github.com/elotl/cloud-instance-provider/pkg/util/yaml"
)

const semverRegexFmt string = `v?([0-9]+)(\.[0-9]+)(\.[0-9]+)?` +
	`(-([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?` +
	`(\+([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?`

var (
	ItzoVersionPath = "/tmp/milpa/itzo_version"
	ItzoURLPath     = "/tmp/milpa/itzo_url"
	semverRegex     = regexp.MustCompile("^" + semverRegexFmt + "$")
)

type MilpaFile struct {
	content     string
	path        string
	permissions string
}

type File struct {
	userData        filewatcher.Watcher
	lastSeenVersion int
	milpaFiles      map[string]MilpaFile
}

func New(path string) *File {
	fw := filewatcher.New(path)
	fw.CheckPeriod = 20 * time.Second
	f := &File{
		userData:        fw,
		lastSeenVersion: fw.Version(),
		milpaFiles:      make(map[string]MilpaFile),
	}
	return f
}

func (f *File) Validate(c string) error {
	type Empty struct{}
	var empty Empty
	yml := []byte(c)
	decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(yml), 16000)
	return decoder.Decode(&empty)
}

func (f *File) ResetInstanceData() {
	f.milpaFiles = make(map[string]MilpaFile)
}

func (f *File) AddMilpaFile(content, path, permissions string) {
	f.milpaFiles[path] = MilpaFile{
		content:     content,
		path:        path,
		permissions: permissions,
	}
}

// Adds an itzo version number to cloud-init file.  If the user
// didn't specify "latest" but they left off the leading 'v'
// then add it on (itzo files are named like: itzo-v1.2.3)
func (f *File) AddItzoVersion(version string) {
	if version == "" {
		return
	} else if version != "latest" &&
		version[0] != 'v' &&
		semverRegex.MatchString(version) {
		version = "v" + version
	}
	f.AddMilpaFile(version, ItzoVersionPath, "0444")
}

func (f *File) AddItzoURL(url string) {
	if url == "" {
		return
	}
	f.AddMilpaFile(url, ItzoURLPath, "0444")
}

func (f *MilpaFile) Content() []byte {
	indented := strings.Replace(f.content, "\n", "\n      ", -1)
	str := fmt.Sprintf(
		"  - content: |\n      %s\n    path: %s\n    permissions: %s\n",
		indented, f.path, f.permissions)
	return []byte(str)
}

func (f *File) MilpaContents() string {
	if len(f.milpaFiles) == 0 {
		return ""
	}
	data := make([]byte, 0, 2048)
	data = append(data, []byte("\nmilpa_files:\n")...)
	for _, mf := range f.milpaFiles {
		data = append(data, mf.Content()...)
	}
	return string(data)
}

func (f *File) Contents() (string, error) {
	userContent := f.userData.Contents()
	if f.lastSeenVersion != f.userData.Version() {
		if err := f.Validate(userContent); err != nil {
			return "", util.WrapError(
				err, "Error validating user cloud-init script")
		}
		f.lastSeenVersion = f.userData.Version()
	}

	milpaContent := f.MilpaContents()
	cloudInitContent := milpaContent + "\n" + userContent
	if len(cloudInitContent) > 16000 {
		return "", fmt.Errorf("Cloud init data length is over 16K")
	}
	return cloudInitContent, nil
}
