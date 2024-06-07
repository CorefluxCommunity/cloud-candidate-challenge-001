package list_droplets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"src/services"
)

func ListDropletsHandler(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("DO_TOKEN")
	if token == "" {
		http.Error(w, "DigitalOcean token not found. Please set the DO_TOKEN environment variable.", http.StatusInternalServerError)
		return
	}
	output, err := services.RunTerraformListDroplets(token)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v, Output: %s", err, string(output)), http.StatusInternalServerError)
		return
	}

	terraformOutput, err := services.GetTerraformOutput("terraform/list-droplets")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting output: %v, Output: %s", err, string(terraformOutput)), http.StatusInternalServerError)
		return
	}

	droplets, err := services.ParseDropletListOutput(terraformOutput)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing output: %v", err), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(droplets); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		return
	}
}
