package filesystem

import (
	"github.com/trace-panic/linux/config"
)

var cfg = config.Config{}

func Init(cfg config.Config) error {
	cfg.SetHomeDirectory(cfg.BaseDirectory + "/linux/home/" + cfg.Username)
	return CreateFileSystem(cfg)
}

func CreateFileSystem(cfg config.Config) error {
	directories := []string{
		cfg.HomeDirectory + "/.config",
	}

	if err := CreateDirs(directories); err != nil {
		return err
	}

	return nil
}
