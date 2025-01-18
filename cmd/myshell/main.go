package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var builtins [2]string = [2]string{"echo", "exit"}

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
		case "type":
			if len(args) == 0 {
				fmt.Fprintln(os.Stderr, "usage: type [command]")
				break
			}
			switch args[0] {
			case "exit", "echo", "type":
				fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", args[0])
			default:
				fmt.Fprintf(os.Stdout, "%s: not found\n", args[0])
			}
		case "exit":
			if len(args) == 0 {
				fmt.Fprintf(os.Stdout, "usage: exit [code]\n")
				break
			}
			code, _ := strconv.Atoi(args[0])
			os.Exit(code)
		case "echo":
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))

		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", userInput)
		}

	}
}
