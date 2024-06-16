# Start from a base image that includes Go installed (e.g., golang:1.18)
FROM golang:1.18 AS build

# Install necessary tools (including Terraform)
RUN apt-get update && apt-get install -y \
    unzip \
 && rm -rf /var/lib/apt/lists/*

# Download and install Terraform
RUN wget https://releases.hashicorp.com/terraform/0.15.5/terraform_0.15.5_linux_amd64.zip && \
    unzip terraform_0.15.5_linux_amd64.zip -d /usr/local/bin/ && \
    rm terraform_0.15.5_linux_amd64.zip

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .
RUN  chmod +x ./terraform_init.sh
RUN  ./terraform_init.sh

# Set the working directory inside the container
WORKDIR /app/src

# Build the Go application
RUN go mod download

RUN go build -o /go_webserver

# Command to run the application
CMD ["/go_webserver"]
