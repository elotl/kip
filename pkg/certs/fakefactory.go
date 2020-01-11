package certs

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/docker/libkv/store/mock"
)

func NewFake() (*CertificateFactory, error) {
	store, err := mock.New(nil, nil)
	if err != nil {
		return nil, err
	}
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	root, err := CreateRootCert(privateKey)
	if err != nil {
		return nil, err
	}

	certFactory := &CertificateFactory{
		Root:    *root,
		key:     *privateKey,
		kvstore: store,
	}
	return certFactory, nil
}
