// S3 VARIABLES
variable "bucket_name" {
  type    = string
  default = "coreflux-challenge"
}



// DIGITAL OCEAN VARIABLES
variable "DIGITALOCEAN_API_TOKEN" {
  type = string
}
variable "JWK_URL" {
  type = string
}
variable "COGNITO_ISSUER" {
  type = string
}

variable "AWS_SECRET_ACCESS_KEY" {
  type = string
}
variable "AWS_ACCESS_KEY_ID" {
  type = string
}
variable "AWS_REGION" {
  type    = string
  default = "eu-west-2"
}
