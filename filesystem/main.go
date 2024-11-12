package filesystem

import (
	"os"
	"os/user"
	"path/filepath"
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

// FileExists checks if a file exists at the given path.
// It returns true if the file exists, otherwise false.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// CreateFile creates a file at the specified path.
// If the file creation fails, it returns an error.
func CreateFile(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0750)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

// CreateFileSecure creates a file at the specified path ensuring the parent directory exists.
// If the directory is missing, it will be created first.
func CreateFileSecure(path string) error {
	dir := filepath.Dir(path)
	if err := CreateDir(dir); err != nil {
		return err
	}

	return CreateFile(path)
}

// SaveFile creates a file at the specified path and writes the provided byte data into it.
// If the file saving fails, it returns an error.
func SaveFile(path string, data []byte) error {
	if err := CreateFile(path); err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_WRONLY, 0750)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// SaveFileSecure ensures the parent directory exists, then creates the file and writes the provided byte data into it.
func SaveFileSecure(path string, data []byte) error {
	dir := filepath.Dir(path)
	if err := CreateDir(dir); err != nil {
		return err
	}

	return SaveFile(path, data)
}

// ReadFile reads the contents of a file at the specified path.
// It returns the byte data read from the file or an error if it fails to read.
func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}
