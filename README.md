![KIP](KipOpenSource-logo.png "KIP")
# Kip, the Kubernetes Cloud Instance Provider

Kip is a [Virtual Kubelet](https://github.com/virtual-kubelet/virtual-kubelet) provider that allows a Kubernetes cluster to transparently launch pods onto their own cloud instances.  The kip pod is run on a cluster and will create a virtual Kubernetes node in the cluster.  When a pod is scheduled onto the Virtual Kubelet, Kip starts a right-sized cloud instance for the pod’s workload and dispatches the pod onto the instance.  When the pod is finished running, the cloud instance is terminated. We call these cloud instances “cells”.

When workloads run on Kip, your cluster size naturally scales with the cluster workload, pods are strongly isolated from each other and the user is freed from managing worker nodes and strategically packing pods onto nodes.  This results in lower cloud costs, improved security and simpler operational overhead.

#### Table of Contents

* [Installation](#installation)
    + [Option 1: Create a Minimal Cluster](#installation-option-1-create-a-minimal-k8s-cluster)
    + [Option 2: Using an Existing Cluster](#installation-option-2-using-an-existing-cluster)
* [Running Pods on Kip](#running-pods-on-kip)
* [Uninstall](#uninstall)
* [Current Status](#current-status)
* [FAQ](#faq)
* [How it Works](#how-it-works)

## Installation

There are two ways to get Kip up and running.

* Option 1: Use the provided terraform scripts to create a new Kubernetes cluster with a single Kip node on AWS or GKE.
* Option 2: Add Kip to an existing kubernetes cluster.

### Installation Option 1: Create a Minimal K8s Cluster

Prequisites:
- An AWS or Google Cloud account
- [Terraform](https://www.terraform.io/downloads.html) (tested with terraform 0.12)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) >= 1.14
- jq and [aws-cli](https://aws.amazon.com/cli/) if using AWS

In [deploy/terraform-aws](deploy/terraform-aws), you will find a terraform config that creates a simple one master, one worker cluster and starts Kip on AWS.

``` bash
cd deploy/terraform-aws
terraform init
cp env.tfvars.example myenv.tfvars
vi myenv.tfvars  # customize variables as necessary
terraform apply -var-file myenv.tfvars
```

On Google Cloud, the config under `deploy/terraform-gcp` works in a similar way, but uses a GKE base cluster.

### Installation Option 2: Using an Existing Cluster

To deploy Kip into an existing cluster, you'll need to setup cloud credentials that allow the Kip provider to manipulate cloud instances, security groups and other cloud resources.

**Step 1: Credentials**

In AWS, Kip can either use API keys supplied in the Kip provider configuration file (`provider.yaml`) or use the instance profile of the machine the Kip pod is running on.

On Google Cloud, a service account key is used for authentication.

**AWS Credentials Option 1 - Configuring AWS API keys:**

You can configure the AWS access key Kip will use in your provider configuration, via changing `accessKeyID` and `secretAccessKey` under the `cloud.aws` section. See below on how to create a kustomize overlay with your custom provider configuration.

**AWS Credentials Option 2 - Instance Profile Credentials:**

In AWS, Kip can use credentials supplied by the [instance profile](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html) attached to the node the pod is dispatched to.  To use an instance profile, create an IAM policy with the [minimum Kip permissions](docs/kip-iam-permissions.md) then apply the instance profile to the node that will run the Kip provider pod.  The Kip pod must run on the cloud instance that the instance profile is attached to.

**GCP Credentials - Service Account private key:**

Add your email and key to `cloud.gce.credentials`. Example:

    cloud:
      gce:
        projectID: "my-project"
        credentials:
          clientEmail: my-account@my-project.iam.gserviceaccount.com
          privateKey: "-----BEGIN PRIVATE KEY-----\n[base64-encoded private key]-----END PRIVATE KEY-----\n"
        zone: us-central1-c
        vpcName: "default"
        subnetName: "default"

**Step 2: Apply the manifests**

The resources in [deploy/manifests/kip](deploy/manifests/kip) create ServiceAccounts, Roles and a StatefulSet to run the provider. [Kip is not stateless](docs/state.md), the manifest will also create a PersistentVolumeClaim to store the provider data.

Once credentials are set up, apply [deploy/manifests/kip/base](deploy/manifests/kip/base) to create the necessary kubernetes resources to support and run the provider:

    $ kubectl apply -k deploy/manifests/kip/base

For rendering the manifests, [kustomize](https://kustomize.io/) is used. You can create your own overlays on top of the base template. For example, to override provider.yaml, Kip's configuration file:

    $ mkdir -p deploy/manifests/kip/overlay/local-config
    $ cp deploy/manifests/kip/base/provider.yaml deploy/manifests/kip/overlay/local-config/provider.yaml
    # Edit your provider configuration file.
    $ vi deploy/manifests/kip/overlay/local-config/provider.yaml
    $ cat > deploy/manifests/kip/overlay/local-config/kustomization.yaml <<EOF
    > apiVersion: kustomize.config.k8s.io/v1beta1
    > kind: Kustomization
    > bases:
    > - ../../base
    > configMapGenerator:
    > - behavior: merge
    >   files:
    >   - provider.yaml
    >   name: kip-config
    >   namespace: kube-system
    EOF
    $ kubectl apply -k deploy/manifests/kip/overlays/local-config

After applying, you should see a new kip pod in the kube-system namespace and a new node named "kip-0" in the cluster.

## Running Pods on Kip

To assign pods to run on the virtual kubelet node, add the following node selector to the pod spec in manifests.

    spec:
      nodeSelector:
        type: virtual-kubelet

If you enabled taints on your virtual node (they are disabled by default in the example manifests; remove `--disable-taint` from the command line flags to enable), add the necessary tolerations too:

    spec:
      tolerations:
      - key: virtual-kubelet.io/provider
        operator: Exists

## Uninstall

If you used the provided terraform config for creating your cluster, you can remove the VPC and the cluster via:

    terraform destroy -var-file <env.tfvars>.

If you deployed Kip in an existing cluster, make sure that you first remove all the pods and deployments that have been created by Kip. Then remove the kip statefulset via:

    kubectl delete -n kube-system statefulset kip

## Current Status

### Features
- [Networking](docs/networking.md), including host network mode, cluster IPs, DNS, HostPorts and NodePorts
- Pods will be started on a cloud instance that matches the pod resource requests/limits. If no requests/limits are present in the pod spec, Kip will fall back to a default cloud instance type specified in [provider-config.yaml](docs/provider-config.md)
- GPU instances
- Logs
- Exec
- Stats
- Readiness/Liveness probes
- Service account token automounts in pods.
- [Security Groups](docs/security_groups.md)
- Attaching instance profiles to Cells via [annotations](docs/annotations.md)
- The following volume types are supported
    - EmptyDir
    - ConfigMap
    - Secret
    - HostPath
    - Projected ConfigMaps and Secrets

### Limitations
- Stateful workloads and PersistentVolumes are not supported.
- No support for updating ConfigMaps and Secrets for running Pods and Cells.
- Virtual-kubelet has limitations on what it supports in the Downward API, e.g. pod.status.podIP is not supported
- VolumeMounts do not support readOnly, subPath and subPathExpr attributes.
- VolumeMount mountPropagation is always Bidirectional
- Unsupported pod attributes:
    - EphemeralContainers
    - ReadinessGates
    - Lifecycle handlers
    - TerminationGracePeriodSeconds
    - ActiveDeadlineSeconds
    - VolumeDevices
    - TerminationMessagePolicy FallbackToLogsOnError is not implemented
    - The following PodSecurityContext fields
        - FSGroup
        - RunAsNonRoot
        - ShareProcessNamespace
        - HostIPC
        - HostPID


We are actively working on adding missing features. One of the main objectives of the project is to provide full support for all Kubernetes features.

## FAQ

**Q.** I’ve seen the name Milpa mentioned in the logs and source code. What is Milpa?

**A.** Kip’s source code was adapted from an earlier project developed at Elotl called Milpa.  We will be migrating away from that name in coming releases.  Milpa started out as a stand alone unikernel (and later container) orchestration system and it was natural to move a subset of its functionality into an open source virtual-kubelet provider.

##
**Q.** How long does it take to start a workload?

**A.** In AWS, instances boot in under a minute, usually pods are dispatched to the instance in about 45 seconds. Depending on the size of the container image, a pod will be running in 60 to 90 seconds.  In our experience, starting pods in Azure can be a bit slower with startup times between 1.5 to 3 minutes.

##
**Q.** Does it work with the Horizontal Pod Autoscaler and Vertical Pod Autoscaler?

**A.** Yes it does.  However, to use the VPA with the provider, the pod must be dispatched to a new cloud instance.

##
**Q.** Are DaemonSets supported?

**A.** Yes, though they might not work the way intended. The pod will start on a separate cloud instance, and not on the node.  It's possible to patch a DaemonSet so it does not get dispatched to the Kip virtual node.

##
**Q.** Are you a [kubernetes conformant](https://github.com/cncf/k8s-conformance) runtime?

**A.** We are not 100% conformant at this time but we are working towards getting as close as possible to conformance.  Currently Kip passes 70-80% of conformance tests but are hoping to get those values above 90% soon.

##
**Q.** What cloud providers does Kip support?

**A.** Kip is currently GA on AWS and pre-alpha on Azure.

##
**Q.** What components make up the Kip system?

**A.** The following repositories are part of the Kip system
* [Itzo](https://github.com/elotl/itzo) containes the cell agent and code for building cell images
* [Tosi](https://github.com/elotl/tosi) for downloading images to cells
* [Cloud-Init](https://github.com/elotl/cloud-init) a minimal cloud-init implementation

##
**Q.** We have our custom built AMI. Can I use it for running cells?

**A.** Yes, take a look at [Bring your Own AMI](docs/cells.md#bring-your-own-ami).

## How it Works
* [Cells](docs/cells.md)
* [Networking](docs/networking.md)
* [Annotations](docs/annotations.md)
* [Security Groups](docs/security-groups.md)
* [Provider Configuration](docs/provider-config.md)
* [IAM Permissions](docs/kip-iam-permissions.md)
* [State](docs/state.md)
