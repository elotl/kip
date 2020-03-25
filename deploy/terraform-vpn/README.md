# VPN Setup

The Terraform configuration example here creates a VPC with a VPN Gateway on AWS, and deploys Kip and a VPN client into a local Kubernetes cluster. A VPN connection will link the local cluster to the VPC, allowing pods and services to reach each other between the two sides. Kip is configured to create pods in the VPC.

For more information on Site-to-Site VPN, see [documentation from AWS](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html).

## Install

For provisioning, you need:
* terraform >= 0.12
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

### Use via Minikube

If you have a Linux environment, you can start minikube via:

    $ minikube start --vm-driver=none

Alternatively, you can spin up a Linux VM, and run minikube in it:

    $ vagrant init ubuntu/bionic64
    $ vagrant ssh

Now you're in the VM. Install dependencies and set up minikube:

    vagrant@ubuntu-bionic:~$ sudo apt-get update && sudo apt-get install -y conntrack docker.io unzip
    vagrant@ubuntu-bionic:~$ curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl && chmod 755 kubectl && sudo mv kubectl /usr/local/bin/
    vagrant@ubuntu-bionic:~$ curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 && chmod +x minikube
    vagrant@ubuntu-bionic:~$ sudo mv minikube /usr/local/bin/
    vagrant@ubuntu-bionic:~$ sudo minikube start --vm-driver=none
    [...]
    vagrant@ubuntu-bionic:~$ sudo chown -R vagrant: /home/vagrant/.{kube,minikube}

Minikube should start without errors. Install terraform:

    vagrant@ubuntu-bionic:~$ curl -LO https://releases.hashicorp.com/terraform/0.12.24/terraform_0.12.24_linux_amd64.zip && unzip terraform_0.12.24_linux_amd64.zip && chmod 755 terraform && sudo mv terraform /usr/local/bin/

Docker by default drops forwarded packets. Enable it:

    vagrant@ubuntu-bionic:~$ sudo iptables -P FORWARD ACCEPT

Clone this repo:

    vagrant@ubuntu-bionic:~$ git clone https://github.com/elotl/kip && cd kip && git checkout vilmos-minikube && cd deploy/terraform-vpn
    vagrant@ubuntu-bionic:/kip/deploy/terraform-vpn$ terraform init
    [...]

Set tunnel1_psk, tunnel2_psk, aws_access_key_id and aws_secret_access_key (check variables.tf for all possible configuration variables):

    vagrant@ubuntu-bionic:/kip/deploy/terraform-vpn$ vi myenv.tfvars

Make sure you also set up AWS access for terraform:

    vagrant@ubuntu-bionic:/kip/deploy/terraform-vpn$ export AWS_ACCESS_KEY_ID=...
    vagrant@ubuntu-bionic:/kip/deploy/terraform-vpn$ export AWS_SECRET_ACCESS_KEY=...
    vagrant@ubuntu-bionic:/kip/deploy/terraform-vpn$ terraform apply -var-file myenv.tfvars
    [...]

Once terraform finishes, you should have a working cluster with a VPN link into an AWS VPC, and kip set up to create pods in the VPC. You can interact with your cluster in your vagrant VM via `kubectl` as the default (vagrant) user.

Don't forget to delete any pods you created, and to remove the VPC and VPN gateway via `terraform destroy -var-file myenv.tfvars` before deleting your vagrant VM.

## Customizing your setup

See variables.tf for all possible configuration variables.
