package functions

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"src/structs"
)

func RunTerraformInit(dir string) ([]byte, error) {
	cmd := exec.Command("terraform", "init")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, fmt.Errorf("error running terraform init: %v, output: %s", err, string(output))
	}
	return output, nil
}

func RunTerraformApply(req structs.CreateDropletRequest, token string) ([]byte, error) {
	dir := "terraform/create-droplet"
	initOutput, err := RunTerraformInit(dir)
	if err != nil {
		return initOutput, fmt.Errorf("error running terraform init: %v, output: %s", err, string(initOutput))
	}
	cmd := exec.Command("terraform", "apply", "-auto-approve",
		fmt.Sprintf("-var=do_token=%s", token),
		fmt.Sprintf("-var=image=%s", req.Image),
		fmt.Sprintf("-var=region=%s", req.Region),
		fmt.Sprintf("-var=size=%s", req.Size),
		fmt.Sprintf("-var=droplet_name=%s", req.Name),
	)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, fmt.Errorf("error running terraform apply: %v, output: %s", err, string(output))
	}
	return output, nil
}

func RunTerraformListDroplets(token string) ([]byte, error) {
	dir := "terraform/list-droplets"
	initOutput, err := RunTerraformInit(dir)
	if err != nil {
		return initOutput, fmt.Errorf("error running terraform init: %v, output: %s", err, string(initOutput))
	}
	cmd := exec.Command("terraform", "apply", "-auto-approve",
		fmt.Sprintf("-var=do_token=%s", token),
	)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, fmt.Errorf("error running terraform apply: %v, output: %s", err, string(output))
	}
	return output, nil
}

func GetTerraformOutput(dir string) ([]byte, error) {
	cmd := exec.Command("terraform", "output", "-json")
	cmd.Dir = dir

	output, err := cmd.Output()
	if err != nil {
		return output, fmt.Errorf("error getting terraform output: %v, output: %s", err, string(output))
	}
	return output, nil
}

func ParseTerraformOutput(output []byte) (structs.CreateDropletOutput, error) {
	var terraformOutput structs.CreateDropletOutput
	err := json.Unmarshal(output, &terraformOutput)
	if err != nil {
		return terraformOutput, fmt.Errorf("error parsing terraform output: %v", err)
	}
	return terraformOutput, nil
}

func ParseDropletListOutput(output []byte) ([]structs.Droplet, error) {
	var terraformOutput map[string]structs.DropletListOutput
	err := json.Unmarshal(output, &terraformOutput)
	if err != nil {
		return nil, fmt.Errorf("error parsing terraform output: %v", err)
	}

	dropletList, ok := terraformOutput["droplet_list"]
	if !ok {
		return nil, fmt.Errorf("error converting terraform output: %v", err)
	}

	return dropletList.Value, nil
}
