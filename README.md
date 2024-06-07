# Coreflux Cloud Team - Candidate Challenge 001 (Go and Terraform)

## Overview

Welcome to the Coreflux Cloud Team candidate challenge 001! In this challenge, you will be tasked with building a small-scale orchestration engine using Go and Terraform. Your goal is to create an HTTP server in Go that processes HTTP requests to provision Terraform resources. This will test your skills in Go, Terraform, asynchronous programming and handling concurrency.

## Requirements

- Write a Terraform module using a provider and resource of your choice, but ideally using the DigitalOcean provider. This module must take some variable(s) as input and return some output(s).
- Go HTTP Server: Implement a basic HTTP server using Go’s net/http package.
- Request Handling: The server should map each request path to a specific Terraform module.
- Request Parsing: Design the server to accept JSON input in the HTTP request body and parse it to the Terraform command.
- Return a Response: Once the Terraform apply returns the output, the server should return it as the response.

## Instructions

### Prerequisites

- Ensure you have Docker installed on your machine
- Obtain a DigitalOcean API token

### Setup

1. **Clone the Repository:**
   ```sh
   git clone <repository-url>
   cd <repository-directory>
2. **Create .env file**
    ```sh
    DO_TOKEN=your_digitalocean_token
3. **Build and run with Docker Compose**
    ```sh
    docker-compose up --build

### Endpoints

1. **/create** - Creates a droplet in DigitalOcean based on the JSON payload provided in the request body.
   - **Method:** POST
   - **Request Body Example:**
   ```json
   {
    "image": "ubuntu-20-04-x64",
    "region": "nyc3",
    "size": "s-1vcpu-1gb",
    "droplet_name": "coreflux"
    }
    ```
    - **Response:** Returns the IP address of the created droplet.    
    

2. **/all** - Returns all the droplets created.
    - **Method:** GET
    - **Response:** JSON array of all droplets.
   
## How It Works

**Endpoint Execution:**
- **/create Endpoint:**
    - When a POST request is made to the `/create` endpoint with a JSON payload, the server parses the JSON input to extract the necessary parameters for creating a droplet.
    - The server then initiates the Terraform module responsible for creating droplets by executing the `terraform init` and `terraform apply` command with the provided parameters.
    - That guarantees that terraform does not initiate modules not needed.
    - Terraform uses the DigitalOcean provider to create the droplet in the specified region with the given characteristics.
    - Once the droplet is created, Terraform outputs the IP address, which the server captures and returns in the HTTP response.


- **/all Endpoint:**
    - When a GET request is made to the `/all` endpoint, the server initiates the Terraform module responsible for listing all created droplets by executing the `terraform init` and `terraform apply` command configured to list droplets.
    - Terraform uses the DigitalOcean provider to retrieve information about all droplets created.
    - The server captures the list of droplets and returns it as a JSON array in the HTTP response.


## Resources

- [Go Documentation](https://go.dev/doc/)
- [Terraform Documentation](https://developer.hashicorp.com/terraform)
- [Terraform Provider Registry](https://registry.terraform.io)
- [Terraform DigitalOcean Provider](https://registry.terraform.io/providers/digitalocean/digitalocean/2.39.1)

## Submission

Please fork this repository and submit your solution as a pull request. Include detailed documentation on how to run your solution and any assumptions or design decisions you made.
Good luck, and we look forward to reviewing your submission!

## Extra credit - If you really want to stand out

Add some auth mechanism to your server and deploy it on the cloud. Ensure it uses TLS with an ACME cert. Extra extra bonus points if your deployment is done using Terraform ;)

## Claim 200$ of DigitalOcean Credits

If you want to use DigitalOcean for this challenge, you can claim 200$ of DigitialOcean credits by clicking on the button below.

[![DigitalOcean Referral Badge](https://web-platforms.sfo2.cdn.digitaloceanspaces.com/WWW/Badge%203.svg)](https://www.digitalocean.com/?refcode=dbb46c5fa238&utm_campaign=Referral_Invite&utm_medium=Referral_Program&utm_source=badge)

