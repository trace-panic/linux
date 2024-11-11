package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/trace-panic/linux/initializer"
)

func ReadInput() string {
	cfg := initializer.GetConfig()
	fmt.Printf("%v@%v", cfg.Username, cfg.Hostname)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func clearConsole() {
	fmt.Print("\033[H\033[2J")
}

func BaseLoop() {
	clearConsole()

	for {
		input := ReadInput()
		fmt.Println(input)
	}
}

func main() {
	if err := initializer.CreateConfig(); err != nil {
		log.Fatal("Error creating configuration: ", err)
		return
	}

	BaseLoop()
}
