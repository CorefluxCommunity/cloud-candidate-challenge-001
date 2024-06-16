package main

import (
	"encoding/json"
	"os"
	"os/exec"
)

// Init all terraform modules before routing to ensure everything is prepared for any request
func terraformInit() error {
	for _, dir := range terraformModules {
		cmd := exec.Command("terraform", "init")
		cmd.Dir = "../terraform/" + dir

		if err := cmd.Run(); err != nil {
			logEvent.Printf("Terraform init failed in directory %s: %v", dir, err)
			return err
		}
	}

	return nil
}

// Sets the needed Environment Variables according to the type of request
func setEnvironmentVariables(input interface{}) error {
	os.Setenv("TF_VAR_do_token", os.Getenv("do_token"))
	switch input := input.(type) {
	case *CreateRequest:
		tags, err := json.Marshal(input.Tags)
		if err != nil {
			logError.Printf("Error marshaling JSON tags: %v", err)
			return err
		}
		os.Setenv("TF_VAR_droplet_name", input.DropletName)
		os.Setenv("TF_VAR_region", input.Region)
		os.Setenv("TF_VAR_size", input.Size)
		os.Setenv("TF_VAR_image", input.Image)
		os.Setenv("TF_VAR_tags", string(tags))
	case *SearchRequest:
		tags, err := json.Marshal(input.TagToFind)
		if err != nil {
			logError.Printf("Error marshaling JSON tags: %v", err)
			return err
		}
		os.Setenv("TF_VAR_tag_to_find", string(tags))
	case *SortRequest:
		os.Setenv("TF_VAR_direction", input.Direction)
	default:
		logError.Println("Unsupported request type")
	}
	return nil
}

// Execute terraform apply of the requested module
// Capture any error executing the command using the error channel
// Send output through output channel
func callTerraformModule(moduleName string, outputChan chan<- string, errChan chan<- error) error {
    dir := "../terraform/" + moduleName
    
    cmd := exec.Command("terraform", "apply", "-auto-approve")
    cmd.Dir = dir
    
    output, err := cmd.CombinedOutput()
    if err != nil {
        logError.Printf("Command %s failed: %v\n", cmd, err)
        errChan <- err
        return err
    }

    outputChan <- string(output)
    
    return nil
}
