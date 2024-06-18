resource "digitalocean_app" "go-server" {
  spec {
    name   = "orchestration-engine"
    region = "lon1"

    service {
      name               = "go-server"
      environment_slug   = "go"
      instance_count     = 1
      instance_size_slug = "professional-xs"
      http_port          = 8000


      health_check {
        initial_delay_seconds = 10
        period_seconds        = 30
        timeout_seconds       = 5
        failure_threshold     = 3
      }

      env {
        key   = "TF_VAR_dyanmodb_table"
        value = var.module.aws_dynamodb_table.table_name
      }
      env {
        key   = "TF_VAR_bucket_name"
        value = var.module.aws_s3_bucket.bucket_name
      }
      env {
        key   = "DIGITALOCEAN_API_TOKEN"
        value = var.DIGITALOCEAN_API_TOKEN
      }
      env {
        key   = "JWK_URL"
        value = var.JWK_URL
      }
      env {
        key   = "COGNITO_ISSUER"
        value = var.COGNITO_ISSUER
      }
      env {
        key   = "AWS_SECRET_ACCESS_KEY"
        value = var.AWS_SECRET_ACCESS_KEY
      }
      env {
        key   = "AWS_ACCESS_KEY_ID"
        value = var.AWS_ACCESS_KEY_ID
      }
      env {
        key   = "AWS_REGION"
        value = var.AWS_REGION

      }
      image {
        registry_type = "DOCR"
        repository    = "coreflux-challenge"
        tag           = "latest"
        deploy_on_push {
          enabled = true
        }
      }
    }
  }
}
