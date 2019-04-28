package common

import (
	"errors"
	"os"
	"path/filepath"
)

const (
	// ServersRoot is the root directory where all of the servers are stored.
	ServersRoot = "servers/"
)

var (
	// ErrServernameAlreadyInUse is thrown when there is already a server with that name
	ErrServernameAlreadyInUse = errors.New("that server name is already in use")
)

// CreateDirIfDoesNotExist creates a directory if it does not already exist.
func CreateDirIfDoesNotExist(dir string) error {
	dir = filepath.FromSlash(dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewServerPath returns the path to a server given its name.
func NewServerPath(name string) (string, error) {
	// Create a path
	rawPath := ServersRoot + name
	abs, err := filepath.Abs(rawPath)
	if err != nil {
		return "", err
	}

	// // Fix later
	// // Check that the server does not already exist
	// if _, err := os.Stat(abs); os.IsNotExist(err) {
	// 	return "", ErrServernameAlreadyInUse
	// }
	return abs, nil

}
