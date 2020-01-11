package etcd

import (
	"io/ioutil"
	"os"
	"sync"

	"github.com/golang/glog"
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
			glog.Fatal("Error removing etcd data directory")
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
