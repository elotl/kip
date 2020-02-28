/*
Copyright 2016 The Kubernetes Authors.
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

package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/kubernetes/kubernetes/pkg/kubelet/network/dns"
	utilvalidation "k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/klog"
)

const (
	etcHostsPath                      = "/etc/hosts"
	managedHostsHeader                = "# Kubernetes-managed hosts file.\n"
	managedHostsHeaderWithHostNetwork = "# Kubernetes-managed hosts file (host network).\n"
)

// CreateEtcHosts creates an /etc/hosts file for the pod.
func CreateEtcHosts(dnsConfigurer *dns.Configurer, podName, namespace, specHostname, specSubdomain string, podIPs []string, hostAliases []api.HostAlias, useHostNetwork bool) ([]byte, error) {
	var hostsFileContent []byte
	var err error

	hostName, hostDomainName, err := generatePodHostNameAndDomain(
		dnsConfigurer, podName, namespace, specHostname, specSubdomain)
	if err != nil {
		return nil, err
	}

	if useHostNetwork {
		// if Pod is using host network, read hosts file from the node's filesystem.
		// `etcHostsPath` references the location of the hosts file on the node.
		// `/etc/hosts` for *nix systems.
		hostsFileContent, err = nodeHostsFileContent(etcHostsPath, hostAliases)
		if err != nil {
			return nil, err
		}
	} else {
		// if Pod is not using host network, create a managed hosts file with Pod IP and other information.
		hostsFileContent = managedHostsFileContent(podIPs, hostName, hostDomainName, hostAliases)
	}

	return hostsFileContent, nil
}

// nodeHostsFileContent reads the content of node's hosts file.
func nodeHostsFileContent(hostsFilePath string, hostAliases []api.HostAlias) ([]byte, error) {
	hostsFileContent, err := ioutil.ReadFile(hostsFilePath)
	if err != nil {
		return nil, err
	}
	var buffer bytes.Buffer
	buffer.WriteString(managedHostsHeaderWithHostNetwork)
	buffer.Write(hostsFileContent)
	buffer.Write(hostsEntriesFromHostAliases(hostAliases))
	return buffer.Bytes(), nil
}

// managedHostsFileContent generates the content of the managed etc hosts based on Pod IPs and other
// information.
func managedHostsFileContent(hostIPs []string, hostName, hostDomainName string, hostAliases []api.HostAlias) []byte {
	var buffer bytes.Buffer
	buffer.WriteString(managedHostsHeader)
	buffer.WriteString("127.0.0.1\tlocalhost\n")                      // ipv4 localhost
	buffer.WriteString("::1\tlocalhost ip6-localhost ip6-loopback\n") // ipv6 localhost
	buffer.WriteString("fe00::0\tip6-localnet\n")
	buffer.WriteString("fe00::0\tip6-mcastprefix\n")
	buffer.WriteString("fe00::1\tip6-allnodes\n")
	buffer.WriteString("fe00::2\tip6-allrouters\n")
	if len(hostDomainName) > 0 {
		// host entry generated for all IPs in podIPs
		// podIPs field is populated for clusters even
		// dual-stack feature flag is not enabled.
		for _, hostIP := range hostIPs {
			buffer.WriteString(fmt.Sprintf("%s\t%s.%s\t%s\n", hostIP, hostName, hostDomainName, hostName))
		}
	} else {
		for _, hostIP := range hostIPs {
			buffer.WriteString(fmt.Sprintf("%s\t%s\n", hostIP, hostName))
		}
	}
	buffer.Write(hostsEntriesFromHostAliases(hostAliases))
	return buffer.Bytes()
}

func hostsEntriesFromHostAliases(hostAliases []api.HostAlias) []byte {
	if len(hostAliases) == 0 {
		return []byte{}
	}

	var buffer bytes.Buffer
	buffer.WriteString("\n")
	buffer.WriteString("# Entries added by HostAliases.\n")
	// for each IP, write all aliases onto single line in hosts file
	for _, hostAlias := range hostAliases {
		buffer.WriteString(fmt.Sprintf("%s\t%s\n", hostAlias.IP, strings.Join(hostAlias.Hostnames, "\t")))
	}
	return buffer.Bytes()
}

// truncatePodHostnameIfNeeded truncates the pod hostname if it's longer than 63 chars.
func truncatePodHostnameIfNeeded(podName, hostname string) (string, error) {
	// Cap hostname at 63 chars (specification is 64bytes which is 63 chars and the null terminating char).
	const hostnameMaxLen = 63
	if len(hostname) <= hostnameMaxLen {
		return hostname, nil
	}
	truncated := hostname[:hostnameMaxLen]
	klog.Errorf("hostname for pod:%q was longer than %d. Truncated hostname to :%q", podName, hostnameMaxLen, truncated)
	// hostname should not end with '-' or '.'
	truncated = strings.TrimRight(truncated, "-.")
	if len(truncated) == 0 {
		// This should never happen.
		return "", fmt.Errorf("hostname for pod %q was invalid: %q", podName, hostname)
	}
	return truncated, nil
}

// generatePodHostNameAndDomain creates a hostname and domain name for a pod,
// given that pod's spec and annotations or returns an error.
func generatePodHostNameAndDomain(dnsConfigurer *dns.Configurer, podName, namespace, specHostname, specSubdomain string) (string, string, error) {
	clusterDomain := dnsConfigurer.ClusterDomain

	hostname := podName
	if len(specHostname) > 0 {
		if msgs := utilvalidation.IsDNS1123Label(specHostname); len(msgs) != 0 {
			return "", "", fmt.Errorf("pod Hostname %q is not a valid DNS label: %s", specHostname, strings.Join(msgs, ";"))
		}
		hostname = specHostname
	}

	hostname, err := truncatePodHostnameIfNeeded(podName, hostname)
	if err != nil {
		return "", "", err
	}

	hostDomain := ""
	if len(specSubdomain) > 0 {
		if msgs := utilvalidation.IsDNS1123Label(specSubdomain); len(msgs) != 0 {
			return "", "", fmt.Errorf("pod Subdomain %q is not a valid DNS label: %s", specSubdomain, strings.Join(msgs, ";"))
		}
		hostDomain = fmt.Sprintf("%s.%s.svc.%s", specSubdomain, namespace, clusterDomain)
	}

	return hostname, hostDomain, nil
}

func GeneratePodHostname(dnsConfigurer *dns.Configurer, podName, namespace, specHostname, specSubdomain string) (string, error) {
	hostname, domainname, err := generatePodHostNameAndDomain(
		dnsConfigurer, podName, namespace, specHostname, specSubdomain)
	if err != nil {
		return "", err
	}
	if domainname != "" {
		return fmt.Sprintf("%s.%s", hostname, domainname), nil
	}
	return hostname, nil
}
