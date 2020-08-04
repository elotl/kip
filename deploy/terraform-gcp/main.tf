terraform {
  required_version = ">= 0.12.0"
}

provider "null" {
  version =  "~> 2.1"
}

provider "random" {
  version =  "~> 2.2"
}

provider "google" {
  version = "~> 3.21"
  project = var.project
  region  = var.region
  zone    = var.zone
}

resource "random_id" "password" {
  byte_length = 16
}

locals {
  password = format("%s", random_id.password.hex)
}

resource "google_container_cluster" "cluster" {
  name = var.cluster-name

  # Create a zonal cluster. In production, you want a regional cluster with
  # multiple masters spread across zones in the region.
  location        = var.zone

  # We can't create a cluster with no node pool defined, but we want to only use
  # separately managed node pools. So we create the smallest possible default
  # node pool and immediately delete it.
  remove_default_node_pool = true
  initial_node_count       = 1

  master_auth {
    username = "client"
    password = local.password

    client_certificate_config {
      issue_client_certificate = true
    }
  }

  ip_allocation_policy {
    cluster_ipv4_cidr_block   = var.pod-cidr
    services_ipv4_cidr_block  = var.service-cidr
  }

  master_authorized_networks_config {
    cidr_blocks {
      # Limit this in production to your client IPs.
      cidr_block = "0.0.0.0/0"
    }
  }
}

resource "google_container_node_pool" "node-pool" {
  name       = "node-pool-${var.cluster-name}"
  location   = var.zone
  cluster    = google_container_cluster.cluster.name
  node_count = 1

  node_config {
    preemptible  = false
    machine_type = "n1-standard-1"

    metadata = {
      disable-legacy-endpoints = "true"
    }

    oauth_scopes = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
      "https://www.googleapis.com/auth/compute",
    ]
  }
}

resource "google_filestore_instance" "filestore" {
  count = var.filestore-enable ? 1 : 0

  name = var.cluster-name
  zone = var.zone
  tier = var.filestore-tier

  file_shares {
    name        = var.filestore-fileshare-name
    capacity_gb = var.filestore-fileshare-capacity-gb
  }

  networks {
    network = "default"
    modes   = ["MODE_IPV4"]
    reserved_ip_range = var.filestore-reserved-ip-range
  }
}

locals {
  kubeconfig = "${path.module}/kubeconfig"
}

resource "null_resource" "kubeconfig" {
  provisioner "local-exec" {
    interpreter = [
      "bash",
      "-x",
      "-c",
    ]
    command = <<-EOF
    echo 'apiVersion: v1' > ${local.kubeconfig}
    echo 'clusters:' >> ${local.kubeconfig}
    echo '- cluster:' >> ${local.kubeconfig}
    echo '    certificate-authority-data: ${google_container_cluster.cluster.master_auth.0.cluster_ca_certificate}' >> ${local.kubeconfig}
    echo '    server: https://${google_container_cluster.cluster.endpoint}' >> ${local.kubeconfig}
    echo '  name: kubernetes' >> ${local.kubeconfig}
    echo 'contexts:' >> ${local.kubeconfig}
    echo '- context:' >> ${local.kubeconfig}
    echo '    cluster: kubernetes' >> ${local.kubeconfig}
    echo '    user: kubernetes-admin' >> ${local.kubeconfig}
    echo '  name: kubernetes-admin@kubernetes' >> ${local.kubeconfig}
    echo 'current-context: kubernetes-admin@kubernetes' >> ${local.kubeconfig}
    echo 'kind: Config' >> ${local.kubeconfig}
    echo 'preferences: {}' >> ${local.kubeconfig}
    echo 'users:' >> ${local.kubeconfig}
    echo '- name: kubernetes-admin' >> ${local.kubeconfig}
    echo '  user:' >> ${local.kubeconfig}
    echo '    client-certificate-data: ${google_container_cluster.cluster.master_auth.0.client_certificate}' >> ${local.kubeconfig}
    echo '    client-key-data: ${google_container_cluster.cluster.master_auth.0.client_key}' >> ${local.kubeconfig}
    EOF
  }

  triggers = {
    cluster_instance_ids = google_container_cluster.cluster.id
  }

  depends_on = [
    google_container_cluster.cluster
  ]
}

resource "null_resource" "client_permissions" {
  provisioner "local-exec" {
    interpreter = [
      "bash",
      "-x",
      "-c",
    ]
    environment = {
      KUBECONFIG = "${path.module}/kubeconfig"
    }
    command = "kubectl apply -f cluster-admin.yaml --username=client --password=${local.password}"
  }

  triggers = {
    cluster_instance_ids = google_container_cluster.cluster.id
  }

  depends_on = [
    google_container_cluster.cluster,
    null_resource.kubeconfig
  ]
}

resource "null_resource" "deploy" {
  count = length(var.kustomize-dir) > 0 ? 1 : 0

  provisioner "local-exec" {
    environment = {
      KUBECONFIG = "${path.module}/kubeconfig"
    }
    interpreter = [
      "bash",
      "-x",
      "-c",
    ]
    # This needs kubectl and kustomize.
    command = "kustomize build ${var.kustomize-dir} | kubectl apply -f -"
  }

  triggers = {
    cluster_instance_ids = google_container_cluster.cluster.id
  }

  depends_on = [
    google_container_cluster.cluster,
    null_resource.client_permissions
  ]
}
