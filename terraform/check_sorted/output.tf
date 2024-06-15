output "sorted_droplet_ids" {
  value       = data.digitalocean_droplets.all_droplets.droplets[*].id
}

output "sorted_droplet_names" {
  value       = data.digitalocean_droplets.all_droplets.droplets[*].name
}

output "sorted_droplet_creation_dates" {
  value       = data.digitalocean_droplets.all_droplets.droplets[*].created_at
}
