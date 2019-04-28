package types

import (
	"errors"
	"time"

	"github.com/xoreo/mcserver-backend/common"
)

var (
	// ErrUnsupportedVersion is thrown when an unsupported server version is given.
	ErrUnsupportedVersion = errors.New("that is not a supported version")
)

// Server holds the metadata for a local Minecraft server.
type Server struct {
	Version     string `json:"version"`     // The server version
	Path        string `json:"path"`        // The local path to the server
	Port        int    `json:"port"`        // The port that the server runs on
	Name        string `json:"name"`        // The name of the server
	TimeCreated string `json:"timeCreated"` // The time that the server was created
}

// NewServer constructs a new server struct.
func NewServer(version, name string) (*Server, error) {
	var url string

	// Determine the server download url
	switch version {
	case "1.12":
		url = common.ServerV112
	case "1.8":
		url = common.ServerV18
	case "1.7.2":
		url = common.ServerV172
	case "1.2.1":
		url = common.ServerV121
	default:
		return nil, ErrUnsupportedVersion
	}

	// Determine the path for the server
	path := common.NewServerPath(name)

	// Create the new server
	newServer := &Server{
		Version:     version,
		Path:        path,
		Port:        port,
		Name:        name,
		TimeCreated: time.Now().String(),
	}

	return newServer, nil

}
