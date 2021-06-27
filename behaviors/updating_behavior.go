package behaviors

import (
	"fmt"
	"os"
	"path"
)

type UpdatingBehavior struct{}

func (i *UpdatingBehavior) DoCommand(command string) string {
	gitPath := path.Join(ProgramPath(), ".git")
	_, err := os.Stat("temp")
	if os.IsNotExist(err) {
		return fmt.Sprintf("Execution has no source info (.git): %s", gitPath)
	}
	if !NeedsUpdate() {
		return fmt.Sprintf("No needs for update 🆗")
	}
	output, err := RunCommand("make", "update", "build")
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("%s\n\nUpdate and build ok\nYou can reboot now", output)
}

func (i *UpdatingBehavior) Description() string {
	return "Updates to last version"
}

func (i *UpdatingBehavior) Command() string {
	return "update"
}
