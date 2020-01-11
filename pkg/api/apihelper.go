package api

import (
	"fmt"
	"reflect"
)

type upgradeFn func(string, string, MilpaObject) (MilpaObject, error)

var (
	storageCodec = JsonCodec{}
	upgradeFuncs = map[string]upgradeFn{}
	versionMap   = map[string]map[string]reflect.Type{}
)

func init() {
	addKnownTypes("v1",
		Pod{},
		PodList{},
		Node{},
		NodeList{},
		Service{},
		ServiceList{},
		Secret{},
		SecretList{},
		Event{},
		EventList{},
		LogFile{},
		LogFileList{},
		Usage{},
		UsageList{},
		UsageReport{},
		Metrics{},
		MetricsList{},
	)
	// When we move to a new API version, register the conversion function for
	// the previous API version here. E.g. if the current API version is v3:
	//
	//     upgradeFuncs["v1"] = v1.upgradeV1ToV2
	//     upgradeFuncs["v2"] = upgradeV2ToV3
	//
	// The idea is to go through each version upgrade one by one, until we
	// reach the latest (current) API version. See upgrade().
}

// Register an API version with its supported types.
func addKnownTypes(version string, types ...MilpaObject) {
	knownTypes := map[string]reflect.Type{}
	for _, obj := range types {
		t := reflect.TypeOf(obj)
		knownTypes[t.Name()] = t
	}
	versionMap[version] = knownTypes
}

// Returns a pointer to the object.
func ptr(obj MilpaObject) MilpaObject {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		return obj
	}
	v2 := reflect.New(v.Type())
	v2.Elem().Set(v)
	return v2.Interface().(MilpaObject)
}

// Returns the kind and apiversion of the struct.
func metainfo(obj MilpaObject) (string, string, error) {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr {
		return "", "", fmt.Errorf("Expected pointer, but got %v",
			v.Type().Name())
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return "", "", fmt.Errorf("Expected struct, but got %v: %v (%#v)",
			v.Kind(), v.Type().Name(), v.Interface())
	}
	kind := v.FieldByName("Kind")
	if !kind.IsValid() {
		return "", "", fmt.Errorf("%v does not have TypeMeta", v)
	}
	apiversion := v.FieldByName("APIVersion")
	if !apiversion.IsValid() {
		return "", "", fmt.Errorf("%v does not have TypeMeta", v)
	}
	return kind.String(), apiversion.String(), nil
}

// Encode encodes a struct or a struct pointer to a JSON string that can be
// persisted in a registry.
func Encode(obj MilpaObject) (data []byte, err error) {
	obj = ptr(obj) // Make sure obj is a pointer.
	kind, apiversion, err := metainfo(obj)
	if err != nil {
		return nil, err
	}
	// Make sure persistance for this type is supported.
	knownTypes, found := versionMap[apiversion]
	if !found {
		return nil, fmt.Errorf("Can't encode %v: unknown API version %s",
			obj, apiversion)
	}
	if _, contains := knownTypes[kind]; !contains {
		return nil, fmt.Errorf("Can't encode %v: unknown type %s", obj, kind)
	}
	return storageCodec.Marshal(obj)
}

// Decode decodes a JSON string to an internal type. The output is a struct
// pointer. This is where we can "upgrade" persisted objects to the latest API
// version.
func Decode(data []byte) (MilpaObject, error) {
	tm := TypeMeta{}
	err := storageCodec.Unmarshal(data, &tm)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling %s into object: %+v",
			string(data), err)
	}
	knownTypes, found := versionMap[tm.APIVersion]
	if !found {
		return nil, fmt.Errorf("Unsupported API verson: %s", tm.APIVersion)
	}
	objType, found := knownTypes[tm.Kind]
	if !found {
		return nil, fmt.Errorf("Unsupported object kind: %+v", tm)
	}
	obj := reflect.New(objType).Interface().(MilpaObject)
	err = storageCodec.Unmarshal(data, obj)
	if err != nil {
		return nil, err
	}
	// If this object is of an old API version, it needs to be upgraded.
	obj, err = upgrade(tm.APIVersion, tm.Kind, obj)
	if err != nil {
		return nil, err
	}
	_, apiversion, err := metainfo(obj)
	if err != nil {
		return nil, err
	}
	if apiversion != LatestAPIVersion {
		return nil, fmt.Errorf(
			"Error upgrading object (v %s) to latest API version (v %s)",
			apiversion, LatestAPIVersion)
	}
	return obj, nil
}

// DecodeInto parses a JSON string and stores it in obj, which needs to be a
// pointer to a struct.
func DecodeInto(data []byte, obj MilpaObject) error {
	decoded, err := Decode(data)
	if err != nil {
		return err
	}
	v := reflect.ValueOf(obj)
	iv := reflect.ValueOf(decoded)
	if !iv.Type().AssignableTo(v.Type()) {
		return fmt.Errorf("%s is not assignable to %s", v.Type(), iv.Type())
	}
	v.Elem().Set(iv.Elem())
	return nil
}

func upgrade(apiversion, kind string, obj MilpaObject) (MilpaObject, error) {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr {
		// Not a pointer.
		value := reflect.New(v.Type())
		value.Elem().Set(v)
		result, err := upgrade(apiversion, kind, value.Interface().(MilpaObject))
		if err != nil {
			return nil, err
		}
		// Return a pointer if obj was a pointer.
		return reflect.ValueOf(result).Elem().Interface().(MilpaObject), nil
	}
	for apiversion != LatestAPIVersion {
		fn, found := upgradeFuncs[apiversion]
		if !found {
			return nil, fmt.Errorf("Can't find upgrade function for %s %s",
				kind, apiversion)
		}
		var err error
		obj, err = fn(apiversion, kind, obj)
		if err != nil {
			return nil, err
		}
		_, apiversion, err = metainfo(obj)
		if err != nil {
			return nil, err
		}
	}
	return obj, nil
}
