terraform {
  required_providers {
    digitalocean = {
      source = "digitalocean/digitalocean"
      version = "2.39.1"
    }
  }
}

provider "digitalocean" {
  token = var.api_token
}

resource "digitalocean_droplet" "web" {
  image = var.image
  name = "test-web-vm"
  region = "nyc1"
  size = "s-1vcpu-1gb"
  monitoring = true
  ipv6 = true
}

