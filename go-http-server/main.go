package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type TerraformRequest struct {
	DigitalOceanToken string   `json:"digitalocean_token"`
	Image             string   `json:"image"`
	Name              string   `json:"name"`
	Region            string   `json:"region"`
	Size              string   `json:"size"`
	SSHKeys           []string `json:"ssh_keys"`
	IPv6              bool     `json:"ipv6"`
	Monitoring        bool     `json:"monitoring"`
	VPCUUID           string   `json:"vpc_uuid"`
}

type TerraformResponse struct {
	DropletIP string `json:"droplet_ip"`
}

func handleTerraform(w http.ResponseWriter, r *http.Request) {
	var req TerraformRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tfDir := "./terraform"
	varsFile := filepath.Join(tfDir, "terraform.tfvars")

	// Write the variables to terraform.tfvars
	varsContent := fmt.Sprintf(`
do_token = "%s"
image = "%s"
name = "%s"
region = "%s"
size = "%s"
ssh_keys = %q
ipv6 = %t
monitoring = %t
vpc_uuid = "%s"
`, req.DigitalOceanToken, req.Image, req.Name, req.Region, req.Size, req.SSHKeys, req.IPv6, req.Monitoring, req.VPCUUID)

	if err := os.WriteFile(varsFile, []byte(varsContent), 0644); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Run `terraform init`
	cmd := exec.Command("terraform", "init")
	cmd.Dir = tfDir
	if output, err := cmd.CombinedOutput(); err != nil {
		log.Printf("Terraform init failed: %s\n%s", err, string(output))
		http.Error(w, "Terraform init failed", http.StatusInternalServerError)
		return
	}

	// Run `terraform apply`
	cmd = exec.Command("terraform", "apply", "-auto-approve")
	cmd.Dir = tfDir
	if output, err := cmd.CombinedOutput(); err != nil {
		log.Printf("Terraform apply failed: %s\n%s", err, string(output))
		http.Error(w, "Terraform apply failed", http.StatusInternalServerError)
		return
	}

	// Run `terraform output -json`
	cmd = exec.Command("terraform", "output", "-json")
	cmd.Dir = tfDir
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Terraform output failed: %s\n%s", err, string(output))
		http.Error(w, "Terraform output failed", http.StatusInternalServerError)
		return
	}

	// Parse the JSON output dynamically
	var tfOutput map[string]interface{}
	if err := json.Unmarshal(output, &tfOutput); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Extract the Droplet IP from the output
	dropletIP := ""
	if dropletInfo, ok := tfOutput["droplet_ip"].(map[string]interface{}); ok {
		if val, ok := dropletInfo["value"].(string); ok {
			dropletIP = val
		}
	}

	// Send the response
	response := TerraformResponse{DropletIP: dropletIP}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/terraform", handleTerraform)
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
