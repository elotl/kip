output "master_ip" {
  value = aws_instance.k8s-master.public_ip
}

output "worker_ip" {
  value = aws_instance.k8s-worker.public_ip
}
