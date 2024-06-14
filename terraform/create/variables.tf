variable "droplet_name" {
  description = "Name of the Droplet"
  default     = "new_droplet"
}

variable "region" {
  description = "DigitalOcean region"
  default     = "nyc1"
}

variable "size" {
  description = "Droplet size"
  default     = "s-1vcpu-1gb"
}

variable "image" {
  description = "Droplet image"
  default     = "ubuntu-20-04-x64"
}

variable "tags" {
  description = "Droplet tags"
  default     = ["tag"]
}
