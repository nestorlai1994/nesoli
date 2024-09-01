package cmdExecutor

import (
	"fmt"
	"os/exec"

	"github.com/nestorlai1994/nesoli/logger"
)

func ExecuteCommand(prog string, args ...string) (string, error) {
	logger.Info("Start executing command: " + fmt.Sprintf("%s %v", prog, args))
	cmd := exec.Command(prog, args...)

	// Capture the command's output
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error("Failed to execute command: "+cmd.String(), err)
		return string(output), err
	}
	logger.Info("Command output: " + string(output))
	return string(output), nil
}
