package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ListDropletsHandler(w http.ResponseWriter) {
	token := "dop_v1_11f72fc8c6223c25925ddeaaaca7f61da2ba8377fd9abb121d920812a90b5dc6"
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
