package commands

import "fmt"

const ascii string = `

`

func Help() error {
	fmt.Printf("%s\n", ascii)
	fmt.Println("There are no help in the void...")
	return nil
}
