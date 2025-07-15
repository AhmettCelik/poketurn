package commands

import "fmt"

func Market() error {
	fmt.Println("Welcome to the market.")
	fmt.Println("Common chest 100 (type 1)")
	fmt.Println("Rare chest 200 (type 2)")
	fmt.Println("Epic chest 300 (type 3)")
	return nil
}