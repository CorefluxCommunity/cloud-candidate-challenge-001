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

variable "DIGITALOCEAN_API_TOKEN" {
  type = string
}


variable "dynamodb_table_name" {
  type = string
}

variable "s3_bucket_name" {
  type = string
}
