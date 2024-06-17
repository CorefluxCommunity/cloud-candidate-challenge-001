variable "do_token" {
  description = "DigitalOcean API token"
  type        = string
}

variable "tag_to_find" {
  description = "Tag to filter droplets"
  type        = list(string)
}
