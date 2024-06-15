package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"os"
)

// Define each struct according to the variables used in each Terraform module
type CreateRequest struct {
	DropletName string   `json:"droplet_name"`
	Region      string   `json:"region"`
	Size        string   `json:"size"`
	Image       string   `json:"image"`
	Tags        []string `json:"tags"`
}

type SearchRequest struct {
	TagToFind	[]string `json:"tag_to_find"`
}

type SortRequest struct {
	Direction	string	`json:"direction"`
}

// Define Terraform Directories here to facilitate future addition if applicable
var terraformModules = []string {"create", "search_by_tag", "check_sorted"}

var logError = log.New(os.Stderr, "", log.LstdFlags)
var logEvent = log.New(os.Stdout, "", log.LstdFlags)

// Webserver entrypoint
// Load the .env file into the Os Environemt Variables (do_token)
// Set Routes for each endpoint
func main() {
	if err := terraformInit(); err != nil {
		logError.Fatalf("Error initiating Terraform: %v", err)
	}
	
	r := setupRouter()

	logEvent.Println("Starting server on :8080...")
	logError.Fatal(http.ListenAndServe(":8080", r))
}

// Init all terraform modules before routing to ensure everything is prepared for any request
func terraformInit() error {
	for _, dir := range terraformModules {
		cmd := exec.Command("terraform", "init")
		cmd.Dir = "./terraform/" + dir

		if err := cmd.Run(); err != nil {
			logEvent.Printf("Terraform init failed in directory %s: %v", dir, err)
			return err
		}
	}

	return nil
}

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

// Sets the needed Environment Variables according to the type of request
func setEnvironmentVariables(input interface{}) error {
	os.Setenv("TF_VAR_do_token", os.Getenv("do_token"))
	switch input := input.(type) {
	case *CreateRequest:
		tags, err := json.Marshal(input.Tags)
		if err != nil {
			logError.Printf("Error marshaling JSON tags: %v", err)
			return err
		}
		os.Setenv("TF_VAR_droplet_name", input.DropletName)
		os.Setenv("TF_VAR_region", input.Region)
		os.Setenv("TF_VAR_size", input.Size)
		os.Setenv("TF_VAR_image", input.Image)
		os.Setenv("TF_VAR_tags", string(tags))
	case *SearchRequest:
		tags, err := json.Marshal(input.TagToFind)
		if err != nil {
			logError.Printf("Error marshaling JSON tags: %v", err)
			return err
		}
		os.Setenv("TF_VAR_tag_to_find", string(tags))
	case *SortRequest:
		os.Setenv("TF_VAR_direction", input.Direction)
	default:
		logError.Println("Unsupported request type")
	}
	return nil
}

// Execute terraform apply of the requested module
// Capture any error executing the command using the error channel
// Send output through output channel
func callTerraformModule(moduleName string, outputChan chan<- string, errChan chan<- error) error {
    dir := "./terraform/" + moduleName
    
    cmd := exec.Command("terraform", "apply", "-auto-approve")
    cmd.Dir = dir
    
    output, err := cmd.CombinedOutput()
    if err != nil {
        logError.Printf("Command %s failed: %v\n", cmd, err)
        errChan <- err
        return err
    }

    outputChan <- string(output)
    
    return nil
}

func handleBasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || !validateCredentials(username, password) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

// Validate username and password
func validateCredentials(username, password string) bool {
	return username == os.Getenv("go_server_user") && password == os.Getenv("go_server_pass")
}