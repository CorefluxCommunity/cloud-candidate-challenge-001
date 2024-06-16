package main

import (
	"net/http"
	"os"
)

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

func validateCredentials(username, password string) bool {
	return username == os.Getenv("go_server_user") && password == os.Getenv("go_server_pass")
}
