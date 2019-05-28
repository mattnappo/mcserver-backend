package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/xoreo/mcserver-backend/common"
)

// Server holds the metadata for a local Minecraft server.
type Server struct {
	Version     string      `json:"version"`     // The server version
	Path        string      `json:"path"`        // The local path to the server
	Port        int         `json:"port"`        // The port that the server runs on
	RAM         int         `json:"ram"`         // The amount of ram to be allocated to the server
	Name        string      `json:"name"`        // The name of the server
	TimeCreated string      `json:"timeCreated"` // The time that the server was created
	Initialized bool        `json:"initialized"` // Whether the server has been initialized or not
	StartScript string      `json:"startScript"` // The path to the start.sh script
	Hash        common.Hash `json:"hash"`        // The hash of the server
}

// NewServer constructs a new server struct.
func NewServer(version, name string, port, ram uint32) (*Server, error) {
	// Determine the path for the server
	path, err := common.NewServerPath(name)
	if err != nil {
		return nil, err
	}

	// Replace spaces with '-' in server name
	name = strings.Replace(name, " ", "-", -1)

	// Create the new server
	newServer := &Server{
		Version:     version,
		Path:        path,
		Port:        int(port),
		RAM:         int(ram),
		Name:        name,
		TimeCreated: time.Now().String(),
		Initialized: false,
		StartScript: "",
	}

	newServer.Hash = common.Sha3(newServer.Bytes())
	fmt.Println(newServer.Hash.String())
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

// Recalculate recalculates the hash of the server.
func (server *Server) Recalculate() {
	server.Hash = common.Sha3(server.Bytes())
}

/* -- END HELPER METHODS -- */
