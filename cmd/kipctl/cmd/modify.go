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

package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/elotl/cloud-instance-provider/pkg/util/yaml"
)

const (
	modifyCreate = iota
	modifyUpdate
	modifyDelete
	modifyDeleteCascade
)

var (
	containersSpec = []byte("\"containers\"")
)

func getFieldValue(obj interface{}, fieldName string) string {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.Struct {
		v := v.FieldByName(fieldName)
		if v.IsValid() {
			return v.String()
		}
	}
	return "<unknown>"
}

// Warn the user if they create a manifest with no units but
// the manifest has a "containers" section. Note, by this point
// b ([]byte) has been converted from yaml to json.
func warnAboutContainers(obj api.MilpaObject, b []byte) {
	warnUser := false
	switch obj := obj.(type) {
	case *api.Pod:
		if len(obj.Spec.Units) == 0 {
			warnUser = true
		}
	}
	if warnUser && bytes.Contains(b, containersSpec) {
		fmt.Printf("Warning: manifest contains no units but does seem to have a %s section. Kip uses \"units\" not containers", string(containersSpec))
	}
}

func modifyFromBytes(client clientapi.KipClient, b []byte, operation int) (string, error) {
	b = bytes.TrimSpace(b)
	if len(b) == 0 || bytes.Equal(b, []byte("null")) {
		return "", nil
	}

	obj, err := api.Decode(b)
	if err != nil {
		return "", util.WrapError(err, "Error decoding object from input file")
	}
	kind := getFieldValue(obj, "Kind")
	name := getFieldValue(obj, "Name")

	if operation == modifyCreate || operation == modifyUpdate {
		warnAboutContainers(obj, b)
	}

	var reply *clientapi.APIReply
	switch operation {
	case modifyCreate:
		r := &clientapi.CreateRequest{
			Manifest: b,
		}
		reply, err = client.Create(context.Background(), r)
	case modifyUpdate:
		r := &clientapi.UpdateRequest{
			Manifest: b,
		}
		reply, err = client.Update(context.Background(), r)
	case modifyDelete, modifyDeleteCascade:
		cascade := operation == modifyDeleteCascade
		r := &clientapi.DeleteRequest{
			Kind:    []byte(kind),
			Name:    []byte(name),
			Cascade: cascade,
		}
		reply, err = client.Delete(context.Background(), r)
	}
	if err != nil {
		return "", util.WrapError(
			err, "Error creating resource %s - %s", kind, name)
	}
	if reply.Status < 200 || reply.Status >= 400 {
		return "", fmt.Errorf("Create %s - %s returned %d: %s",
			kind, name, reply.Status, reply.Body)
	}
	if len(reply.Warning) != 0 {
		fmt.Fprintf(os.Stderr, "Warning: %s\n", string(reply.Warning))
	}
	return name, nil
}

func modify(client clientapi.KipClient, appManifestFile string, operation int) []error {
	r, err := os.Open(appManifestFile)
	dieIfError(err, "Error opening manifest file")
	decoder := yaml.NewYAMLOrJSONDecoder(r, 16000)
	errors := make([]error, 0)
	for {
		raw := json.RawMessage{}
		if err := decoder.Decode(&raw); err != nil {
			if err == io.EOF {
				break
			}
			wrappedErr := fmt.Errorf("Error decoding portion of file: %v", err)
			errors = append(errors, wrappedErr)
			continue
		}

		name, err := modifyFromBytes(client, raw, operation)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			errors = append(errors, err)
			continue
		}
		fmt.Printf("%s\n", name)
	}
	return errors
}
