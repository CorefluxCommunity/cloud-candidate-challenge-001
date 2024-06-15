#!/bin/bash

echo "\n\n=================================================="
echo "Testing create Module"
echo "=================================================="
curl -X POST http://localhost:8080/create \
-H "Content-Type: application/json" \
-d '{"droplet_name": "droplet13", "region": "nyc1", "size": "s-1vcpu-1gb", "image": "ubuntu-20-04-x64", "tags": ["tag"]}'

echo "\n\n=================================================="
echo "Testing check_sorted Module"
echo "=================================================="
curl -X POST http://localhost:8080/check_sorted \
-H "Content-Type: application/json" \
-d '{"direction": "asc"}'


echo "\n\n=================================================="
echo "Testing search_by_tag Module"
echo "=================================================="
curl -X POST http://localhost:8080/search_by_tag \
-H "Content-Type: application/json" \
-d '{"tag_to_find": ["droplet2"]}'

