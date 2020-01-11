package etcd

import "github.com/docker/libkv/store"

type Storer interface {
	Put(key string, value []byte, options *store.WriteOptions) error

	// Get a value given its key
	Get(key string) (*store.KVPair, error)

	// Delete the value at the specified key
	Delete(key string) error

	// Verify if a Key exists in the store
	Exists(key string) (bool, error)

	// List the content of a given prefix
	List(directory string) ([]*store.KVPair, error)

	// Atomic CAS operation on a single value.
	// Pass previous = nil to create a new key.
	AtomicPut(key string, value []byte, previous *store.KVPair, options *store.WriteOptions) (bool, *store.KVPair, error)
}
