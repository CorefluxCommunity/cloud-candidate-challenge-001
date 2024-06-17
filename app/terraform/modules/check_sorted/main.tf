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

data "digitalocean_droplets" "all_droplets" {
  sort {
    key = "name"
    direction = var.direction
  }
}
