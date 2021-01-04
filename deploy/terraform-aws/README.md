# Provision a Simple Test Cluster

The Terraform config here can be used to provision a simple test cluster with kip.

![VPC with KIP](kip_tf_dev_env.png "VPC with KIP")

## Getting Started

You need:

* an AWS account configured (check that e.g. `aws iam get-user` works)
* Terraform >= 0.12
* aws-cli
* jq

Initialize Terraform and verify the configuration:

    $ terraform init
    [... lots of output ...]
    $ terraform plan
    [... lots of output ...]

[Upload][0] your SSH public key to AWS, and remember its name. Make
sure the private key for this public key is your default SSH key:
`$HOME/.ssh/id_rsa` (NB: investigate the SSH key situation to see if
there’s an easier way to get this set-up correctly.) In this example
we’ll name our key-pair “alice-key”. Come up with a unique cluster
name like “alice-cluster”. Finally put these values into a tfvars file:

[0]: https://console.aws.amazon.com/ec2/v2/home?region=us-east-1#KeyPairs:

    $ cat > myenv.tfvars <<EOF
    ssh_key_name = "alice-key"
    cluster_name = "alice-cluster"
    EOF

There are other configuration variables available, there’s an example tfvars
file env.tfvars.example. You can then apply your config:

    terraform apply -var-file myenv.tfvars

This will create a new VPC and a one-node Kubernetes cluster in it with kip, and show the public IP address of the node when done:

    [...]

    Apply complete! Resources: 13 added, 0 changed, 0 destroyed.

    Outputs:

    node-ip = 34.201.59.101

You can now ssh into the instance using the username "ubuntu", and the ssh key you set in your environment file. (It takes a a minute or two for the instance to bootstrap). On the instance, you can use kubectl to interact with your new cluster:

    $ kubectl get nodes
    NAME                          STATUS   ROLES    AGE   VERSION
    ip-10-0-26-113.ec2.internal   Ready    master   67s   v1.17.3
    kip-provider-0                Ready    agent    13s   v1.14.0-vk-v0.0.1-125-g3b2cc98

If you haven't set an existing ssh key in your configuration, a new ssh key has been created.
You can extract and use it via:

    $ terraform show -json | jq -r '.values.root_module.resources | .[] | select(.address=="tls_private_key.ssh-key") | .values.private_key_pem' > /tmp/id_rsa
    $ ssh -i /tmp/id_rsa ubuntu@<public IP of node>

AmazonSystemManagedInstanceCore permissions are attached to the node instance.
You should be able to see node instance in AWS Console under Systems Manager -> Instances and Nodes -> Managed Instances. There you can start new SSH session.
You can also connect to instance using AWS CLI:
`$ aws ssm start-session --target <instance ID in i-XXX format>`

## Run a Pod via Virtual Kubelet

The node taint in kip is disabled in the manifest, so Kubernetes will try to run all pods via the virtual node.

If you decide to enable the taint on the virtual node (via removing the `--disable-taint` command line flag), you will need to add a toleration and/or node selector for pods that are meant to run via kip:

    spec:
      nodeSelector:
        type: virtual-kubelet
      tolerations:
      - key: virtual-kubelet.io/provider
        operator: Exists
