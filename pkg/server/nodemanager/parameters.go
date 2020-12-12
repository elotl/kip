package nodemanager

import (
	"regexp"

	"github.com/elotl/kip/pkg/certs"
	"github.com/elotl/kip/pkg/util"
	"gopkg.in/yaml.v2"
)

const (
	semverRegexFmt string = `v?([0-9]+)(\.[0-9]+)(\.[0-9]+)?` +
		`(-([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?` +
		`(\+([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?`
	ParameterCACertificate     = "ca.crt"
	ParameterServerCertificate = "server.crt"
	ParameterServerKey         = "server.key"
	ParameterItzoVersion       = "itzo_version"
	ParameterItzoURL           = "itzo_url"
	ParameterCellConfig        = "cell_config.yaml"
)

var (
	semverRegex = regexp.MustCompile("^" + semverRegexFmt + "$")
)

type InstanceConfig struct {
	ItzoVersion string
	ItzoURL     string
	CellConfig  map[string]string
}

func createNodeCertAndKey(certificateFactory *certs.CertificateFactory) (map[string]string, error) {
	params := make(map[string]string)

	cert, key, err := certificateFactory.CreateNodeCertAndKey()
	if err != nil {
		return nil, util.WrapError(err, "creating node cert")
	}

	certBytes, err := certs.MarshalCert(cert)
	if err != nil {
		return nil, util.WrapError(err, "serializing node cert")
	}
	params[ParameterServerCertificate] = string(certBytes)

	keyBytes, err := certs.MarshalKey(key)
	if err != nil {
		return nil, util.WrapError(err, "serializing node key")
	}
	params[ParameterServerKey] = string(keyBytes)

	rootCertBytes, err := certs.MarshalCert(&certificateFactory.Root)
	if err != nil {
		return nil, util.WrapError(err, "serializing root cert")
	}
	params[ParameterCACertificate] = string(rootCertBytes)

	return params, nil
}

func getInstanceParameters(certificateFactory *certs.CertificateFactory, config InstanceConfig) (map[string]string, error) {
	params, err := createNodeCertAndKey(certificateFactory)
	if err != nil {
		return nil, err
	}

	version := config.ItzoVersion
	if len(version) > 0 {
		if version != "latest" &&
			version[0] != 'v' &&
			semverRegex.MatchString(version) {
			version = "v" + version
		}
		params[ParameterItzoVersion] = version
	}

	url := config.ItzoURL
	if len(url) > 0 {
		params[ParameterItzoURL] = url
	}

	cellconf := config.CellConfig
	if len(cellconf) > 0 {
		buf, err := yaml.Marshal(cellconf)
		if err != nil {
			return nil, util.WrapError(err, "serializing cell_config")
		}
		params[ParameterCellConfig] = string(buf)
	}

	return params, nil
}

func getMarshalledInstanceParameters(certificateFactory *certs.CertificateFactory, config InstanceConfig) (string, error) {
	params, err := getInstanceParameters(certificateFactory, config)
	if err != nil {
		return "", util.WrapError(err, "getInstanceParameters()")
	}
	buf, err := yaml.Marshal(params)
	if err != nil {
		return "", util.WrapError(err, "marshalling instance parameters")
	}
	return string(buf), nil
}
