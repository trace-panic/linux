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

func InitWorkflow() {
	fmt.Print("Enter computer hostname (e.g., dell): ")
	hostname := ReadInput()

	fmt.Print("Enter your username (e.g., patrick): ")
	username := ReadInput()

	cfg.SetHostname(hostname)
	cfg.SetUsername(username)
	cfg.SetOSVersion("0.0.1")
}

func Init() error {
	hmdir := filesystem.OSHomeDir()
	cfg.SetBaseDirectory(hmdir + "/linux")
	cfg.SetHomeDirectory(cfg.BaseDirectory + "/home/" + cfg.Username)

	new := !filesystem.DirExists(cfg.HomeDirectory + "/.config")
	if new {
		InitWorkflow()
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
