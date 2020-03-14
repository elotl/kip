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

package yaml

import (
	"io"
	"reflect"
	"time"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/util"
	"github.com/mitchellh/mapstructure"
)

func DetectUnknownKeys(input io.Reader, into interface{}, bufferSize int) error {
	ifaceMap := map[string]interface{}{}
	decoder := NewYAMLOrJSONDecoder(input, bufferSize)
	err := decoder.Decode(&ifaceMap)
	if err != nil {
		return util.WrapError(err, "Unable to verify manifest keys")
	}
	stringToDateTimeHook := func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t == reflect.TypeOf(api.Time{}) && f == reflect.TypeOf("") {
			return time.Parse(time.RFC3339, data.(string))
		}
		return data, nil
	}
	cfg := &mapstructure.DecoderConfig{
		DecodeHook: stringToDateTimeHook,
		// Setting WeaklyTypedInput makes some unknown types work
		// (e.g. map[string][]byte) we probably don't want to use that
		// flag for other purposes other than validating keys
		WeaklyTypedInput: true,
		ErrorUnused:      true,
		TagName:          "json",
		Result:           into,
	}
	mapDecoder, err := mapstructure.NewDecoder(cfg)
	if err != nil {
		return util.WrapError(err, "Unable to verify manifest keys")
	}
	err = mapDecoder.Decode(ifaceMap)
	return err
}
