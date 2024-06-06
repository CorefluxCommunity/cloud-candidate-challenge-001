package functions

import (
	"awesomeProject/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateDropletHandler(w http.ResponseWriter, r *http.Request) {
	var req structs.CreateDropletRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Region == "" || req.Size == "" || req.Image == "" {
		http.Error(w, "All fields must be provided", http.StatusBadRequest)
		return
	}

	token := "dop_v1_11f72fc8c6223c25925ddeaaaca7f61da2ba8377fd9abb121d920812a90b5dc6"
	output, err := RunTerraformApply(req, token)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v, Output: %s", err, string(output)), http.StatusInternalServerError)
		return
	}

	terraformOutput, err := GetTerraformOutput("terraform/create-droplet")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting output: %v, Output: %s", err, string(terraformOutput)), http.StatusInternalServerError)
		return
	}

	parsedOutput, err := ParseTerraformOutput(terraformOutput)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing output: %v", err), http.StatusInternalServerError)
		return
	}

	ipAddress := parsedOutput.DropletIP.Value
	fmt.Fprintln(w, ipAddress)
}
