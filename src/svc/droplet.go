package svc

import (
	"fmt"
	"os"
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

func (s *DropletService) CreateDir(req terraform.DropletRequest) error {
	if !req.IsValid() {
		return fmt.Errorf("invalid request")
	}
	tempDir, err := os.MkdirTemp("", fmt.Sprintf("droplet-%s-%s", req.Name, req.Region))
	if err != nil {
		return err
	}
	s.Dir = tempDir
	return nil

}
func (s *DropletService) CreateTfvars(req terraform.DropletRequest) error {
	// Parse the request and create terraform files
	if !req.IsValid() {
		return fmt.Errorf("invalid request")
	}
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

func (s *DropletService) CreateStaticFiles() error {
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
