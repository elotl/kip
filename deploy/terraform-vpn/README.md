# VPN Setup

The Terraform configuration example here creates a VPC with a VPN Gateway on AWS, and deploys Kip and a VPN client into a local Kubernetes cluster. A VPN connection will link the local cluster to the VPC, allowing pods and services to reach each other between the two sides. Kip is configured to create pods in the VPC.

For more information on Site-to-Site VPN, see [documentation from AWS](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html).

## Install

For provisioning, you need:
* terraform >= 0.12
* kustomize >= 3.0.0
* kubectl >= 1.14

The VPN client uses IPsec, and needs a few kernel modules available on the worker node:
* xfrm4_tunnel
* tunnel4
* ipcomp
* xfrm_ipcomp
* esp4
* ah4
* af_key
* ip_tunnel
* ip_vti

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` to tear down the VPC (but leave the local cluster unchanged). You have to delete the pods running in the VPC first via `kubectl delete`.

### Use BGP instead of static routes

By default, the VPN client will add static routes for the VPC CIDR, and add CIDRs used by Kubernetes on the AWS side (see the `local_cidrs` variable). If you want to use BGP for advertising routes, disable static routes:

    static_routes_only=false

It is also recommended to run the pod without host network mode in this case:

    vpn_hostnetwork=false

Check the IP of the VPN client pod, and use that as a BGP peer to receive AWS routes and advertise routes from the cluster. For example, if you use calico and the aws-vpn-client pod IP is 192.168.120.71:

    cat <<EOF | calicoctl apply -f -
    apiVersion: projectcalico.org/v3
    kind: BGPPeer
    metadata:
      name: aws-vpn-client
    spec:
      peerIP: 192.168.120.71
      asNumber: 65220
    EOF

See `variables.tf` for all related configuration parameters and their descriptions.

### Using via Minikube

The official minikube iso does not have a few kernel modules needed by the VPN client, so e.g. using the VirtualBox driver and installing the VPN won't work. If you have a Linux environment, you can use the `docker` or the `none` drivers instead:

    $ minikube start --vm-driver=none

or

    $ minikube start --vm-driver=docker

Alternatively, you can use also an ISO we built that includes the missing kernel modules:

    $ minikube start --iso-url=https://kip-builds.s3.amazonaws.com/minikube.iso

### Teardown

Don't forget to delete any pods you have created. You can check them:

    $ kubectl get pods -A --field-selector spec.nodeName=kip
    NAMESPACE     NAME                       READY   STATUS    RESTARTS   AGE
    default       debug-vk                   1/1     Running   0          3h16m
    kube-system   coredns-66bff467f8-vctcc   1/1     Running   0          3h50m
    kube-system   kube-proxy-48ggm           1/1     Running   0          3h51m

Then you can remove the VPC and VPN gateway via:

    $ terraform destroy -var-file myenv.tfvars

## Customizing your setup

See variables.tf for all possible configuration variables.
