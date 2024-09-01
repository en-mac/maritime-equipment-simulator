provider "aws" {
  region  = var.aws_region
  profile = "default"  # Use the default profile configured in AWS CLI
}
