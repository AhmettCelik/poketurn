package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AhmettCelik/poketurn/internal/commands"
)

func main() {
	isGameOn := true
	scanner := bufio.NewScanner(os.Stdin)
	cmds := map[string]commands.CliCommand{}
	commands.LoadDefaults(cmds)

	fmt.Println("Welcome to poketurn!")
	fmt.Println("type help to see commands.")

	for isGameOn {
		fmt.Print(">")
		scanner.Scan()
		input := scanner.Text()

		cmd, exists := cmds[input]
		if !exists {
			fmt.Println("Unknown command")
			return
		}

		err := cmd.Callback()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
