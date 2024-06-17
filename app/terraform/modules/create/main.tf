terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

provider "digitalocean" {
  token = var.do_token
}

resource "digitalocean_droplet" "web" {
  count  = 3
  name   = "${var.droplet_name}-droplet-${count.index}"
  region = var.region
  size   = var.size
  image  = var.image
  tags   = concat(var.tags, [ "droplet${count.index}" ])
}