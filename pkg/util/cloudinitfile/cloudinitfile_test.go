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
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/util/filewatcher"
	"github.com/stretchr/testify/assert"
)

func TestValidateValid(t *testing.T) {
	data := `
apiVersion: v1
kind: Pod
`
	cif := New("")
	err := cif.Validate(data)
	assert.NoError(t, err)
}

func TestValidateInvalid(t *testing.T) {
	data := `
apiVersion: v1
      kind: Pod
`
	cif := New("")
	err := cif.Validate(data)
	assert.Error(t, err)
}

func TestWriteMilpaFile(t *testing.T) {
	content := `a man a plan
a canal
panama`
	path := "/usr/local/bin"
	permissions := "0600"
	m := MilpaFile{
		content:     content,
		path:        path,
		permissions: permissions,
	}
	expected := `  - content: |
      a man a plan
      a canal
      panama
    path: /usr/local/bin
    permissions: 0600
`
	assert.Equal(t, expected, string(m.Content()))
}

func TestWriteContent(t *testing.T) {
	userContent := `write_files:
-   encoding: b64
    content: CiMgVGhpcyBmaWxlIGNvbnRyb2xzIHRoZSBzdGF0ZSBvZiBTRUxpbnV4...
    owner: root:root
    path: /etc/sysconfig/selinux
    permissions: '0644'`
	cif := &File{
		userData: &filewatcher.FakeWatcher{
			FakeContents: userContent,
			FakeVersion:  2,
		},
		milpaFiles: make(map[string]MilpaFile),
	}

	content := "A programmer, a plan, whatever dude..."
	path := "/usr/local/bin"
	permissions := "0600"
	cif.AddMilpaFile(content, path, permissions)

	expected := `
milpa_files:
  - content: |
      A programmer, a plan, whatever dude...
    path: /usr/local/bin
    permissions: 0600

` + userContent
	cloudInitContent, err := cif.Contents()
	assert.NoError(t, err)
	assert.Equal(t, expected, string(cloudInitContent))
}

func TestAddItzoFuncs(t *testing.T) {
	cif := &File{
		userData:   &filewatcher.FakeWatcher{},
		milpaFiles: make(map[string]MilpaFile),
	}
	cif.AddItzoVersion("")
	cif.AddItzoURL("")
	cloudInitContent, err := cif.Contents()
	assert.NoError(t, err)
	assert.Equal(t, "\n", string(cloudInitContent))
	versionString := "v1.0.0"
	cif.AddItzoVersion(versionString)
	cloudInitContent, err = cif.Contents()
	assert.NoError(t, err)
	expected := fmt.Sprintf(`
milpa_files:
  - content: |
      %s
    path: %s
    permissions: 0444

`, versionString, ItzoVersionPath)
	assert.Equal(t, expected, cloudInitContent)
	cif.AddItzoVersion("1.0.0")
	cloudInitContent, err = cif.Contents()
	assert.NoError(t, err)
	assert.Equal(t, expected, cloudInitContent)

	cif.ResetInstanceData()
	urlString := "http://my-bucket.s3.com"
	cif.AddItzoURL(urlString)
	expected = fmt.Sprintf(`
milpa_files:
  - content: |
      %s
    path: %s
    permissions: 0444

`, urlString, ItzoURLPath)
	cloudInitContent, err = cif.Contents()
	assert.NoError(t, err)
	assert.Equal(t, expected, cloudInitContent)

}
