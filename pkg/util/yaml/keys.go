package yaml

import (
	"io"
	"reflect"
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/util"
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
