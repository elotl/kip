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
