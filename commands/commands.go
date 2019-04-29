package scripts

import (
	"errors"

	"github.com/xoreo/mcserver-backend/common"
	"github.com/xoreo/mcserver-backend/types"
)

var (
	// ErrUnsupportedVersion is thrown when an unsupported server version is given.
	ErrUnsupportedVersion = errors.New("that is not a supported version")
)

func downloadServerJar(url string) error {
	return nil
}

// CreateNewServer creates a new server.
func CreateNewServer(server *types.Server) error {
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
	err := downloadServerJar(url)
	if err != nil {
		return err
	}

	return nil
}

// StartServer starts a server.
func StartServer(server *types.Server) {

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
