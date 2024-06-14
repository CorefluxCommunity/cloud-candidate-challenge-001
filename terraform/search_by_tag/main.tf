terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

data "digitalocean_droplets" "all_droplets" {
  filter {
    key = "tags"
    values = var.tag_to_find
  }
}
