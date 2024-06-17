#!/bin/sh

# Loop through each subdirectory in /app/deploy and run terraform init
for dir in /app/terraform/modules/*; do
    if [ -d "$dir" ]; then
        echo "Initializing Terraform in $dir..."
        terraform -chdir="$dir" init
    fi
done