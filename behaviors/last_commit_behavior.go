package behaviors

import "os/exec"

type LastCommitBehavior struct {

}
func (i *LastCommitBehavior) DoCommand(command string) string {
	out,err := exec.Command("git", "log","-n","1").Output()
	if err!=nil{
		return err.Error()
	}
	return string(out)
}

func (i *LastCommitBehavior) Description() string{
	return "Shows last commit info"
}
func (i *LastCommitBehavior) Command() string{
	return "version"
}

