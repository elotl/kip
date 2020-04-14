## Pod Annotations

Kip supports a number of annotations to customize the cloud instance that pods are launched onto.

**pod.elotl.co/instance-type**


Kip will attempt to run a pod on the cheapest cloud instance that satisfies the resouce requests and limits specified by the pod.  Use the `instance-type` annotation to specify a specific cloud instance type the pod will run on.  The annotation will override requested container limits.

```yaml
annotations:
  pod.elotl.co/instance-type: c5.large
```

**pod.elotl.co/launch-type**

Kip will run pods using OnDemand instances.  Set the `launch-type` annotation to `Spot` to run a pod on a spot instance

```yaml
annotations:
  pod.elotl.co/launch-type: spot
```

** pod.elotl.co/private-ip-only

In AWS, Kip cells will get a public IP address if the cell is run in a subnet connected to an internet gateway.  Set the private-ip-only annotation to "true" instruct kip to run the pod on a cell without a public IP address.  The subnet must be configured to allow downloading the itzo binary and container images without a public IP on the cell.

```yaml
annotations:
  pod.elotl.co/private-ip-only: 'true'
```

**pod.elotl.co/security-groups**

Use the `security-groups` annotation to set one or more security groups on the cloud instance the pod is running on.  If multiple security groups are specified, they should be separated by a comma.  Each cloud instance started by Kip cell has one security group assigned to it by the Kip controller.  In most AWS accounts, instances are limited to 5 security groups.  In those setups, that would leave room for 4 more security groups to be assigned to the cloud instance.

```yaml
annotations:
  pod.elotl.co/security-groups: sg-0011a33dcc0da8151, sg-0026179f4abedb34a
```

**pod.elotl.co/instance-profile**

AWS Instance profiles can be attached to the instances backing Kip cells.  At this time, instance profiles must be specified by using the full ARN of the profile

```yaml
annotations:
  pod.elotl.co/instance-profile: "arn:aws:iam::11123456789:instance-profile/kip-s3-full-access-role"
```