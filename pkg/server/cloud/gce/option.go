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

package gce

import (
	"context"
	"encoding/json"
	"fmt"

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

type WithPrivateIPOnly bool

func (w WithPrivateIPOnly) Apply(c *gceClient) error {
	privateIPOnly := bool(w)
	c.usePublicIPs = !privateIPOnly
	return nil
}

type WithCredentialsFile string

func (w WithCredentialsFile) Apply(c *gceClient) error {
	// envVar := "GOOGLE_APPLICATION_CREDENTIALS"
	// if os.Getenv(envVar) == "" {
	// 	err := os.Setenv(envVar, string(w))
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	c.service, err = compute.NewService(ctx, option.WithCredentialsFile(string(w)))
	return err
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

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	c.service, err = compute.NewService(ctx, option.WithCredentialsJSON(b))
	return err
}
