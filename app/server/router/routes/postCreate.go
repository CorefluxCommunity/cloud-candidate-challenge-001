package routes

import (
    "net/http"
    "encoding/json"
    "os"
    "fmt"
    
    "go_webserver/loggers"
    "go_webserver/terraform"
    "go_webserver/terraform/schemas"
    "go_webserver/server/middleware"
)

// Handles the create route by passing through the correct middlewares
func PostCreateHandler() http.Handler {
	loggers.LogEvent.Println("Route Created: postCreate")

	var fHandler http.Handler = http.HandlerFunc(createModuleResolver)

	fHandler = middleware.MwBasicAuth(fHandler)

	return fHandler
}

func createModuleResolver(w http.ResponseWriter, r *http.Request) {
	var input schemas.CreateRequest
	terraform.HandleTerraformResponse(w, r, "create", &input, setCreateEnvironment)
}

// Sets the needed Environment Variables for the "create" module
func setCreateEnvironment(input interface{}) error {
	createRequest, ok := input.(*schemas.CreateRequest)
	if !ok {
		return fmt.Errorf("invalid input type")
	}
    
    os.Setenv("TF_VAR_do_token", os.Getenv("do_token"))
	tags, err := json.Marshal(createRequest.Tags)
	if err != nil {
		loggers.LogError.Printf("Error marshaling JSON tags: %v", err)
		return err
	}
	os.Setenv("TF_VAR_droplet_name", createRequest.DropletName)
	os.Setenv("TF_VAR_region", createRequest.Region)
	os.Setenv("TF_VAR_size", createRequest.Size)
	os.Setenv("TF_VAR_image", createRequest.Image)
	os.Setenv("TF_VAR_tags", string(tags))
	
	return nil
}
