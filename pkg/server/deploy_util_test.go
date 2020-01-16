package server

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"io"
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
)

func TestMakeDeployPackage(t *testing.T) {
	contents := map[string]packageFile{
		"file1":         packageFile{data: []byte("file1"), mode: 0777},
		"path/to/file2": {data: []byte("file2"), mode: 0400},
	}
	buf, err := makeDeployPackage(contents)
	assert.NoError(t, err)
	gzr, err := gzip.NewReader(bufio.NewReader(buf))
	assert.NoError(t, err)
	defer gzr.Close()
	tr := tar.NewReader(gzr)
	tfContents := make(map[string]packageFile)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		assert.NoError(t, err)
		if header.Typeflag == tar.TypeReg {
			data := make([]byte, header.Size)
			read_so_far := int64(0)
			for read_so_far < header.Size {
				n, err := tr.Read(data[read_so_far:])
				if err == io.EOF {
					break
				}
				assert.NoError(t, err)
				read_so_far += int64(n)
			}

			tfContents[header.Name[7:]] = packageFile{
				data: data,
				mode: int32(header.Mode),
			}
		}
	}
	assert.Equal(t, contents, tfContents)
}

func TestGetConfigMapFiles(t *testing.T) {
	trueVal := true
	tests := []struct {
		name          string
		vol           api.ConfigMapVolumeSource
		cm            v1.ConfigMap
		isErr         bool
		expectedFiles map[string]packageFile
	}{
		{
			name: "optional is skipped",
			vol: api.ConfigMapVolumeSource{
				LocalObjectReference: api.LocalObjectReference{"optional"},
				Optional:             &trueVal,
			},
			cm:            v1.ConfigMap{},
			isErr:         false,
			expectedFiles: map[string]packageFile{},
		},
	}
	for _, tc := range tests {
		files, err := getConfigMapFiles(&tc.vol, &tc.cm)
		if tc.isErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedFiles, files)
		}
	}
}
