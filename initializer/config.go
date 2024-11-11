package initializer

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"
)

type Config struct {
	Username string
	Hostname string
}

var config Config

func GetConfig() *Config {
	return &config
}

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func CreateConfig() error {
	fmt.Print("Enter your username (e.g., patrick): ")
	username := ReadInput()

	fmt.Print("Enter computer hostname (e.g., dell): ")
	hostname := ReadInput()

	config = Config{
		Username: username,
		Hostname: hostname,
	}

	usr, err := user.Current()
	if err != nil {
		return err
	}

	baseDir := usr.HomeDir + "/linux"
	homeDir := baseDir + "/home/" + username
	tempDir := baseDir + "/temp"
	varsDir := baseDir + "/vars"
	directories := []string{homeDir, tempDir, varsDir}

	for _, dir := range directories {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0750); err != nil {
				return err
			}
			fmt.Printf("Directory created: %s\n", dir)
		} else {
			fmt.Printf("Directory already exists: %s\n", dir)
		}
	}

	fmt.Println("Operation complete. Continuing...")
	time.Sleep(3 * time.Second)

	return nil
}
