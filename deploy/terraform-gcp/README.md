# Provision a GCP Test Cluster

The Terraform config here can be used to provision a simple test cluster
with Kip on GCP.

## Getting Started

You need:

* a GCP account configured and the necessary services enabled
* kustomize >= 3.0.0
* kubectl >= 1.14
* Terraform >= 0.12
* google-cloud-sdk >= 321 (older versions may work)

Then create your configuration:

    $ cat > myenv.tfvars <<EOF
    # name of GCP Project
    project = "my-project"

    # Cluster name.
    cluster_name = "alice"

    # Region and zone to create resources in.
    region = "us-central1"
    zone = "us-central1-c"
    EOF

There’s a sample tfvars file in deploy/terraform-gcp/env.tfvars.example.

Run apply to provision the Kubernetes dataplane and the virtual-kubelet:

    terraform apply -var-file myenv.tfvars

This will create a new GKE cluster and deploy Kip:

    $ kubectl get nodes
    NAME                                        STATUS   ROLES    AGE    VERSION
    gke-vk-node-pool-vk-9e5f3d39-44rg           Ready    <none>   171m   v1.14.10-gke.36
    kip-provider-0                              Ready    agent    108s

## Run a Pod via Virtual Kubelet

The node taint in kip is disabled in the manifest, so Kubernetes will try
to run all pods via the virtual node.

To enable taint on the virtual node: remove the `--disable-taint` command
line flag. Then add a toleration or node selector for pods that are meant
to run via kip:

    spec:
      nodeSelector:
        type: virtual-kubelet
      tolerations:
      - key: virtual-kubelet.io/provider
        operator: Exists
