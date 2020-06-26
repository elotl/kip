variable "cluster-name" {
  default     = "vk"
  description = "A name for the cluster and its associated resources."
}

variable "project" {
  type        = string
  description = "The GCP project where resources will be created."
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
  default     = ""
  description = "The CIDR in the cluster used for pod IP addresses. By default, a CIDR will be allocated automatically."
}

variable "service-cidr" {
  default     = ""
  description = "The CIDR in the cluster used for service IP addresses. By default, a CIDR will be allocated automatically."
}

variable "kustomize-dir" {
  default     = "../manifests/kip/overlays/gcp"
  description = "A kustomization directory that will be applied once the cluster is created. Leave it empty to disable."
}

variable "filestore-enable" {
  type        = bool
  default     = false
  description = "Enable to create a filestore storage for this cluster. The filestore will contain one fileshare."
}

variable "filestore-tier" {
  type        = string
  default     = "STANDARD"
  description = "The service tier of the instance. Possible values are: TIER_UNSPECIFIED, STANDARD, PREMIUM."
}

variable "filestore-fileshare-capacity-gb" {
  type        = number
  default     = 1024
  description = "File share capacity in GiB. This must be at least 1024 GiB for the standard tier, or 2560 GiB for the premium tier."
}

variable "filestore-fileshare-name" {
  type        = string
  default     = "data"
  description = "The name of the file share to be created. It must be 16 characters or less."
}

variable "filestore-reserved-ip-range" {
  type        = string
  default     = "10.20.0.0/29"
  description = "A /29 CIDR block that identifies the range of IP addresses reserved for this filestore."
}
