terraform {
  required_version = "~> 0.12.0"
}

provider "external" {
  version = "~> 1.2"
}

provider "random" {
  version = "~> 2.3"
}

provider "template" {
  version = "~> 2.1"
}

provider "tls" {
  version = "~> 2.2"
}

provider "aws" {
  version = "~> 3.0"
  region  = var.region
}

locals {
  k8s_cluster_tags = {
    "Name"                                      = "${var.cluster_name}"
    "kubernetes.io/cluster/${var.cluster_name}" = "owned"
  }
}

data "aws_availability_zones" "available_azs" {
  state             = "available"
  exclude_zone_ids  = var.excluded_azs
}

resource "random_shuffle" "azs" {
  input        = data.aws_availability_zones.available_azs.names
  result_count = 1
}

resource "aws_vpc" "main" {
  cidr_block           = var.vpc_cidr
  enable_dns_hostnames = true

  provisioner "local-exec" {
    when        = destroy
    command     = "${path.module}/cleanup-vpc.sh ${self.id} ${var.cluster_name}"
    interpreter = ["/bin/bash", "-c"]
    environment = {
      "AWS_REGION"         = var.region
      "AWS_DEFAULT_REGION" = var.region
    }
  }

  tags = local.k8s_cluster_tags
}

resource "aws_internet_gateway" "gw" {
  vpc_id = aws_vpc.main.id

  provisioner "local-exec" {
    when        = destroy
    command     = "${path.module}/cleanup-vpc.sh ${self.vpc_id} ${var.cluster_name}"
    interpreter = ["/bin/bash", "-c"]
    environment = {
      "AWS_REGION"         = var.region
      "AWS_DEFAULT_REGION" = var.region
    }
  }

  tags = local.k8s_cluster_tags
}

resource "aws_subnet" "subnet" {
  vpc_id                  = aws_vpc.main.id
  cidr_block              = cidrsubnet(var.vpc_cidr, 4, 1)
  availability_zone       = element(random_shuffle.azs.result, 0)
  map_public_ip_on_launch = true

  tags = local.k8s_cluster_tags
}

resource "aws_route_table" "route_table" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gw.id
  }

  depends_on = [aws_internet_gateway.gw]

  tags = local.k8s_cluster_tags

  lifecycle {
    ignore_changes = [route]
  }
}

resource "aws_route_table_association" "route_table_to_subnet" {
  subnet_id      = aws_subnet.subnet.id
  route_table_id = aws_route_table.route_table.id
}

resource "aws_efs_file_system" "efs" {
  count = var.efs_enable ? 1 : 0
  performance_mode = var.efs_performance_mode
  provisioned_throughput_in_mibps = var.efs_provisioned_throughput_in_mibps
  throughput_mode = var.efs_throughput_mode
  tags = {
    "Name" = "${var.cluster_name}"
  }
}

resource "aws_efs_mount_target" "efs" {
  count = var.efs_enable ? 1 : 0
  file_system_id = aws_efs_file_system.efs[0].id
  subnet_id      = aws_subnet.subnet.id
  security_groups = [
    aws_security_group.kubernetes.id
  ]
}

resource "aws_security_group" "kubernetes" {
  name        = "kubernetes"
  description = "Main kubernetes security group"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = [var.vpc_cidr]
  }

  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = [var.pod_cidr]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = local.k8s_cluster_tags
}

resource "aws_iam_role" "k8s_node" {
  name               = "k8s-node-${var.cluster_name}"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF

  tags = local.k8s_cluster_tags
}

resource "aws_iam_policy" "k8s_node_k8s_policy" {
  name = "k8s-node-${var.cluster_name}-k8s"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "autoscaling:DescribeAutoScalingGroups",
        "autoscaling:DescribeLaunchConfigurations",
        "autoscaling:DescribeTags",
        "ec2:AttachVolume",
        "ec2:AuthorizeSecurityGroupIngress",
        "ec2:CreateRoute",
        "ec2:CreateSecurityGroup",
        "ec2:CreateTags",
        "ec2:CreateVolume",
        "ec2:DeleteRoute",
        "ec2:DeleteSecurityGroup",
        "ec2:DeleteVolume",
        "ec2:DescribeInstances",
        "ec2:DescribeInstances",
        "ec2:DescribeRegions",
        "ec2:DescribeRegions",
        "ec2:DescribeRouteTables",
        "ec2:DescribeSecurityGroups",
        "ec2:DescribeSubnets",
        "ec2:DescribeVolumes",
        "ec2:DescribeVpcs",
        "ec2:DetachVolume",
        "ec2:ModifyInstanceAttribute",
        "ec2:ModifyVolume",
        "ec2:RevokeSecurityGroupIngress",
        "ecr:BatchCheckLayerAvailability",
        "ecr:BatchGetImage",
        "ecr:DescribeRepositories",
        "ecr:GetAuthorizationToken",
        "ecr:GetDownloadUrlForLayer",
        "ecr:GetRepositoryPolicy",
        "ecr:ListImages",
        "ecs:ListTaskDefinitions",
        "ec2:DescribeElasticGpus",
        "ecs:DescribeClusters",
        "ec2:ModifyNetworkInterfaceAttribute",
        "ec2:ModifyVpcAttribute",
        "ecs:ListTasks",
        "ec2:AssignPrivateIpAddresses",
        "ecs:RunTask",
        "ecs:DeregisterTaskDefinition",
        "ecs:ListAccountSettings",
        "ecs:CreateCluster",
        "ec2:DescribeVpcAttribute",
        "ec2:DeleteNetworkInterface",
        "ec2:UnassignPrivateIpAddresses",
        "ecs:RegisterTaskDefinition",
        "ec2:DescribeAddresses",
        "ec2:DetachNetworkInterface",
        "ec2:DescribeSpotPriceHistory",
        "ec2:AttachNetworkInterface",
        "ec2:DescribeTags",
        "ecs:PutAccountSetting",
        "ecs:DescribeTasks",
        "ec2:CreateNetworkInterface",
        "ec2:RequestSpotInstances",
        "ecs:StopTask",
        "elasticloadbalancing:AddTags",
        "elasticloadbalancing:AddTags",
        "elasticloadbalancing:ApplySecurityGroupsToLoadBalancer",
        "elasticloadbalancing:AttachLoadBalancerToSubnets",
        "elasticloadbalancing:ConfigureHealthCheck",
        "elasticloadbalancing:CreateListener",
        "elasticloadbalancing:CreateLoadBalancer",
        "elasticloadbalancing:CreateLoadBalancerListeners",
        "elasticloadbalancing:CreateLoadBalancerPolicy",
        "elasticloadbalancing:CreateTargetGroup",
        "elasticloadbalancing:DeleteListener",
        "elasticloadbalancing:DeleteLoadBalancer",
        "elasticloadbalancing:DeleteLoadBalancerListeners",
        "elasticloadbalancing:DeleteTargetGroup",
        "elasticloadbalancing:DeregisterInstancesFromLoadBalancer",
        "elasticloadbalancing:DeregisterTargets",
        "elasticloadbalancing:DescribeListeners",
        "elasticloadbalancing:DescribeLoadBalancerAttributes",
        "elasticloadbalancing:DescribeLoadBalancerPolicies",
        "elasticloadbalancing:DescribeLoadBalancers",
        "elasticloadbalancing:DescribeTargetGroups",
        "elasticloadbalancing:DescribeTargetHealth",
        "elasticloadbalancing:DetachLoadBalancerFromSubnets",
        "elasticloadbalancing:ModifyListener",
        "elasticloadbalancing:ModifyLoadBalancerAttributes",
        "elasticloadbalancing:ModifyTargetGroup",
        "elasticloadbalancing:RegisterInstancesWithLoadBalancer",
        "elasticloadbalancing:RegisterTargets",
        "elasticloadbalancing:SetLoadBalancerPoliciesForBackendServer",
        "elasticloadbalancing:SetLoadBalancerPoliciesOfListener",
        "iam:CreateServiceLinkedRole",
        "kms:DescribeKey"
      ],
      "Resource": [
        "*"
      ]
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "k8s-policy-attachment" {
  policy_arn = aws_iam_policy.k8s_node_k8s_policy.arn
  role = aws_iam_role.k8s_node.name
}

resource "aws_iam_policy" "k8s_node_kip_policy" {
  name = "k8s-node-${var.cluster_name}-kip"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
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
        "ec2:RevokeSecurityGroupIngress",
        "ec2:RunInstances",
        "ec2:TerminateInstances",
        "ecr:BatchGetImage",
        "ecr:GetAuthorizationToken",
        "ecr:GetDownloadUrlForLayer",
        "iam:PassRole"
      ],
      "Resource": [
        "*"
      ]
    }
  ]
}
EOF
}


resource "aws_iam_role_policy_attachment" "kip-policy-attachment" {
  policy_arn = aws_iam_policy.k8s_node_kip_policy.arn
  role = aws_iam_role.k8s_node.name
}

resource "aws_iam_role_policy_attachment" "ssm-policy-attachment" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
  role = aws_iam_role.k8s_node.name
}

resource "aws_iam_instance_profile" "k8s_node" {
  name = "k8s-node-${var.cluster_name}"
  role = aws_iam_role.k8s_node.name
}

resource "random_id" "k8stoken_prefix" {
  byte_length = 3
}

resource "random_id" "k8stoken-suffix" {
  byte_length = 8
}

locals {
  k8stoken      = format(
    "%s.%s",
    random_id.k8stoken_prefix.hex,
    random_id.k8stoken-suffix.hex,
  )
  kustomize_dir = "%{ if substr(var.kustomize_dir, 0, 1) == "." }${path.module}/${var.kustomize_dir}%{ else }${var.kustomize_dir}%{ endif }"
  kip_manifest  = length(local.kustomize_dir) > 0 ? base64encode(data.external.manifest[0].result.output) : ""
}

data "external" "manifest" {
  count   = length(local.kustomize_dir) > 0 ? 1 : 0
  program = ["bash", "-c", "set -e; set -o pipefail; kustomize build ${local.kustomize_dir} | jq -s -R '{\"output\": .}'"]
}

data "template_file" "node_userdata" {
  template = file("${path.module}/node.sh")

  vars = {
    k8stoken      = local.k8stoken
    k8s_version   = var.k8s_version
    pod_cidr      = var.pod_cidr
    service_cidr  = var.service_cidr
    kip_manifest  = local.kip_manifest
  }
}

data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-bionic-18.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical.
}

locals {
  node_ami = length(var.node_ami) > 0 ? var.node_ami : data.aws_ami.ubuntu.id
  create_ssh_key = length(var.ssh_key_name) > 0 ? false : true
  ssh_key_name = local.create_ssh_key ? aws_key_pair.ssh_key.0.key_name : var.ssh_key_name
}

resource "tls_private_key" "ssh_key" {
  count = local.create_ssh_key ? 1 : 0
  algorithm = "RSA"
}

resource "aws_key_pair" "ssh_key" {
  count = local.create_ssh_key ? 1 : 0
  key_name   = "ssh-key-${var.cluster_name}"
  public_key = tls_private_key.ssh_key.0.public_key_openssh

  tags = local.k8s_cluster_tags
}

resource "aws_instance" "k8s_node" {
  ami                         = local.node_ami
  instance_type               = "t3.medium"
  subnet_id                   = aws_subnet.subnet.id
  user_data                   = data.template_file.node_userdata.rendered
  key_name                    = local.ssh_key_name
  associate_public_ip_address = true
  vpc_security_group_ids      = [aws_security_group.kubernetes.id]
  iam_instance_profile        = aws_iam_instance_profile.k8s_node.id
  source_dest_check           = false

  root_block_device {
    volume_size = var.node_disk_size
  }

  provisioner "remote-exec" {
    connection {
      type        = "ssh"
      user        = "ubuntu"
      host        = self.public_ip
      private_key = local.create_ssh_key ? tls_private_key.ssh_key.0.private_key_pem : null
    }
    inline = [
      "timeout 600 bash -x -c 'echo Waiting for cluster to come up; while true; do kubectl cluster-info && kubectl get svc kubernetes && exit 0; sleep 5; done'",
    ]
  }

  depends_on = [aws_internet_gateway.gw, aws_key_pair.ssh_key]

  tags = merge(local.k8s_cluster_tags,
    {"Name" = "${var.cluster_name}-node"})
}
