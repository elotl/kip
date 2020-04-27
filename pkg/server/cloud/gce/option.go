package gce

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

type ClientOption interface {
	Apply(*gceClient) error
}

type WithZone string

func (w WithZone) Apply(c *gceClient) error {
	c.zone = string(w)
	var err error
	c.region, err = zoneToRegion(c.zone)
	return err
}

// type WithProject string

// func (w WithProject) Apply(c *gceClient) error {
// 	c.projectID = string(w)
// 	return nil
// }

type WithVPCName string

func (w WithVPCName) Apply(c *gceClient) error {
	c.vpcName = string(w)
	return nil
}

type WithSubnetName string

func (w WithSubnetName) Apply(c *gceClient) error {
	c.subnetName = string(w)
	return nil
}

type WithCredentialsFile string

func (w WithCredentialsFile) Apply(c *gceClient) error {
	envVar := "GOOGLE_APPLICATION_CREDENTIALS"
	if os.Getenv(envVar) == "" {
		err := os.Setenv(envVar, string(w))
		if err != nil {
			return err
		}
	}
	return nil
}

type withCredentials struct {
	clientEmail string
	privateKey  string
}

func WithCredentials(email, key string) withCredentials {
	return withCredentials{
		clientEmail: email,
		privateKey:  key,
	}
}

func (w withCredentials) Apply(c *gceClient) error {
	if c.projectID == "" {
		return fmt.Errorf("project ID not specified")
	}

	creds := struct {
		Type        string `json:"type"`
		ClientEmail string `json:"client_email"`
		PrivateKey  string `json:"private_key"`
		ProjectID   string `json:"project_id"`
	}{
		Type:        "service_account",
		ProjectID:   c.projectID,
		ClientEmail: w.clientEmail,
		PrivateKey:  w.privateKey,
	}
	b, err := json.Marshal(creds)
	if err != nil {
		return err
	}

	ctx := context.Background()
	service, err := compute.NewService(ctx, option.WithCredentialsJSON(b))
	if err != nil {
		return err
	}
	c.service = service
	return nil
}
