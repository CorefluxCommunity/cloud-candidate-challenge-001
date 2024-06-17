#!/bin/bash

echo "Enter the Digital Ocean App URL:"
read app_url
echo "Enter the webserver username:"
read go_username
echo "Enter the webserver password:"
read go_password

# Ensure the URL does not end with a slash
app_url=${app_url%/}

echo "=================================================="
echo "Testing create Module"
echo "=================================================="
curl -u $go_username:$go_password -X POST $app_url/create \
-H "Content-Type: application/json" \
-d '{"droplet_name": "cf-challenge", "region": "nyc1", "size": "s-1vcpu-1gb", "image": "ubuntu-20-04-x64", "tags": ["tag"]}'

echo "=================================================="
echo "Testing check_sorted Module"
echo "=================================================="
curl -u $go_username:$go_password -X POST $app_url/check_sorted \
-H "Content-Type: application/json" \
-d '{"direction": "asc"}'

echo "=================================================="
echo "Testing search_by_tag Module"
echo "=================================================="
curl -u $go_username:$go_password -X POST $app_url/search_by_tag \
-H "Content-Type: application/json" \
-d '{"tag_to_find": ["droplet2"]}'

echo "=================================================="
echo "Testing unauthorized access"
echo "=================================================="
curl -X POST $app_url/search_by_tag \
-H "Content-Type: application/json" \
-d '{"tag_to_find": ["droplet2"]}'