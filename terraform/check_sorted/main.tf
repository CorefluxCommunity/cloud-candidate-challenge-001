terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

data "digitalocean_droplets" "all_droplets" {
  sort {
    key = "created_at"
    direction = var.direction
  }
}
