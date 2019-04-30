package types

import (
	"encoding/json"
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
	Initialized bool   `json:"created"`     // Whether the server has been initialized or not
}

// NewServer constructs a new server struct.
func NewServer(version, name string, port int) (*Server, error) {
	// Determine the path for the server
	path, err := common.NewServerPath(name)
	if err != nil {
		return nil, err
	}

	// Create the new server
	newServer := &Server{
		Version:     version,
		Path:        path,
		Port:        port,
		Name:        name,
		TimeCreated: time.Now().String(),
		Initialized: false,
	}

	return newServer, nil
}

/* -- BEGIN HELPER METHODS -- */

// Bytes returns the raw bytes of the marshalled server.
func (server *Server) Bytes() []byte {
	json, _ := json.MarshalIndent(*server, "", "  ")
	return json
}

// String returns a string of the marshalled server.
func (server *Server) String() string {
	json, _ := json.MarshalIndent(*server, "", "  ")
	return string(json)
}

/* -- END HELPER METHODS -- */
