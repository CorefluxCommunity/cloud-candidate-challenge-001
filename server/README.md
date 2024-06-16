# Cloud Backend Challenge

## Goal
- Write a Terraform module using a provider and resource of your choice, but ideally using the DigitalOcean provider. This module must take some variable(s) as input and return some output(s).
- Go HTTP Server: Implement a basic HTTP server using Go’s net/http package.
- Request Handling: The server should map each request path to a specific Terraform module.
- Request Parsing: Design the server to accept JSON input in the HTTP request body and parse it to the Terraform command.
- Return a Response: Once the Terraform apply returns the output, the server should return it as the response.
- Add some auth mechanism to your server and deploy it on the cloud.
- Ensure it uses TLS with an ACME cert.
- Deploy using Terraform

## API Endpoints

### POST /droplet
  **Description:** Create a droplet instance on Digital Ocean using terraform and the provided configuration input

  **Accepted Data**
  - `image`: Operating System Image to be installed on virtual machine (STRING) (REQUIRED)
  - `name`: Unique Hostname to refer to your virtual machine (STRING) (REQUIRED)
  - `region`: Region your VM will be deployed (STRING) (REQUIRED)
  - `size`: Size memory and computer resources (STRING) (REQUIRED)
  - `monitoring`: Collect and graph expanded system-level metrics, track performance, and set up alerts instantly within the control panel (BOOLEAN) (DEFAULT FALSE)
  - `ipv6`: If you wish to provide an Ipv6 address (BOOLEAN) (DEFAULT FALSE)

  **Returned Data:**
  - `droplet_id`: Unique identifier of the droplet instance created (STRING)
  - `droplet_ip_address`: Ipv4 address of the droplet instance created (STRING)
