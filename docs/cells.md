## Cells

Cells are the name we’ve given to the cloud instances that Kip runs pods on.  The cells run Alpine Linux and a lightweight agent called [Itzo](https://github.com/elotl/itzo) that performs the task of running pods on the cells.  More information about cells can be found in the itzo repository.

To boot cells, unless an instance-type annotation is present in the pod, kip will choose the cheapest instance type that satisfies the resource requirements of the pod and will fall back to the `defaultInstanceType` in the `provider.yaml` file if no resources are specified.  Server certificates are passed to the cell via instance user data and those certificates allow the kip controller to connect to the cell.

Cells can be customized by specifying a [cloud-init](cloud-init.md) file that will be applied when the instance boots.

### Bring your Own Image

We maintain images that are optimized for cells and come with our tools pre-installed. However, if you need a custom-built image or you have a dependency only available for certain Linux distributions, you can either add our cell agent to your image builds, or use cloud-init to fetch it and start it at boot time.

You can check out our image build scripts and configuration files [here](https://github.com/elotl/kip-cell-image).

If you want to use cloud-init to download our cell agent, first update your provider config:

    cells:
      cloudInitFile: /etc/kip/cloudinit.yaml
      bootImageSpec:
        owners: "099720109477"
        filters: name=ubuntu/images/hvm-ssd/ubuntu-bionic-18.04-amd64-server-*

Add a cloud-init section to the provider configmap, e.g. for ubuntu or debian:

    cloudinit.yaml: |
      runcmd:
        - apt-get update && apt-get install -y ipset iproute2 iptables
        - wget -O /usr/local/bin/itzo https://itzo-kip-download.s3.amazonaws.com/itzo-latest
        - wget -O /usr/local/bin/kube-router https://milpa-builds.s3.amazonaws.com/kube-router
        - wget -O /usr/local/bin/tosi https://tosi.s3.amazonaws.com/tosi
        - chmod 755 /usr/local/bin/*
        - mount --make-rprivate /
        - /usr/local/bin/itzo -v=2 > /var/log/itzo.log 2>&1 &

Finally, restart the provider:

    $ kubectl delete pod -n kube-system -l app=kip-provider

### Secondary/Alias IP Handling

If you build your custom image or use cloud-init to bootstrap cell instances, you will also need to ensure that the operating system does not configure secondary or alias IP addresses automatically. A secondary or alias IP address (depending on the cloud provider) is used for the pod that runs on the instance, and configuring it automatically on the main interface will result in pod networking issues.

On our default images this is done when the image is built, but on custom images you might have take extra steps. Please refer to your cloud provider and operating system documentation on the exact steps.

You can find more information on networking [here](networking.md).

### Cell Configuration via Cloud Parameters

By default, cell configuration (agent certificate and key, cell config yaml, etc) will be sent to cells by Kip via user data in cloud-init.

To use a more secure channel instead, set `cells.useCloudParameterStore` in your provider config:

    cells:
      ...
      useCloudParameterStore: true

This is currently only supported on AWS.

Make sure that Kip has the necessary IAM permissions to use SSM. The [default IAM policy](kip-iam-permissions.md) has this.

You can also limit cells to ensure they are only able to read their own parameters, for example:

    {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Sid": "",
                "Effect": "Deny",
                "Action": "ssm:GetParametersByPath",
                "Resource": "arn:aws:ssm:<region>:<account ID>:parameter/kip/cells/*"
            },
            {
                "Sid": "",
                "Effect": "Allow",
                "Action": [
                    "ssm:GetParameters",
                    "ssm:GetParameter"
                ],
                "Resource": "arn:aws:ssm:<region>:<account ID>:parameter/kip/cells/*",
                "Condition": {
                    "StringEquals": {
                        "aws:ResourceTag/AWSUserID": "${aws:userid}"
                    }
                }
            },
            {
                "Sid": "",
                "Effect": "Deny",
                "Action": [
                    "ssm:GetParameters",
                    "ssm:GetParameter"
                ],
                "Resource": "arn:aws:ssm:<region>:<account ID>:parameter/kip/cells/*",
                "Condition": {
                    "StringNotEquals": {
                        "aws:ResourceTag/AWSUserID": "${aws:userid}"
                    }
                }
            }
        ]
    }

This policy needs to be attached to the role used by the instance profile cells use via `cells.defaultIAMPermissions`. Kip tags parameters in SSM with the AWS user ID (which is <role ID from instance profile>:<instance ID>), so cells will only be able to read parameters that are tagged with their user ID.
