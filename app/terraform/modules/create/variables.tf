variable "do_token" {
  description = "DigitalOcean API token"
  type        = string
}

variable "droplet_name" {
  description = "Name of the Droplet"
  type        = string
  default     = "new"
}

variable "region" {
  description = "DigitalOcean region"
  type        = string
  default     = "nyc1"
}

variable "size" {
  description = "Droplet size"
  type        = string
  default     = "s-1vcpu-1gb"
}

variable "image" {
  description = "Droplet image"
  type        = string
  default     = "ubuntu-20-04-x64"
}

variable "tags" {
  description = "Droplet tags"
  type        = list(string)
  default     = ["tag"]
}