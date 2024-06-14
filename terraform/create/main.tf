terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

resource "digitalocean_droplet" "web" {
  name   = var.droplet_name
  region = var.region
  size   = var.size
  image  = var.image
  tags   = var.tags
}
