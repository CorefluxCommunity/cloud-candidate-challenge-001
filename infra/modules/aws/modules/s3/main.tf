resource "aws_s3_bucket" "coreflux" {
  bucket = var.bucket_name

  tags = {
    Name        = "coreflux"
    Environment = "Dev"
  }
}

variable "bucket_name" {
  type    = string
  default = "coreflux-challenge"
}

output "bucket_name" {
  value = aws_s3_bucket.coreflux.id
}
