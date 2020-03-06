# Kip, the Kubernetes Cloud Instance Provider

Kip is a [Virtual Kubelet](https://github.com/virtual-kubelet/virtual-kubelet) provider that allows a Kubernetes cluster to transparently launch pods onto their own cloud instances.  When a pod is scheduled onto the Virtual Kubelet Kip starts a right sized cloud instance for the pod’s workload and dispatches the pod onto the instance.  When the pod is finished running, the cloud instance is terminated. We call these cloud instances “cells”.

When workloads run on Kip, your cluster size naturally scales with the cluster workload, pods are strongly isolated from each other and the user is freed from managing worker nodes and strategically packing pods onto nodes.  This results in lower cloud costs, improved security and simpler operational overhead.

#### Table of Contents

* [Installation](#installation)
    + [Option 1: Create a Minimal Cluster](#installation-option-1-create-a-minimal-k8s-cluster)
    + [Option 2: Using an Existing Cluster](#installation-option-2-using-an-existing-cluster)
* [Running Pods on Virtual Kubelet](#running-pods-on-virtual-kubelet)
* [Uninstall](#uninstall)
* [Current Status](#current-status)
* [FAQ](#faq)
* [How it Works](#how-it-works)

## Installation
There are two ways to get Kip up and running in AWS.

* Option 1: Use the provided terraform scripts to create a new VPC and a new kubernetes cluster.
* Option 2: Add Kip to an existing kubernetes cluster.

### Installation Option 1: Create a Minimal K8s Cluster

Prequisites:
- An AWS account
- [Terraform](https://www.terraform.io/downloads.html) (tested with terraform 0.12)
- [aws-cli](https://aws.amazon.com/cli/)
- jq

In [deploy/terraform](deploy/terraform), you will find a terraform config that creates a simple one master, one worker cluster and starts a Kip deployment.

``` bash
    cd deploy/terraform
    cp env.tfvars.example myenv.tfvars
    vi myenv.tfvars  # customize variables as necessary
    terraform apply -var-file myenv.tfvars
```

### Installation Option 2: Using an Existing Cluster

To deploy Kip into an existing cluster, you'll need to setup cloud credentials for the Kip provider to use then create a ConfigMap containing [deploy/virtual-kubelet-config.yaml](deploy/virtual-kubelet-config.yaml).  After creating the configuration, apply [deploy/virtual-kubelet.yaml](deploy/virtual-kubelet.yaml) to create the necessary kubernetes resources to support and run the provider.

**Step 1: Credentials**

Kip requires credentials to manipulate cloud instances, security groups and other cloud resources.  In AWS, those credentials can be supplied via API keys or by an instance profile.

**Credentials Option 1 - Configuring AWS API keys:**

Use `envsubst` to apply virtual-kubelet-config.yaml with your keys and apply the virtual-kubelet-config.yaml ConfigMap:

    export aws_access_key_id="<your aws access key>"
    export aws_secret_access_key="<your aws secret access key>"
    envsubst '$aws_access_key_id,aws_secret_access_key' < deploy/virtual-kubelet-config.yaml | kubectl apply -f -

Alternatively you can edit the file by hand, replacing the two templated environment variables with your access keys and apply the manifest.

**Credentials Option 2 - Instance Profile Credentials:**

Kip can use credentials from an instance profile of the VM it is running on.  This setup assumes the user has already provisioned a node with the correct [instance profile](#docs/kip-iam-permissions.md).  To use instance profile credentials, leave the `accessKeyID` and `secretAccessKey` blank in virtual-kubelet-config.yaml.

    export aws_access_key_id=""
    export aws_secret_access_key=""
    envsubst '$aws_access_key_id,aws_secret_access_key' < deploy/virtual-kubelet-config.yaml | kubectl apply -f -

**Step 2: Apply virtual-kubelet.yaml**

The resources in virtual-kubelet.yaml create ServiceAccounts Roles and a virtual-kubelet Deployment to run the provider. [Kip is not stateless](docs/state.md), virtual-kubelet.yaml will also create a PersistentVolume to store the provider data.

    kubectl apply -f deploy/virtual-kubelet.yaml

## Running Pods on Virtual Kubelet

To assign pods to run on the virtual kubelet node, add the following node selector and toleration to the pod spec in manifests.

    spec:
      nodeSelector:
        type: virtual-kubelet
      tolerations:
      - key: virtual-kubelet.io/provider
        operator: Exists

## Uninstall

If you used the provided terraform config for creating your cluster, you can remove the VPC and the cluster via:

    terraform destroy -var-file <env.tfvars>.

If you deployed Kip in an existing cluster, make sure that you first remove all the pods and deployments that have been created by Kip. Then remove the virtual-kubelet deployment via:

    kubectl delete -n kube-system deployment virtual-kubelet

## Current Status

### Features
- [Networking](docs/networking.md), including host network mode, cluster IPs, DNS, HostPorts and NodePorts
- Pods will be started on a cloud instance that matches the pod resource requests/limits. If no requests/limits are present in the pod spec, kip will fall back to a default cloud instance type specified in [cloud-instance-provider.yaml](docs/provider-config.md)
- GPU instances
- Logs
- Exec
- Stats
- Readiness/Liveness probes
- Service account token automounts in pods.
- [Security Groups](docs/security_groups.md)
- Attaching instance profiles to Cells
- The following volume types are supported
    - EmptyDir
    - ConfigMap
    - Secret
    - HostPath
    - Projected ConfigMap and Secrets

### Limitations
- Stateful workloads and PersistentVolumes are not supported.
- No support for updating ConfigMaps and Secrets for running Pods and Cells.
- Virtual-kubelet has limitations on what it supports in the Downward API, e.g. pod.status.podIP is not supported
- Unsupported pod attributes:
    - ReadinessGates
    - Lifecycle handlers
    - TerminationGracePeriodSeconds
    - ActiveDeadlineSeconds
    - Subdomain
    - HostAliases
    - The following PodSecurityContext fields
        - FSGroup
        - RunAsNonRoot
        - ShareProcessNamespace
        - HostIPC
        - HostPID
- Volume mount propagation is always bidirectional

We are actively working on adding missing features. One of the main objectives of the project is to provide full support for all Kubernetes features.

## FAQ

**Q.** I’ve seen the name Milpa mentioned in various places (logs, directories, tags). What is Milpa?

**A.** Kip’s source code was adapted from an earlier project developed at Elotl called Milpa.  We will be migrating away from that name in coming releases.  Milpa started out as a stand alone unikernel (and later container) orchestration system and it was natural to move a subset of its functionality into an open source virtual-kubelet provider.

##
**Q.** How long does it take to start a workload?

**A.** In AWS, instances boot in under a minute, usually pods are dispatched to the instance in about 45 seconds. Depending on the size of the container, a pod will be running in 60 to 90 seconds.  In our experience, starting pods in Azure can be a bit slower with startup times between 1.5 to 3 minutes.

##
**Q.** Does it work with the Horizontal Pod Autoscaler and Vertical Pod Autoscaler?

**A.** Yes it does.  However, to use the VPA with the provider, the pod must be dispatched to a new cloud instance.

##
**Q.** Are DaemonSets supported?

**A.** Yes, though they might not work the way intended. The pod will start on a separate cloud instance, and not on the node.

##
**Q.** Are you a [kubernetes conformant](https://github.com/cncf/k8s-conformance) runtime?

**A.** We are not 100% conformant at this time but we are working towards getting as close as possible to conformant.  Currently we are 70-80% conformant but are hoping to get those values above 90% soon.

##
**Q.** What cloud providers does Kip support

**A.** Kip is currently GA on AWS and alpha on Azure.

##
**Q.** What components make up the Kip system?

**A.** The following repositories are part of the Kip system
* [Itzo](https://github.com/elotl/itzo) containes the cell agent and code for building cell images
* [Tosi](https://github.com/elotl/tosi) for downloading images to cells
* [Cloud-Init](https://github.com/elotl/cloud-init) a minimal cloud-init implementation

## How it Works
* [Cells](docs/cells.md)
* [Networking](docs/networking.md)
* [Annotations](docs/annotations.md)
* [Security Groups](docs/security_groups.md)
* [Provider Configuration](docs/provider-config.md)
* [IAM Permissions](docs/kip-iam-permissions.md)
* [State](docs/state.md)