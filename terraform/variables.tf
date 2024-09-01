variable "aws_region" {
  description = "The AWS region to deploy resources"
  default     = "us-west-2"
}

variable "environment" {
  description = "The environment name (dev, prod, etc.)"
  default     = "dev"
}

variable "postgres_db_name" {
  description = "The name of the PostgreSQL database"
}

variable "postgres_username" {
  description = "The username for the PostgreSQL database"
}

variable "postgres_password" {
  description = "The password for the PostgreSQL database"
  sensitive   = true
}

variable "rabbitmq_username" {
  description = "The username for RabbitMQ"
}

variable "rabbitmq_password" {
  description = "The password for RabbitMQ"
  sensitive   = true
}

variable "key_name" {
  description = "The name of the SSH key pair to use for EKS worker nodes"
}
