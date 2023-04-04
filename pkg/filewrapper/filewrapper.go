package filewrapper

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	if fileInfo.IsDir() {
		return true
	}
	return false
}

// ExpandTild expands the tilde character in the specified path to the current user's home directory path.
func ExpandTild(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}

	// Get information about the current user.
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	// Replace the tilde character with the home directory path.
	return strings.Replace(path, "~", usr.HomeDir, 1), nil
}

// AbsPath returns the absolute path of the specified path, expanding the tilde character if necessary.
func AbsPath(path string) (string, error) {
	var err error
	if strings.HasPrefix(path, "~") {
		path, err = ExpandTild(path)
		if err != nil {
			return "", err
		}
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return absPath, nil
}

// CreateDirIfNotExists If we provide a file name/dir name it creates the specified directory if it does not exist.
func CreateDirIfNotExists(path string) error {
	// If the directory already exists, return.
	path = filepath.Dir(path)
	if IsDirectory(path) {
		return nil
	}

	// Create the directory.
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// GetFileNameFromPath returns the file name from the specified path.
func GetFileNameFromPath(path string) string {
	return filepath.Base(path)
}
