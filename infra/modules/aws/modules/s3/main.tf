resource "aws_s3_bucket" "coreflux" {
  bucket = var.bucket_name

  tags = {
    Name        = "coreflux"
    Environment = "Dev"
  }
}

variable "bucket_name" {
  type    = string
  default = "cloud-challenge-s3-state"
}

output "bucket_name" {
  value = aws_s3_bucket.coreflux.id
}
