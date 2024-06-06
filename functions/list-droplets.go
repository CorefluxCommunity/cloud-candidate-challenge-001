package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func ListDropletsHandler(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("DO_TOKEN")
	output, err := RunTerraformListDroplets(token)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v, Output: %s", err, string(output)), http.StatusInternalServerError)
		return
	}

	terraformOutput, err := GetTerraformOutput("terraform/list-droplets")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting output: %v, Output: %s", err, string(terraformOutput)), http.StatusInternalServerError)
		return
	}

	droplets, err := ParseDropletListOutput(terraformOutput)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing output: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(droplets)
}
