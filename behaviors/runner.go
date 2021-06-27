package behaviors

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type Runner struct {
	behaviors map[string]CommandBehavior
}

func (r *Runner) AddBehavior(behavior CommandBehavior) {
	command := behavior.Command()
	r.behaviors[command] = behavior
}

func (r *Runner) Run(command string) string {
	for cmd, behavior := range r.behaviors {
		if cmd == command {
			return behavior.DoCommand(command)
		}
	}
	return r.Help()
}

func (r *Runner) Help() string {
	var help = "Commands\n\n"
	for cmd, behavior := range r.behaviors {
		help += fmt.Sprintf("%s : %s\n", cmd, behavior.Description())
	}

	return help
}

func InitRunner() *Runner {
	behaviors := make(map[string]CommandBehavior)
	return &Runner{
		behaviors,
	}
}

func ProgramPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func RunCommand(command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = ProgramPath()
	var output, err = cmd.Output()
	if err != nil {
		log.Fatalf("Error when running command %s: %s\n", command, err.Error())
	}
	return output, err
}
