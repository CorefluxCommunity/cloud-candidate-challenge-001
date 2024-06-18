terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "2.39.1"
    }
  }
}

terraform {
  backend "s3" {
    bucket         = var.bucket_name
    key            = var.dynamodb_table
    dynamodb_table = "terraform-state-lock-dynamo"
    region         = var.aws_region
    encrypt        = true
  }
}

provider "digitalocean" {
  token = var.api_token
}

resource "digitalocean_droplet" "web" {
  image      = var.image
  name       = var.name
  region     = var.region
  size       = var.size
  monitoring = var.monitoring
  ipv6       = var.ipv6
}

