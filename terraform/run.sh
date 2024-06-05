#!/bin/bash

set -e

for dir in */; do
    if [ -d "$dir" ]; then
        cd "$dir"
        terraform init
        cd ..
    fi
done
