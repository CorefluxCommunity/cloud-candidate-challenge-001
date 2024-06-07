package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"src/structs"
	"strings"
)

func HandleHTTPResponse(w http.ResponseWriter, ch <-chan string) {
	result := <-ch
	if strings.HasPrefix(result, "Error:") {
		http.Error(w, result, http.StatusInternalServerError)
	} else {
		if _, err := fmt.Fprintln(w, result); err != nil {
			http.Error(w, fmt.Sprintf("Error writing response: %v", err), http.StatusInternalServerError)
		}
	}
}

func GetDigitalOceanToken() (string, error) {
	token := os.Getenv("DO_TOKEN")
	if token == "" {
		return "", fmt.Errorf("DigitalOcean token not found. Please set the DO_TOKEN environment variable")
	}
	return token, nil
}

func DecodeRequest(r *http.Request) (structs.CreateDropletRequest, error) {
	var req structs.CreateDropletRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return req, err
	}
	if req.Name == "" || req.Region == "" || req.Size == "" || req.Image == "" {
		return req, fmt.Errorf("All fields must be provided")
	}
	return req, nil
}

func GetDropletIPAddress() (string, error) {
	terraformOutput, err := GetTerraformOutput("terraform/create-droplet")
	if err != nil {
		return "", fmt.Errorf("error getting output: %v, output: %s", err, string(terraformOutput))
	}
	parsedOutput, err := ParseTerraformOutput(terraformOutput)
	if err != nil {
		return "", fmt.Errorf("error parsing output: %v", err)
	}
	return parsedOutput.DropletIP.Value, nil
}

func runTerraformCommand(dir string, args ...string) ([]byte, error) {
	cmd := exec.Command("terraform", args...)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, fmt.Errorf("error running terraform %s: %v, output: %s", args[0], err, string(output))
	}
	return output, nil
}

func RunTerraformApply(req structs.CreateDropletRequest, token string) ([]byte, error) {
	dir := "terraform/create-droplet"
	if initOutput, err := runTerraformCommand(dir, "init"); err != nil {
		return initOutput, err
	}

	args := []string{
		"apply", "-auto-approve",
		fmt.Sprintf("-var=do_token=%s", token),
		fmt.Sprintf("-var=image=%s", req.Image),
		fmt.Sprintf("-var=region=%s", req.Region),
		fmt.Sprintf("-var=size=%s", req.Size),
		fmt.Sprintf("-var=droplet_name=%s", req.Name),
	}
	return runTerraformCommand(dir, args...)
}

func RunTerraformListDroplets(token string) ([]byte, error) {
	dir := "terraform/list-droplets"
	if initOutput, err := runTerraformCommand(dir, "init"); err != nil {
		return initOutput, err
	}
	return runTerraformCommand(dir, "apply", "-auto-approve", fmt.Sprintf("-var=do_token=%s", token))
}

func GetTerraformOutput(dir string) ([]byte, error) {
	return runTerraformCommand(dir, "output", "-json")
}

func ParseTerraformOutput(output []byte) (structs.CreateDropletOutput, error) {
	var terraformOutput structs.CreateDropletOutput
	if err := json.Unmarshal(output, &terraformOutput); err != nil {
		return terraformOutput, fmt.Errorf("error parsing terraform output: %v", err)
	}
	return terraformOutput, nil
}

func ParseDropletListOutput(output []byte) ([]structs.Droplet, error) {
	var terraformOutput map[string]structs.DropletListOutput
	if err := json.Unmarshal(output, &terraformOutput); err != nil {
		return nil, fmt.Errorf("error parsing terraform output: %v", err)
	}
	dropletList, ok := terraformOutput["droplet_list"]
	if !ok {
		return nil, fmt.Errorf("droplet_list key not found in output")
	}
	return dropletList.Value, nil
}
