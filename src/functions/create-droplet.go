package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"src/structs"
)

func CreateDropletHandler(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("DO_TOKEN")
	if token == "" {
		http.Error(w, "DigitalOcean token not found. Please set the DO_TOKEN environment variable.", http.StatusInternalServerError)
		return
	}

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
