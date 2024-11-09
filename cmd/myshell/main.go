package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
		} else if command == "echo" {
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))
		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
		}

	}
}
