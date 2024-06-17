package postCreate

import (
    "net/http"
    "encoding/json"
    "os"
    
    "go_webserver/loggers"
    "go_webserver/terraform"
    "go_webserver/terraform/schemas"
    "go_webserver/server/middleware"
)

func PostCreate() http.Handler {
	loggers.LogEvent.Println("Route Created: postCreate")

	var fHandler http.Handler = http.HandlerFunc(createModule)

	fHandler = middleware.MwBasicAuth(fHandler)

	return fHandler
}

func createModule(w http.ResponseWriter, r *http.Request) {
	var input schemas.CreateRequest
	moduleName := "create"

	loggers.LogEvent.Printf("Incoming request: %s %s", r.Method, r.URL.Path)
	
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Failed to parse JSON request body", http.StatusBadRequest)
		return
	}

	// Channels for receiving Terraform output and errors
    outputChan := make(chan string)
    errChan := make(chan error)
    
    // Go routine to set environment variables and execute the corresponding Terraform module, this way we can send concurrent requests
    go func(moduleName string, input schemas.CreateRequest, outputChan chan<- string, errChan chan<- error) {
        // Defer close channels immmediately to avoid any problems later on
        defer close(outputChan)
        defer close(errChan)
        
        if err := setCreateEnvironment(input); err != nil {
            loggers.LogError.Printf("Error setting env variables: %v\n", err)
            errChan <- err
            return
        }
        
        if err := terraform.CallTerraformModule(moduleName, outputChan, errChan); err != nil {
            errChan <- err
            return
        }
    }(moduleName, input, outputChan, errChan)
    
	loggers.LogEvent.Printf("Executing Terraform module '%s'...", moduleName)

    // Wait for Terraform output or error
    select {
    case output := <-outputChan:
        loggers.LogEvent.Printf("Terraform module '%s' execution completed!", moduleName)
		w.WriteHeader(http.StatusOK)
        w.Write([]byte("Terraform execution completed:\n"))
        w.Write([]byte(output)) // Terraform output is sent to client
    case err := <-errChan:
        loggers.LogError.Printf("Failed to execute Terraform module '%s': %v\n", moduleName, err)
        errorMsg := "Failed to execute Terraform module '" + moduleName + "': " + err.Error()
        http.Error(w, errorMsg, http.StatusInternalServerError)
    }
}

// Sets the needed Environment Variables for the "create" module
func setCreateEnvironment(input schemas.CreateRequest) error {
	os.Setenv("TF_VAR_do_token", os.Getenv("do_token"))
	tags, err := json.Marshal(input.Tags)
	if err != nil {
		loggers.LogError.Printf("Error marshaling JSON tags: %v", err)
		return err
	}
	os.Setenv("TF_VAR_droplet_name", input.DropletName)
	os.Setenv("TF_VAR_region", input.Region)
	os.Setenv("TF_VAR_size", input.Size)
	os.Setenv("TF_VAR_image", input.Image)
	os.Setenv("TF_VAR_tags", string(tags))
	
	return nil
}
