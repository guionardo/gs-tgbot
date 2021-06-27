package behaviors

import (
	"os/exec"
)

type ShutdownBehavior struct {
}

func (i *ShutdownBehavior) DoCommand(command string) string {
	out := exec.Command("shutdown", "-h now")

	defer out.Run()
	return "Shutting down..."
}

func (i *ShutdownBehavior) Description() string{
	return "Shuts down raspberry pi"
}
func (i *ShutdownBehavior) Command() string{
	return "shutdown"
}