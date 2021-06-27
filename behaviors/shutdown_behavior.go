package behaviors

import (
	"fmt"
	"os/exec"
)

type ShutdownBehavior struct {
}

func (i *ShutdownBehavior) DoCommand(command string) string {
	err := exec.Command("shutdown", "-h", "now").Run()
	if err != nil {
		return fmt.Sprintf("Error on shutdown: %s", err.Error())
	}

	return "Shutting down..."
}

func (i *ShutdownBehavior) Description() string {
	return "Shuts down raspberry pi"
}
func (i *ShutdownBehavior) Command() string {
	return "shutdown"
}
