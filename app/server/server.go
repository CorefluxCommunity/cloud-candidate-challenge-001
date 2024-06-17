package server

import (
	"net/http"
	
	"go_webserver/server/router"
	"go_webserver/loggers"
)

// Start the server by setting up the router
func Start() {
	server := http.NewServeMux()

	loggers.LogEvent.Println("Starting server...")
	
	router.HandleRoutes(server)
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		loggers.LogError.Printf("Http Server error %v\n", err)
	}
}
