terraform {
	required_providers {
		digitalocean = {
			source  = "digitalocean/digitalocean"
			version = "~> 2.0"
		}
	}
}

provider "digitalocean" {
	token = var.do_token
}

resource "digitalocean_app" "go-webserver" {
	spec {
	name   = "go-webserver"
	region = "ams"

	env {
		key = "do_token"
		value = var.do_token
	}
	env {
		key = "go_server_user"
		value = var.go_server_user
	}
	env {
		key = "go_server_pass"
		value = var.go_server_pass
	}

	service {
		name               = "go-service"
		environment_slug   = "go"
		instance_count     = 1
		instance_size_slug = "professional-xs"

		image {
		registry_type 	= "DOCKER_HUB"
		registry 		= "palzap"
		repository 		= "webserver"
		}
	}
	}
}