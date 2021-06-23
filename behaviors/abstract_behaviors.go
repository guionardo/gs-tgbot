package behaviors

type CommandBehavior interface {
	DoCommand(command string) string
}
