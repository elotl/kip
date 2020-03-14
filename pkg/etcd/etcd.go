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
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/compactor"
	"github.com/coreos/etcd/embed"
	"github.com/coreos/etcd/etcdserver/api/v3client"
	"github.com/coreos/etcd/pkg/transport"
	"github.com/docker/libkv/store"
	"github.com/elotl/kip/pkg/util"
	"golang.org/x/sys/unix"
	"k8s.io/klog"
)

var (
	DefaultEtcdDataDir = "/opt/kip/data"
	contextDeadline    = 10 * time.Second
)

type EtcdServer struct {
	ConfigFile string
	DataDir    string
	Proc       *embed.Etcd
	Client     *SimpleEtcd
}

func ensureEtcdDataDir(dataDir string) error {
	// if it exists, cool, if it doesn't exist, set it up
	errMsg := fmt.Sprintf("Could not create kip storage directory at %s, please verify the directory exists and is writable by kip. The error was", dataDir)
	_, err := os.Stat(dataDir)
	if os.IsNotExist(err) {
		klog.V(2).Infof("Creating kip data directory at %s", dataDir)
		err := os.MkdirAll(dataDir, 0750)
		if err != nil {
			return util.WrapError(err, errMsg)
		}
	} else if err != nil {
		return util.WrapError(err, errMsg)
	}

	err = unix.Access(dataDir, unix.W_OK)
	if err != nil {
		return util.WrapError(err, errMsg)
	}
	return nil
}

func (s *EtcdServer) reconcileDataDirectoryValues(cfg *embed.Config) error {
	if cfg.Dir == "" && s.DataDir == "" {
		s.DataDir = DefaultEtcdDataDir
	}
	if cfg.Dir == "" && s.DataDir != "" {
		cfg.Dir = s.DataDir
	}
	if s.DataDir == "" && cfg.Dir != "" {
		s.DataDir = cfg.Dir
	}
	if s.DataDir != cfg.Dir {
		msg := fmt.Sprintf(`Two different values have been specified for the etcd data directory:
  provider.yaml etcd.dataDir value: %s
  etcd.configFile.data-dir value: %s`, s.DataDir, cfg.Dir)
		return fmt.Errorf(msg)
	}
	return nil
}

func (s *EtcdServer) Start(quit <-chan struct{}, wg *sync.WaitGroup) error {
	var cfg *embed.Config
	var err error
	if s.ConfigFile != "" {
		cfg, err = embed.ConfigFromFile(s.ConfigFile)
		if err != nil {
			return util.WrapError(err, "Error creating etcd configuration")
		}
	} else {
		cfg = embed.NewConfig()
		cfg.LPUrls = []url.URL{}
		cfg.LCUrls = []url.URL{}
	}
	if cfg.AutoCompactionMode == "" {
		klog.V(2).Info("Setting etcd compaction mode to periodic")
		cfg.AutoCompactionMode = compactor.ModePeriodic
	}
	if cfg.AutoCompactionMode == compactor.ModePeriodic &&
		cfg.AutoCompactionRetention == "" {
		cfg.AutoCompactionRetention = "1"
		klog.V(2).Info("Setting etcd compaction interval to 1 hour")
	}

	err = s.reconcileDataDirectoryValues(cfg)
	if err != nil {
		return err
	}
	err = ensureEtcdDataDir(cfg.Dir)
	if err != nil {
		return err
	}

	s.Proc, err = embed.StartEtcd(cfg)
	if err != nil {
		return util.WrapError(err, "Error starting etcd")
	}
	select {
	case <-s.Proc.Server.ReadyNotify():
		klog.V(2).Info("Etcd server is ready to serve requests")
	case <-time.After(60 * time.Second):
		s.Proc.Server.Stop()
		s.Proc.Close()
		return fmt.Errorf("Etcd took too long to start!")
	}

	apiClient := v3client.New(s.Proc.Server)
	s.Client = &SimpleEtcd{
		client:   apiClient,
		external: false,
	}
	wg.Add(1)
	go func() {
		<-quit
		// if we don't pause, clients will crash, it's a bad look.
		pause := 2 * time.Second
		klog.V(2).Infof("Pausing for %ds before shutting down etcd...", int(pause.Seconds()))
		time.Sleep(pause)
		s.Proc.Server.Stop()
		s.Proc.Close()
		wg.Done()
	}()
	return nil
}

func NewEtcdClient(endpoints []string, certFile, keyFile, caFile string) (*SimpleEtcd, error) {
	tlsInfo := transport.TLSInfo{
		CertFile: certFile,
		KeyFile:  keyFile,
		CAFile:   caFile,
	}
	tlsConfig, err := tlsInfo.ClientConfig()
	if err != nil {
		return nil, util.WrapError(
			err, "Error creating TLS configuration for etcd")
	}
	if len(certFile) == 0 && len(keyFile) == 0 && len(caFile) == 0 {
		tlsConfig = nil
	}
	cfg := clientv3.Config{
		Endpoints: endpoints,
		TLS:       tlsConfig,
	}
	client, err := clientv3.New(cfg)
	if err != nil {
		return nil, util.WrapError(err, "Error creating new etcd client")
	}
	etcd := &SimpleEtcd{
		client:   client,
		external: true,
	}
	return etcd, nil
}

type SimpleEtcd struct {
	client   *clientv3.Client
	external bool
}

func (s *SimpleEtcd) Client() *clientv3.Client {
	return s.client
}

func (s *SimpleEtcd) External() bool {
	return s.external
}

// Normalize the key for usage in Etcd
func normalize(key string) string {
	return strings.TrimPrefix(key, "/")
}

// I'm pretty sure we can get away with serializable calls here
// K8s allows it but it's off by default.  Might want to make
// this a config flag (k8s flag: etcd-quorum-read)
func (s *SimpleEtcd) Get(key string) (pair *store.KVPair, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), contextDeadline)
	defer cancel()
	result, err := s.client.Get(ctx, normalize(key), clientv3.WithSerializable())
	if err != nil {
		return nil, err
	}
	if len(result.Kvs) == 0 {
		return nil, store.ErrKeyNotFound
	}

	pair = &store.KVPair{
		Key:       key,
		Value:     result.Kvs[0].Value,
		LastIndex: uint64(result.Kvs[0].Version),
	}

	return pair, nil
}

func (s *SimpleEtcd) Put(key string, value []byte, opts *store.WriteOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), contextDeadline)
	defer cancel()
	return s.putHelper(ctx, key, value, opts)
}

// We have one call in server.go that will block until we can write to an etcd
// cluster so we'll call putHelper with a background context.
func (s *SimpleEtcd) PutNoTimeout(key string, value []byte, opts *store.WriteOptions) error {
	return s.putHelper(context.Background(), key, value, opts)
}

func (s *SimpleEtcd) putHelper(ctx context.Context, key string, value []byte, opts *store.WriteOptions) error {
	var err error
	o := []clientv3.OpOption{}
	if opts != nil && opts.TTL > 0 {
		ctxLease, cancel := context.WithCancel(ctx)
		lease, err := s.client.Grant(ctxLease, int64(opts.TTL.Seconds()))
		cancel()
		if err != nil {
			return err
		}
		o = []clientv3.OpOption{clientv3.WithLease(lease.ID)}
	}
	_, err = s.client.Put(ctx, normalize(key), string(value), o...)
	return err
}

func (s *SimpleEtcd) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), contextDeadline)
	defer cancel()
	resp, err := s.client.Delete(ctx, normalize(key))
	if err != nil {
		return err
	}
	if resp.Deleted == 0 {
		return store.ErrKeyNotFound
	}
	return err
}

func (s *SimpleEtcd) Exists(key string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), contextDeadline)
	defer cancel()
	result, err := s.client.Get(ctx, normalize(key))
	if err != nil {
		return false, err
	}
	if len(result.Kvs) == 0 {
		return false, nil
	}
	return true, nil
}

// I'd like to get away with serializable here.  We typically write to the
// same store that we're reading from.  If this becomes a problem we can
// remove the WithSerializable()
func (s *SimpleEtcd) List(directory string) ([]*store.KVPair, error) {
	ctx, cancel := context.WithTimeout(context.Background(), contextDeadline)
	defer cancel()
	result, err := s.client.Get(
		ctx, normalize(directory), clientv3.WithPrefix(), clientv3.WithSerializable())
	if err != nil {
		return nil, err
	}
	if len(result.Kvs) == 0 {
		return nil, store.ErrKeyNotFound
	}
	kv := make([]*store.KVPair, result.Count)
	for i := int64(0); i < result.Count; i++ {
		kv[i] = &store.KVPair{
			Key:       string(result.Kvs[i].Key),
			Value:     result.Kvs[i].Value,
			LastIndex: uint64(result.Kvs[i].Version),
		}
	}

	return kv, nil
}

func (s *SimpleEtcd) AtomicPut(key string, value []byte, previous *store.KVPair, opts *store.WriteOptions) (bool, *store.KVPair, error) {
	o := []clientv3.OpOption{}
	if opts != nil && opts.TTL > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), contextDeadline)
		defer cancel()
		lease, err := s.client.Grant(ctx, int64(opts.TTL.Seconds()))
		cancel()
		if err != nil {
			return false, nil, err
		}
		o = []clientv3.OpOption{clientv3.WithLease(lease.ID)}
	}
	var txn clientv3.Txn
	ctx, cancel := context.WithTimeout(context.Background(), contextDeadline)
	defer cancel()
	if previous != nil {
		txn = s.client.Txn(ctx).If(
			clientv3.Compare(clientv3.Version(key), "=", int64(previous.LastIndex)))
	} else {
		txn = s.client.Txn(ctx).If(
			clientv3.Compare(clientv3.Version(key), "=", 0))
	}
	txn = txn.Then(clientv3.OpPut(key, string(value), o...))
	putresp, err := txn.Commit()
	if err != nil {
		if putresp != nil {
			return putresp.Succeeded, nil, err
		} else {
			return false, nil, err
		}
	} else if !putresp.Succeeded {
		return false, nil, store.ErrKeyModified
	}

	// Etcd gives us the old version of the key, just don't return anything
	return putresp.Succeeded, nil, err
}

func (s *SimpleEtcd) DeleteTree(keyPrefix string) error {
	ctx, cancel := context.WithTimeout(context.Background(), contextDeadline)
	defer cancel()
	resp, err := s.client.Delete(ctx, normalize(keyPrefix), clientv3.WithPrefix())
	if err != nil {
		return err
	}
	if resp.Deleted == 0 {
		return store.ErrKeyNotFound
	}
	return err
}
