package tarutil

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"

	"github.com/golang/glog"
)

func CreatePackage(hostRootfs string, paths []string) (io.Reader, error) {
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	if hostRootfs == "" {
		hostRootfs = "/"
	}
	for i := range paths {
		path := filepath.Clean(filepath.Join(hostRootfs, paths[i]))
		err := filepath.Walk(path,
			func(fp string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				tfp := removePathPrefix(fp, hostRootfs)
				return AddFile(tw, fp, tfp)
			})
		if err != nil {
			return nil, err
		}
	}
	return bufio.NewReader(&buf), nil
}

func AddFile(tw *tar.Writer, source, target string) error {
	glog.Infof("Adding file %s->%s to package\n", source, target)
	fi, err := os.Lstat(source)
	if err != nil {
		glog.Errorf("Error LStat()ing %s: %v", source, err)
		return err
	}
	sldest := ""
	if fi.Mode()&os.ModeSymlink != 0 {
		// Check what the symlink points to.
		sldest, err = os.Readlink(source)
		if err != nil {
			glog.Errorf("Error Readlink() %s: %v", source, err)
			return err
		}
	}
	if fi.Mode()&os.ModeSocket != 0 {
		// Sockets are unsupported in archive/tar.
		return nil
	}
	header, err := tar.FileInfoHeader(fi, sldest)
	if err != nil {
		glog.Errorf("Error creating tar header for %s: %v", source, err)
		return err
	}
	// Files/directories are inside a top-level directory called "ROOTFS"
	// in Milpa packages.
	header.Name = filepath.Join(".", "ROOTFS", target)
	if err = tw.WriteHeader(header); err != nil {
		glog.Errorf("Error writing tar header for %s->%s: %v",
			source, target, err)
		return err
	}
	if !fi.Mode().IsRegular() {
		// Directory, link, hardlink, etc. No file content.
		return nil
	}
	file, err := os.Open(source)
	if err != nil {
		glog.Errorf("Error trying to open %s: %v", source, err)
		return err
	}
	defer file.Close()
	n, err := io.CopyN(tw, file, fi.Size())
	if err != nil {
		glog.Errorf("Error copying contents of %s->%s into tarball: %v",
			source, target, err)
		return err
	}
	glog.Infof("Copied %d bytes for %s->%s\n", n, source, target)
	return nil
}

func hasPathPrefix(path, prefix string) bool {
	if len(path) == len(prefix) {
		return path == prefix
	}
	if prefix == "" {
		return true
	}
	if len(path) > len(prefix) {
		if prefix[len(prefix)-1] == '/' || path[len(prefix)] == '/' {
			return path[:len(prefix)] == prefix
		}
	}
	return false
}

func removePathPrefix(path, prefix string) string {
	if prefix == "" || prefix == "/" || !hasPathPrefix(path, prefix) {
		return path
	}
	plen := len(prefix)
	if prefix[len(prefix)-1] == '/' {
		plen--
	}
	return path[plen:]
}
