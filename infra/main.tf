terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "2.39.1"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}
provider "aws" {
  region     = var.AWS_REGION
  access_key = var.AWS_ACCESS_KEY_ID
  secret_key = var.AWS_SECRET_ACCESS_KEY
}

provider "digitalocean" {
  token = var.DIGITALOCEAN_API_TOKEN
}

module "aws_dynamodb_table" {
  source         = "./modules/aws/modules/dynamodb"
  dynamodb_table = var.dynamodb_table
}
module "aws_s3_bucket" {
  source      = "./modules/aws/modules/s3"
  bucket_name = var.bucket_name
}

module "digitalocean_app" {
  source                 = "./modules/digitalocean"
  AWS_ACCESS_KEY_ID      = var.AWS_ACCESS_KEY_ID
  AWS_SECRET_ACCESS_KEY  = var.AWS_SECRET_ACCESS_KEY
  AWS_REGION             = var.AWS_REGION
  COGNITO_ISSUER         = var.COGNITO_ISSUER
  JWK_URL                = var.JWK_URL
  DIGITALOCEAN_API_TOKEN = var.DIGITALOCEAN_API_TOKEN
}

