package behaviors

import (
	"fmt"
	"os/exec"
)

type RebootBehavior struct {
}

func (i *RebootBehavior) DoCommand(command string) string {
	err := exec.Command("shutdown", "-r +1").Run()
	if err != nil {
		return fmt.Sprintf("Error on rebooting: %s", err.Error())
	}

	return "Rebooting in one minute..."
}

func (i *RebootBehavior) Description() string {
	return "Reboots raspberry pi"
}

func (i *RebootBehavior) Command() string {
	return "reboot"
}
