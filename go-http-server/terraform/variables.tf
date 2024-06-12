variable "do_token" {
  description = "The DigitalOcean API token."
  type        = string
}

variable "image" {
  description = "The image ID or slug to use for the Droplet."
  type        = string
  default     = "ubuntu-20-04-x64"
}

variable "name" {
  description = "The name of the Droplet."
  type        = string
  default     = "example-droplet"
}

variable "region" {
  description = "The region to create the Droplet in."
  type        = string
  default     = "nyc3"
}

variable "size" {
  description = "The size of the Droplet."
  type        = string
  default     = "s-1vcpu-1gb"
}

variable "ssh_keys" {
  description = "A list of SSH key IDs or fingerprints to enable in the format."
  type        = list(string)
  default     = []
}

variable "ipv6" {
  description = "Enable IPv6"
  type        = bool
  default     = false
}

variable "monitoring" {
  description = "Enable monitoring"
  type        = bool
  default     = false
}

variable "vpc_uuid" {
  description = "The VPC UUID to create the Droplet in."
  type        = string
  default     = ""
}
