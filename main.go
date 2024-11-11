package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("user@myLinuxSim:~$ ")
		input, _ := reader.ReadString('\n')
		command := strings.TrimSpace(input)

		switch command {
		case "exit":
			fmt.Println("Goodbye!")
			return
		case "ls":
			fmt.Println("Listing files...")
		case "cd":
			fmt.Println("Changing directory...")
		default:
			fmt.Printf("Command not found: %s\n", command)
		}
	}
}
