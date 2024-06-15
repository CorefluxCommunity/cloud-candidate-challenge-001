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
    bucket         = "cloud-challenge-s3-state"
    key            = "states/terraform.tfstate"
    dynamodb_table = "terraform-state-lock-dynamo"
    region         = "eu-west-2"
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

