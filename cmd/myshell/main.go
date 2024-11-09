package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var builtins = []string{
	"exit",
	"echo",
	"type",
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		inputParts := strings.Fields(strings.TrimSpace(input))
		if len(inputParts) == 0 {
			continue
		}
		command := inputParts[0]
		args := inputParts[1:]
		if command == "exit" {
			os.Exit(0)
			continue
		}
		if command == "echo" {
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))
			continue
		}
		if command == "type" {
			if slices.Contains(builtins, args[0]) {
				fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", args[0])
			} else {
				fmt.Fprintf(os.Stdout, "%s: not found\n", args[0])
			}
			continue
		}
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)

	}
}
