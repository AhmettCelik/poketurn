package models

type Command struct {
	Name        string
	Description string
	Callback    func() error
}
