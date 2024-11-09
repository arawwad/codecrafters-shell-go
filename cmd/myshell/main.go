package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

var builtins = []string{
	"exit",
	"echo",
	"type",
	"pwd",
	"cd",
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
		if command == "pwd" {
			pwdCommand()
			continue
		}
		if command == "cd" {
			cdCommand(args[0])
			continue
		}
		if path, ok := getPath(command); ok {
			cmd := exec.Command(path, args...)
			output, err := cmd.CombinedOutput()

			if err != nil {
				fmt.Fprintf(os.Stdout, err.Error())
			}

			if _, err = os.Stdout.Write(output); err != nil {
				fmt.Fprintf(os.Stdout, err.Error())
			}

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

	if slices.Contains(builtins, commandName) {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", commandName)
		return
	}
	if path, ok := getPath(commandName); ok {
		fmt.Fprintf(os.Stdout, "%s is %s\n", commandName, path)
		return
	}
	fmt.Fprintf(os.Stdout, "%s: not found\n", commandName)
}

func getPath(commandName string) (string, bool) {
	paths := strings.Split(os.Getenv("PATH"), ":")

	for _, path := range paths {
		pathFiles, _ := os.ReadDir(path)
		for _, dirEntry := range pathFiles {
			if dirEntry.Name() == commandName {
				return fmt.Sprintf("%s/%s", path, commandName), true
			}
		}
	}

	return "", false
}

func pwdCommand() {
	pwd, err := os.Getwd()
	if err != nil {
		println(err)
	} else {
		println(pwd)
	}
}

func cdCommand(path string) {
	if err := os.Chdir(path); err != nil {
		fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", path)
	}
}
