package router

import (
	"net/http"
	
	"go_webserver/server/router/routes/postCheckSorted"
	"go_webserver/server/router/routes/postCreate"
	"go_webserver/server/router/routes/postSearchByTag"
)

func HandleRoutes(server *http.ServeMux) {
	server.Handle("/create", postCreate.PostCreate())
	
	server.Handle("/search_by_tag", postSearchByTag.PostSearchByTag())
	
	server.Handle("/check_sorted", postCheckSorted.PostCheckSorted())
}