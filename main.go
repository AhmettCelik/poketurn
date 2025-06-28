package main

import (
	"bufio"
	"fmt"
	"os"
)

func start(isGameOn bool, scanner bufio.Scanner) {
	for isGameOn {
		fmt.Println("Welcome to poketurn!")
		fmt.Println("type help to see commands.")
		fmt.Print(">")
		scanner.Scan()
		input := scanner.Text()
		fmt.Print(input)
	}
}

func main() {
	isGameOn := true
	scanner := bufio.NewScanner(os.Stdin)
	start(isGameOn, *scanner)
}
