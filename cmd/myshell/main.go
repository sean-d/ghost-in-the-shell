package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		commands := strings.Fields(userInput)
		command := commands[0]
		args := commands[1:]

		switch command {
		case "exit":
			code, _ := strconv.Atoi(args[0])
			os.Exit(code)
		case "echo":
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))

		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", userInput)
		}

	}
}
