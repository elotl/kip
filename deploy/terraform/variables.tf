variable "ssh-key-name" {
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

variable "worker-disk-size" {
  default = 15
}

variable "blacklisted-azs" {
  type    = list(string)
  default = ["use1-az3"]
}

variable "worker-ami" {
  default = ""
}

variable "virtual-kubelet-manifest" {
  default = "../virtual-kubelet.yaml"
}
