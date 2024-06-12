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
  image  = var.image
  name   = var.name
  region = var.region
  size   = var.size
  ssh_keys = var.ssh_keys
}

output "droplet_ip" {
  value = digitalocean_droplet.web.ipv4_address
}
