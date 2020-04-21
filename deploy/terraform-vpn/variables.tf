variable "region" {
  type        = string
  default     = "us-east-1"
  description = "The AWS region to use."
}

variable "aws_access_key_id" {
  type        = string
  description = "The AWS access key id Kip will use for creating cells."
}

variable "aws_secret_access_key" {
  type        = string
  description = "The AWS secret access key Kip will use for creating cells."
}

variable "name" {
  type        = string
  default     = "cloud-burst"
  description = "A name that will be used to tag AWS resources."
}

variable "client_ip" {
  type        = string
  default     = ""
  description = "The VPN connection needs a source IP. If left empty, it will be auto-detected."
}

variable "vpc_cidr" {
  type        = string
  default     = "10.10.0.0/16"
  description = "The CIDR to use for the VPC."
}

variable "azs" {
  type        = list(string)
  default     = ["us-east-1a", "us-east-1b", "us-east-1c"]
  description = "Availability zones used for subnets in the VPC."
}

variable "local_cidrs" {
  type        = list(string)
  default     = ["192.168.0.0/16", "172.16.0.0/12", "10.0.2.0/24", "10.244.0.0/16"]
  description = "These CIDRs will be routed back from the VPC via the VPN connection."
}

variable "tunnel1_inside_cidr" {
  type        = string
  default     = "169.254.10.20/30"
  description = "A link-local /30 CIDR that will be used for the first VPN tunnel."
}

variable "tunnel2_inside_cidr" {
  type        = string
  default     = "169.254.30.40/30"
  description = "A link-local /30 CIDR that will be used for the second VPN tunnel."
}

variable "tunnel1_psk" {
  type        = string
  description = "The pre-shared key for the first VPN tunnel."
}

variable "tunnel2_psk" {
  type        = string
  description = "The pre-shared key for the second VPN tunnel."
}

variable "deploy_to_kubernetes" {
  type        = bool
  default     = true
  description = "Whether the generated Kubernetes resources will be applied via kubectl. Disable if you only need the kustomization/ directory generated, and you plan to apply it separately. If enabled, it needs kubectl >= 1.14."
}

variable "static_routes_only" {
  type        = bool
  default     = true
  description = "If true, the CIDRs set in local_cidrs will be propagated into the VPC subnet route table. Enable if you plan NOT to use BGP over the tunnel, and static routes to the local cluster are all that is needed. Disable if BGP is used via the VPN tunnel. You can still use BGP inside your Kubernetes cluster, and peer with the VPN pod even if you enable static routes via the VPN tunnel (see k8s_asn below)."
}

variable "amazon_side_asn" {
  type        = number
  default     = 64620
  description = "The ASN for the AWS side, if BGP over the VPN connection is used. This is set in the VPN gateway as 'AmazonSideAsn'."
}

variable "bgp_asn" {
  type        = number
  default     = 65220
  description = "The ASN for the VPN client side, if BGP over the VPN connection is used. This is set in the AWS customer gateway as 'BgpAsn'."
}

variable "k8s_asn" {
  type        = number
  default     = 64512
  description = "If BGP is used for route distribution in your Kubernetes cluster, set this to the ASN used in the cluster."
}

variable "k8s_bgp_peer_ips" {
  type        = string
  default     = ""
  description = "If you would like the BGP agent to actively initiate connections to peers in Kubernetes, set the whitespace-separated list of peer IPs here. Example: \"172.17.0.2 172.17.0.3\"."
}

variable "k8s_bgp_dynamic_neighbor_prefix" {
  type        = string
  default     = ""
  description = "Kubernetes BGP peers from this CIDR will be accepted. Set it to the prefix where the Kubernetes BGP agents can their IP addresses from. Example: 172.17.0.0/16."
}

variable "enable_bgp_agent" {
  type        = bool
  default     = false
  description = "Whether to run a BGP agent in the VPN client pod. If it is enabled, it will set up BGP via the VPN tunnel using bgp_asn for its own ASN and amazon_side_asn for its own ASN. It will also enable BGP connection from the Kubernetes cluster, expecting k8s_asn to be used as the ASN in cluster. You can also use an external BGP agent, and enable host network mode for the VPN client pod instead."
}

variable "vpn_hostnetwork" {
  type        = bool
  default     = true
  description = "Whether to run the VPN client pod in host network mode, thus directly adding routes on the worker node host (useful for a simple static route setup). If you plan on using BGP with the VPN tunnel, set it to false, and instead add the VPN client pod as a BGP peer."
}
