package commands

import (
	"errors"
	// "fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/xoreo/mcserver-backend/common"
	"github.com/xoreo/mcserver-backend/types"
)

var (
	// ErrUnsupportedVersion is thrown when an unsupported server version is given.
	ErrUnsupportedVersion = errors.New("that is not a supported version")

	// ErrServerHasNotBeenInitialized is thrown when a server's metadata exists but the server has not actually been initialized on the local machine.
	ErrServerHasNotBeenInitialized = errors.New("that server has not actually been initialized yet. Initialize it with InitializeServer()")
)

// GenerateStartScript generates the script that launches the server.
func GenerateStartScript(server types.Server) []byte {
	ramstr := strconv.Itoa(server.RAM) // Convert the ram to a string

	path := filepath.Join(server.Path, server.Version)
	script := `#!/bin/bash
cd ` + path + `
java -Xms` + ramstr + `M -Xmx` + ramstr + `M -jar ` + path + `/` + server.Version + `.jar nogui`
	return []byte(script)
}

// InitializeServer initializes a new server onto the local machine.
func InitializeServer(server *types.Server) error {
	var url string
	dServer := *server // Make a copy of the pointer (dereference for convenience)

	// Determine the pre-made server download url
	switch dServer.Version {
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
	zipPath, err := common.DownloadServer(url, dServer.Path, dServer.Version)
	if err != nil {
		return err
	}

	// Unzip the downloaded file
	_, err = common.Unzip(zipPath, dServer.Path)
	if err != nil {
		return err
	}

	// Create amd set the paths to be used later
	workingPath := filepath.Join(dServer.Path, dServer.Version)
	startScriptPath := filepath.Join(workingPath, "start.sh")
	server.StartScript = startScriptPath // Set the script path for the server

	// Create start script for the server
	script := GenerateStartScript(dServer)

	// Install the script
	err = ioutil.WriteFile(startScriptPath, script, 0644)
	if err != nil {
		return err
	}

	// Generate a service
	service, err := GenerateService(*server)
	if err != nil {
		return err
	}

	// Install the generated service
	err = InstallService(service, server.Name)
	if err != nil {
		return err
	}

	server.Initialized = true // Set the server's initialized state to true
	return nil
}

// Execute executes a systemctl command (start, stop, restart, status) on a certain server.
func Execute(command string, server types.Server) (string, error) {
	// Check to make sure that the server is initialized
	if !server.Initialized || server.StartScript == "" {
		return "", ErrServerHasNotBeenInitialized
	}

	// Make sure that it is a valid command
	switch command {
	case "start":
		break
	case "stop":
		break
	case "status":
		break
	case "restart":
		break
	default:
		return "", errors.New("that is not a valid command to execute")
	}

	// Execute the systemctl command
	cmd := exec.Command("/bin/systemctl", command, server.Name)

	// Get the output
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// EnterServer launches a shell of the server console.
func EnterServer(server *types.Server) {

}

// EditProperties is used to edit a server property (such as max build height or default gamemode).
func EditProperties(server types.Server, property, newValue string) error {
	return nil
}
