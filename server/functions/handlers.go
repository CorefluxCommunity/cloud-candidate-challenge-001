package functions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/structs"
)

func ApplyTerraform(w http.ResponseWriter, r *http.Request) {
	fmt.Println("A iniciar")
	var req structs.TerraformRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	applyOut, err := RunTerraformApply(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Terraform apply error: %s\nOutput: %s\n", err, string(applyOut))
		return
	}

	outputOut, err := GetTerraformOutput()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Terraform output error: %s\nOutput: %s\n", err, string(outputOut))
		return
	}

	output, err := ParseTerraformOutput(outputOut)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := structs.TerraformResponse{
		DropletIP: output.DropletIP.Value,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
