package commands

import (
	"os"

	"github.com/xoreo/mcserver-backend/types"
)

func generateService(server types.Server) error {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return err
	}

	service := `[Unit]
	Description=` + server.Version + ` ` + ` Server (` + server.Name + `)
	[Service]
	User=ubuntu
	WorkingDirectory=` + workingDirectory + `
	ExecStart=/home/ubuntu/workspace/my-webapp
	SuccessExitStatus=143
	TimeoutStopSec=10
	Restart=on-failure
	RestartSec=5
	[Install]
	WantedBy=multi-user.target`

	return nil
}
