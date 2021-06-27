package behaviors

import (
	"os/exec"
)

type RebootBehavior struct {
}

func (i *RebootBehavior) DoCommand(command string) string {
	out := exec.Command("shutdown", "-r now")

	defer out.Run()
	return "Rebooting down..."
}

func (i *RebootBehavior) Description() string {
	return "Reboots raspberry pi"
}

func (i *RebootBehavior) Command() string {
	return "reboot"
}
