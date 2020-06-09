provider "aws" {
  region = var.region
}

data "http" "client_ip" {
  url = "http://ipv4.icanhazip.com"
}

locals {
  client_ip           = length(var.client_ip) > 0 ? var.client_ip : chomp(data.http.client_ip.body)
  client_cidr         = "${local.client_ip}/32"
  extra_cidrs         = concat([local.client_cidr], var.local_cidrs)
  vpc_public_subnets  = [cidrsubnet(var.vpc_cidr, 4, 1), cidrsubnet(var.vpc_cidr, 4, 2)]
}

module "vpn_gateway" {
  source = "terraform-aws-modules/vpn-gateway/aws"
  version = "~> 2.0"

  create_vpn_connection                     = true

  vpn_connection_static_routes_only         = var.static_routes_only
  vpn_connection_static_routes_destinations = var.local_cidrs

  vpn_gateway_id      = module.vpc.vgw_id
  customer_gateway_id = length(module.vpc.cgw_ids) > 0 ? module.vpc.cgw_ids[0] : ""

  vpc_id                       = module.vpc.vpc_id
  vpc_subnet_route_table_ids   = module.vpc.public_route_table_ids
  vpc_subnet_route_table_count = length(module.vpc.public_subnets)

  tunnel1_inside_cidr   = var.tunnel1_inside_cidr
  tunnel2_inside_cidr   = var.tunnel2_inside_cidr
  tunnel1_preshared_key = var.tunnel1_psk
  tunnel2_preshared_key = var.tunnel2_psk
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 2.0"

  name = var.name
  cidr = var.vpc_cidr
  azs  = var.azs
  public_subnets  = local.vpc_public_subnets
  enable_nat_gateway = false
  enable_vpn_gateway = true
  amazon_side_asn    = var.amazon_side_asn

  customer_gateways = {
    cgw = {
      bgp_asn    = var.bgp_asn
      ip_address = local.client_ip
      type       = "ipsec.1"
      tags       = "Name=${var.name}"
    }
  }

  tags = {
    Name        = var.name
  }
}

resource "local_file" "vpn-deployment-yaml" {
  content         = templatefile("${path.module}/kustomization/vpn-deployment.yaml.tmpl", {
    dnspolicy=var.vpn_hostnetwork ? "ClusterFirstWithHostNet" : "ClusterFirst",
    hostnetwork=var.vpn_hostnetwork ? "true" : "false",
    enable_bgp_agent=var.enable_bgp_agent,
  })
  filename        = "${path.module}/kustomization/vpn-deployment.yaml"
  file_permission = "0644"
}

resource "local_file" "kustomization-yaml" {
  sensitive_content = templatefile("${path.module}/kustomization/kustomization.yaml.tmpl", {
    aws_access_key_id=var.aws_access_key_id,
    aws_secret_access_key=var.aws_secret_access_key,
  })
  filename          = "${path.module}/kustomization/kustomization.yaml"
  file_permission   = "0600"
}

resource "local_file" "aws-vpn-client-env" {
  sensitive_content = templatefile("${path.module}/kustomization/aws-vpn-client.env.tmpl", {
    tunnel1_address=module.vpn_gateway.vpn_connection_tunnel1_address,
    tunnel1_cgw_inside_address=module.vpn_gateway.vpn_connection_tunnel1_cgw_inside_address,
    tunnel1_vgw_inside_address=module.vpn_gateway.vpn_connection_tunnel1_vgw_inside_address,
    tunnel1_psk=var.tunnel1_psk,
    tunnel2_address=module.vpn_gateway.vpn_connection_tunnel2_address,
    tunnel2_cgw_inside_address=module.vpn_gateway.vpn_connection_tunnel2_cgw_inside_address,
    tunnel2_vgw_inside_address=module.vpn_gateway.vpn_connection_tunnel2_vgw_inside_address,
    tunnel2_psk=var.tunnel2_psk,
    vpc_cidr=var.static_routes_only ? var.vpc_cidr : "",
    bgp_asn=var.bgp_asn,
    k8s_asn=var.k8s_asn,
    amazon_side_asn=var.amazon_side_asn,
    k8s_bgp_peer_ips=var.k8s_bgp_peer_ips,
    k8s_bgp_dynamic_neighbor_prefix=var.k8s_bgp_dynamic_neighbor_prefix,
  })
  filename          = "${path.module}/kustomization/aws-vpn-client.env"
  file_permission   = "0600"
}

resource "local_file" "provider-yaml" {
  sensitive_content = templatefile("${path.module}/kustomization/provider.yaml.tmpl", {
    name=var.name,
    region=var.region,
    vpc_id=module.vpc.vpc_id,
    subnet_id=module.vpc.public_subnets[0],
    extra_cidrs=local.extra_cidrs,
  })
  filename          = "${path.module}/kustomization/provider.yaml"
  file_permission   = "0600"
}

resource "local_file" "command-extra-args-yaml" {
  content = templatefile("${path.module}/kustomization/command-extra-args.yaml.tmpl", {
    local_dns=var.local_dns,
  })
  filename          = "${path.module}/kustomization/command-extra-args.yaml"
  file_permission   = "0644"
}

resource "local_file" "node-local-dns-yaml" {
  content = templatefile("${path.module}/kustomization/node-local-dns.yaml.tmpl", {
    cluster_domain=var.cluster_domain,
    local_dns=var.local_dns,
    kube_dns=var.kube_dns,
  })
  filename          = "${path.module}/kustomization/node-local-dns.yaml"
  file_permission   = "0644"
}

resource "null_resource" "deploy" {
  count = var.deploy_to_kubernetes ? 1 : 0

  provisioner "local-exec" {
    # This needs kustomize and kubectl.
    command = "kustomize build ${path.module}/kustomization/ | kubectl apply -f -"
    interpreter = [
      "bash",
      "-x",
      "-c",
    ]
  }

  triggers = {
    vpn-deployment-yaml=local_file.vpn-deployment-yaml.content,
    node-local-dns-yaml=local_file.node-local-dns-yaml.content
    aws-vpn-client-env=local_file.aws-vpn-client-env.sensitive_content,
    kustomization-yaml=local_file.kustomization-yaml.sensitive_content,
    command-extra-args-yaml=local_file.command-extra-args-yaml.content,
    provider-yaml=local_file.provider-yaml.sensitive_content,
  }

  depends_on = [
    local_file.vpn-deployment-yaml,
    local_file.node-local-dns-yaml,
    local_file.aws-vpn-client-env,
    local_file.kustomization-yaml,
    local_file.command-extra-args-yaml,
    local_file.provider-yaml,
  ]
}
