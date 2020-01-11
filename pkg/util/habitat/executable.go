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
