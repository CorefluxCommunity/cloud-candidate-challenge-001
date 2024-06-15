variable "aws_region" {
  default = "eu-west-2"
  type    = string
}

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
