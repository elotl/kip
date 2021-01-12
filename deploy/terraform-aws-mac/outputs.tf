output "node_ip" {
  value       = aws_instance.k8s_node.public_ip
  description = "Node IP address"
}

output "efs_ip" {
  value       = aws_efs_mount_target.efs.*.ip_address
  description = "EFS IP address"
}

output "ssh_key" {
  value       = length(tls_private_key.ssh_key) > 0 ? tls_private_key.ssh_key[0].private_key_pem : ""
  description = "SSH key"
  sensitive   = true
}
