package svc

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Desgue/cloud-candidate-challenge-001/src/droplet"
)

type DropletService struct {
	Dir    string
	Main   string
	Output string
	Tfvars string
}

func NewDropletService() *DropletService {
	return &DropletService{Dir: "../terraform"}
}

func (s *DropletService) CreateDroplet(req droplet.DropletRequest) (*droplet.DropletResponse, error) {
	log.Println("Creating DigitalOcean Droplet")
	var err error
	if !req.IsValid() {
		return nil, fmt.Errorf("invalid request")
	}

	err = s.runTerraformInit()
	if err != nil {
		return nil, err
	}
	err = s.runTerraformApply(req)
	if err != nil {
		return nil, err
	}

	return s.terraformOutput()
}

func (s DropletService) runTerraformInit() error {
	log.Println("Running terraform init")
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

func (s DropletService) runTerraformApply(req droplet.DropletRequest) error {
	args := []string{
		"apply",
		"-auto-approve",
		fmt.Sprintf(`
		-var="api_token=%s"
		-var="image=%s"
		-var="name=%s"
		-var="region=%s"
		-var="size=%s"
		-var="monitoring=%t"
		-var="ipv6=%t"
		`,
			req.Token,
			req.Image,
			req.Name,
			req.Region,
			req.Size,
			req.Monitoring,
			req.Ipv6,
		),
	}
	cmd := exec.Command("terraform", args...)
	log.Printf("Running %s %s", cmd.Path, cmd.Args[3])
	cmd.Dir = s.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (s DropletService) terraformOutput() (*droplet.DropletResponse, error) {
	var dropletResponse *droplet.DropletResponse
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
