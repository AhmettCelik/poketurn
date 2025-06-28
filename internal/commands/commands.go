package commands

type CliCommand struct {
	name        string
	description string
	callback    func(name string) error
}
