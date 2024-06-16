#!/bin/bash

# Check if terraform.tfvars already exists
if [ -f terraform.tfvars ]; then
    echo "terraform.tfvars already exists. Do you want to overwrite it? (yes/no)"
    read response
    if [ "$response" != "yes" ]; then
        echo "Aborting."
        exit 1
    fi
fi

# Collect user inputs
echo "Enter your Digital Ocean API Token:"
read do_token
echo "Enter your desired Webserver username:"
read go_server_user
echo "Enter your desired Webserver password:"
read go_server_pass

# Write to terraform.tfvars
cat > terraform.tfvars <<EOF
do_token = "$do_token"
go_server_user = "$go_server_user"
go_server_pass = "$go_server_pass"
EOF

echo "terraform.tfvars has been created."