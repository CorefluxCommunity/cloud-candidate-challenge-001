variable "do_token" {
  description = "DigitalOcean API token"
  type        = string
}

variable "direction" {
  description = "Sort direction"
  type        = string
  default     = "desc"
}