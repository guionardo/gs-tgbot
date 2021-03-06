package behaviors

import (
	"strings"
)

type LastCommitBehavior struct {
}

func (i *LastCommitBehavior) DoCommand(command string) string {
	out, err := RunCommand("git", "log", "-n", "1")
	if err != nil {
		return err.Error()
	}

	updInfo := "is updated š"
	if NeedsUpdate() {
		updInfo = "needs update ā "
	}

	return string(out) + "\nš¤ = " + updInfo
}
func NeedsUpdate() bool {
	out, err := RunCommand("git", "remote", "show", "origin")
	if err != nil {
		return false
	}

	for _, line := range strings.Split(string(out), "\n") {
		if strings.Contains(line, "out of date") {
			return true
		}
	}
	return false
}

func (i *LastCommitBehavior) Description() string {
	return "Shows last commit info"
}
func (i *LastCommitBehavior) Command() string {
	return "version"
}
