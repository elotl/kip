package gce

import (
	"fmt"
	"os"

	"cloud.google.com/go/compute/metadata"
	"google.golang.org/api/compute/v1"
)

type gceClient struct {
	service      *compute.Service
	controllerID string
	nametag      string
	projectID    string
	region       string
	zone         string
	vpcName      string
	subnetName   string
	vpcCIDR      string // Unsure if this is needed?
	subnetCIDR   string // Unsure if this is needed?

	usePublicIPs bool
	// bootSecurityGroupIDs []string
	// cloudStatus          *cloud.AZSubnetStatus
}

func NewGCEClient(controllerID, nametag, projectID string, opts ...ClientOption) (*gceClient, error) {
	client := &gceClient{
		service:      nil,
		controllerID: controllerID,
		nametag:      nametag,
		projectID:    projectID,
	}
	var err error
	if client.projectID == "" {
		client.projectID, err = client.autodetectProject()
		if err != nil {
			return nil, err
		}
	}
	for _, opt := range opts {
		err = opt.Apply(client)
		if err != nil {
			return nil, err
		}
	}
	if client.service == nil {
		client.service, err = serviceFromEnvironment()
		if err != nil {
			return nil, err
		}
	}
	if client.zone == "" {
		client.region, client.zone, err = client.autodetectRegionAndZone()
		if err != nil {
			return nil, err
		}
	}
	err = client.setupVPC(client.vpcName)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Try to get credentials from environment variables or from
// the environment the machine is running in
func serviceFromEnvironment() (*compute.Service, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (c *gceClient) autodetectProject() (string, error) {
	envVal := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if envVal != "" {
		return envVal, nil
	} else if metadata.OnGCE() {
		md := newMetadataClient()
		projectID, _ := md.ProjectID()
		if projectID != "" {
			return projectID, nil
		}
	}
	return "", fmt.Errorf("Could not get GCE project ID from environment or from instance metadata service")
}
