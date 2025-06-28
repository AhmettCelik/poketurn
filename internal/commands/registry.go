package commands

func Register(cmds map[string]CliCommand, name string, description string, callback func(name string) error) {
	cmds[name] = CliCommand{
		name:        name,
		description: description,
		callback:    callback,
	}
}

func LoadDefaults(cmds map[string]CliCommand) {
	Register(cmds, "help", "Shows commands and their descriptions.", Help)
}
