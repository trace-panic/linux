package commands

import (
	"fmt"
	"os"

	"github.com/trace-panic/linux/config"
)

// RunCommand interprets the input string and runs the appropriate command.
func RunCommand(input string) {
	switch input {
	case "pwd":
		pwd()
	case "exit":
		exit()
	default:
		fmt.Println("Command not found:", input)
	}
}

func pwd() {
	cfg := config.GetConfig()
	fmt.Println(cfg.CWD)
}

func exit() {
	cfg := config.GetConfig()
	cfg.SaveConfig()
	os.Exit(0)
}
