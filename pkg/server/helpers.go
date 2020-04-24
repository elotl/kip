package server

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"

	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/util"
	"github.com/elotl/kip/pkg/util/k8s"
	"github.com/elotl/kip/pkg/util/k8s/eventrecorder"
	"github.com/kubernetes/kubernetes/pkg/kubelet/network/dns"
	"github.com/virtual-kubelet/node-cli/manager"
	v1 "k8s.io/api/core/v1"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/klog"
)

func createResolverFile(nameservers, searches []string) (string, error) {
	tmpf, err := ioutil.TempFile("", "resolv-conf")
	if err != nil {
		klog.Warningf("creating resolver tempfile: %v", err)
		return "", err
	}
	defer tmpf.Close()
	for _, ns := range nameservers {
		tmpf.Write([]byte(fmt.Sprintf("nameserver %s\n", ns)))
	}
	if len(searches) > 0 {
		searchList := strings.Join(searches, " ")
		tmpf.Write([]byte(fmt.Sprintf("search %s\n", searchList)))
	}
	resolverConfig := tmpf.Name()
	return resolverConfig, nil
}

func createDNSConfigurer(kubernetesNodeName, clusterDNS, clusterDomain string, cloudClient cloud.CloudClient, rm *manager.ResourceManager) (*dns.Configurer, error) {
	loggingEventRecorder := eventrecorder.NewLoggingEventRecorder(4)
	nodeRef := &v1.ObjectReference{
		Kind:       "Node",
		APIVersion: "v1",
		Name:       kubernetesNodeName,
	}
	nameservers, searches, err := cloudClient.GetDNSInfo()
	if err != nil {
		return nil, util.WrapError(err, "getting cloud DNS info")
	}
	klog.V(2).Infof("host nameservers %v searches %v", nameservers, searches)
	resolverConfig, err := createResolverFile(nameservers, searches)
	ip := net.ParseIP(clusterDNS)
	if ip == nil || ip.IsUnspecified() {
		services, err := rm.ListServices()
		if err != nil {
			return nil, util.WrapError(err, "looking up kube-dns service")
		}
		for _, svc := range services {
			if svc.Name != "kube-dns" || svc.Namespace != "kube-system" {
				continue
			}
			ip = net.ParseIP(svc.Spec.ClusterIP)
		}
	}
	if ip == nil || ip.IsUnspecified() {
		return nil, fmt.Errorf("missing or misconfigured kube-dns service")
	}
	return dns.NewConfigurer(
		loggingEventRecorder,
		nodeRef,
		nil,
		[]net.IP{ip},
		clusterDomain,
		resolverConfig,
	), nil
}

func createNetworkAgentKubeconfig(kubernetesNodeName, networkAgentSecret, serverURL string, rm *manager.ResourceManager) (*clientcmdapi.Config, error) {
	kc, err := k8s.CreateNetworkAgentKubeconfig(
		rm, serverURL, networkAgentSecret)
	if err != nil {
		return nil, util.WrapError(err, "creating network-agent kubeconfig")
	}
	if err := k8s.ValidateKubeconfig(kc); err != nil {
		return nil, util.WrapError(err, "validating network-agent kubeconfig")
	}
	return kc, err
}
