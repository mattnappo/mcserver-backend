package commands

import (
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/xoreo/mcserver-backend/types"
)

// GenerateService generates a Linux service file that runs the Minecraft server.
func GenerateService(server types.Server) (string, error) {
	// Get the current working directory
	workingDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Get the current user
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	service := `[Unit]
Description=` + server.Version + ` Server (` + server.Name + `)
[Service]
User=` + user.Username + `
WorkingDirectory=` + workingDirectory + `
ExecStart=` + server.StartScript + `
SuccessExitStatus=143
TimeoutStopSec=10
Restart=on-failure
RestartSec=5
[Install]
WantedBy=multi-user.target`

	return service, nil
}

// InstallService installs a service to the system.
func InstallService(service, name string) (string, error) {
	serviceName := name + ".service" // For convenience
	// Init the path to the system services (Linux only)
	path := filepath.Join("/etc/systemd/system/", serviceName)

	// Write the service to the file
	err := ioutil.WriteFile(path, []byte(service), 0664)
	if err != nil {
		return "", err
	}

	// Execute the necessary commands to register the daemon
	exec.Command("/bin/sh", "sudo systemctl daemon-reload")

	return path, nil
}

// UninstallService uninstalls a service from the system.
func UninstallService(serviceName string) (string, error) {
	command := exec.Command("/bin/sh", "sudo rm /etc/systemd/system/"+serviceName+".service") // Delete the service

	output, err := command.Output() // Get the output
	if err != nil {
		return "", err
	}

	return string(output), nil // Return the output
}
