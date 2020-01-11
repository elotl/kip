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
	assert.Len(t, kc.config.AuthInfos, 1)
	for _, authInfo := range kc.config.AuthInfos {
		assert.Equal(t, authInfo.Token, string(origToken))
	}
	newToken := "new-token"
	err = ioutil.WriteFile(tokenFile, []byte(newToken), 0644)
	assert.NoError(t, err)
	err = kc.Refresh()
	assert.NoError(t, err)
	assert.Len(t, kc.config.AuthInfos, 1)
	for _, authInfo := range kc.config.AuthInfos {
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
	for _, cluster := range loadedKc.config.Clusters {
		cluster.LocationOfOrigin = ""
	}
	for _, context := range loadedKc.config.Contexts {
		context.LocationOfOrigin = ""
	}
	for _, authInfo := range loadedKc.config.AuthInfos {
		authInfo.LocationOfOrigin = ""
	}
	kcJson, err := kc.toJSON()
	assert.NoError(t, err)
	loadedKcJson, err := loadedKc.toJSON()
	assert.NoError(t, err)
	assert.Equal(t, kcJson, loadedKcJson)
}
