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
  blacklisted_zone_ids = var.blacklisted-azs
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
}

resource "aws_internet_gateway" "gw" {
  vpc_id = aws_vpc.main.id

  provisioner "local-exec" {
    when        = destroy
    command     = "./cleanup-vpc.sh ${self.id} ${var.cluster-name}"
    interpreter = ["/bin/bash", "-c"]
    environment = {
      "AWS_REGION"         = var.region
      "AWS_DEFAULT_REGION" = var.region
    }
  }
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

resource "aws_iam_role" "k8s-master" {
  name               = "k8s-master-${var.cluster-name}"
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
}

resource "aws_iam_role_policy" "k8s-master" {
  name = "k8s-master-${var.cluster-name}"
  role = aws_iam_role.k8s-master.id
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
        "ec2:DescribeInstances",
        "ec2:DescribeNetworkInterfaces",
        "ec2:DescribeRegions",
        "ec2:DescribeRouteTables",
        "ec2:DescribeSecurityGroups",
        "ec2:DescribeSubnets",
        "ec2:DescribeTags",
        "ec2:DescribeVolumes",
        "ec2:DescribeVpcs",
        "ec2:DetachNetworkInterface",
        "ec2:DetachVolume",
        "ec2:ModifyInstanceAttribute",
        "ec2:ModifyNetworkInterfaceAttribute",
        "ec2:ModifyVolume",
        "ec2:RevokeSecurityGroupIngress",
        "ec2:UnassignPrivateIpAddresses",
        "ecr:GetAuthorizationToken",
        "ecr:BatchCheckLayerAvailability",
        "ecr:GetDownloadUrlForLayer",
        "ecr:GetRepositoryPolicy",
        "ecr:DescribeRepositories",
        "ecr:ListImages",
        "ecr:BatchGetImage",
        "elasticloadbalancing:AddTags",
        "elasticloadbalancing:AttachLoadBalancerToSubnets",
        "elasticloadbalancing:ApplySecurityGroupsToLoadBalancer",
        "elasticloadbalancing:CreateLoadBalancer",
        "elasticloadbalancing:CreateLoadBalancerPolicy",
        "elasticloadbalancing:CreateLoadBalancerListeners",
        "elasticloadbalancing:ConfigureHealthCheck",
        "elasticloadbalancing:DeleteLoadBalancer",
        "elasticloadbalancing:DeleteLoadBalancerListeners",
        "elasticloadbalancing:DescribeLoadBalancers",
        "elasticloadbalancing:DescribeLoadBalancerAttributes",
        "elasticloadbalancing:DetachLoadBalancerFromSubnets",
        "elasticloadbalancing:DeregisterInstancesFromLoadBalancer",
        "elasticloadbalancing:ModifyLoadBalancerAttributes",
        "elasticloadbalancing:RegisterInstancesWithLoadBalancer",
        "elasticloadbalancing:SetLoadBalancerPoliciesForBackendServer",
        "elasticloadbalancing:AddTags",
        "elasticloadbalancing:CreateListener",
        "elasticloadbalancing:CreateTargetGroup",
        "elasticloadbalancing:DeleteListener",
        "elasticloadbalancing:DeleteTargetGroup",
        "elasticloadbalancing:DescribeListeners",
        "elasticloadbalancing:DescribeLoadBalancerPolicies",
        "elasticloadbalancing:DescribeTargetGroups",
        "elasticloadbalancing:DescribeTargetHealth",
        "elasticloadbalancing:ModifyListener",
        "elasticloadbalancing:ModifyTargetGroup",
        "elasticloadbalancing:RegisterTargets",
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

resource "aws_iam_instance_profile" "k8s-master" {
  name = "k8s-master-${var.cluster-name}"
  role = aws_iam_role.k8s-master.name
}

resource "aws_iam_role" "k8s-worker" {
  name               = "k8s-worker-${var.cluster-name}"
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
}

resource "aws_iam_role_policy" "k8s-worker" {
  name = "k8s-worker-${var.cluster-name}"
  role = aws_iam_role.k8s-worker.id
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "ec2",
      "Effect": "Allow",
      "Action": [
        "ec2:AssignPrivateIpAddresses",
        "ec2:AttachNetworkInterface",
        "ec2:AuthorizeSecurityGroupIngress",
        "ec2:CreateNetworkInterface",
        "ec2:CreateRoute",
        "ec2:CreateSecurityGroup",
        "ec2:CreateTags",
        "ec2:CreateVolume",
        "ec2:DeleteNetworkInterface",
        "ec2:DeleteRoute",
        "ec2:DeleteSecurityGroup",
        "ec2:DescribeAddresses",
        "ec2:DescribeElasticGpus",
        "ec2:DescribeImages",
        "ec2:DescribeInstances",
        "ec2:DescribeNetworkInterfaces",
        "ec2:DescribeRouteTables",
        "ec2:DescribeSecurityGroups",
        "ec2:DescribeSpotPriceHistory",
        "ec2:DescribeSubnets",
        "ec2:DescribeTags",
        "ec2:DescribeVolumes",
        "ec2:DescribeVpcAttribute",
        "ec2:DescribeVpcs",
        "ec2:DetachNetworkInterface",
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
        "ecr:BatchGetImage",
        "ecr:GetAuthorizationToken",
        "ecr:GetDownloadUrlForLayer",
        "ecr:GetDownloadUrlForLayer",
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
        "ecs:StopTask"
      ],
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_instance_profile" "k8s-worker" {
  name = "k8s-worker-${var.cluster-name}"
  role = aws_iam_role.k8s-worker.name
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

data "template_file" "master-userdata" {
  template = file("master.sh")

  vars = {
    k8stoken                  = local.k8stoken
    k8s_version               = var.k8s-version
    pod_cidr                  = var.pod-cidr
    service_cidr              = var.service-cidr
    virtual_kubelet_manifest  = base64encode(file(var.virtual-kubelet-manifest))
  }
}

data "template_file" "worker-userdata" {
  template = file("worker.sh")

  vars = {
    k8stoken        = local.k8stoken
    k8s_version     = var.k8s-version
    masterIP        = aws_instance.k8s-master.private_ip
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

resource "aws_instance" "k8s-master" {
  ami                         = data.aws_ami.ubuntu.id
  instance_type               = "t3.medium"
  subnet_id                   = aws_subnet.subnet.id
  user_data                   = data.template_file.master-userdata.rendered
  key_name                    = var.ssh-key-name
  associate_public_ip_address = true
  vpc_security_group_ids      = [aws_security_group.kubernetes.id]
  iam_instance_profile        = aws_iam_instance_profile.k8s-master.id
  source_dest_check           = false

  depends_on = [aws_internet_gateway.gw]

  tags = merge(local.k8s_cluster_tags,
    {"Name" = "${var.cluster-name}-master"})
}

locals {
  worker_ami = length(var.worker-ami) > 0 ? var.worker-ami : data.aws_ami.ubuntu.id
}

resource "aws_instance" "k8s-worker" {
  ami                         = local.worker_ami
  instance_type               = "t3.medium"
  subnet_id                   = aws_subnet.subnet.id
  user_data                   = data.template_file.worker-userdata.rendered
  key_name                    = var.ssh-key-name
  associate_public_ip_address = true
  vpc_security_group_ids      = [aws_security_group.kubernetes.id]
  iam_instance_profile        = aws_iam_instance_profile.k8s-worker.id
  source_dest_check           = false

  root_block_device {
    volume_size = var.worker-disk-size
  }

  depends_on = [aws_internet_gateway.gw]

  tags = merge(local.k8s_cluster_tags,
    {"Name" = "${var.cluster-name}-worker"})
}
