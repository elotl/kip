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
	"io/ioutil"
	"os"
	"sync"

	"k8s.io/klog"
)

func SetupEmbeddedEtcdTest() (*SimpleEtcd, func(), error) {
	wg := &sync.WaitGroup{}
	quit := make(chan struct{})
	dataDir, err := ioutil.TempDir(os.TempDir(), "etcdtest")
	if err != nil {
		return nil, func() {}, err
	}
	closer := func() {
		quit <- struct{}{}
		if err := os.RemoveAll(dataDir); err != nil {
			klog.Fatal("Error removing etcd data directory")
		}
	}
	db := EtcdServer{
		DataDir: dataDir,
	}
	err = db.Start(quit, wg)
	if err != nil {
		return nil, closer, err
	}
	return db.Client, closer, nil
}
