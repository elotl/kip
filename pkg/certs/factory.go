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

package certs

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"

	"github.com/docker/libkv/store"
	"github.com/elotl/kip/pkg/etcd"
	"github.com/elotl/kip/pkg/util"
	"k8s.io/klog"
)

const (
	CertificatePath                 string = "/milpa/certificate"
	CertificateDirectoryPlaceholder string = "/milpa/certificate/."
	RootCertPath                    string = "/milpa/certificate/root.crt"
	RootKeyPath                     string = "/milpa/certificate/root.key"
)

type CertificateFactory struct {
	Root    x509.Certificate
	key     ecdsa.PrivateKey
	kvstore etcd.Storer
}

func New(kvstore etcd.Storer) (*CertificateFactory, error) {
	certFactory := &CertificateFactory{
		kvstore: kvstore,
	}
	certFactory.kvstore.Put(CertificateDirectoryPlaceholder, []byte("."), nil)
	err := certFactory.GetRootFromStore()
	if err == store.ErrKeyNotFound {
		klog.V(2).Infof("Initializing Milpa root certificate")
		certFactory.InitRootCert()
	} else if err != nil {
		return nil, util.WrapError(err, "Error creating Milpa PKI")
	}
	return certFactory, nil
}

func (fac *CertificateFactory) GetRootFromStore() error {
	certPair, err := fac.kvstore.Get(RootCertPath)
	if err != nil {
		return err
	}
	cert, err := UnmarshalCert(certPair.Value)
	if err != nil {
		return util.WrapError(err, "Error unmarshalling root cert")
	}
	fac.Root = *cert

	keyPair, err := fac.kvstore.Get(RootKeyPath)
	if err != nil {
		return err
	}
	key, err := UnmarshalKey(keyPair.Value)
	if err != nil {
		return util.WrapError(err, "Error unmarshalling root key cert")
	}
	fac.key = *key
	return nil
}

func (fac *CertificateFactory) InitRootCert() error {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return util.WrapError(err, "Error creating root key")
	}
	keyData, err := MarshalKey(privateKey)
	if err != nil {
		return util.WrapError(err, "Error serializing root key")
	}
	err = fac.kvstore.Put(RootKeyPath, keyData, nil)
	if err != nil {
		return util.WrapError(err, "Error writing root key to storage")
	}
	rootCert, err := CreateRootCert(privateKey)
	if err != nil {
		return util.WrapError(err, "Error creating root certificate")
	}
	certData, err := MarshalCert(rootCert)
	if err != nil {
		return util.WrapError(err, "Error serializing root cert")
	}
	err = fac.kvstore.Put(RootCertPath, certData, nil)
	if err != nil {
		return util.WrapError(err, "Error writing root cert to storage")
	}

	fac.Root = *rootCert
	fac.key = *privateKey

	return nil
}

func (fac *CertificateFactory) CreateNodeCertAndKey() (*x509.Certificate, *ecdsa.PrivateKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, util.WrapError(err, "Error generating node key")
	}
	serverCert, err := CreateCert(&fac.Root, &fac.key, &privateKey.PublicKey, ServerCert)
	if err != nil {
		return nil, nil, util.WrapError(err, "Error generating node certificate")
	}
	return serverCert, privateKey, nil
}

func (fac *CertificateFactory) CreateClientCert() (*tls.Certificate, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, util.WrapError(err, "Error generating client key")
	}
	cert, err := CreateCert(&fac.Root, &fac.key, &privateKey.PublicKey, ClientCert)
	if err != nil {
		return nil, util.WrapError(err, "Error generating node certificate")
	}

	var tlsCert tls.Certificate
	tlsCert.Certificate = append(tlsCert.Certificate, cert.Raw)
	tlsCert.PrivateKey = privateKey
	return &tlsCert, nil
}
