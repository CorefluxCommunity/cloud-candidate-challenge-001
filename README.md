# Coreflux Cloud Team - Candidate Challenge 001 (Go and Terraform)

## The challenge

The challenge was to build a small-scale orchestration engine using Go and Terraform. The engine needed to handle HTTP requests, process them to provision Terraform resources, and return the output as a response. This project aimed to test skills in Go programming, Terraform usage, asynchronous operations, and concurrency handling. Check the challengeREADME.md for detailed information on the challenge.

## Key Components

### Terraform Modules:
* Implemented 3 Terraform modules to be called on certain endpoints, I chose to implement these modules to play around with different functionalities offered by Terraform:
    * create
        ```
        This module creates 3 Droplets and automatically assigns different Tags to each one.
        Input: Name, Region, Size, Image and Tags
        Output: Droplets ID, IPv4 Address and Tags
        ```
    * search_by_tag
        ```
        This module filters the requested data (all Droplets) by Tags.
        Input: Tags
        Output: Droplets ID, Name and IPv4 Address.
        ```
    * check_sorted 
        ```
        This module sorts the requested data (all Droplets) by Name.
        Input: Sort direction
        Output: Droplets ID, Name and Creation Date.
        ```

### Go HTTP Server:
* Developed a basic HTTP server using Go's net/http package.
* Mapped each request path to a specific Terraform module.
* Designed to accept JSON input in the HTTP request body, which contains parameters for Terraform.

### Extra:
* Implemented a basic authentication mechanism on the webserver.
* Used Digital Ocean App Platform to deploy the go webserver. This way it automatically assigns and manages TLS certificates and ensures that the application is served securely over HTTPS.
* Used Terraform for the deployment process.
* Included several scripts and a Makefile to make deployment and testing more automated.

## Learning and Skills Tested

* Go Programming: Developed skills in Go programming, especially on how Go offers good functionality for asynchronous programming.
* Terraform Usage: Learned about Terraform and infrastructure as code, as well as its declarative configuration and state management.
* Cloud deployment: Used Terraform with DigitalOcean as a provider to interact programmatically with DigitalOcean resources, including Droplets and App Platform.

## How to Use
### Prerequisites

Ensure you have the following:
* [Terraform](https://developer.hashicorp.com/terraform/install?product_intent=terraform)
* [Make utility](https://www.gnu.org/software/make/#download)
* [DigitalOcean API Token](https://cloud.digitalocean.com/account/api/tokens)

### Main commands
#### Deploy the App to DigitalOcean's App Platform

```
make all
```
* Will prompt you for your DigitalOcean API token and to set a username and a password for the webserver.

#### Execute tests to verify the functionality of the web server::
```
make test
```
* Will prompt you for the URL of the deployed APP (you can get it from the output of the make all command), username and password to access the webserver.
* This command will test all endpoints in succession and also an Unauthorized access to make sure authentication is working.

### Other options
#### Generate terraform.tfvars File required for Terraform variables
```
make vars
```

#### Initialize Terraform in the deployment directory (deploy)
```
make init
```

#### To see the execution plan without making any changes, run
```
make plan
```

#### Deploy the go-webserver application to DigitalOcean App Platform. This command also applies changes automatically:
```
make apply
```
#### Update the Docker Hub repository with the most recent changes to the go-webserver (if applicable)
```
make update_image
```

#### Remove generated terraform.tfvars file
```
make clean
```

#### Completely clean up Terraform files (including state files, lock files, and temporary outputs):

``` 
make fclean
```

### Further testing
If you want to further test the deployed webserver you can issue your own curl requests:
```
curl -u <USERNAME>:<PASSWORD> -X POST <APP_URL>/<ENDPOINT> \
-H "Content-Type: application/json" \
-d '{<ADJUST ACCORDINGLY>}'
```
