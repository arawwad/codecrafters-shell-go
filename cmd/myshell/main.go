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
		prompt()

		command, args, ok := getInput()

		if !ok {
			continue
		}

		if command == "exit" {
			exitCommand()
		}
		if command == "echo" {
			echoCommand(args)
			continue
		}
		if command == "type" {
			typeCommand(args[0])
			continue
		}
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}

func prompt() {
	fmt.Fprint(os.Stdout, "$ ")
}

func getInput() (string, []string, bool) {
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return "", nil, false
	}
	inputParts := strings.Fields(strings.TrimSpace(input))
	if len(inputParts) == 0 {
		return "", nil, false
	}
	return inputParts[0], inputParts[1:], true
}

func exitCommand() {
	os.Exit(0)
}

func echoCommand(args []string) {
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))
}

func typeCommand(commandName string) {

	paths := strings.Split(os.Getenv("PATH"), ":")

	if slices.Contains(builtins, commandName) {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", commandName)
		return
	}
	for _, path := range paths {
		pathFiles, _ := os.ReadDir(path)
		for _, dirEntry := range pathFiles {
			if dirEntry.Name() == commandName {
				fmt.Fprintf(os.Stdout, "%s is %s/%s\n", commandName, path, commandName)
				return
			}
		}
	}
	fmt.Fprintf(os.Stdout, "%s: not found\n", commandName)
}
