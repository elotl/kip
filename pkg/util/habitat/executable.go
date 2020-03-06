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

package habitat

import (
	"os"
	"path/filepath"
	"strings"
)

// Taken from https://github.com/kardianos/osext/blob/master/osext_procfs.go
// (i didn't want to vendor the whole damn thing and it never gets updated...)
func Executable() (string, error) {
	const deletedTag = " (deleted)"
	execpath, err := os.Readlink("/proc/self/exe")
	if err != nil {
		return execpath, err
	}
	execpath = strings.TrimSuffix(execpath, deletedTag)
	execpath = strings.TrimPrefix(execpath, deletedTag)
	return execpath, nil
}

func ExecutableFolder() (string, error) {
	p, err := Executable()
	if err != nil {
		return "", err
	}

	return filepath.Dir(p), nil
}
