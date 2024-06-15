package svc

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Desgue/cloud-candidate-challenge-001/src/terraform"
	"github.com/Desgue/cloud-candidate-challenge-001/src/terraform/droplet"
)

type DropletService struct {
	Dir    string
	Main   string
	Output string
	Tfvars string
}

func (s DropletService) CreateDroplet(req terraform.DropletRequest) (*terraform.DropletResponse, error) {
	var err error
	if !req.IsValid() {
		return nil, fmt.Errorf("invalid request")
	}
	err = s.createDir(req)
	if err != nil {
		return nil, err
	}
	err = s.createTfvars(req)
	if err != nil {
		return nil, err
	}
	err = s.createStaticFiles()
	if err != nil {
		return nil, err
	}
	err = s.runTerraformInit()
	if err != nil {
		return nil, err
	}
	err = s.runTerraformApply()
	if err != nil {
		return nil, err
	}

	return s.terraformOutput()
}

func (s *DropletService) Cleanup() error {
	err := os.RemoveAll(s.Dir)
	if err != nil {
		return err
	}
	s.Dir = ""
	s.Main = ""
	s.Output = ""
	s.Tfvars = ""
	return nil
}

func (s *DropletService) createDir(req terraform.DropletRequest) error {
	tempDir, err := os.MkdirTemp("", fmt.Sprintf("droplet-%s-%s", req.Name, req.Region))
	if err != nil {
		return err
	}
	s.Dir = tempDir
	return nil

}
func (s *DropletService) createTfvars(req terraform.DropletRequest) error {
	// Parse the request and create terraform files
	fileName := fmt.Sprintf("droplet-%s-%s.tfvars", req.Name, req.Region)
	fileContent := fmt.Sprintf(
		droplet.TfvarsModel,
		req.Token,
		req.Image,
		req.Name,
		req.Region,
		req.Size,
		req.Monitoring,
		req.Ipv6,
	)

	err := os.WriteFile(filepath.Join(s.Dir, fileName), []byte(fileContent), 0644)
	if err != nil {
		return err
	}
	s.Tfvars = fileName
	return nil
}

func (s *DropletService) createStaticFiles() error {
	err := os.WriteFile(filepath.Join(s.Dir, "main.tf"), []byte(droplet.MainModel), 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(s.Dir, "output.tf"), []byte(droplet.OutputModel), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s DropletService) runTerraformInit() error {
	cmd := exec.Command("terraform", "init")
	cmd.Dir = s.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (s DropletService) runTerraformApply() error {
	varFile := fmt.Sprintf(`-var-file="%s"`, s.Tfvars)
	cmd := exec.Command("terraform", "apply", "-auto-approve", varFile)
	cmd.Dir = s.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (s DropletService) terraformOutput() (*terraform.DropletResponse, error) {
	var dropletResponse *terraform.DropletResponse
	cmd := exec.Command("terraform", "output", "-json")
	cmd.Dir = s.Dir
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(output, dropletResponse)
	if err != nil {
		return nil, err
	}
	return dropletResponse, nil
}
