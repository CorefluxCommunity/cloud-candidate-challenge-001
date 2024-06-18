output "droplet_ip_address" {
  value = digitalocean_droplet.web.ipv4_address
}

output "droplet_id" {
  value = digitalocean_droplet.web.id
}

output "droplet_status" {
  value = digitalocean_droplet.web.status
}

output "droplet_created_at" {
  value = digitalocean_droplet.web.created_at
}
