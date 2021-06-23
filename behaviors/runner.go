package behaviors

type Runner struct {
	behaviors map[string]CommandBehavior
}

func (r *Runner) AddBehavior(command string, behavior CommandBehavior) {
	r.behaviors[command] = behavior
}

func (r *Runner) Run(command string) string {
	for cmd, behavior := range r.behaviors {
		if cmd == command {
			return behavior.DoCommand(command)
		}
	}
	return ""
}

func InitRunner() *Runner {
	behaviors := make(map[string]CommandBehavior)
	return &Runner{
		behaviors: behaviors,
	}
}
