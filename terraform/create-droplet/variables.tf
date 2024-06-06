variable "do_token" {
  description = "DigitalOcean API token"
  type        = string
}

variable "image" {
  description = "Droplet image"
  type        = string
  default     = "ubuntu-20-04-x64"
}

variable "region" {
  description = "Droplet region"
  type        = string
  default     = "nyc1"
}

variable "size" {
  description = "Droplet size"
  type        = string
  default     = "s-1vcpu-1gb"
}

variable "droplet_name" {
  description = "Droplet name"
  type        = string
  default     = "example-droplet"
}
