provider "aws" {
  region = var.aws_region
}

# ECR Repository
resource "aws_ecr_repository" "maritime_equipment_simulator_repo" {
  name = "maritime-equipment-simulator"
}

# VPC
module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "3.0.0"

  name = "maritime-equipment-simulator-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["us-west-2a", "us-west-2b", "us-west-2c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  enable_nat_gateway = true

  tags = {
    Terraform = "true"
    Environment = var.environment
  }
}

# EKS Cluster
module "eks" {
  source          = "terraform-aws-modules/eks/aws"
  version         = "17.24.0" # Updated to the latest stable version
  cluster_name    = "maritime-equipment-simulator-cluster"
  cluster_version = "1.20"

  vpc_id     = module.vpc.vpc_id
  subnet_ids = module.vpc.private_subnets

  node_groups = {
    maritime-equipment-simulator-workers = {
      desired_capacity = 2
      max_capacity     = 3
      min_capacity     = 1

      instance_type = "t3.medium"
      key_name      = var.key_name # Ensure this key pair exists in your AWS account
    }
  }

  tags = {
    Terraform = "true"
    Environment = var.environment
  }
}

# RDS PostgreSQL
module "rds" {
  source            = "terraform-aws-modules/rds/aws"
  version           = "5.0.2" # Updated to the latest stable version
  engine            = "postgres"
  instance_class    = "db.t3.micro"
  allocated_storage = 20
  name              = var.postgres_db_name
  username          = var.postgres_username
  password          = var.postgres_password

  vpc_security_group_ids = [module.vpc.default_security_group_id]
  subnet_ids             = module.vpc.private_subnets
  publicly_accessible    = false

  tags = {
    Terraform = "true"
    Environment = var.environment
  }
}

# RabbitMQ (Amazon MQ)
resource "aws_mq_broker" "rabbitmq" {
  broker_name       = "maritime-equipment-simulator-rabbitmq"
  engine_type       = "RABBITMQ"
  engine_version    = "3.8.19"
  host_instance_type = "mq.t3.micro"
  publicly_accessible = false

  user {
    username = var.rabbitmq_username
    password = var.rabbitmq_password
  }

  logs {
    general = true
  }

  tags = {
    Terraform = "true"
    Environment = var.environment
  }
}
