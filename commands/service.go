package commands

import (
	"os"
	"os/user"

	"github.com/xoreo/mcserver-backend/types"
)

// GenerateService generates a Linux service file that runs the Minecraft server.
func GenerateService(server types.Server) (string, error) {
	// Make sure that the server has been initialized.
	if !server.Initialized || server.StartScript == "" {
		return "", ErrServerHasNotBeenInitialized
	}

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
	Description=` + server.Version + ` ` + ` Server (` + server.Name + `)
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
