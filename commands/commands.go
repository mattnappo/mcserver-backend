package commands

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/xoreo/mcserver-backend/common"
	"github.com/xoreo/mcserver-backend/types"
)

var (
	// ErrUnsupportedVersion is thrown when an unsupported server version is given.
	ErrUnsupportedVersion = errors.New("that is not a supported version")

	// ErrServerHasNotBeenInitialized is thrown when a server's metadata exists but the server has not actually been initialized on the local machine.
	ErrServerHasNotBeenInitialized = errors.New("that server has not actually been initialized yet. Initialize it with InitializeServer()")
)

func downloadServerJar(url, localPath, version string) (string, error) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Create the directory
	err = common.CreateDirIfDoesNotExist(localPath)
	if err != nil {
		return "", err
	}

	// Create the file
	zipPath := filepath.Join(localPath, version+".zip")
	out, err := os.Create(zipPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}
	return zipPath, nil
}

// InitializeServer initializes a new server onto the local machine.
func InitializeServer(server *types.Server) error {
	var url string

	// Determine the pre-made server download url
	switch (*server).Version {
	case "1.12":
		url = common.ServerV112
	case "1.8":
		url = common.ServerV18
	case "1.7.2":
		url = common.ServerV172
	case "1.2.1":
		url = common.ServerV121
	default:
		return ErrUnsupportedVersion
	}

	// Download the pre-made server
	zipPath, err := downloadServerJar(url, (*server).Path, (*server).Version)
	if err != nil {
		return err
	}

	// Unzip the downloaded file
	_, err = common.Unzip(zipPath, (*server).Path)
	if err != nil {
		return err
	}

	server.Created = true // Set the server's created state = true
	return nil
}

// StartServer starts a server.
func StartServer(server *types.Server) error {
	if !(*server).Created {
		return ErrServerHasNotBeenInitialized
	}
	startScriptPath := filepath.Join(server.Path, server.Version, "start.sh")
	fmt.Println(startScriptPath)
	return nil
}

// RestartServer restarts a server.
func RestartServer(server *types.Server) {

}

// StopServer stops a server.
func StopServer(server *types.Server) {

}

// EnterServer enters the shell of the server.
func EnterServer(server *types.Server) {

}
