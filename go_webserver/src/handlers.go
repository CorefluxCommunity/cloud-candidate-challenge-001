package main

import (
	"encoding/json"
	"net/http"
)

// Set up routes dynamically to allow for possible future addition of Terraform Modules
func setupRouter() http.Handler {
	r := http.NewServeMux()

	// Protect Terraform endpoints with Basic Auth
	for _, route := range terraformModules {
		r.HandleFunc("/"+route, handleBasicAuth(handleRoute))
	}
	return r
}

// Dependeding on the URL provided process requests accordingly
func handleRoute(w http.ResponseWriter, r *http.Request) {
	var moduleName string
	var input interface{}
	
	logEvent.Printf("Incoming request: %s %s", r.Method, r.URL.Path)
	
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	
	for _, module := range terraformModules {
		if r.URL.Path == "/"+module {
			moduleName = module
			switch module {
			case "create":
				input = &CreateRequest{}
			case "search_by_tag":
				input = &SearchRequest{}
			case "check_sorted":
				input = &SortRequest{}
			}		
			processRequest(w, r, moduleName, input)
			return
		}
	}
	
	http.Error(w, "Module not found", http.StatusNotFound)
}

func processRequest(w http.ResponseWriter, r *http.Request, moduleName string, input interface{}) {
    switch input := input.(type) {
    case *CreateRequest:
        if err := json.NewDecoder(r.Body).Decode(input); err != nil {
            http.Error(w, "Failed to parse JSON request body", http.StatusBadRequest)
            return
        }
    case *SearchRequest:
        if err := json.NewDecoder(r.Body).Decode(input); err != nil {
            http.Error(w, "Failed to parse JSON request body", http.StatusBadRequest)
            return
        }
    case *SortRequest:
        if err := json.NewDecoder(r.Body).Decode(input); err != nil {
            http.Error(w, "Failed to parse JSON request body", http.StatusBadRequest)
            return
        }
    default:
        http.Error(w, "Unsupported request type", http.StatusBadRequest)
        return
    }
    
    // Channels for receiving Terraform output and errors
    outputChan := make(chan string)
    errChan := make(chan error)
    
    // Defer close channels immmediately to avoid any problems later on
    defer close(outputChan)
    defer close(errChan)
    
    // Go routine to set environment variables and execute the corresponding Terraform module, this way we can send concurrent requests
    go func(moduleName string, input interface{}, outputChan chan<- string, errChan chan<- error) {
        if err := setEnvironmentVariables(input); err != nil {
            logError.Printf("Error setting env variables: %v\n", err)
            errChan <- err
            return
        }
        
        if err := callTerraformModule(moduleName, outputChan, errChan); err != nil {
            errChan <- err
            return
        }
    }(moduleName, input, outputChan, errChan)
    
	logEvent.Printf("Executing Terraform module '%s'...", moduleName)

    // Wait for Terraform output or error
    select {
    case output := <-outputChan:
        logEvent.Printf("Terraform module '%s' execution completed!", moduleName)
		w.WriteHeader(http.StatusOK)
        w.Write([]byte("Terraform execution completed:\n"))
        w.Write([]byte(output)) // Terraform output is sent to client
    case err := <-errChan:
        logError.Printf("Failed to execute Terraform module '%s': %v\n", moduleName, err)
        errorMsg := "Failed to execute Terraform module '" + moduleName + "': " + err.Error()
        http.Error(w, errorMsg, http.StatusInternalServerError)
    }
}