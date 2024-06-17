package middleware

import (
	"net/http"
	"os"
)

// Handles the Basic authentication to the server
func MwBasicAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request,) {
		username, password, ok := r.BasicAuth()
		if !ok || !validateCredentials(username, password) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func validateCredentials(username string, password string) bool {
	return username == os.Getenv("go_server_user") && password == os.Getenv("go_server_pass")
}
