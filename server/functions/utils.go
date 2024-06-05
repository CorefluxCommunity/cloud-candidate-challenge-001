package functions

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"server/structs"
)

func RunTerraformApply(req structs.TerraformRequest) ([]byte, error) {
	cmd := exec.Command("terraform", "apply", "-auto-approve",
		fmt.Sprintf("-var=do_token=%s", req.Token),
		fmt.Sprintf("-var=image=%s", req.Image),
		fmt.Sprintf("-var=region=%s", req.Region),
		fmt.Sprintf("-var=size=%s", req.Size),
		fmt.Sprintf("-var=droplet_name=%s", req.DropletName),
	)
	cmd.Dir = "/terraform/create-droplet" // Certifique-se de que este caminho esteja correto
	return cmd.CombinedOutput()
}

func GetTerraformOutput() ([]byte, error) {
	cmd := exec.Command("terraform", "output", "-json")
	cmd.Dir = "/terraform/create-droplet" // Certifique-se de que este caminho esteja correto
	return cmd.Output()
}

func ParseTerraformOutput(output []byte) (structs.TerraformOutput, error) {
	var terraformOutput structs.TerraformOutput
	err := json.Unmarshal(output, &terraformOutput)
	return terraformOutput, err
}
