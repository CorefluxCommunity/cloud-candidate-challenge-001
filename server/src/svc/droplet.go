package svc

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Desgue/cloud-candidate-challenge-001/src/domain"
)

type DropletService struct {
	dir    string
	Main   string
	Output string
	Tfvars string
}

func NewDropletService() *DropletService {
	return &DropletService{dir: "../terraform"}
}

func (s *DropletService) CreateDroplet(req domain.DropletRequest) (*domain.DropletResponse, error) {
	log.Println("Creating DigitalOcean Droplet")
	var err error
	if !req.IsValid() {
		return nil, fmt.Errorf("invalid or missing request fields")
	}
	errCh := make(chan error, 1)
	resCh := make(chan *domain.DropletResponse, 1)
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

func (s DropletService) DeleteDroplet() error {
	errCh := make(chan error)

	go func() {
		err := s.runTerraformInit()
		if err != nil {
			errCh <- err
			return
		}
		err = s.runTerraformDestroy()
		if err != nil {
			errCh <- err
			return
		}
		errCh <- nil
	}()

	err := <-errCh
	return err

}

func (s DropletService) runTerraformInit() error {
	log.Println("Running terraform init")
	awsEnv := NewAwsEnv()
	cmd := exec.Command("terraform", "init")
	cmd.Env = append(
		cmd.Env,
		fmt.Sprintf("AWS_ACCESS_KEY_ID=%s", awsEnv.AccessKeyID),
		fmt.Sprintf("AWS_SECRET_ACCESS_KEY=%s", awsEnv.SecretAccessKey),
	)

	if awsEnv.SessionToken != "" {
		cmd.Env = append(cmd.Env, fmt.Sprintf("AWS_SESSION_TOKEN=%s", awsEnv.SessionToken))
	}
	cmd.Dir = s.dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println("Error running Terraform Init")
		return err
	}
	return nil
}

func (s DropletService) runTerraformApply(req domain.DropletRequest) error {
	log.Println("Running terraform apply")
	token := os.Getenv("DIGITALOCEAN_API_TOKEN")
	awsRegion := os.Getenv("AWS_REGION")
	args := []string{
		"apply",
		"-auto-approve",
		fmt.Sprintf(`-var=aws_region=%s`, awsRegion),
		fmt.Sprintf(`-var=api_token=%s`, token),
		fmt.Sprintf(`-var=image=%s`, req.Image),
		fmt.Sprintf(`-var=name=%s`, req.Name),
		fmt.Sprintf(`-var=region=%s`, req.Region),
		fmt.Sprintf(`-var=size=%s`, req.Size),
		fmt.Sprintf(`-var=monitoring=%t`, req.Monitoring),
		fmt.Sprintf(`-var=ipv6=%t`, req.Ipv6),
	}
	cmd := exec.Command("terraform", args...)
	cmd.Dir = s.dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("Error running creating droplet: %s", err.Error())
		return err
	}
	return nil
}

func (s DropletService) runTerraformDestroy() error {
	log.Println("Running terraform apply -destroy")
	token := os.Getenv("DIGITALOCEAN_API_TOKEN")
	args := []string{
		"destroy",
		"-auto-approve",
		fmt.Sprintf("-var=api_token=%s", token),
	}
	cmd := exec.Command("terraform", args...)
	cmd.Dir = s.dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("Error running terraform apply -destroy\n%s", err.Error())
		return err
	}
	return nil
}

func (s DropletService) terraformOutput() (*domain.DropletResponse, error) {
	log.Println("Running terraform output")
	cmd := exec.Command("terraform", "output", "-json")
	cmd.Dir = s.dir
	output, err := cmd.Output()
	if err != nil {
		log.Println("Error running terraform output")
		return nil, err
	}
	var dropletResponse domain.DropletResponse
	err = json.Unmarshal(output, &dropletResponse)
	if err != nil {
		log.Println("fail to unmarshal response")
		return nil, err
	}
	return &dropletResponse, nil
}
