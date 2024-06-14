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

module "create_droplet" {
  source = "./create"
  droplet_name = "example-droplet1"
  region = "nyc3"
  size = "s-1vcpu-1gb"
  image = "ubuntu-20-04-x64"
  tags = ["default"]
}

module "search_by_tag" {
  source       = "./search_by_tag"
  tag_to_find  = ["default"]
}

module "check_sorted" {
  source = "./check_sorted"
  direction = "asc"
}
