package gce

import (
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/compute/metadata"
)

const (
	metadataTimeout = 4 * time.Second
	interfacesPath  = "instance/network-interfaces/"
)

func newMetadataClient() *metadata.Client {
	timeout := time.Duration(metadataTimeout)
	client := &http.Client{
		Timeout: timeout,
	}
	return metadata.NewClient(client)
}

func getMetadataTrimmed(c *metadata.Client, suffix string) (s string, err error) {
	s, err = c.Get(suffix)
	s = strings.TrimSpace(s)
	return
}

func getMetadataLines(c *metadata.Client, suffix string) ([]string, error) {
	j, err := c.Get(suffix)
	if err != nil {
		return nil, err
	}
	s := strings.Split(strings.TrimSpace(j), "\n")
	for i := range s {
		s[i] = strings.TrimSpace(s[i])
	}
	return s, nil
}
