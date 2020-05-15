# Provision a GCP Test Cluster

The Terraform config here can be used to provision a simple test cluster with Kip on GCP.

## Getting Started

You need:
* a GCP account configured and the necessary services enabled
* Terraform >= 0.12

Create an overlay for deploying Kip:

    cp -rf ../manifests/virtual-kubelet/overlays/gcp ../manifests/virtual-kubelet/overlays/local-my-gcp
    # Fill in authentication details for your GCP project.
    vi ../manifests/virtual-kubelet/overlays/local-my-gcp/provider.yaml

You can then apply your config:

    cp env.tfvars.example myenv.tfvars
    # Change kustomize-dir to ../manifests/virtual-kubelet/overlays/local-my-gcp
    vi myenv.tfvars
    terraform apply -var-file myenv.tfvars

This will create a new GKE cluster and deploy Kip:

    $ kubectl get nodes
    NAME                                        STATUS   ROLES    AGE    VERSION
    gke-vk-node-pool-vk-9e5f3d39-44rg           Ready    <none>   171m   v1.14.10-gke.36
    virtual-kubelet                             Ready    agent    108s

## Run a Pod via Virtual Kubelet

The node taint in virtual-kubelet is disabled in the manifest, so Kubernetes will try to run all pods via the virtual node.

If you decide to enable the taint on the virtual node (via removing the `--disable-taint` command line flag), you will need to add a toleration and/or node selector for pods that are meant to run via virtual-kubelet:

    spec:
      nodeSelector:
        type: virtual-kubelet
      tolerations:
      - key: virtual-kubelet.io/provider
        operator: Exists
