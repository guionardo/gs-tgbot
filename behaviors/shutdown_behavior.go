package behaviors

import (
	"fmt"
	"os/exec"
)

type ShutdownBehavior struct {
}

func (i *ShutdownBehavior) DoCommand(command string) string {
	err := exec.Command("sudo","shutdown", "-h", "+1").Run()
	if err != nil {
		return fmt.Sprintf("Error on shutdown: %s", err.Error())
	}

	return "Shutting down in one minute..."
}

func (i *ShutdownBehavior) Description() string {
	return "Shuts down raspberry pi"
}
func (i *ShutdownBehavior) Command() string {
	return "shutdown"
}
