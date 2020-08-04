package main

import (
	"bufio"
	"fmt"
	"os"
)

type metaCommandResult string

const (
	metaCommandSuccess             = "metaCommandSuccess"
	metaCommandUnrecognizedCommand = "metaCommandUnrecognizedCommand"
)

type prepareCommandResult string

const (
	prepareSuccess               = "prepareSuccess"
	prepareUnrecognizedStatement = "prepareUnrecognizedStatement"
)

type statementType string

const (
	statementInsert = "Insert"
	statementSelect = "Select"
)

type statement struct {
	statementType statementType
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printPrompt()
		scanner.Scan()
		input := scanner.Text()

		statement := statement{}

		if input[0] == '.' {
			switch execComand(input) {
			case metaCommandSuccess:
				continue
			case metaCommandUnrecognizedCommand:
				fmt.Printf("Unrecognized command '%s'.\n", input)
				continue
			}
		} else {
			switch prepareStatement(input, &statement) {
			case prepareSuccess:
				executeStatement(&statement)
				continue
			case prepareUnrecognizedStatement:
				fmt.Printf("Unrecognized keyword at start of '%s'.\n", input)
				continue
			}
		}
	}
}

func printPrompt() {
	fmt.Printf("anadb > ")
}

func execComand(command string) metaCommandResult {
	switch command {
	case ".exit":
		fmt.Println("😘")
		os.Exit(0)
	default:
		return metaCommandUnrecognizedCommand
	}

	return metaCommandSuccess
}

func prepareStatement(inputStatement string, statement *statement) prepareCommandResult {
	if len(inputStatement) < 6 {
		return prepareUnrecognizedStatement
	}

	if inputStatement[:6] == "insert" {
		statement.statementType = statementInsert
		return prepareSuccess
	}

	if inputStatement[:6] == "select" {
		statement.statementType = statementSelect
		return prepareSuccess
	}

	return prepareUnrecognizedStatement
}

func executeStatement(statement *statement) {
	switch statement.statementType {
	case statementInsert:
		fmt.Println("Executando INSERT")
	case statementSelect:
		fmt.Println("Executando SELECT")
	}
}
