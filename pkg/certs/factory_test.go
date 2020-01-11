package certs

import (
	"crypto/x509"
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestInitRootCert(t *testing.T) {
	name, closer := util.MakeTempFileName("milpatf")
	defer closer()
	kvstore := registry.CreateKVStore(name)
	fac, err := New(kvstore)
	assert.NoError(t, err)
	_, err = fac.CreateClientCert()
	assert.NoError(t, err)
}

func TestEncodeDecodeKeyAndCert(t *testing.T) {
	fac, err := NewFake()
	assert.NoError(t, err)
	cert, key, err := fac.CreateNodeCertAndKey()
	assert.NoError(t, err)
	certBits, err := MarshalCert(cert)
	assert.NoError(t, err)
	keyBits, err := MarshalKey(key)
	assert.NoError(t, err)
	certAgain, err := UnmarshalCert(certBits)
	assert.NoError(t, err)
	keyAgain, err := UnmarshalKey(keyBits)
	assert.NoError(t, err)
	assert.Equal(t, cert, certAgain)
	assert.Equal(t, key, keyAgain)
}

func TestVerifyCertWorks(t *testing.T) {
	fac, err := NewFake()
	assert.NoError(t, err)
	cert, _, err := fac.CreateNodeCertAndKey()
	roots := x509.NewCertPool()
	roots.AddCert(&fac.Root)
	opts := x509.VerifyOptions{
		Roots: roots,
	}
	_, err = cert.Verify(opts)
	assert.NoError(t, err, "failed to verify certificate")
}

func TestVerifyCertFails(t *testing.T) {
	fac, err := NewFake()
	assert.NoError(t, err)
	fac2, err := NewFake()
	assert.NoError(t, err)
	cert, _, err := fac2.CreateNodeCertAndKey()
	assert.NoError(t, err)
	roots := x509.NewCertPool()
	roots.AddCert(&fac.Root)
	opts := x509.VerifyOptions{
		Roots: roots,
	}
	_, err = cert.Verify(opts)
	assert.Error(t, err, "Certificate verification should fail")
}
