package certs

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"
)

type CertUsageType int

const (
	ClientCert CertUsageType = iota
	ServerCert
)

func createCertHelper(certTemplate, parentCert *x509.Certificate, pubKey, privKey interface{}) (*x509.Certificate, error) {
	certBytes, err := x509.CreateCertificate(
		rand.Reader, certTemplate, parentCert, pubKey, privKey)
	if err != nil {
		return nil, fmt.Errorf("Could not create certificate: %s", err)
	}

	newCert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		return nil, fmt.Errorf("Could not parse certificate: %s", err)
	}
	return newCert, nil
}

func CreateRootCert(privateKey *ecdsa.PrivateKey) (*x509.Certificate, error) {
	name := pkix.Name{CommonName: "Milpa Server CA"}
	certTemplate := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      name,
		NotBefore:    time.Now().Add(-600 * time.Second).UTC(),
		NotAfter:     time.Now().Add(time.Duration(10*24*365) * time.Hour).UTC(),
		IsCA:         true,
		BasicConstraintsValid: true,
		KeyUsage: x509.KeyUsageKeyEncipherment |
			x509.KeyUsageDataEncipherment |
			x509.KeyUsageDigitalSignature |
			x509.KeyUsageKeyAgreement |
			x509.KeyUsageCertSign,
		Issuer: name,
	}
	return createCertHelper(
		&certTemplate, &certTemplate, privateKey.Public(), privateKey)
}

func CreateCert(rootCert *x509.Certificate, rootPrivateKey *ecdsa.PrivateKey, serverPublicKey *ecdsa.PublicKey, certUsage CertUsageType) (*x509.Certificate, error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, fmt.Errorf("Error creating cert serial number: %s", err)
	}
	name := "MilpaServer"
	organization := "Elotl"
	commonName := "MilpaNode"
	extKeyUsage := []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}
	if certUsage == ClientCert {
		extKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}
	}
	certTemplate := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization:       []string{organization},
			OrganizationalUnit: []string{name},
			Country:            []string{"United States"},
			CommonName:         commonName, // used for client verification
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  extKeyUsage,
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	return createCertHelper(
		certTemplate, rootCert, serverPublicKey, rootPrivateKey)
}

func MarshalCert(cert *x509.Certificate) (text []byte, err error) {
	buf := &bytes.Buffer{}
	err = pem.Encode(
		buf,
		&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: cert.Raw,
		})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func MarshalKey(privateKey *ecdsa.PrivateKey) (text []byte, err error) {
	keyBits, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	err = pem.Encode(
		buf,
		&pem.Block{
			Type:  "EC PRIVATE KEY",
			Bytes: keyBits,
		})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func SaveCert(filename string, cert *x509.Certificate) error {
	certOut, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	b, err := MarshalCert(cert)
	if err != nil {
		return err
	}
	_, err = certOut.Write(b)
	if err != nil {
		return err
	}
	certOut.Close()
	return nil
}

func SaveKey(filename string, privateKey *ecdsa.PrivateKey) error {
	keyOut, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	b, err := MarshalKey(privateKey)
	if err != nil {
		return err
	}
	_, err = keyOut.Write(b)
	if err != nil {
		return err
	}
	keyOut.Close()
	return nil
}

func UnmarshalCert(b []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(b)
	if block == nil {
		return nil, fmt.Errorf("failed to decode text certificate")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: %v", err)
	}
	return cert, nil
}

func UnmarshalKey(b []byte) (*ecdsa.PrivateKey, error) {
	keyDERBlock, _ := pem.Decode(b)
	if keyDERBlock == nil {
		return nil, fmt.Errorf("failed to decode text private key")
	}
	key, err := x509.ParseECPrivateKey(keyDERBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}
	return key, nil
}

func LoadKey(filename string) (*ecdsa.PrivateKey, error) {
	pemContents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading file %s: %v", filename, err)
	}
	return UnmarshalKey(pemContents)
}

func LoadCert(filename string) (*x509.Certificate, error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading file %s: %v", filename, err)
	}
	return UnmarshalCert(contents)
}

func NewTLSConfig(cacert, cert, key, serverName string, isServer bool) (*tls.Config, error) {
	certificate, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return nil, fmt.Errorf("error loading key pair %s %s: %v",
			cert, key, err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(cacert)
	if err != nil {
		return nil, fmt.Errorf("error reading CA certificate %s: %s",
			cacert, err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, fmt.Errorf("error adding CA certificate to pool")
	}
	cfg := tls.Config{
		Certificates: []tls.Certificate{certificate},
	}
	if isServer {
		cfg.ClientAuth = tls.RequireAndVerifyClientCert
		cfg.ClientCAs = certPool
	} else {
		cfg.ServerName = serverName
		cfg.RootCAs = certPool
	}
	return &cfg, nil
}

func NewServerTLSConfig(cacert, cert, key string) (*tls.Config, error) {
	return NewTLSConfig(cacert, cert, key, "", true)
}

func NewClientTLSConfig(cacert, cert, key, serverName string) (*tls.Config, error) {
	return NewTLSConfig(cacert, cert, key, serverName, false)
}
