## Provider Configuration

Kip is configured using a yaml file.  The easiest way to get that file into kip’s virtual-kubelet pod is to use a ConfigMap and then supply the location of the file to kip via the `--provider-config` flag.  A sample config file with documentation comments for each section is shown below:

```yaml
---
# Like all the resources in k8s, this configuration file is versioned.
apiVersion: v1

# configures the cloud provider for kip
cloud:

  # Only one cloud provider can be enabled in kip. Currently
  # only AWS is supported but we are working on support for
  # Azure

  aws:

    # The AWS region where kip will create instances. Example:
    # `us-east-1`. The environment variable `AWS_REGION` can also be
    # used instead.
    region: us-east-1

    # The AWS access key ID Kip will use for interacting with the
    # AWS API. The environment variable `AWS_ACCESS_KEY_ID` can also
    # be used instead.
    accessKeyID: FILL_IN

    # The AWS secret access key. The environment variable
    # `AWS_SECRET_ACCESS_KEY` can also be used instead.
    secretAccessKey: FILL_IN

    # This is the VPC ID in which instances will be created.
    # "default" will selecct the default VPC. If empty, kip will
    # attempt to detect if it is running on an instance inside a VPC
    # (via AWS metadata) and will use its current VPC.
    # vpcID: ''

    # the AWS subnet ID of the subnet kip will launch pods into.  if
    # blank, kip will detect its current subnet (via AWS metadata)
    # and will launch pods into that subnet.
    # subnetID: ''

# the etcd section controls how kip stores its state, either using
# an external etcd cluster or using an embedded etcd database.
etcd:
  # Settings for running kip from an embedded etcd server.
  internal:

    # The dataDir is the directory that will be used by etcd for
    # storage. The kip user must have write access to the directory.
    # Defaults to `/opt/kip/data`. The data directory can also be
    # specified in the `configFile` but should not be specified in both
    # locations.
    dataDir: "/opt/kip/data"

  # Path to an etcd configuration file that will be used to further
  # customize the behavior of etcd. All etcd command-line flags are
  # supported. For more information, see the documentation available
  # on the etcd website.
  # configFile: /opt/kip/etc/etcd.yml

# the kubelet section describes the specs of the virtual-kubelet node
# that will be created by kip. This is the advertised size of the
# kubernetes node, not the size of of the physical instance the
# virtual-kubelet is running on. Values are resource quantities,
# similar to those found in a container's Resources configuration.

kubelet:
  cpu: "64"
  memory: "512Gi"
  pods: "400"

# the cells section configures parameters that affect kip cells
cells:
  # nametag is a name that will be added onto cloud tags to help
  # identify cloud resources that belong to this controller. Because
  # of restrictions across clouds, the nametag must be a valid dns
  # label (the name must start with a lower case letter followed by
  # lowercase letters numbers or dashes) and must have a length of 34
  # characters or less.
  nametag: kip

  # defaultInstanceType specifies the the default cloud instance type
  # Kip will use when creating pods, if the pod does not set an
  # instance type. Examples t3.micro, c4.8xlarge, p2.xlarge
  defaultInstanceType: t3.nano

  # defaultVolumeSize specifies the default size for the root volume
  # of cells in bytes.  To set the root volume to 8 GiB, specify the
  # value as "8Gi".  Must be 1Gi or larger, defaults to 5Gi.
  defaultVolumeSize: "5Gi"

  # bootImageSpec is a dictionary of cloud-specific image properties for
  # specifying the boot image to use for cells.
  # Valid fields on AWS are:
  #   - owners, which is a space separated list of AWS account IDs, "self", or
  #     an AWS owner alias such as "amazon" or "aws-marketplace".
  #   - executableUsers, which is a space separated list of users with explicit
  #     launch permissions. You can specify an AWS account ID, "self" or "all".
  #   - filters, which lets you filter results. Example:
  #     "name=elotl-kip-* is-public=true". This is equivalent to using:
  #     $ aws ec2 describe-images \
  #     --filters Name=name,Values=elotl-kip-*,Name=is-public,Values=true
  #   - imageIDs, one or more space separated image IDs. Example:
  #     "ami-07cfdf0f7c08293fd"
  # If there are multiple matching images, the latest one will be used.
  # See
  # https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-images.html
  # for more details.
  # The below settings are the default if no bootImageSpec is specified.
  bootImageSpec:
    owners: "689494258501"
    filters: "name=elotl-kip-*"

  # cloudInitFile specifies a path to a cloudInitFile that will be
  # used to provision all cells that Kip boots. Kip will detect
  # modifications to this file. Cells started afte a modification are
  # made will get the updated cloudInit file.
  #
  # cloudInitFile: /etc/virtual-kubelet/cloudinit.yml

  # standbyCells is used to speicfy pools of standby cells kip will
  # keep so pods created can be dispatched to cells quickly.
  # standbyCells should be a list of `standbyCell` values that contain
  # three parameters
  #  * instanceType: the name of the cloud instance type
  #  * spot: whether the instance should be a spot instance
  #  * count: the number of standby instances of this types
  #
  # standbyCells:
  #   - instanceType: "t2.micro"
  #     count: 2
  #     spot: false

  # extraSecurityGroups contains the IDs of additional security groups
  # that will be attached to kip cells.  In AWS those groups looks
  # like: sg-av3sp192jur
  #
  # extraSecurityGroups:
  #   - sg-246810

  # itzo configures the version of the itzo agent to use and where
  # itzo will be downloaded from.  You should only customize this if
  # you have built your own itzo agent or you would like to pin your
  # itzo agent version to a particular version.

  # itzo:
  #   version: 532
  #   url: "http://itzo-kip-download.s3.amazonaws.com"

  # statusPeriod controls how often (in seconds) kip will query cells
  # for their status.

  # statusInterval: 5

  # the healthCheck section is used to configure how cells are
  # health checked.  Kip supports two methods of determining cell health:
  # 1) Status health checks that ensure cells are responding to
  #    status queries from kip.
  # 2) Cloud API healthchecks supplement status checks by also
  #    querying the cloud API to determine cell health based on the
  #    status of the cell's backing compute instance.

  # status is the default healthcheck mode and is the more robust of
  # the two methods. cloudAPI health checks are useful for
  # configurations where where kip running outside the cloud network
  # and connected to cells through the internet. Kip will query the
  # cloud provider's API to determine if a cell is running. cloudAPI
  # healthchecks provide less confidence the cells are functioning
  # correctly but allows cells to survive temporary network
  # connectivity issues between the kip controller and the cells.
  # Only one of status or cloudAPI healthcheck can be specified in the
  # healthcheck section. When cloudAPI health checks are used, status
  # queries are still used to check cell status.

  # healthcheck:
  #   status:
  #     healthyTimeout: 90
  #   cloudAPI:
  #     healthyTimeout: 180
  #     # interval between checks must be >= 10, default is 60
  #     interval: 60
```
