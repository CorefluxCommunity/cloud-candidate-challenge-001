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
	errCh := make(chan error, 1)
	resCh := make(chan *droplet.DropletResponse, 1)

	go func() {
		defer close(errCh)
		defer close(resCh)
		err = s.runTerraformInit()
		if err != nil {
			errCh <- err
			return
		}
		err = s.runTerraformApply(req)
		if err != nil {
			errCh <- err
			return
		}

		res, err := s.terraformOutput()
		if err != nil {
			errCh <- err
			return
		}

		resCh <- res

	}()

	select {
	case err := <-errCh:
		return nil, err
	case res := <-resCh:
		return res, nil
	}

}

func (s DropletService) runTerraformInit() error {
	log.Println("Running terraform init")
	awsEnv := NewAwsEnv()
	cmd := exec.Command("terraform", "init")
	cmd.Env = []string{
		fmt.Sprintf("AWS_ACCESS_KEY_ID=%s", awsEnv.AccessKeyID),
		fmt.Sprintf("AWS_SECRET_ACCESS_KEY=%s", awsEnv.SecretAccessKey),
	}
	if awsEnv.SessionToken != "" {
		cmd.Env = append(cmd.Env, fmt.Sprintf("AWS_SESSION_TOKEN=%s", awsEnv.SessionToken))
	}
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
		fmt.Sprintf(`-var=api_token=%s`, req.Token),
		fmt.Sprintf(`-var=image=%s`, req.Image),
		fmt.Sprintf(`-var=name=%s`, req.Name),
		fmt.Sprintf(`-var=region=%s`, req.Region),
		fmt.Sprintf(`-var=size=%s`, req.Size),
		fmt.Sprintf(`-var=monitoring=%t`, req.Monitoring),
		fmt.Sprintf(`-var=ipv6=%t`, req.Ipv6),
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
	cmd := exec.Command("terraform", "output", "-json")
	cmd.Dir = s.Dir
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var dropletResponse droplet.DropletResponse
	err = json.Unmarshal(output, &dropletResponse)
	if err != nil {
		log.Println("fail to unmarshal response")
		return nil, err
	}
	return &dropletResponse, nil
}