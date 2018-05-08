package console



var commands = []Command{
	new(CommandHelp),
	new(CommandCPUProf),
	new(CommandProf),
}

type Command interface {
	// must goroutine safe
	name() string
	// must goroutine safe
	help() string
	// must goroutine safe
	run(args []string) string
}
