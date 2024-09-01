output "ecr_repository_url" {
  description = "ECR repository URL"
  value       = aws_ecr_repository.maritime_equipment_simulator_repo.repository_url
}

output "eks_cluster_name" {
  description = "EKS Cluster Name"
  value       = module.eks.cluster_id
}

output "eks_cluster_endpoint" {
  description = "EKS Cluster Endpoint"
  value       = module.eks.cluster_endpoint
}

output "rds_endpoint" {
  description = "RDS PostgreSQL Endpoint"
  value       = module.rds.this_db_instance_endpoint
}

output "rabbitmq_broker_endpoint" {
  description = "RabbitMQ Broker Endpoint"
  value       = aws_mq_broker.rabbitmq.primary_endpoint
}
