package behaviors

import (
	"log"
	"os/exec"
)

type LinuxBehavior struct {
}

func (i *LinuxBehavior) DoCommand(command string) string {
	out, err := exec.Command("lsb_release", "-a").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
