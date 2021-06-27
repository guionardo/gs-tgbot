package behaviors

import "fmt"

type Runner struct {
	behaviors    map[string]CommandBehavior
}

func (r *Runner) AddBehavior(behavior CommandBehavior) {
	command:=behavior.Command()
	r.behaviors[command] = behavior
}

func (r *Runner) Run(command string) string {
	if command == "help" {
		return r.Help()
	}
	for cmd, behavior := range r.behaviors {
		if cmd == command {
			return behavior.DoCommand(command)
		}
	}
	return ""
}

func (r *Runner) Help() string {
	var help = "Commands\n\n"
	for cmd,behavior:=range r.behaviors{
		help+=fmt.Sprintf("%s : %s\n",cmd,behavior.Description())
	}

	return help
}

func InitRunner() *Runner {
	behaviors := make(map[string]CommandBehavior)
	return &Runner{
		behaviors,
	}
}
