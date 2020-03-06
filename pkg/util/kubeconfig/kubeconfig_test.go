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

package kubeconfig

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestConfig(t *testing.T) (*Kubeconfig, string) {
	tmpdir, err := ioutil.TempDir("", "test-kubeconfig-")
	assert.NoError(t, err)
	tokenFile := filepath.Join(tmpdir, "token")
	err = ioutil.WriteFile(tokenFile, []byte("token"), 0644)
	assert.NoError(t, err)
	rootCAFile := filepath.Join(tmpdir, "ca")
	err = ioutil.WriteFile(rootCAFile, []byte("CA-data"), 0644)
	assert.NoError(t, err)
	kc, err := NewFromToken("user", "cluster", "http://serv.er:8080", tokenFile, rootCAFile)
	assert.NoError(t, err)
	assert.NotNil(t, kc)
	return kc, tmpdir
}

func TestNewFromToken(t *testing.T) {
	_, tmpdir := createTestConfig(t)
	defer os.RemoveAll(tmpdir)
}

func TestRefresh(t *testing.T) {
	kc, tmpdir := createTestConfig(t)
	defer os.RemoveAll(tmpdir)
	tokenFile := filepath.Join(tmpdir, "token")
	origToken, err := ioutil.ReadFile(tokenFile)
	assert.NoError(t, err)
	assert.Len(t, kc.Config.AuthInfos, 1)
	for _, authInfo := range kc.Config.AuthInfos {
		assert.Equal(t, authInfo.Token, string(origToken))
	}
	newToken := "new-token"
	err = ioutil.WriteFile(tokenFile, []byte(newToken), 0644)
	assert.NoError(t, err)
	err = kc.Refresh()
	assert.NoError(t, err)
	assert.Len(t, kc.Config.AuthInfos, 1)
	for _, authInfo := range kc.Config.AuthInfos {
		assert.Equal(t, authInfo.Token, newToken)
	}
}

func TestJson(t *testing.T) {
	kc, tmpdir := createTestConfig(t)
	defer os.RemoveAll(tmpdir)
	assert.NotNil(t, kc)
	kcFile := filepath.Join(tmpdir, "kubeconfig")
	err := kc.WriteToFile(kcFile)
	assert.NoError(t, err)
	loadedKc, err := LoadFromFile(kcFile)
	assert.NoError(t, err)
	for _, cluster := range loadedKc.Config.Clusters {
		cluster.LocationOfOrigin = ""
	}
	for _, context := range loadedKc.Config.Contexts {
		context.LocationOfOrigin = ""
	}
	for _, authInfo := range loadedKc.Config.AuthInfos {
		authInfo.LocationOfOrigin = ""
	}
	kcJson, err := kc.toJSON()
	assert.NoError(t, err)
	loadedKcJson, err := loadedKc.toJSON()
	assert.NoError(t, err)
	assert.Equal(t, kcJson, loadedKcJson)
}
