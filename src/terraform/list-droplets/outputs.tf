output "droplet_list" {
  value = [
    for droplet in data.digitalocean_droplets.all.droplets :
    {
      id     = droplet.id
      name   = droplet.name
      ipv4   = droplet.ipv4_address
      region = droplet.region
      size   = droplet.size
      image  = droplet.image
    }
  ]
}