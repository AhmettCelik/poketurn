package commands

import "github.com/AhmettCelik/poketurn/internal/models"

func Register(cmds map[string]models.Command, name string, description string, callback func() error) {
	cmds[name] = models.Command{
		Name:        name,
		Description: description,
		Callback:    callback,
	}
}

func LoadDefaults(cmds map[string]models.Command) {
	Register(cmds, "help", "Shows commands and their descriptions.", Help)
	Register(cmds, "market", "opens market.", Market)
}
