package main

import (
	"log"
	"net/http"
	"os"
)

// Define Terraform Directories here to facilitate future additions if applicable
var terraformModules = []string {"create", "search_by_tag", "check_sorted"}

var logError = log.New(os.Stderr, "", log.LstdFlags)
var logEvent = log.New(os.Stdout, "", log.LstdFlags)

// Webserver entrypoint
// Set Routes for each endpoint
func main() {
	if err := terraformInit(); err != nil {
		logError.Fatalf("Error initiating Terraform: %v", err)
	}
	
	r := setupRouter()

	logEvent.Println("Starting server on :8080...")
	logError.Fatal(http.ListenAndServe(":8080", r))
}
