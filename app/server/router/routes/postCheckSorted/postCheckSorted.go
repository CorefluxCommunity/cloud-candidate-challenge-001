package postCheckSorted

import (
    "net/http"
    "encoding/json"
    "os"
    
    "go_webserver/loggers"
    "go_webserver/terraform"
    "go_webserver/terraform/schemas"
    "go_webserver/server/middleware"
)

func PostCheckSorted() http.Handler {
	loggers.LogEvent.Println("Route Created: postCheckSorted")

	var fHandler http.Handler = http.HandlerFunc(checkSortedModule)

	fHandler = middleware.MwBasicAuth(fHandler)

	return fHandler
}

func checkSortedModule(w http.ResponseWriter, r *http.Request) {
	var input schemas.SortRequest
	moduleName := "check_sorted"

	loggers.LogEvent.Printf("Incoming request: %s %s", r.Method, r.URL.Path)
	
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Failed to parse JSON request body", http.StatusBadRequest)
		return
	}

	// Channels for receiving Terraform output and errors
    outputChan := make(chan string)
    errChan := make(chan error)
    
    // Go routine to set environment variables and execute the corresponding Terraform module, this way we can send concurrent requests
    go func(moduleName string, input schemas.SortRequest, outputChan chan<- string, errChan chan<- error) {
        // Defer close channels immmediately to avoid any problems later on
        defer close(outputChan)
        defer close(errChan)
        
        if err := setSortEnvironment(input); err != nil {
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

// Sets the needed Environment Variables for the "check_sorted" module
func setSortEnvironment(input schemas.SortRequest) error {
	os.Setenv("TF_VAR_do_token", os.Getenv("do_token"))
	os.Setenv("TF_VAR_direction", input.Direction)
	
	return nil
}