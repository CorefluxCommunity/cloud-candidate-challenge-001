
resource "aws_dynamodb_table" "coreflux-sate-lock" {
  name           = var.dynamodb_table
  billing_mode   = "PROVISIONED"
  read_capacity  = 10
  write_capacity = 10
  hash_key       = "LockID"
  attribute {
    name = "LockID"
    type = "S"
  }

  tags = {
    Name        = "coreflux"
    Environment = "Dev"
  }
}


variable "dynamodb_table" {
  type    = string
  default = "coreflux-state-lock"
}
output "table_name" {
  value = aws_dynamodb_table.coreflux-sate-lock.id
}
