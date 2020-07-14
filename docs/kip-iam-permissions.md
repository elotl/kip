## Minimum IAM Permissions

The following policy covers all permissions Kip requires in order to run in AWS.

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "ec2",
            "Effect": "Allow",
            "Action": [
                "ec2:ModifyVolume",
                "ec2:AuthorizeSecurityGroupIngress",
                "ec2:DescribeInstances",
                "ec2:TerminateInstances",
                "ec2:CreateTags",
                "ec2:DeleteRoute",
                "ecr:GetAuthorizationToken",
                "ec2:DescribeDhcpOptions",
                "ec2:RunInstances",
                "ec2:DescribeSecurityGroups",
                "ec2:RevokeSecurityGroupIngress",
                "ec2:DescribeImages",
                "ec2:DescribeNetworkInterfaces",
                "ec2:DescribeAvailabilityZones",
                "ec2:ModifyInstanceCreditSpecification",
                "ec2:CreateRoute",
                "ec2:DescribeVpcs",
                "ec2:CreateSecurityGroup",
                "ec2:DescribeVolumes",
                "ec2:DeleteSecurityGroup",
                "ec2:ModifyInstanceAttribute",
                "ec2:DescribeSubnets",
                "ec2:DescribeRouteTables",
                "ec2:AssociateIamInstanceProfile",
                "iam:PassRole"
            ],
            "Resource": "*"
        }
    ]
}
```
