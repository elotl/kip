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

package tarutil

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePackage(t *testing.T) {
	file1, err := ioutil.TempFile("", "milpa")
	assert.NoError(t, err)
	defer file1.Close()
	file2, err := ioutil.TempFile("", "milpa")
	assert.NoError(t, err)
	defer file2.Close()
	// Create a broken symlink.
	symlink1 := file2.Name() + ".lnk"
	err = os.Symlink(file2.Name(), symlink1)
	assert.NoError(t, err)
	os.Remove(file2.Name())
	filenames := []string{file1.Name(), symlink1}
	// Unix sockets are not supported, but they should not cause an error.
	sockfile, err := ioutil.TempFile("", "milpa.sock")
	assert.NoError(t, err)
	defer sockfile.Close()
	os.Remove(sockfile.Name())
	sock, err := net.Listen("unix", sockfile.Name())
	assert.NoError(t, err)
	defer sock.Close()
	buf, err := CreatePackage("", filenames)
	assert.NoError(t, err)
	gzr, err := gzip.NewReader(buf)
	assert.NoError(t, err)
	defer gzr.Close()
	tr := tar.NewReader(gzr)
	nfiles := 0
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		assert.NoError(t, err)
		name := strings.Replace(header.Name, "ROOTFS", "", 1)
		assert.Contains(t, filenames, name)
		nfiles++
	}
	assert.Equal(t, len(filenames), nfiles)
}

func TestCreatePackageError(t *testing.T) {
	filenames := []string{"/does/not/exist"}
	buf, err := CreatePackage("", filenames)
	assert.Error(t, err)
	assert.Nil(t, buf)
}

//func hasPathPrefix(path, prefix string) bool {
func TestHasPathPrefix(t *testing.T) {
	testCases := []struct {
		path   string
		prefix string
		result bool
	}{
		{"/a/b/c.d", "/a/b", true},
		{"/a/b/c.d", "/a/b/", true},
		{"/a/b/c.d", "/a/bb/", false},
		{"/a/b/c.d", "/a/b/c.d/", false},
		{"/a/b/c.d", "/a/b/c.d/e", false},
		{"/a/b/c.d", "/a/b/c.d/e/", false},
		{"/a/b/c.d", "", true},
		{"/a/b/c.d", "a", false},
		{"/a/b/c.d", "a/b", false},
		{"a/b/c.d", "a/b", true},
		{"a/b/c.d", "a/b/", true},
		{"a/b/c.d", "a/bb/", false},
		{"a/b/c.d", "", true},
		{"a/b/c.d", "a", true},
		{"a/b/c.d", "a/", true},
		{"a/b/c.d", "a/b", true},
		{"a/b/c.d", "a/b/", true},
	}
	for i, tc := range testCases {
		res := hasPathPrefix(tc.path, tc.prefix)
		assert.Equal(t,
			tc.result,
			res,
			fmt.Sprintf("Test case %d failed %q has prefix %q returned %v",
				i+1, tc.path, tc.prefix, res))
	}
}

//func removePathPrefix(path, prefix string) string
func TestRemovePathPrefix(t *testing.T) {
	testCases := []struct {
		path   string
		prefix string
		result string
	}{
		{"/a/b/c.d", "/a/b", "/c.d"},
		{"/a/b/c.d", "/a/b/", "/c.d"},
		{"/a/b/c.d", "/a/bb/", "/a/b/c.d"},
		{"/a/b/c.d", "/a/b/c.d/", "/a/b/c.d"},
		{"/a/b/c.d", "/a/b/c.d/e", "/a/b/c.d"},
		{"/a/b/c.d", "/a/b/c.d/e/", "/a/b/c.d"},
		{"/a/b/c.d", "", "/a/b/c.d"},
		{"/a/b/c.d", "a", "/a/b/c.d"},
		{"/a/b/c.d", "a/b", "/a/b/c.d"},
		{"a/b/c.d", "a/b", "/c.d"},
		{"a/b/c.d", "a/b/", "/c.d"},
		{"a/b/c.d", "a/bb/", "a/b/c.d"},
		{"a/b/c.d", "", "a/b/c.d"},
		{"a/b/c.d", "a", "/b/c.d"},
		{"a/b/c.d", "a/", "/b/c.d"},
		{"a/b/c.d", "a/b", "/c.d"},
		{"a/b/c.d", "a/b/", "/c.d"},
	}
	for i, tc := range testCases {
		res := removePathPrefix(tc.path, tc.prefix)
		assert.Equal(t,
			tc.result,
			res,
			fmt.Sprintf("Test case %d failed %q removing prefix %q returned %q",
				i+1, tc.path, tc.prefix, res))
	}
}
