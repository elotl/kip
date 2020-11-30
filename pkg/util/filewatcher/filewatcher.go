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

package filewatcher

import (
	"io/ioutil"
	"os"
	"time"

	"k8s.io/klog"
)

// Watches a file on the local filesystem pointed to by path. This
// will refresh the cached file contents periodically if the file has
// changed. CheckPeriod serves to ensure that we don't stat the file
// incessantly.
type Watcher interface {
	Contents() string
	Version() int
}

type File struct {
	path         string
	modTime      time.Time
	fileContents string
	statTime     time.Time
	CheckPeriod  time.Duration
	version      int
}

func New(path string) *File {
	fw := &File{
		path:        path,
		CheckPeriod: 10 * time.Second,
		version:     1,
	}
	fw.refresh()
	return fw
}

func (fw *File) Version() int {
	return fw.version
}

func (fw *File) refresh() {
	if fw.path == "" {
		return
	}
	now := time.Now()
	if fw.statTime.Add(fw.CheckPeriod).Before(now) {
		info, err := os.Stat(fw.path)
		if err != nil {
			klog.Warningf("Error getting file info at %s: %s", fw.path, err)
		}
		if info.ModTime().After(fw.modTime) {
			c, err := ioutil.ReadFile(fw.path)
			if err != nil {
				klog.Warningf("Error reading contents of file at %s: %s", fw.path, err)
				return
			}
			fw.version += 1
			fw.fileContents = string(c)
			fw.modTime = info.ModTime()
		}
		fw.statTime = now
	}
}

func (fw *File) Contents() string {
	fw.refresh()
	return fw.fileContents
}
