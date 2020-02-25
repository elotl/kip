# Kip, the Cloud Instance Provider

## What it is

Kip is a virtual-kubelet provider that allows a kubernetes cluster to transparently launch pods onto their own cloud instances.  When a pod is scheduled onto the virtual-kubelet Kip starts a right sized cloud instance for the pod’s workload and dispatches the pod onto the instance.  When the pod is finished running, the cloud instance is terminated. We call these cloud instances “cells”.

When workloads run on Kip, your cluster size naturally scales with the cluster workload, pods are strongly isolated from each other and the user is freed from managing worker nodes and strategically packing pods onto nodes.  This results in lower cloud costs, improved security and simpler operational overhead.

## Getting Started

### Quickstart

What you need:
- an AWS account
- terraform (tested with terraform 0.12)
- aws-cli
- jq

In [deploy/terraform](deploy/terraform), you will find a terraform config that creates a simple one master, one worker cluster and starts a Kip deployment.

    cd deploy/terraform
    cp env.tfvars.example myenv.tfvars
    vi myenv.tfvars
    terraform apply -var-file myenv.tfvars

### Using an Existing Cluster

To deploy Kip into an existing cluster, it’s necessary to create a configuration file for the provider (virtual-kubelet-config.yaml) and create a couple of kubernetes resources to support the provider (virtual-kubelet.yaml)

Step 1: Configuring cloud credentials for the provider
Kip requires credentials to start, stop and control cloud instances, security groups and other cloud resources.  In AWS, those credentials can be supplied via API keys or instance profiles.

Step 1, Option 1: Configuring AWS API keys
Use envsubst to apply virtual-kubelet-config.yaml with your keys and apply the virtual-kubelet-config.yaml ConfigMap:

    export aws_access_key_id=’’
    export aws_secret_access_key=’’
    envsubst '$aws_access_key_id,aws_secret_access_key' < deploy/virtual-kubelet-config.yaml | kubectl apply -f -

Alternatively you can edit the file by hand, replacing the two templated environment variables with your access keys and apply the manifest.

Step 1, Option 2: Instance Profile Credentials 

Kip can use credentials from an instance profile of the VM it is running on.  This setup assumes the user has already provisioned a node with the correct instance profile (TODO link).  To use instance profile credentials, remove the `accessKeyID` and `secretAccessKey` from virtual-kubelet-config.yaml.

<TODO: shell script to remove and apply virtual-kubelet-config.yaml>

Step 2: Apply virtual-kubelet.yaml

The manifests in virtual-kubelet.yaml create ServiceAccounts Roles and a virtual-kubelet Deployment to run the provider. Kip is not stateless, the provider data is stored on a PersistentVolume.
  
    kubectl apply -f deploy/virtual-kubelet.yaml 

## Run Pods on Virtual Kubelet

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

If you deployed Kip in an existing cluster, make sure that you first remove all the pods and deployments that have been created via it. Then remove the virtual-kubelet deployment via:

    kubectl delete -n kube-system deployment virtual-kubelet

## Current Status

Features
- Networking, including host network mode, cluster IPs, DNS, HostPorts and NodePorts (TODO link to networking section)
- Pods will be started on a cloud instance that matches the pod resource requests/limits. If no requests/limits are present in the pod spec, cloud-instance-provider will fall back to a default cloud instance type.
- GPU instances
- Logs
- Stats
- Readiness/Liveness probes
- Service account token automounts in pods.
- Security groups (TODO: group created by the provider, how to add additional groups)
- Annotations TODO
- Cells (TODO link to itzo repo, explain how they work)
- Virtual-kubelet-config.yaml parameters
- The following volume types are supported: EmptyDir, ConfigMap, Secret, HostPath.
- TODO: do we support private registries right now?

Limitations
- Stateful workloads and PersistentVolumes are not supported.
- No support for updating ConfigMaps and Secrets.
- Virtual-kubelet has limitations on what it supports in the Downward API, e.g. pod.status.podIP is not supported.
- Unsupported pod attributes:
- ReadinessGates
- Lifecycle handlers
- TerminationGracePeriodSeconds
- ActiveDeadlineSeconds
- Subdomain
- HostAliases
- The following PodSecurityContext fields: FSGroup, RunAsNonRoot, ShareProcessNamespace, HostIPC, HostPID

We are actively working on adding missing features. One of the main objectives of the project is to provide full support for all Kubernetes features.

Kip is currently GA on AWS and alpha on Azure.

## FAQ

Q. How long does it take to start a workload?
A. In AWS, instances boot in under a minute, usually pods are dispatched to the instance in about 45 seconds. Depending on the size of the container, a pod will be running in 60 to 90 seconds.  In our experience, starting pods in Azure can be a bit slower with startup times between 1.5 to 3 minutes.

Q. Does it work with the Horizontal Pod Autoscaler and Vertical Pod Autoscaler?
A. Yes it does.  However, to use the VPA with the provider, the pod must be dispatched to a new cloud instance.

Q. Are DaemonSets supported?
A. Yes, though they might not work the way intended. The pod will start on a separate cloud instance, and not on the node.

## How it Works

### [Networking](docs/networking.md)
