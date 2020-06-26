# Provision a GCP Test Cluster

The Terraform config here can be used to provision a simple test cluster with Kip on GCP.

## Getting Started

You need:
* a GCP account configured and the necessary services enabled
* kustomize >= 3.0.0
* kubectl >= 1.14
* Terraform >= 0.12

You can then apply your config:

    cp env.tfvars.example myenv.tfvars
    vi myenv.tfvars # You can change settings for your cluster here.
    terraform apply -var-file myenv.tfvars

This will create a new GKE cluster and deploy Kip:

    $ kubectl get nodes
    NAME                                        STATUS   ROLES    AGE    VERSION
    gke-vk-node-pool-vk-9e5f3d39-44rg           Ready    <none>   171m   v1.14.10-gke.36
    kip-provider-0                              Ready    agent    108s

## Run a Pod via Virtual Kubelet

The node taint in kip is disabled in the manifest, so Kubernetes will try to run all pods via the virtual node.

If you decide to enable the taint on the virtual node (via removing the `--disable-taint` command line flag), you will need to add a toleration and/or node selector for pods that are meant to run via kip:

    spec:
      nodeSelector:
        type: virtual-kubelet
      tolerations:
      - key: virtual-kubelet.io/provider
        operator: Exists
