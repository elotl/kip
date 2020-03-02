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

package server

import (
	"bytes"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/util/yaml"
)

func VersionAndKind(m []byte) (string, string, error) {
	var typeMeta api.TypeMeta
	decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(m), 8000)
	err := decoder.Decode(&typeMeta)
	if err != nil {
		return "", "", err
	}
	return typeMeta.APIVersion, typeMeta.Kind, nil
}
