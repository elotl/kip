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

package cloudinitfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/coreos/yaml"
	cc "github.com/elotl/cloud-init/config"
	"github.com/stretchr/testify/assert"
)

func ciTmpFile(t *testing.T, contents string) (string, func()) {
	tempFile, err := ioutil.TempFile("", "kip-cloud-init")
	if err != nil {
		t.FailNow()
	}
	tempFile.Write([]byte(contents))
	return tempFile.Name(), func() { os.Remove(tempFile.Name()) }
}

func TestNewValid(t *testing.T) {
	data := `
apiVersion: v1
kind: Pod
`
	path, closer := ciTmpFile(t, data)
	defer closer()
	_, err := New(path)
	assert.NoError(t, err)
}

func TestNewInvalid(t *testing.T) {
	data := `
apiVersion: v1
      kind: Pod
`
	path, closer := ciTmpFile(t, data)
	defer closer()
	_, err := New(path)
	assert.Error(t, err)
}

func loadCloudConfigFromString(s string) (ucc cc.CloudConfig, err error) {
	err = yaml.Unmarshal([]byte(s), &ucc)
	return ucc, err
}

func TestWriteContent(t *testing.T) {
	userContent := `write_files:
- encoding: b64
  content: CiMgVGhpcyBmaWxlIGNvbnRyb2xzIHRoZSBzdGF0ZSBvZiBTRUxpbnV4...
  owner: root:root
  path: /etc/sysconfig/selinux
  permissions: "0644"`
	ucc, err := loadCloudConfigFromString(userContent)
	assert.NoError(t, err)
	cif := &File{
		userData: ucc,
		kipFiles: make(map[string]cc.File),
	}

	content := "A programmer, a plan, whatever dude..."
	path := "/usr/local/bin"
	permissions := "0600"
	cif.AddKipFile(content, path, permissions)

	expected := string(cloudInitHeader) + userContent + `
- content: A programmer, a plan, whatever dude...
  owner: root
  path: /usr/local/bin
  permissions: "0600"
`
	cloudInitContent, err := cif.Contents()
	assert.NoError(t, err)
	assert.Equal(t, expected, string(cloudInitContent))
}

func TestAddItzoFuncs(t *testing.T) {
	cif, err := New("")
	assert.NoError(t, err)
	cif.AddItzoVersion("")
	cif.AddItzoURL("")
	cloudInitContent, err := cif.Contents()
	assert.NoError(t, err)
	expected := string(cloudInitHeader) + "{}\n"
	assert.Equal(t, expected, string(cloudInitContent))

	versionString := "v1.0.0"
	cif.AddItzoVersion(versionString)
	cloudInitContent, err = cif.Contents()
	assert.NoError(t, err)
	expected = string(cloudInitHeader) + fmt.Sprintf(`write_files:
- content: %s
  owner: root
  path: %s
  permissions: "0444"
`, versionString, ItzoVersionPath)
	assert.Equal(t, expected, string(cloudInitContent))

	cif.AddItzoVersion("1.0.0")
	cloudInitContent, err = cif.Contents()
	assert.NoError(t, err)
	assert.Equal(t, expected, string(cloudInitContent))

	cif.ResetInstanceData()
	versionString = "dev"
	cif.AddItzoVersion(versionString)
	expected = string(cloudInitHeader) + fmt.Sprintf(`write_files:
- content: %s
  owner: root
  path: %s
  permissions: "0444"
`, versionString, ItzoVersionPath)
	cloudInitContent, err = cif.Contents()
	assert.NoError(t, err)
	assert.Equal(t, expected, string(cloudInitContent))

	cif.ResetInstanceData()
	versionString = "314"
	cif.AddItzoVersion(versionString)
	expected = string(cloudInitHeader) + fmt.Sprintf(`write_files:
- content: "%s"
  owner: root
  path: %s
  permissions: "0444"
`, versionString, ItzoVersionPath)
	cloudInitContent, err = cif.Contents()
	assert.NoError(t, err)
	assert.Equal(t, expected, string(cloudInitContent))

	cif.ResetInstanceData()
	urlString := "http://my-bucket.s3.com"
	cif.AddItzoURL(urlString)
	expected = string(cloudInitHeader) + fmt.Sprintf(`write_files:
- content: %s
  owner: root
  path: %s
  permissions: "0444"
`, urlString, ItzoURLPath)
	cloudInitContent, err = cif.Contents()
	assert.NoError(t, err)
	assert.Equal(t, expected, string(cloudInitContent))
}
