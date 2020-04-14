## Networking

Kip allocates two IP addresses for each cell: one for management communication (between the provider and a small agent running on the instance), and one for the pod. Unless the pod has hostNetwork enabled, a new Linux network namespace is created for the pod with the second IP. Both IP addresses come from the VPC address space — fortunately, even the tiniest cloud instances are allowed to allocate at least two IP addresses. This design ensures that the pod can’t interfere with management communications.

As for network interoperability between regular pods and virtual-kubelet pods, we recommend the native CNI plugin that integrates with the cloud provider VPC, i.e. the aws-vpc-cni plugin on AWS. That way both virtual-kubelet pods and regular pods will get their IP addresses from the VPC address space, and the VPC network will take care of routing.

If you would like to use another CNI plugin for some reason, that will also work as long as the cloud controller is configured to create cloud routes with the PodCIDR allocated to nodes and the CNI plugin used in the cluster is able to use the PodCIDR (most CNI plugins can do this).
Currently, Kip needs to run in host network mode. Since NodePorts are managed by the service proxy running on Kubernetes nodes, they also work seamlessly. Iptables rules for HostPort mappings are created and maintained by Kip.

The cloud instances hosting pods created via Kip also run a combined service proxy and network policy agent, which is kube-router. BGP is disabled in kube-router, since the cloud routes take care of routing to the CIDRs used for regular pods on the nodes; only service proxying and firewalling are enabled.

This way, pods can reach Kubernetes services either via IPVS or iptables NAT rules, created by kube-router, and network policies are also enforced (via iptables filtering rules, also maintained by kube-router). Essentially, this is the same setup as on Kubernetes nodes. Kube-router has a very low memory and CPU footprint, thus even small cloud instances can dedicate almost all their resources to the application(s) running in the pod.

## Public and Private Subnets in AWS

If running inside a cloud network, Kip will deploy pods onto cells that are located in the same subnet and region as the instance where Kip is running.  If Kip is runnning outside the cloud network, Kip's [provider.yaml](#provider-configuration) file will need to specify the subnet where cells will run.

All Kip cells have private network addresses and some will also have a publicly accessable address depending on the subnet Kip deploys to and annotations specified in the Kip pod.  In AWS, a public subnet is defined as subnet with a associated route table entry that points at an internet gateway.  Subnets without a routable internet gateway are considered private subnets.

Kip cells running in public subnets are assigned a public address by default while Kip cells in private subnets will never have a public address.  The user can force cells in public subnets to only have a private address by setting the following annotation on the pod: `pod.elotl.co/private-ip-only: "true"`.

Kip pods without public addresses and access to the internet must be able to download an itzo binary from S3 and be able to download any container images used in the pod spec.  The easiest way to ensure access to S3 from a private subnet in AWS is to setup a VPC endpoint to allow access to `com.amazonaws.us-east-1.s3`.  Likewise, a VPC endpoint can be used to allow access to ECR registries.

Itzo images can be downloaded from alternative locations by specifying a custom url for `cells.itzo.url` in [provider.yaml](#provider-configuration).  This allows cells to download the itzo binary from a webserver or other endpoint inside the user's cloud network.