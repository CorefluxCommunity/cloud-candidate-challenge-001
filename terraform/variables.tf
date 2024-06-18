// Variables have a irrelevant default due to a terrform issue that demands to variables to be passed when using the destroy command. 
// To workaraound all variables must have a default value, an empty string does not count.
// To avoid an unwanted droplet creation in case the user forget an config variable i chose not to provide useful default values and threat the error instead.


// Env Variables 

variable "api_token" {
  default = "default"
  type    = string
}

variable "aws_region" {
  default = "eu-west-2"
  type    = string
}

variable "bucket_name" {
  type    = string
  default = "cloud-challenge-s3-state"

}


// Droplet variables provided by the user

variable "image" {
  default = "default"
  type    = string
}
variable "name" {
  default = "default"
  type    = string
}
variable "region" {
  default = "default"
  type    = string
}
variable "size" {
  default = "default"
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
