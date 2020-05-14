variable "cluster-name" {
  default     = "vk"
  description = "A name for the cluster and its associated resources."
}

variable "region" {
  default     = "us-west1"
  description = "The region to create resources in."
}

variable "zone" {
  default     = "us-west1-a"
  description = "The zone inside the region to create resources in."
}

variable "pod-cidr" {
  default     = "10.20.0.0/16"
  description = "The CIDR in the cluster used for pod IP addresses."
}

variable "service-cidr" {
  default     = "10.96.0.0/16"
  description = "The CIDR in the cluster used for service IP addresses."
}

variable "kustomize-dir" {
  default     = "../manifests/virtual-kubelet/base"
  description = "A kustomization directory that will be applied once the cluster is created. Leave it empty to disable."
}
