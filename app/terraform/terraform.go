package terraform

import (
	"os/exec"
    "net/http"
    "encoding/json"

	"go_webserver/loggers"
)

// Execute terraform apply of the requested module
// Capture any error executing the command using the error channel
// Send output through output channel
func CallTerraformApply(moduleName string, outputChan chan<- string, errChan chan<- error) error {
    dir := "./terraform/modules/" + moduleName
    
    cmd := exec.Command("terraform", "apply", "-auto-approve")
    cmd.Dir = dir
    
    output, err := cmd.CombinedOutput()
    if err != nil {
        loggers.LogError.Printf("Command %s failed: %v\n", cmd, err)
        errChan <- err
        return err
    }

    outputChan <- string(output)
    
    return nil
}

// Applies go routines to enable concurrency on requests
func HandleTerraformResponse(w http.ResponseWriter, r *http.Request, moduleName string, input interface{}, setEnvFunc func(interface{}) error) {
	loggers.LogEvent.Printf("Incoming request: %s %s", r.Method, r.URL.Path)

	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		http.Error(w, "Failed to parse JSON request body", http.StatusBadRequest)
		return
	}

	// Channels for receiving Terraform output and errors
	outputChan := make(chan string)
	errChan := make(chan error)

	// Go routine to set environment variables and execute the corresponding Terraform module
	go func(moduleName string, input interface{}, outputChan chan<- string, errChan chan<- error) {
		defer close(outputChan)
		defer close(errChan)

		if err := setEnvFunc(input); err != nil {
			loggers.LogError.Printf("Error setting env variables: %v\n", err)
			errChan <- err
			return
		}

		if err := CallTerraformApply(moduleName, outputChan, errChan); err != nil {
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
