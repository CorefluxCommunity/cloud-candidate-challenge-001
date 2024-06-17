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
- DIGITALOCEAN_API_TOKEN=<YOUR_DIGITAL_OCEAN_API_KEY>
- COGNITO_ISSUER=<YOUR_USER_POOL_ISSUER_URL>
- JWK_URL=<YOUR_USER_POOL_VALIDATION_URL>
- AWS_ACCESS_KEY_ID=<YOUR_AWS_ACCESS_KEY_ID>
- AWS_SECRET_ACCESS_KEY=<YOUR_AWS_SECRET_ACESS_KEY>
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

## Deploying on Digital Ocean
### Deploy considerations
I looked into two ways of the deploying the application on Digital Ocean, one was to use a [Sample Web Application](https://github.com/do-community/terraform-sample-digitalocean-architectures/tree/master/01-minimal-web-db-stack) provided by DO as a terraform production ready blueprint to deploy a web server or use the APP Platform solution. 
I chose to deploy on Digital Ocean APP Platform for the follwoing reasons:
- Mantain the scope of the assignment and make use of the $200 credit provided
- Maintain the deployment simple and pratical, as this is no more then a PoC I believed it was an overkill to deploy a web server with load balancer and multiples droplets.
- Lack of knowledge about more complex devops and deployment features which could expose me to unecessary risks.

### Deployment Process
I chose to use terraform as my deployment mechanism, even though I experimented with manual deploying on the Digital Ocean website before, I could compare both process and felt way more confortable using terraform as its straightfoward and simple process.

To deploy the application you will need to push your recently built docker image to a registry, I chose the Digital Ocean Registry but you can push to the Docker Hub if you wish.
To do that I suggest you follow the specific documentation for the method you chose.

To deploy the application make sure you have all your env variables set with a TF_VAR_ prefix so terraform can search for them automatically.
Make sure to change the ./deploy/main.tf image repository to match where you hosted your docker image.

Inside the ./deploy folder run the following 3 commands
```
terraform init

// Review the plan
terraform plan

// If the env variables are set with the TF_VAR_ prefix then you dont need to pass any variable to the command line.
terraform apply --auto-aprove
```


## Terraform Modules
One of the goals of the exercise was to provide an API to provision Terraform resources, each API route should point to a single Module.
I chose to create a Droplet module because I believe the objective is to build a Proof of Concept that can be later adapted to different needs, so with no particular use case presented a single Droplet should suffice.
That said, I tough it was better to build a solution with a few considerations:
- Multiple Users from the same organization can interact with the Module
- Concurrent calls to the same Module can happen
With that in mind I decided to host the terraform state in a S3 bucket so the state can be easly accessed and lock the state when its currently being altered. For that terraform also needs a dynamodb table to keep record of the locked Ids so it does not perform concurrent opeartions and thus avoiding data loss.

Setting up this mechanism can ensure the API can scale if needed and its concurrent safe.



## API
All endpoints required authorization to access. Auth is handled by AWS Cognito using client credentials. With the proper credentials one must authenticate with cognito to obtain a access token valid for 1 hour. That token must be provided in an Authentication header to access all routes. Further explanation on the Auth section of this document.

### Feature design and reasoning
The goal of this project is to provide a proof of concept orchestration engine that receive terraform inputs in form of JSON to parse and give to terraform to provide a instance, in this case a droplet in Digital Ocean.
The current limitation with this current design is not being a multi tenant solution. Each route is connected to a module and a specific state, so that each call to the route is acting on the same Droplet.
To implement a multi tenant solution where multiple users can manage their own instances would require a more complex system with a database integration and a different authentication model to manage separate sessions.
The reasoning behind this decision was the complexity and time limitation for the amount of features proposed, so to deliver a working PoC that can be further enhanced to accomodate more features I chose to develop a single tenant solution.

### POST /droplet
  **Description:** Create a droplet instance on Digital Ocean using terraform and the provided configuration input

  **Accepted Data**
  - `image`: Operating System Image to be installed on virtual machine (STRING) (REQUIRED)
  - `name`: Unique Hostname to refer to your virtual machine (STRING) (REQUIRED)
  - `region`: Region your VM will be deployed (STRING) (REQUIRED)
  - `size`: Size memory and computer resources (STRING) (REQUIRED)
  - `monitoring`: Collect and graph expanded system-level metrics, track performance, and set up alerts instantly within the control panel (BOOLEAN) (DEFAULT FALSE)
  - `ipv6`: If you wish to provide an Ipv6 address (BOOLEAN) (DEFAULT FALSE)
  - Example JSON input
  ```
  {
    "image": "ubuntu-24-04-x64",
    "name":"test",
    "region":"lon1",
    "size":"s-1vcpu-1gb",
    "monitoring": true,
    "ipv6": true
}
```


  **Returned Data:**
  - `droplet_id`: Unique identifier of the droplet instance created (STRING)
  - `droplet_ip_address`: Ipv4 address of the droplet instance created (STRING)
  - `droplet_status`: The Status of the Droplet (STRING)
  - `droplet_created_at`: Creating date of the Droplet (STRING)

### DELETE /droplet
**Description:** Destroy the droplet using terraform -destroy flag


## Authentication and Authorization
Auth is handled by Aws cognito in a client credentials grant type with machine to machine auth flow in mid.

### Design and Reasoning
I chose Aws Cognito because of my familiarity with it, the ease of use to set up Auth and the low cost allowing me to develop and test my solution basically for free.
The current Auth flow works like the following:
- User is provided with Client_id and Client_secret by the User_pool Admin
- User then Authenticate with the Aws Cognito User Pool using the credentials and receive a Access token in return. 
- Access token is valid for 1 hour and must be pass to every request in the Authentication header of the request in the following format "Authentication:Bearer <ACCESS_TOKEN>"

