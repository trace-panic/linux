package config

import (
	"errors"
	"sync"
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
	instance *Config
	once     sync.Once
)

// GetConfig returns the singleton instance of Config.
// It initializes the instance if it doesn't already exist.
func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
	})
	return instance
}

func (c *Config) SetUsername(username string) error {
	if len(username) < 1 {
		return errors.New("Missing username")
	}
	c.Username = username
	return nil
}

func (c *Config) SetHostname(hostname string) error {
	if len(hostname) < 1 {
		return errors.New("Missing hostname")
	}
	c.Hostname = hostname
	return nil
}

func (c *Config) SetCWD(cwd string) error {
	if len(cwd) < 1 {
		return errors.New("Missing CWD")
	}
	c.CWD = cwd
	return nil
}

func (c *Config) SetBaseDirectory(baseDir string) error {
	if len(baseDir) < 1 {
		return errors.New("Missing directory")
	}
	c.BaseDirectory = baseDir
	return nil
}

func (c *Config) SetHomeDirectory(homeDir string) error {
	if len(homeDir) < 1 {
		return errors.New("Missing directory")
	}
	c.HomeDirectory = homeDir
	return nil
}

func (c *Config) SetOSVersion(osVersion string) {
	c.OSVersion = osVersion
}
