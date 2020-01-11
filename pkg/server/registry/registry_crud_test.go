package registry

import (
	"reflect"
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/stretchr/testify/assert"
)

type registryCRUDTester struct {
	t      *testing.T
	regs   map[string]Registryer
	closer func()
}

type modifyObjFunc func(o api.MilpaObject) api.MilpaObject
type equalObjFunc func(lhs api.MilpaObject, rhs api.MilpaObject) bool

func getStringField(obj interface{}, field string) string {
	r := reflect.ValueOf(obj)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func getLenField(obj interface{}, field string) int {
	r := reflect.ValueOf(obj)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Len()
}

func checkUpdatedMapField(t *testing.T, obj api.MilpaObject, field, k, v string) {
	r := reflect.ValueOf(obj)
	f := reflect.Indirect(r).FieldByName(field)
	valVal := f.MapIndex(reflect.ValueOf(k))
	assert.Equal(t, v, valVal.String(), "expected to find %s=%s in %s", k, v, field)
}

func modifyMapField(obj api.MilpaObject, field, k, v string) api.MilpaObject {
	r := reflect.ValueOf(obj)
	f := reflect.Indirect(r).FieldByName(field)
	mapType := reflect.MapOf(reflect.TypeOf(k), reflect.TypeOf(v))
	newMap := reflect.MakeMap(mapType)
	keyVal := reflect.ValueOf(k)
	valVal := reflect.ValueOf(v)
	newMap.SetMapIndex(keyVal, valVal)
	f.Set(newMap)
	return obj
}

func (r *registryCRUDTester) registryTestType(a, b api.MilpaObject, mod modifyObjFunc, eq equalObjFunc) {
	objType := getStringField(a, "Kind")
	aName := getStringField(a, "Name")
	bName := getStringField(b, "Name")
	reg, ok := r.regs[objType]
	assert.True(r.t, ok)
	createdA, err := reg.Create(a)
	assert.NoError(r.t, err, "Error creating object")

	createdB, err := reg.Create(b)
	assert.NoError(r.t, err, "Error creating object")
	assert.True(r.t, eq(a, createdA) && eq(b, createdB), "Create returns unequal values for %s", objType)

	getA, err := reg.Get(aName)
	assert.NoError(r.t, err, "Get returned an error")
	assert.True(r.t, eq(a, getA), "Get returns an unequal value for %s", objType)

	items, err := reg.List()
	assert.NoError(r.t, err, "List returned an error for %s", objType)
	assert.Equal(r.t, 2, getLenField(items, "Items"), "List returned wrong number of items")

	modA := mod(getA)
	modA = modifyMapField(modA, "Labels", "env", "dev")
	modA = modifyMapField(modA, "Annotations", "foo", "bar")
	updatedA, err := reg.Update(modA)
	assert.NoError(r.t, err, "Update returned an error for %s", objType)

	checkUpdatedMapField(r.t, updatedA, "Labels", "env", "dev")
	checkUpdatedMapField(r.t, updatedA, "Annotations", "foo", "bar")
	assert.True(r.t, eq(modA, updatedA), "update didn't correctly update the object for %s", objType)

	_, err = reg.Delete(aName)
	assert.NoError(r.t, err, "Delete returned an error for %s", objType)
	_, err = reg.Delete(bName)
	assert.NoError(r.t, err, "Delete returned an error for %s", objType)
	items, err = reg.List()
	assert.NoError(r.t, err, "List returned an errror for %s", objType)
	// Special case: Pods remain in the terminated state so we expect
	// to still have 2 terminated pods after deleting
	if _, ok := a.(*api.Pod); ok {
		assert.Equal(r.t, 2, getLenField(items, "Items"), "Unexpected number of %s objects after delete", objType)
	} else {
		assert.Equal(r.t, 0, getLenField(items, "Items"), "Unexpected number of %s objects after delete", objType)
	}
}

func createCRUDTestRegistries(t *testing.T) *registryCRUDTester {
	// create all the reg and return a map of type to registry
	// also return a function that'll close everything
	podRegistry, c1 := SetupTestPodRegistry()

	regs := map[string]Registryer{
		"Pod": podRegistry,
	}

	return &registryCRUDTester{
		t:      t,
		regs:   regs,
		closer: c1,
	}
}

func TestRegistryCRUD(t *testing.T) {
	t.Parallel()

	regTester := createCRUDTestRegistries(t)
	defer regTester.closer()

	// Pod
	regTester.registryTestType(
		api.GetFakePod(),
		api.GetFakePod(),
		func(o api.MilpaObject) api.MilpaObject {
			p := o.(*api.Pod)
			p.Spec.InstanceType = "t0.large"
			return p
		},
		func(lhs api.MilpaObject, rhs api.MilpaObject) bool {
			a := lhs.(*api.Pod)
			b := rhs.(*api.Pod)
			return reflect.DeepEqual(a.Spec, b.Spec)
		},
	)
}
