package server

import (
	"flag"
	"testing"
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/etcd"
	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/stretchr/testify/assert"
)

var runFunctional = flag.Bool("functional", false, "run functional tests")

func TestFilterReplyObject(t *testing.T) {
	t.Parallel()
	// create a registry
	// add an object
	// make sure we get the correct form of the object
	MaxEventListSize = 2
	er, closer := registry.SetupTestEventRegistry()
	defer closer()

	e := api.GetFakeEvent()
	_, err := er.CreateEvent(e)
	assert.NoError(t, err)
	objList, err := er.List()
	assert.NoError(t, err)
	filtered := filterReplyObject(objList)
	eList := filtered.(*api.EventList)
	assert.Len(t, eList.Items, 1)
	var lastTime api.Time
	for i := 0; i < MaxEventListSize+5; i++ {
		time.Sleep(10 * time.Millisecond)
		e := api.GetFakeEvent()
		_, err := er.CreateEvent(e)
		assert.NoError(t, err)
		e, _ = er.GetEvent(&e.InvolvedObject, e.Name)
		lastTime = e.CreationTimestamp
	}
	objList, err = er.List()
	assert.NoError(t, err)
	filtered = filterReplyObject(objList)
	eList = filtered.(*api.EventList)
	assert.Len(t, eList.Items, MaxEventListSize)
	assert.Equal(t, lastTime, eList.Items[MaxEventListSize-1].CreationTimestamp)
}

func TestEnsureRegionUnchanged(t *testing.T) {
	if !(*runFunctional) {
		t.Skip("skipping region change functional tests")
	}
	etcdClient, closer, err := etcd.SetupEmbeddedEtcdTest()
	assert.NoError(t, err)
	defer closer()
	err = ensureRegionUnchanged(etcdClient, "us-east-1")
	assert.NoError(t, err)
	err = ensureRegionUnchanged(etcdClient, "us-west-2")
	assert.Error(t, err)
	err = ensureRegionUnchanged(etcdClient, "us-east-1")
	assert.NoError(t, err)
}
