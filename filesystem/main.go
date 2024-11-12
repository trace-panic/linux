package filesystem

import (
	"os"
	"os/user"
)

// DirExists checks if a directory exists at the given path
// It returns either true or false based on dirs existance
func DirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// HomeDir returns the home directory of current parent OS CompUser
// Panics if it fails, this is critical "syscall" to parent OS
func OSHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		// Save to logs once we have a logs system
		panic(err)
	}

	return usr.HomeDir
}

// CreateDir creates a single directory if it doesn't already exist.
// It returns an error if it fails to create the directory.
func CreateDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0750); err != nil {
			return err
		}
	}
	return nil
}

// CreateDirs accepts a list of directories and creates each one by calling CreateDir.
// It returns an error if any directory creation fails.
func CreateDirs(directories []string) error {
	for _, dir := range directories {
		if err := CreateDir(dir); err != nil {
			return err
		}
	}
	return nil
}
