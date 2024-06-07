package create_droplet

import (
	"fmt"
	"net/http"
	"src/services"
)

func CreateDropletHandler(w http.ResponseWriter, r *http.Request) {
	token, err := services.GetDigitalOceanToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req, err := services.DecodeRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := services.RunTerraformApply(req, token)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v, Output: %s", err, string(output)), http.StatusInternalServerError)
		return
	}

	ipAddress, err := services.GetDropletIPAddress()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := fmt.Fprintln(w, ipAddress); err != nil {
		http.Error(w, fmt.Sprintf("Error writing response: %v", err), http.StatusInternalServerError)
		return
	}
}
