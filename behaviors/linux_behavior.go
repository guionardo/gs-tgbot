package behaviors

import (
	"fmt"
	"log"
	"os/exec"
)

type LinuxBehavior struct {
}

func getCommand(command string, args ...string) string {
	out, err := exec.Command(command, args...).Output()
	if err != nil {
		log.Printf("linux %s error: %s", command, err.Error())
		return fmt.Sprintf("ERROR %s : %s", command, err.Error())
	} else {
		log.Printf("linux %s: %s", command, string(out))
		return fmt.Sprintf("%s: %s", command, string(out))
	}
}

func getLinux() string {
	return getCommand("lsb_release", "-a")
}

func getUptime() string {
	return getCommand("uptime")
}

func getUsers() string {
	return getCommand("w")
}
func (i *LinuxBehavior) DoCommand(command string) string {
	out := fmt.Sprintf("%s\n%s\n%s", getLinux(), getUptime(), getUsers())

	return out
}

func (i *LinuxBehavior) Description() string {
	return "Shows linux info"
}
func (i *LinuxBehavior) Command() string {
	return "linux"
}
