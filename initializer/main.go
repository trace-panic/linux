package initializer

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/trace-panic/linux/config"
	"github.com/trace-panic/linux/console"
	"github.com/trace-panic/linux/filesystem"
)

func CreateFileSytem(cfg config.Config) error {
	directories := []string{
		cfg.HomeDirectory + "/.config",
	}

	if err := filesystem.CreateDirs(directories); err != nil {
		return err
	}

	return nil
}

func Init() error {
	cfg := config.GetConfig()
	hmdir := filesystem.OSHomeDir()

	fmt.Print("Enter computer hostname (e.g., dell): ")
	hostname := console.ReadInput()

	fmt.Print("Enter your username (e.g., patrick): ")
	username := console.ReadInput()

	// Potential issues when username is empty
	cfg.Hostname = hostname
	cfg.Username = username
	cfg.OSVersion = "0.0.1"
	cfg.BaseDirectory = filepath.Join(hmdir, "linux")
	cfg.HomeDirectory = filepath.Join(cfg.BaseDirectory, "home", cfg.Username)
	cfg.CWD = cfg.HomeDirectory

	if err := CreateFileSytem(*cfg); err != nil {
		return err
	}

	cfg.SaveConfig()

	fmt.Println("Setup complete. Continuing...")
	time.Sleep(2 * time.Second)
	console.ClearConsole()
	return nil
}
