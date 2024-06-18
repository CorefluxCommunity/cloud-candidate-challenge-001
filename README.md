# Index

1. [Cloud Backend Challenge](#cloud-backend-challenge)
   1. [Goal](#goal)
   2. [Deploying on Digital Ocean](#deploying-on-digital-ocean)
      1. [Deploy Considerations](#deploy-considerations)
      2. [Deployment Process](#deployment-process)
   3. [Folder Structure and Module Interaction](#folder-structure-and-module-interaction)
      1. [Root Directory](#root-directory)
      2. [Modules Directory](#modules-directory)
   4. [Running Locally](#running-locally)
      1. [Docker](#docker)
      2. [Go Build](#go-build)
   5. [API](#api)
      1. [Feature Design and Reasoning](#feature-design-and-reasoning)
      2. [POST /droplet](#post-droplet)
      3. [DELETE /droplet](#delete-droplet)
   6. [Authentication and Authorization](#authentication-and-authorization)
      1. [Design and Reasoning](#design-and-reasoning)

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

## Deploying on Digital Ocean
### Deploy considerations
I explored two ways to deploy the application on Digital Ocean. One option was to use a [Sample Web Application](https://github.com/do-community/terraform-sample-digitalocean-architectures/tree/master/01-minimal-web-db-stack) provided by Digital Ocean as a Terraform production-ready blueprint to deploy a web server. The other option was to use the APP Platform solution. I chose to deploy on the Digital Ocean APP Platform for the following reasons:

1. **Maintain the Scope of the Assignment**: The $200 credit provided by Digital Ocean can be effectively utilized with the APP Platform.
2. **Simplicity and Practicality**: Since this is a Proof of Concept (PoC), deploying a web server with a load balancer and multiple droplets would be overkill.
3. **Avoid Unnecessary Risks**: Due to my lack of knowledge about more complex DevOps and deployment features, choosing a simpler solution helps avoid potential risks.

### Deployment Process
I chose to use Terraform as my deployment mechanism. After experimenting with manual deployment on the Digital Ocean website, I found Terraform to be more straightforward and simpler.

To deploy the application, you need to push your recently built Docker image to a registry. I chose the Digital Ocean Registry, but you can also use Docker Hub if you prefer. For detailed instructions, please refer to the specific documentation for the registry you choose.

To deploy the application, ensure you have all your environment variables set with a TF_VAR_ prefix so Terraform can automatically detect them. Make sure to update the ./deploy/main.tf file to match the image repository where you hosted your Docker image.

### Folder Structure and Module Interaction
Our Terraform configuration is organized into a clear hierarchy to manage infrastructure deployment efficiently. Below is an overview of each directory and its role:
```
./infra
|
|_ _ _ modules
|        |
|        |_ _ _ aws
|        |        |
|        |        |_ _ _ modules
|        |                 |
|        |                 |_ _ _ s3
|        |                 |        |_ _ _ main.tf
|        |                 |
|        |                 |_ _ _ dynamodb
|        |                          |_ _ _ main.tf
|        |
|        |_ _ _ digitalocean
|                 |_ _ _ main.tf
|
|_ _ _ main.tf
|_ _ _ variables.tf
|_ _ _ output.tf
```
### Root Directory
- **main.tf:** This file is the entry point for Terraform. It includes references to the various modules and defines the overall infrastructure configuration.
- **variables.tf:** Contains all the variable definitions used throughout the infrastructure. It ensures that values are parameterized and can be easily configured.
- **output.tf:** Specifies the outputs from the Terraform run, making it easier to access important information (e.g., resource IDs, URLs) after deployment.

### Modules Directory
- **aws/:** This directory contains sub-modules related to AWS services.
- **modules/s3/main.tf:** Manages the creation and configuration of S3 buckets. It is responsible for defining bucket properties and policies.
- **modules/dynamodb/main.tf:** Handles the creation and setup of the DynamoDB table responsible for persisting the state lock ID of the API created droplets.
- **digitalocean/:** Contains the configuration for deploying the server resources on Digital Ocean.

Inside the ./infra folder, run the following three commands:
```
terraform init

// Review the plan
terraform plan

// If the env variables are set then you dont need to pass any variable to the command line.
terraform apply --auto-aprove
```


## Running Locally

### Docker
To run the server locally one must create its own Cognito User Pool and S3 bucket with Amazon to configure the application. More information about this topic can be found in the AWS Documentation.
As the server is not meant to be run locally, but it is deployed on the cloud and meant to be tested there, some aspects of the configuration wont be fully covered in this doc.

Clone the repository locally
```
git clone https://github.com/Desgue/cloud-candidate-challenge-001.git
```

Build the docker image from root folder
```
 docker build -t <IMAGE_NAME> .
```

Run the image with the necessary ENV Variables:
- TF_VAR_dyanmodb_table=<DYNAMO_TABLE_TO_LOCK_STATE>
- TF_VAR_bucket_name=<BUCKET_TO_SAVE_TFSTATE>
- DIGITALOCEAN_API_TOKEN=<YOUR_DIGITAL_OCEAN_API_KEY>
- COGNITO_ISSUER=<YOUR_USER_POOL_ISSUER_URL>
- JWK_URL=<YOUR_USER_POOL_VALIDATION_URL>
- AWS_ACCESS_KEY_ID=<YOUR_AWS_ACCESS_KEY_ID>
- AWS_SECRET_ACCESS_KEY=<YOUR_AWS_SECRET_ACESS_KEY>
- AWS_REGION=<REGION_OF
```go
docker run -it --rm -p 8000:8000 <IMAGE_NAME> -e DIGITALOCEAN_API_TOKEN -e COGNITO_ISSUER -e JWK_URL -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY

// Or use a .env file to provide the variables
docker run --env-file path/to/.env -it --rm -p 8000:8000 <IMAGE_NAME>
```

After having your image runing locally authenticate with Cognito
- Send a POST request to **https://<USER_POOL_NAME>.auth.<YOUR_REGION>.amazoncognito.com/oauth2/token**
- The request must contain the client credentials on the Authorization header or at the body.
- With the access token provided in the response pass it to the your api calls with the following header
- Authorization:Bearer <ACCESS_TOKEN>
- You can find more information about how to authenticate with Aws at the [Cognito Documentation](https://docs.aws.amazon.com/cognito/latest/developerguide/token-endpoint.html)

### Go Build
If you dont wish to run the docker container you can always build the application locally and run it as a binary.
Make sure you have your env variables configured locally or in a .env file located at the "./server" folder.

```cmd
// Inside the folder /server
go build -o <BINARY_NAME>
```
Execute the binary or if you dont want to build just

```cmd
// inside the folder /server
go run main.go
```


## API


### Feature design and reasoning
One of the goals of the exercise was to provide an API to provision Terraform resources. Each API route should point to a single module. I chose to create a Droplet module because I believe the objective is to build a proof of concept that can be later adapted to different needs. Since no particular use case was presented, a single Droplet should suffice.

That said, I thought it was better to build a solution with a few considerations:
- Multiple users from the same organization can interact with the module.
- Concurrent calls to the same module can happen.

With that in mind, I decided to host the Terraform state in an S3 bucket so the state can be easily accessed and locked when it is currently being altered. For that, Terraform also needs a DynamoDB table to keep a record of the locked IDs, preventing concurrent operations and thus avoiding data loss.

Setting up this mechanism can ensure the API can scale if needed and is concurrent safe.
All endpoints required authorization to access. Auth is handled by AWS Cognito using client credentials. With the proper credentials one must authenticate with cognito to obtain a access token valid for 1 hour. That token must be provided in an Authentication header to access all routes. Further explanation on the [Auth](#authentication-and-authorization) section of this document.

Currently, the design is limited to a single-tenant solution. Each API route is connected to a specific module and state, meaning that each call to the route acts on the same droplet.

To implement a multi-tenant solution, where multiple users can manage their own instances, a more complex system would be required. This system would need database integration and a different authentication model to manage separate user sessions.

The reasoning behind this decision is based on the complexity and time limitations given the number of proposed features. To deliver a functional proof of concept that can be further enhanced to accommodate more features, I chose to develop a single-tenant solution.

This approach ensures that the project remains manageable and deliverable within the constraints, while providing a solid foundation for future enhancements to support multi-tenancy.

### POST /droplet

**Description:** Create a droplet instance on Digital Ocean using Terraform and the provided configuration input.

**Accepted Data:**
- `image`: Operating System Image to be installed on virtual machine (STRING) (REQUIRED)
- `name`: Unique Hostname to refer to your virtual machine (STRING) (REQUIRED)
- `region`: Region where your VM will be deployed (STRING) (REQUIRED)
- `size`: Size of memory and compute resources (STRING) (REQUIRED)
- `monitoring`: Collect and graph expanded system-level metrics, track performance, and set up alerts instantly within the control panel (BOOLEAN) (DEFAULT FALSE)
- `ipv6`: Whether to provide an IPv6 address (BOOLEAN) (DEFAULT FALSE)

**Example JSON input:**
```json
{
  "image": "ubuntu-24-04-x64",
  "name": "test",
  "region": "lon1",
  "size": "s-1vcpu-1gb",
  "monitoring": true,
  "ipv6": true
}
```
**Returned Data:**
- `droplet_id`: Unique identifier of the droplet instance created (STRING)
- `droplet_ip_address`: IPv4 address of the droplet instance created (STRING)
- `droplet_status`: Status of the Droplet (STRING)
- `droplet_created_at`: Creation date of the Droplet (STRING)


### DELETE /droplet

**Description:** Destroy the droplet using Terraform's `-destroy` flag.

This endpoint will destroy the droplet provisioned by the POST /droplet endpoint.



## Authentication and Authorization



### Design and Reasoning

Authentication is handled by AWS Cognito using the client credentials grant type with machine-to-machine authentication flow.

I chose the client credentials grant type because I am developing a single-tenant solution where a shared credential can be used across the team. This approach simplifies access management within the organization. It's crucial to note that using the client credentials grant type requires trusting the client application. For a multi-tenant solution or when dealing with untrusted users, I would opt for the authorization code grant, which provides a more secure authentication flow.

**Current Authentication Flow:**
1. **Client Registration:** Users are provided with a `client_id` and `client_secret` by the User Pool Admin.
   
2. **Authentication:** Users authenticate with the AWS Cognito User Pool using these credentials and receive an access token in return.

3. **Access Token:** The access token is valid for 1 hour and must be passed with every request in the `Authorization` header of the request, formatted as `Authorization: Bearer <ACCESS_TOKEN>`.

This authentication mechanism ensures secure access to the API endpoints while maintaining simplicity and cost-effectiveness during development and testing phases.


