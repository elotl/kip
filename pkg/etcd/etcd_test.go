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

package etcd

import (
	"fmt"
	"testing"
	"time"

	"github.com/docker/libkv/store"
	"github.com/stretchr/testify/assert"
)

// We group a lot of ops here since I'd rather not start and stop a
// bunch of etcd processes over and over again during testing
func TestEtcdOperations(t *testing.T) {
	c, closer, err := SetupEmbeddedEtcdTest()
	if err != nil {
		fmt.Println(err)
		assert.FailNow(t, "Error setting up embedded etcd for testing")
	}
	defer closer()
	exists, err := c.Exists("foo")
	assert.NoError(t, err)
	assert.False(t, exists)
	_, err = c.Get("foo")
	assert.Error(t, err)
	assert.Equal(t, store.ErrKeyNotFound, err)
	fooVal := []byte("val1")
	err = c.Put("foo", fooVal, nil)
	assert.NoError(t, err)
	exists, err = c.Exists("foo")
	assert.NoError(t, err)
	assert.True(t, exists)
	p, err := c.Get("foo")
	assert.NoError(t, err)
	assert.Equal(t, fooVal, p.Value)
	err = c.Delete("foo")
	assert.NoError(t, err)
	exists, err = c.Exists("foo")
	assert.NoError(t, err)
	assert.False(t, exists)

	barVal := []byte("barval")
	_ = c.Put("foo/bar", barVal, nil)
	bazVal := []byte("bazval")
	_ = c.Put("/foo/baz", bazVal, nil)
	vals, err := c.List("foo/")
	assert.NoError(t, err)
	assert.Len(t, vals, 2)
	assert.Equal(t, barVal, vals[0].Value)
	assert.Equal(t, bazVal, vals[1].Value)

	// set with ttl takes too much time (like, 5s), don't run it..
	testTTL := false
	if testTTL {
		ttl := time.Second * 1
		err = c.Put("withttl", fooVal, &store.WriteOptions{TTL: ttl})
		assert.NoError(t, err)
		exists, err = c.Exists("withttl")
		assert.NoError(t, err)
		assert.True(t, exists)
		time.Sleep(5 * time.Second)
		_, err = c.Get("withttl")
		assert.Error(t, err)
		assert.Equal(t, store.ErrKeyNotFound, err)
	}

	// test transaction with atomic creates
	worked, _, err := c.AtomicPut("atom", []byte("atomVal"), nil, nil)
	assert.True(t, worked)
	assert.NoError(t, err)
	worked, _, err = c.AtomicPut("atom", []byte("UpdatedValue"), nil, nil)
	assert.False(t, worked)
	assert.Equal(t, store.ErrKeyModified, err)

	// test transaction with atomic but changed behind our backs
	p, err = c.Get("atom")
	assert.NoError(t, err)
	err = c.Put("atom", []byte("sneakyChangeValue"), nil)
	assert.NoError(t, err)
	worked, _, err = c.AtomicPut("atom", []byte("UpdatedValue"), p, nil)
	assert.False(t, worked)
	assert.Equal(t, store.ErrKeyModified, err)

	// Test transaction where it works
	p, err = c.Get("atom")
	assert.NoError(t, err)
	worked, _, err = c.AtomicPut("atom", []byte("UpdatedValue"), p, nil)
	assert.True(t, worked)
	assert.NoError(t, err)
	p, err = c.Get("atom")
	assert.NoError(t, err)
	assert.Equal(t, []byte("UpdatedValue"), p.Value)

	// Test DeleteRange
	// make a hirearchy of a few keys with keys around that hirearchy
	// and make sure we don't delete the keys we want to keep
	keepValues := []string{
		"milpa",
		"milpb/nested",
		"milpasomething",
		"nilpa",
		"elotl/",
		"elotl/nested",
		"/this/beginswithslash",
	}
	doomedValues := []string{
		"milpa/first/value",
		"milpa/first/valueagain",
		"milpa/first/.",
		"milpa/second/value",
	}
	allValues := append(keepValues, doomedValues...)
	for _, k := range allValues {
		err = c.Put(k, []byte(k), nil)
		assert.NoError(t, err)
	}
	c.DeleteTree("milpa/")
	for _, k := range keepValues {
		kv, err := c.Get(k)
		assert.NoError(t, err)
		assert.Equal(t, k, kv.Key)
		assert.Equal(t, []byte(k), kv.Value)
	}
	for _, k := range doomedValues {
		_, err := c.Get(k)
		assert.Error(t, err)
		assert.Equal(t, store.ErrKeyNotFound, err)
	}

}
