package commands

func Register(cmds map[string]CliCommand, name string, description string, callback func() error) {
	cmds[name] = CliCommand{
		Name:        name,
		Description: description,
		Callback:    callback,
	}
}

func LoadDefaults(cmds map[string]CliCommand) {
	Register(cmds, "help", "Shows commands and their descriptions.", Help)
}
