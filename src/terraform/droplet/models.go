package droplet

var MainModel = `
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
    
`
var OutputModel = `
output "droplet_ip_address" {
	value = digitalocean_droplet.web.ipv4_address
  }
  
  output "droplet_id"{
	  value = digitalocean_droplet.web.id
  }
`
var TfvarsModel = `
api_token  = "%s"
image      = "%s"
name       = "%s"
region     = "%s"
size       = "%s"
monitoring = %t
ipv6       = %t
`
