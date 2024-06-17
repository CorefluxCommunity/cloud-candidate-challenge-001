package routes

import (
    "net/http"
    "os"
    "fmt"
    
    "go_webserver/loggers"
    "go_webserver/terraform"
    "go_webserver/terraform/schemas"
    "go_webserver/server/middleware"
)

// Handles the check_sorted route by passing through the correct middlewares
func PostCheckSortedHandler() http.Handler {
	loggers.LogEvent.Println("Route Created: postCheckSorted")

	var fHandler http.Handler = http.HandlerFunc(checkSortedModuleResolver)

	fHandler = middleware.MwBasicAuth(fHandler)

	return fHandler
}

func checkSortedModuleResolver(w http.ResponseWriter, r *http.Request) {
	var input schemas.SortRequest
	terraform.HandleTerraformResponse(w, r, "check_sorted", &input, setSortEnvironment)
}

// Sets the needed Environment Variables for the "check_sorted" module
func setSortEnvironment(input interface{}) error {
	sortRequest, ok := input.(*schemas.SortRequest)
	if !ok {
        loggers.LogEvent.Println(ok)
		return fmt.Errorf("input is not of type SortRequest")
	}

	os.Setenv("TF_VAR_do_token", os.Getenv("do_token"))
	os.Setenv("TF_VAR_direction", sortRequest.Direction)
	
	return nil
}
