# Minimum IAM permissions
Permissions required for using SSM store:
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "",
            "Effect": "Deny",
            "Action": "ssm:GetParametersByPath",
            "Resource": "arn:aws:ssm:us-west-2:689494258501:parameter/kip/cells/*"
        },
        {
            "Sid": "",
            "Effect": "Allow",
            "Action": [
                "ssm:GetParameters",
                "ssm:GetParameter"
            ],
            "Resource": "arn:aws:ssm:us-west-2:689494258501:parameter/kip/cells/*",
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
            "Resource": "arn:aws:ssm:us-west-2:689494258501:parameter/kip/cells/*",
            "Condition": {
                "StringNotEquals": {
                    "aws:ResourceTag/AWSUserID": "${aws:userid}"
                }
            }
        }
    ]
}
```
and minimum KIP permissions:
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "ec2",
            "Effect": "Allow",
            "Action": [
                "ec2:AssociateIamInstanceProfile",
                "ec2:AuthorizeSecurityGroupIngress",
                "ec2:CreateRoute",
                "ec2:CreateSecurityGroup",
                "ec2:CreateTags",
                "ec2:DeleteRoute",
                "ec2:DeleteSecurityGroup",
                "ec2:DescribeAvailabilityZones",
                "ec2:DescribeDhcpOptions",
                "ec2:DescribeIamInstanceProfileAssociations",
                "ec2:DescribeImages",
                "ec2:DescribeInstances",
                "ec2:DescribeNetworkInterfaces",
                "ec2:DescribeRouteTables",
                "ec2:DescribeSecurityGroups",
                "ec2:DescribeSubnets",
                "ec2:DescribeVolumes",
                "ec2:DescribeVolumesModifications",
                "ec2:DescribeVpcs",
                "ec2:ModifyInstanceAttribute",
                "ec2:ModifyInstanceCreditSpecification",
                "ec2:ModifyVolume",
                "ec2:ReplaceIamInstanceProfileAssociation",
                "ec2:RevokeSecurityGroupIngress",
                "ec2:RunInstances",
                "ec2:TerminateInstances",
                "ecr:BatchGetImage",
                "ecr:GetAuthorizationToken",
                "ecr:GetDownloadUrlForLayer",
                "iam:GetInstanceProfile",
                "iam:PassRole",
                "ssm:AddTagsToResource",
                "ssm:DeleteParameters",
                "ssm:GetParameter",
                "ssm:GetParameters",
                "ssm:GetParametersByPath",
                "ssm:PutParameter"
            ],
            "Resource": "*"
        }
    ]
}
```
# Recommended IAM permissions with debugging capabilities
Two above and `AmazonSSMManagedInstanceCore`.
