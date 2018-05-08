package console



// help
type CommandHelp struct{}

func (c *CommandHelp) name() string {
	return "help"
}

func (c *CommandHelp) help() string {
	return "this help text"
}

func (c *CommandHelp) run([]string) string {
	output := "Commands:\r\n"
	for _, c := range commands {
		output += c.name() + " - " + c.help() + "\r\n"
	}
	output += "exit - exit console"

	return output
}
