output "app_default_ingress" {
  description = "The default URL to access the app"
  value       = digitalocean_app.go-webserver.default_ingress
}

# you can check this in the Droplet list
output "app_live_url" {
  description = "The live URL of the app"
  value       = digitalocean_app.go-webserver.live_url
}
