package main

import (
	"fmt"
	"net"

	certutil "k8s.io/client-go/util/cert"
	"k8s.io/client-go/util/keyutil"
	"k8s.io/klog"
)

func ensureCert(hostName, certFile, keyFile string, ips []net.IP) error {
	ok, err := certutil.CanReadCertAndKey(certFile, keyFile)
	if ok {
		klog.V(2).Infof("found server cert %q and key %q", certFile, keyFile)
	}
	if err != nil {
		klog.Warningf(
			"verifying server cert %q and key %q: %v", certFile, keyFile, err)
	}
	cert, key, err := certutil.GenerateSelfSignedCertKey(hostName, ips, nil)
	if err != nil {
		return fmt.Errorf("unable to generate self signed cert: %v", err)
	}
	if err := certutil.WriteCert(certFile, cert); err != nil {
		return err
	}
	if err := keyutil.WriteKey(keyFile, key); err != nil {
		return err
	}
	klog.V(2).Infof("using self-signed cert %q %q", certFile, keyFile)
	return nil
}
