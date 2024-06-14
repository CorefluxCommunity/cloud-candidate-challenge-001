# Define the different outputs

output "droplet_id" {
  description = "The IP address of the DigitalOcean Droplet"
  value       = module.create_droplet.droplet_id
}

output "droplet_ip" {
  description = "The IP address of the DigitalOcean Droplet"
  value       = module.create_droplet.droplet_ipv4_address
}

output "filtered_droplet_ids" {
  value = module.search_by_tag.droplet_ids
}

output "filtered_droplet_names" {
  value = module.search_by_tag.droplet_names
}

output "filtered" {
  value = module.search_by_tag.droplet_ips
}

output "sorted_droplet_ids" {
  value = module.search_by_tag.droplet_ids
}

output "sorted_droplet_names" {
  value = module.search_by_tag.droplet_names
}

output "sorted_droplet_ips" {
  value = module.search_by_tag.droplet_ips
}