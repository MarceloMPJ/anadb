package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		printPrompt()
		if readNextLine(scanner) {
			execComand(scanner.Text())
		}
	}
}

func printPrompt() {
	fmt.Printf("anadb > ")
}

func readNextLine(scanner *bufio.Scanner) bool {
	return scanner.Scan()
}

func execComand(command string) {
	switch command {
	case ".exit":
		fmt.Println("Saindo :*")
		os.Exit(0)
	default:
		fmt.Printf("Unrecognized command '%s'.\n", command)
	}
}
