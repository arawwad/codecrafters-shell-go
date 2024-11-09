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
		} else {
			command := strings.TrimSpace(input)
			fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
		}

	}
}
