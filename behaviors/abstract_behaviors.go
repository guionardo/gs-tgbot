package behaviors

type CommandBehavior interface {
	DoCommand(command string) string
	Description() string
	Command() string
}
