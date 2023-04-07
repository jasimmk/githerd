package execapi

import (
	"os/exec"
)

// RunShellExec runs a command with shell and returns the output and error
func RunShellExec(command, workDir string) ([]byte, error) {

	cmd := exec.Command("sh", "-c", command)
	cmd.Dir = workDir
	return cmd.CombinedOutput()
}
