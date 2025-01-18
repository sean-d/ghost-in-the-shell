package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		// Wait for user input
		fmt.Fprint(os.Stdout, "$ ")

		reader := bufio.NewReader(os.Stdin)
		userInput, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}

		userInput = strings.TrimSpace(userInput)

		switch userInput {
		case "exit 0":
			os.Exit(0)
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", userInput)
		}

	}
}
