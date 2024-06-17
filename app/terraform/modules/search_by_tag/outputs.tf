output "droplet_ids" {
  value = [for droplet in data.digitalocean_droplets.all_droplets.droplets : droplet.id]
}

output "droplet_names" {
  value = [for droplet in data.digitalocean_droplets.all_droplets.droplets : droplet.name]
}

output "droplet_ips" {
  value = [for droplet in data.digitalocean_droplets.all_droplets.droplets : droplet.ipv4_address]
}
