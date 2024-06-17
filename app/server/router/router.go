package router

import (
	"net/http"
	
	"go_webserver/server/router/routes"
)

// Define the routes and pass each one to the correct handler
func HandleRoutes(server *http.ServeMux) {
	server.Handle("/create", routes.PostCreateHandler())
	
	server.Handle("/search_by_tag", routes.PostSearchByTagHandler())
	
	server.Handle("/check_sorted", routes.PostCheckSortedHandler())
}
