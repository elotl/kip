## Networking

Kip allocates two IP addresses for each cell: one for management communication (between the provider and a small agent running on the instance), and one for the pod. Unless the pod has hostNetwork enabled, a new Linux network namespace is created for the pod with the second IP. Both IP addresses come from the VPC address space — fortunately, even the tiniest cloud instances are allowed to allocate at least two IP addresses. This design ensures that the pod can’t interfere with management communications.

As for network interoperability between regular pods and virtual-kubelet pods, we recommend the native CNI plugin that integrates with the cloud provider VPC, i.e. the aws-vpc-cni plugin on AWS. That way both virtual-kubelet pods and regular pods will get their IP addresses from the VPC address space, and the VPC network will take care of routing.

If you would like to use another CNI plugin for some reason, that will also work as long as the cloud controller is configured to create cloud routes with the PodCIDR allocated to nodes and the CNI plugin used in the cluster is able to use the PodCIDR (most CNI plugins can do this).
Currently, Kip needs to run in host network mode. Since NodePorts are managed by the service proxy running on Kubernetes nodes, they also work seamlessly. Iptables rules for HostPort mappings are created and maintained by Kip.

The cloud instances hosting pods created via Kip also run a combined service proxy and network policy agent, which is kube-router. BGP is disabled in kube-router, since the cloud routes take care of routing to the CIDRs used for regular pods on the nodes; only service proxying and firewalling are enabled.

This way, pods can reach Kubernetes services either via IPVS or iptables NAT rules, created by kube-router, and network policies are also enforced (via iptables filtering rules, also maintained by kube-router). Essentially, this is the same setup as on Kubernetes nodes. Kube-router has a very low memory and CPU footprint, thus even small cloud instances can dedicate almost all their resources to the application(s) running in the pod.
