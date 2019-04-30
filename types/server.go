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
	RAM         int    `json:"ram"`         // The amount of ram to be allocated to the server
	Name        string `json:"name"`        // The name of the server
	TimeCreated string `json:"timeCreated"` // The time that the server was created
	Initialized bool   `json:"created"`     // Whether the server has been initialized or not
}

// NewServer constructs a new server struct.
func NewServer(version, name string, port, ram uint32) (*Server, error) {
	// Determine the path for the server
	path, err := common.NewServerPath(name)
	if err != nil {
		return nil, err
	}

	// Create the new server
	newServer := &Server{
		Version:     version,
		Path:        path,
		Port:        int(port),
		RAM:         int(ram),
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
