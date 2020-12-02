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
// without restarting the serer.  Also validates a users cloud-init
// data and has helpers for managing milpa data
package cloudinitfile

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/elotl/kip/pkg/util"
	"github.com/go-yaml/yaml"
)

var (
	itzoDir          = "/tmp/itzo"
	cloudInitHeader  = []byte("#cloud-config\n")
	maxCloudInitSize = 16000
)

type File struct {
	userData CloudConfig
	kipFiles map[string]CloudInitFile
}

func New(path string) (*File, error) {
	var userData CloudConfig
	if path != "" {
		var err error
		userData, err = loadUserCloudConfig(path)
		if err != nil {
			return nil, util.WrapError(err, "Could not load user's cloud config file at %s", path)
		}
	}
	f := &File{
		userData: userData,
		kipFiles: make(map[string]CloudInitFile),
	}
	return f, nil
}

func (f *File) ResetInstanceData() {
	f.kipFiles = make(map[string]CloudInitFile)
}

func (f *File) AddKipFile(content, path, permissions string) {
	if !strings.Contains(path, string(filepath.Separator)) {
		path = filepath.Join(itzoDir, path)
	}
	f.kipFiles[path] = CloudInitFile{
		Content:            content,
		Path:               path,
		Owner:              "root",
		RawFilePermissions: permissions,
	}
}

func loadUserCloudConfig(path string) (ucc CloudConfig, err error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return ucc, err
	}
	err = yaml.Unmarshal([]byte(contents), &ucc)
	return ucc, err
}

func (f *File) Contents() ([]byte, error) {
	mergedConfig := f.userData
	mergedFiles := make([]CloudInitFile, 0, len(f.userData.WriteFiles)+len(f.kipFiles))
	mergedFiles = append(mergedFiles, f.userData.WriteFiles...)
	for _, wf := range f.kipFiles {
		mergedFiles = append(mergedFiles, wf)
	}
	mergedConfig.WriteFiles = mergedFiles
	mergedContent, err := yaml.Marshal(mergedConfig)
	if err != nil {
		return nil, err
	}
	cloudInitContent := cloudInitHeader
	cloudInitContent = append(cloudInitContent, mergedContent...)
	if len(cloudInitContent) > maxCloudInitSize {
		return nil, fmt.Errorf("Cloud init data length is over 16K")
	}
	return cloudInitContent, nil
}
