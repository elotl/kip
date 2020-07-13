provider "aws" {
  region = var.region
}

locals {
  k8s_cluster_tags = {
    "Name"                                      = "${var.cluster-name}"
    "kubernetes.io/cluster/${var.cluster-name}" = "owned"
  }
}

data "aws_availability_zones" "available-azs" {
  state                = "available"
  exclude_zone_ids = var.blacklisted-azs
}

resource "random_shuffle" "azs" {
  input        = data.aws_availability_zones.available-azs.names
  result_count = 1
}

resource "aws_vpc" "main" {
  cidr_block           = var.vpc-cidr
  enable_dns_hostnames = true

  provisioner "local-exec" {
    when        = destroy
    command     = "./cleanup-vpc.sh ${self.id} ${var.cluster-name}"
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
    command     = "./cleanup-vpc.sh ${self.vpc_id} ${var.cluster-name}"
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
  cidr_block              = cidrsubnet(var.vpc-cidr, 4, 1)
  availability_zone       = element(random_shuffle.azs.result, 0)
  map_public_ip_on_launch = true

  tags = local.k8s_cluster_tags
}

resource "aws_route_table" "route-table" {
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

resource "aws_route_table_association" "route-table-to-subnet" {
  subnet_id      = aws_subnet.subnet.id
  route_table_id = aws_route_table.route-table.id
}

resource "aws_efs_file_system" "efs" {
  count = var.efs-enable ? 1 : 0
  performance_mode = var.efs-performance-mode
  provisioned_throughput_in_mibps = var.efs-provisioned-throughput-in-mibps
  throughput_mode = var.efs-throughput-mode
  tags = {
    "Name" = "${var.cluster-name}"
  }
}

resource "aws_efs_mount_target" "efs" {
  count = var.efs-enable ? 1 : 0
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
    cidr_blocks = [var.vpc-cidr]
  }

  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = [var.pod-cidr]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = local.k8s_cluster_tags
}

resource "aws_iam_role" "k8s-node" {
  name               = "k8s-node-${var.cluster-name}"
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

resource "aws_iam_role_policy" "k8s-node" {
  name = "k8s-node-${var.cluster-name}"
  role = aws_iam_role.k8s-node.id
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
        "ec2:AssignPrivateIpAddresses",
        "ec2:AttachNetworkInterface",
        "ec2:AttachVolume",
        "ec2:AuthorizeSecurityGroupIngress",
        "ec2:CreateNetworkInterface",
        "ec2:CreateRoute",
        "ec2:CreateSecurityGroup",
        "ec2:CreateTags",
        "ec2:CreateVolume",
        "ec2:DeleteNetworkInterface",
        "ec2:DeleteRoute",
        "ec2:DeleteSecurityGroup",
        "ec2:DeleteVolume",
        "ec2:DescribeAddresses",
        "ec2:DescribeAvailabilityZones",
        "ec2:DescribeDhcpOptions",
        "ec2:DescribeElasticGpus",
        "ec2:DescribeImages",
        "ec2:DescribeInstances",
        "ec2:DescribeNetworkInterfaces",
        "ec2:DescribeRegions",
        "ec2:DescribeRouteTables",
        "ec2:DescribeSecurityGroups",
        "ec2:DescribeSpotPriceHistory",
        "ec2:DescribeSubnets",
        "ec2:DescribeTags",
        "ec2:DescribeVolumes",
        "ec2:DescribeVpcAttribute",
        "ec2:DescribeVpcs",
        "ec2:DetachNetworkInterface",
        "ec2:DetachVolume",
        "ec2:ModifyInstanceAttribute",
        "ec2:ModifyInstanceCreditSpecification",
        "ec2:ModifyNetworkInterfaceAttribute",
        "ec2:ModifyVolume",
        "ec2:ModifyVpcAttribute",
        "ec2:RequestSpotInstances",
        "ec2:RevokeSecurityGroupIngress",
        "ec2:RunInstances",
        "ec2:TerminateInstances",
        "ec2:UnassignPrivateIpAddresses",
        "ecr:BatchCheckLayerAvailability",
        "ecr:BatchGetImage",
        "ecr:DescribeRepositories",
        "ecr:GetAuthorizationToken",
        "ecr:GetDownloadUrlForLayer",
        "ecr:GetRepositoryPolicy",
        "ecr:ListImages",
        "ecs:CreateCluster",
        "ecs:DeregisterTaskDefinition",
        "ecs:DescribeClusters",
        "ecs:DescribeTasks",
        "ecs:ListAccountSettings",
        "ecs:ListTaskDefinitions",
        "ecs:ListTasks",
        "ecs:PutAccountSetting",
        "ecs:RegisterTaskDefinition",
        "ecs:RunTask",
        "ecs:StopTask",
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

resource "aws_iam_instance_profile" "k8s-node" {
  name = "k8s-node-${var.cluster-name}"
  role = aws_iam_role.k8s-node.name
}

resource "random_id" "k8stoken-prefix" {
  byte_length = 3
}

resource "random_id" "k8stoken-suffix" {
  byte_length = 8
}

locals {
  k8stoken = format(
    "%s.%s",
    random_id.k8stoken-prefix.hex,
    random_id.k8stoken-suffix.hex,
  )
}

data "external" "manifest" {
  program = ["bash", "-c", "set -e; set -o pipefail; kustomize build ${var.kustomize-dir} | jq -s -R '{\"output\": .}'"]
}

data "template_file" "node-userdata" {
  template = file("node.sh")

  vars = {
    k8stoken                  = local.k8stoken
    k8s_version               = var.k8s-version
    pod_cidr                  = var.pod-cidr
    service_cidr              = var.service-cidr
    kip_manifest  = base64encode(data.external.manifest.result.output)
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
  node_ami = length(var.node-ami) > 0 ? var.node-ami : data.aws_ami.ubuntu.id
  create_ssh_key = length(var.ssh-key-name) > 0 ? false : true
  ssh_key_name = local.create_ssh_key ? aws_key_pair.ssh-key.0.key_name : var.ssh-key-name
}

resource "tls_private_key" "ssh-key" {
  count = local.create_ssh_key ? 1 : 0
  algorithm = "RSA"
}

resource "aws_key_pair" "ssh-key" {
  count = local.create_ssh_key ? 1 : 0
  key_name   = "ssh-key-${var.cluster-name}"
  public_key = tls_private_key.ssh-key.0.public_key_openssh

  tags = local.k8s_cluster_tags
}

resource "aws_instance" "k8s-node" {
  ami                         = local.node_ami
  instance_type               = "t3.medium"
  subnet_id                   = aws_subnet.subnet.id
  user_data                   = data.template_file.node-userdata.rendered
  key_name                    = local.ssh_key_name
  associate_public_ip_address = true
  vpc_security_group_ids      = [aws_security_group.kubernetes.id]
  iam_instance_profile        = aws_iam_instance_profile.k8s-node.id
  source_dest_check           = false

  root_block_device {
    volume_size = var.node-disk-size
  }

  provisioner "remote-exec" {
    connection {
      type        = "ssh"
      user        = "ubuntu"
      host        = self.public_ip
      private_key = local.create_ssh_key ? tls_private_key.ssh-key.0.private_key_pem : null
    }
    inline = [
      "timeout 600 bash -x -c 'echo Waiting for cluster to come up; while true; do kubectl cluster-info && kubectl get svc kubernetes && exit 0; sleep 5; done'",
    ]
  }

  depends_on = [aws_internet_gateway.gw, aws_key_pair.ssh-key]

  tags = merge(local.k8s_cluster_tags,
    {"Name" = "${var.cluster-name}-node"})
}
