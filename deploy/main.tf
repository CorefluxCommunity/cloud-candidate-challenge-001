terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "2.39.1"
    }
  }
}

provider "digitalocean" {
  token = var.DIGITALOCEAN_API_TOKEN
}


resource "digitalocean_app" "go-server" {
  spec {
    name   = "orchestration-engine"
    region = "lon1"

    service {
      name               = "go-server"
      environment_slug   = "go"
      instance_count     = 1
      instance_size_slug = "professional-xs"


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
