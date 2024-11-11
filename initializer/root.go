package initializer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Username      string
	Hostname      string
	RootDirectory string
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

	fmt.Print("Enter root directory name (e.g., filesystem): ")
	rootDirectory := ReadInput()

	config = Config{
		Username:      username,
		Hostname:      hostname,
		RootDirectory: rootDirectory,
	}

	if err := os.Mkdir(rootDirectory, 0750); err != nil {
		return err
	}

	return nil
}
