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

	"github.com/coreos/yaml"
	cc "github.com/elotl/cloud-init/config"
	"github.com/elotl/cloud-instance-provider/pkg/util"
)

var (
	ItzoVersionPath  = "/tmp/milpa/itzo_version"
	ItzoURLPath      = "/tmp/milpa/itzo_url"
	cloudInitHeader  = []byte("#cloud-config\n")
	maxCloudInitSize = 16000
)

type File struct {
	userData   cc.CloudConfig
	milpaFiles map[string]cc.File
}

func New(path string) (*File, error) {
	var userData cc.CloudConfig
	if path != "" {
		var err error
		userData, err = loadUserCloudConfig(path)
		if err != nil {
			return nil, util.WrapError(err, "Could not load user's cloud config file at %s", path)
		}
	}
	f := &File{
		userData:   userData,
		milpaFiles: make(map[string]cc.File),
	}
	return f, nil
}

func (f *File) ResetInstanceData() {
	f.milpaFiles = make(map[string]cc.File)
}

func (f *File) AddMilpaFile(content, path, permissions string) {
	f.milpaFiles[path] = cc.File{
		Content:            content,
		Path:               path,
		Owner:              "root",
		RawFilePermissions: permissions,
	}
}

func loadUserCloudConfig(path string) (ucc cc.CloudConfig, err error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return ucc, err
	}
	err = yaml.Unmarshal([]byte(contents), &ucc)
	return ucc, err
}

// Adds an itzo version number to cloud-init file.  If the user
// didn't specify "latest" but they left off the leading 'v'
// then add it on (itzo files are named like: itzo-v1.2.3)
func (f *File) AddItzoVersion(version string) {
	if version == "" {
		return
	} else if version != "latest" && version[0] != 'v' {
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

func (f *File) Contents() ([]byte, error) {
	mergedConfig := f.userData
	mergedFiles := make([]cc.File, 0, len(f.userData.WriteFiles)+len(f.milpaFiles))
	mergedFiles = append(mergedFiles, f.userData.WriteFiles...)
	for _, wf := range f.milpaFiles {
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
