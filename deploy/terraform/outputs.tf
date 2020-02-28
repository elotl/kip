output "node-ip" {
  value = aws_instance.k8s-node.public_ip
}
