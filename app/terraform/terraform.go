package terraform

import (
	"os/exec"

	"go_webserver/loggers"
)

// Execute terraform apply of the requested module
// Capture any error executing the command using the error channel
// Send output through output channel
func CallTerraformModule(moduleName string, outputChan chan<- string, errChan chan<- error) error {
    dir := "./terraform/modules/" + moduleName
    
    cmd := exec.Command("terraform", "apply", "-auto-approve")
    cmd.Dir = dir
    
    output, err := cmd.CombinedOutput()
    if err != nil {
        loggers.LogError.Printf("Command %s failed: %v\n", cmd, err)
        errChan <- err
        return err
    }

    outputChan <- string(output)
    
    return nil
}
