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
	"regexp"

	"github.com/elotl/kip/pkg/util"
	"github.com/go-yaml/yaml"
)

const semverRegexFmt string = `v?([0-9]+)(\.[0-9]+)(\.[0-9]+)?` +
	`(-([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?` +
	`(\+([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?`

var (
	itzoDir          = "/tmp/itzo"
	ItzoVersionPath  = itzoDir + "/itzo_version"
	ItzoURLPath      = itzoDir + "/itzo_url"
	CellConfigPath   = itzoDir + "/cell_config.yaml"
	cloudInitHeader  = []byte("#cloud-config\n")
	maxCloudInitSize = 16000
	semverRegex      = regexp.MustCompile("^" + semverRegexFmt + "$")
)

type CloudInitFile struct {
	Encoding           string `yaml:"encoding,omitempty" valid:"^(base64|b64|gz|gzip|gz\\+base64|gzip\\+base64|gz\\+b64|gzip\\+b64)$"`
	Content            string `yaml:"content,omitempty"`
	Owner              string `yaml:"owner,omitempty"`
	Path               string `yaml:"path,omitempty"`
	RawFilePermissions string `yaml:"permissions,omitempty" valid:"^0?[0-7]{3,4}$"`
}

type User struct {
	Name                 string   `yaml:"name,omitempty"`
	PasswordHash         string   `yaml:"passwd,omitempty"`
	SSHAuthorizedKeys    []string `yaml:"ssh_authorized_keys,omitempty"`
	SSHImportGithubUser  string   `yaml:"coreos_ssh_import_github,omitempty"       deprecated:"trying to fetch from a remote endpoint introduces too many intermittent errors"`
	SSHImportGithubUsers []string `yaml:"coreos_ssh_import_github_users,omitempty" deprecated:"trying to fetch from a remote endpoint introduces too many intermittent errors"`
	SSHImportURL         string   `yaml:"coreos_ssh_import_url,omitempty"          deprecated:"trying to fetch from a remote endpoint introduces too many intermittent errors"`
	GECOS                string   `yaml:"gecos,omitempty"`
	Homedir              string   `yaml:"homedir,omitempty"`
	NoCreateHome         bool     `yaml:"no_create_home,omitempty"`
	PrimaryGroup         string   `yaml:"primary_group,omitempty"`
	Groups               []string `yaml:"groups,omitempty"`
	NoUserGroup          bool     `yaml:"no_user_group,omitempty"`
	System               bool     `yaml:"system,omitempty"`
	NoLogInit            bool     `yaml:"no_log_init,omitempty"`
	Shell                string   `yaml:"shell,omitempty"`
}

// CloudConfig encapsulates the entire cloud-config configuration file and maps
// directly to YAML. Fields that cannot be set in the cloud-config (fields
// used for internal use) have the YAML tag '-' so that they aren't marshalled.
type CloudConfig struct {
	SSHAuthorizedKeys []string        `yaml:"ssh_authorized_keys,omitempty"`
	WriteFiles        []CloudInitFile `yaml:"write_files,omitempty"`
	Hostname          string          `yaml:"hostname,omitempty"`
	Users             []User          `yaml:"users,omitempty"`
	RunCmd            []string        `yaml:"runcmd,omitempty"`
	// this one is legacy, can be removed when no more kip controllers use it
	MilpaFiles []CloudInitFile `yaml:"milpa_files,omitempty"`
	// Todo: add additional parameters supported by traditional cloud-init
}

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
	f.AddKipFile(version, ItzoVersionPath, "0444")
}

func (f *File) AddItzoURL(url string) {
	if url == "" {
		return
	}
	f.AddKipFile(url, ItzoURLPath, "0444")
}

func (f *File) AddCellConfig(cfg map[string]string) {
	if len(cfg) == 0 {
		return
	}
	buf, err := yaml.Marshal(cfg)
	if err != nil {
		return
	}
	f.AddKipFile(string(buf), CellConfigPath, "0444")
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
