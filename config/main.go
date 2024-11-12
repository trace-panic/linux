package config

import (
	"path/filepath"
	"sync"

	"github.com/trace-panic/linux/filesystem"
	"github.com/trace-panic/linux/jsonfs"
)

type Config struct {
	CWD           string
	Hostname      string
	Username      string
	OSVersion     string
	BaseDirectory string
	HomeDirectory string
}

var (
	configInstance *Config
	configLock     sync.Mutex
)

// GetConfig returns the singleton instance of Config.
// It initializes the instance if it doesn't already exist.
func GetConfig() *Config {
	configLock.Lock()
	defer configLock.Unlock()

	if configInstance == nil {
		configInstance = &Config{}
	}

	return configInstance
}

func (c *Config) SaveConfig() error {
	path := filepath.Join(c.HomeDirectory, ".config", "config.json")

	data, err := jsonfs.MarshalJSON(c)
	if err != nil {
		return err
	}

	if err := filesystem.SaveFileSecure(path, data); err != nil {
		return err
	}

	return nil
}

func (c *Config) LoadConfig() error {
	path := filepath.Join(c.HomeDirectory, ".config", "config.json")

	data, err := filesystem.ReadFile(path)
	if err != nil {
		return err
	}

	if err := jsonfs.UnmarshalJSON(data, c); err != nil {
		return err
	}

	return nil
}
