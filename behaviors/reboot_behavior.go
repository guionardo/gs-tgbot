package behaviors

import (
	"log"
	"os/exec"
)

type RebootBehavior struct {
}

func (i *RebootBehavior) DoCommand(command string) string {
	err := exec.Command("shutdown", "-r now").Run()
	if err!=nil{
		log.Printf("Error on rebooting: %s",err.Error())
	}

	return "Rebooting..."
}

func (i *RebootBehavior) Description() string {
	return "Reboots raspberry pi"
}

func (i *RebootBehavior) Command() string {
	return "reboot"
}
