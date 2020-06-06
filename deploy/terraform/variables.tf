variable "ssh-key-name" {
  default = ""
}

variable "cluster-name" {
  default = "vk"
}

variable "region" {
  default = "us-east-1"
}

variable "vpc-cidr" {
  default = "10.0.0.0/16"
}

variable "pod-cidr" {
  default = "172.20.0.0/16"
}

variable "service-cidr" {
  default = "10.96.0.0/12"
}

variable "k8s-version" {
  default = ""
}

variable "node-disk-size" {
  default = 15
}

variable "blacklisted-azs" {
  type    = list(string)
  default = ["use1-az3"]
}

variable "node-ami" {
  default = ""
}

variable "kustomize-dir" {
  default = "../manifests/kip/base"
}

variable "efs-enable" {
  type    = bool
  default = false
  description = "Create an EFS volume cells and/or pods in the VPC can mount and use."
}

variable "efs-performance-mode" {
  type        = string
  default     = "generalPurpose"
  description = "Optional. The file system performance mode. Can be either generalPurpose or maxIO."
}

variable "efs-provisioned-throughput-in-mibps" {
  type        = number
  default     = 0
  description = "Optional. The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with efs-throughput-mode set to provisioned."
}

variable "efs-throughput-mode" {
  type        = string
  default     = "bursting"
  description = "Optional. Throughput mode for the file system. Defaults to bursting. Valid values: bursting, provisioned. When using provisioned, also set efs-provisioned-throughput-in-mibps."
}
