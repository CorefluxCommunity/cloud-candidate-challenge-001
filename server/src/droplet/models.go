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

var VariablesModel = `
variable "api_token" {
	default = ""
	type    = string
  }
  
  variable "image" {
	default = ""
	type    = string
  }
  variable "name" {
	default = ""
	type    = string
  }
  variable "region" {
	default = ""
	type    = string
  }
  variable "size" {
	default = ""
	type    = string
  }
  variable "monitoring" {
	default = true
	type    = bool
  }
  variable "ipv6" {
	default = true
	type    = bool
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
ipv6 = %t
`
