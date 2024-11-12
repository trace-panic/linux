package initializer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/trace-panic/linux/config"
	"github.com/trace-panic/linux/filesystem"
	"github.com/trace-panic/linux/writter"
)

var cfg = config.Config{}

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func InitWorkflow() error {
	fmt.Print("Enter computer hostname (e.g., dell): ")
	hostname := ReadInput()

	fmt.Print("Enter your username (e.g., patrick): ")
	username := ReadInput()

	if err := cfg.SetHostname(hostname); err != nil {
		return err
	}
	if err := cfg.SetUsername(username); err != nil {
		return err
	}
	cfg.SetOSVersion("0.0.1")
	return nil
}

func Init() error {
	hmdir := filesystem.OSHomeDir()
	if err := cfg.SetBaseDirectory(hmdir + "/linux"); err != nil {
		return err
	}
	if err := cfg.SetHomeDirectory(cfg.BaseDirectory + "/home/" + cfg.Username); err != nil {
		return err
	}

	new := !filesystem.DirExists(cfg.HomeDirectory + "/.config")
	if new {
		if err := InitWorkflow(); err != nil {
			return err
		}
	} else {
		// Load existing config
		fmt.Println("Loading existing config")
	}

	if err := filesystem.Init(cfg); err != nil {
		return err
	}

	fmt.Println("Setup complete. Continuing...")
	time.Sleep(2 * time.Second)
	writter.ClearConsole()
	return nil
}
