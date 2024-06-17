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

// Handles the search_by_tag route by passing through the correct middlewares
func PostSearchByTagHandler() http.Handler {
	loggers.LogEvent.Println("Route Created: postSearchByTag")

	var fHandler http.Handler = http.HandlerFunc(searchByTagModuleResolver)

	fHandler = middleware.MwBasicAuth(fHandler)

	return fHandler
}

func searchByTagModuleResolver(w http.ResponseWriter, r *http.Request) {
	var input schemas.SearchRequest
	terraform.HandleTerraformResponse(w, r, "search_by_tag", &input, setSearchEnvironment)
}

// Sets the needed Environment Variables for the "search_by_tag" module
func setSearchEnvironment(input interface{}) error {
    searchRequest, ok := input.(*schemas.SearchRequest)
	if !ok {
		return fmt.Errorf("invalid input type")
	}
    
    os.Setenv("TF_VAR_do_token", os.Getenv("do_token"))
	tags, err := json.Marshal(searchRequest.TagToFind)
		if err != nil {
			loggers.LogError.Printf("Error marshaling JSON tags: %v", err)
			return err
		}
	os.Setenv("TF_VAR_tag_to_find", string(tags))
	
	return nil
}
