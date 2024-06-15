terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "2.39.1"
    }
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

