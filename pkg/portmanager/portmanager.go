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

package portmanager

import (
	"fmt"
	"io"
	"net"
	"strings"

	"k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/util/iptables"
)

const (
	hostPortsDnatChain  iptables.Chain = "VK_HOSTPORT_DNAT"
	hostPortsMasqChain  iptables.Chain = "VK_HOSTPORT_MASQ"
	hostPortsDnatPrefix                = "VK_HP_DNT_"
	hostPortsMasqPrefix                = "VK_HP_MSQ_"
)

type PortManager struct {
	iptables   iptables.Interface
	listeners  map[string][]io.Closer
	portOpener func(port int32, proto string) (io.Closer, error)
}

func openPort(port int32, proto string) (io.Closer, error) {
	switch proto {
	case "tcp":
		s, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			return nil, err
		}
		return s, nil
	case "udp":
		addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
		if err != nil {
			return nil, err
		}
		s, err := net.ListenUDP("udp", addr)
		if err != nil {
			return nil, err
		}
		return s, nil
	default:
		return nil, fmt.Errorf("unknown protocol %q", proto)
	}
}

func NewPortManager(ipt iptables.Interface) *PortManager {
	return &PortManager{
		iptables:   ipt,
		listeners:  make(map[string][]io.Closer),
		portOpener: openPort,
	}
}

func createSuffix(ip string) string {
	return strings.Replace(ip, ".", "-", -1)
}

func podDnatChain(ip string) iptables.Chain {
	return iptables.Chain(hostPortsDnatPrefix + createSuffix(ip))
}

func podMasqChain(ip string) iptables.Chain {
	return iptables.Chain(hostPortsMasqPrefix + createSuffix(ip))
}

func (pm *PortManager) ensureHostPortChains(podIp string) error {
	// DNAT.
	_, err := pm.iptables.EnsureChain(iptables.TableNAT, hostPortsDnatChain)
	if err != nil {
		return err
	}
	_, err = pm.iptables.EnsureChain(iptables.TableNAT, podDnatChain(podIp))
	if err != nil {
		return err
	}
	_, err = pm.iptables.EnsureRule(iptables.Prepend, iptables.TableNAT, iptables.ChainPrerouting, "-j", string(hostPortsDnatChain))
	if err != nil {
		return err
	}
	_, err = pm.iptables.EnsureRule(iptables.Append, iptables.TableNAT, hostPortsDnatChain, "-j", string(podDnatChain(podIp)))
	if err != nil {
		return err
	}
	// MASQ.
	_, err = pm.iptables.EnsureChain(iptables.TableNAT, hostPortsMasqChain)
	if err != nil {
		return err
	}
	_, err = pm.iptables.EnsureChain(iptables.TableNAT, podMasqChain(podIp))
	if err != nil {
		return err
	}
	_, err = pm.iptables.EnsureRule(iptables.Prepend, iptables.TableNAT, iptables.ChainPostrouting, "-j", string(hostPortsMasqChain))
	if err != nil {
		return err
	}
	_, err = pm.iptables.EnsureRule(iptables.Append, iptables.TableNAT, hostPortsMasqChain, "-j", string(podMasqChain(podIp)))
	if err != nil {
		return err
	}
	return nil
}

func getProtoAndPorts(cp v1.ContainerPort) (string, string, string) {
	proto := strings.ToLower(string(cp.Protocol))
	srcport := fmt.Sprintf("%d", cp.HostPort)
	dstport := fmt.Sprintf("%d", cp.ContainerPort)
	return proto, srcport, dstport
}

func filterPortMappings(cps []v1.ContainerPort) (ret []v1.ContainerPort) {
	for _, cp := range cps {
		if cp.HostPort > 0 && cp.ContainerPort > 0 {
			ret = append(ret, cp)
		}
	}
	return ret
}

func (pm *PortManager) AddPodPortMappings(podIp string, cps []v1.ContainerPort) error {
	err := pm.ensureHostPortChains(podIp)
	if err != nil {
		return err
	}
	cps = filterPortMappings(cps)
	// Open ports locally. This serves as a safety mechanism to ensure only one
	// pod gets a specific port+protocol pair.
	listeners := make([]io.Closer, len(cps))
	for i, mapping := range cps {
		hp := mapping.HostPort
		proto := strings.ToLower(string(mapping.Protocol))
		listener, err := pm.portOpener(hp, proto)
		if err != nil {
			for _, l := range listeners[:i] {
				l.Close()
			}
			return err
		}
		listeners[i] = listener
	}
	pm.listeners[podIp] = listeners
	// Add redirect rules.
	for _, mapping := range cps {
		proto, srcport, dstport := getProtoAndPorts(mapping)
		err = pm.addHostPort(podIp, proto, srcport, dstport)
		if err != nil {
			pm.RemovePodPortMappings(podIp)
			return err
		}
	}
	return nil
}

func (pm *PortManager) RemovePodPortMappings(podIp string) {
	for _, listener := range pm.listeners[podIp] {
		if listener != nil {
			listener.Close()
		}
	}
	_ = pm.ensureHostPortChains(podIp)
	dnat := podDnatChain(podIp)
	masq := podMasqChain(podIp)
	_ = pm.iptables.DeleteRule(iptables.TableNAT, hostPortsDnatChain, "-j", string(dnat))
	_ = pm.iptables.FlushChain(iptables.TableNAT, dnat)
	_ = pm.iptables.DeleteChain(iptables.TableNAT, dnat)
	_ = pm.iptables.DeleteRule(iptables.TableNAT, hostPortsMasqChain, "-j", string(masq))
	_ = pm.iptables.FlushChain(iptables.TableNAT, masq)
	_ = pm.iptables.DeleteChain(iptables.TableNAT, masq)
}

func (pm *PortManager) addHostPort(ip, proto, srcport, dstport string) error {
	proto = strings.ToLower(proto)
	// Redirect traffic hitting the host port.
	dst := fmt.Sprintf("%s:%s", ip, dstport)
	_, err := pm.iptables.EnsureRule(iptables.Append, iptables.TableNAT, podDnatChain(ip), "-p", proto, "--dport", srcport, "-m", "addrtype", "--dst-type", "LOCAL", "-j", "DNAT", "--to", dst)
	if err != nil {
		return err
	}
	// Masquerade traffic going to the destination of the porthost mapping.
	// Otherwise, the destination would send reply packets directly back to the
	// source (instead of replies coming back from this node).
	_, err = pm.iptables.EnsureRule(iptables.Append, iptables.TableNAT, podMasqChain(ip), "-p", proto, "--dport", dstport, "--dst", ip, "-j", "MASQUERADE")
	if err != nil {
		return err
	}
	return nil
}
