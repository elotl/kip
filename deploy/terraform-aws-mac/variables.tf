variable "ssh_key_name" {
  default = ""
}

variable "cluster_name" {
  default = "vk"
}

variable "region" {
  default = "us-east-1"
}

variable "vpc_cidr" {
  default = "10.0.0.0/16"
}

variable "pod_cidr" {
  default = "172.20.0.0/16"
}

variable "service_cidr" {
  default = "10.96.0.0/12"
}

variable "k8s_version" {
  default = ""
}

variable "node_disk_size" {
  default = 15
}

variable "excluded_azs" {
  type    = list(string)
  default = ["use1-az3"]
}

variable "node_ami" {
  default = ""
}

variable "kustomize_dir" {
  # TODO change after merging to master
  default = "github.com/elotl/kip/deploy/manifests/kip/overlays/mac?ref=af1a227407e29184987d3c73c12686521a465dd4"
}

variable "efs_enable" {
  type    = bool
  default = false
  description = "Create an EFS volume cells and/or pods in the VPC can mount and use."
}

variable "efs_performance_mode" {
  type        = string
  default     = "generalPurpose"
  description = "Optional. The file system performance mode. Can be either generalPurpose or maxIO."
}

variable "efs_provisioned_throughput_in_mibps" {
  type        = number
  default     = 0
  description = "Optional. The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with efs_throughput_mode set to provisioned."
}

variable "efs_throughput_mode" {
  type        = string
  default     = "bursting"
  description = "Optional. Throughput mode for the file system. Defaults to bursting. Valid values: bursting, provisioned. When using provisioned, also set efs_provisioned_throughput_in_mibps."
}
