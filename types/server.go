// Package types implements the core types to this backend such a the server database.
package types

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/xoreo/mcserver-backend/common"
)

// CoreProperties represets the core properties of a server
type CoreProperties struct {
	Name    string `json:"name"`    // The name of the server
	Version string `json:"version"` // The version of the server
	Port    int    `json:"port"`    // The port of the server
	RAM     int    `json:"ram"`     // The amount of ram of the server
	ID      string `json:"id"`      // The ID of the server
}

// Server holds the metadata for a local Minecraft server.
type Server struct {
	Name    string `json:"name"`    // The name of the server
	Version string `json:"version"` // The server version
	Port    int    `json:"port"`    // The port that the server runs on
	RAM     int    `json:"ram"`     // The amount of ram to be allocated to the server

	TimeCreated string `json:"time_created"` // The time that the server was created
	Initialized bool   `json:"initialized"`  // Whether the server has been initialized or not

	Properties *Properties `json:"properties"` // The contents of the server.properties file

	Path        string `json:"path"`         // The local path to the server
	ServicePath string `json:"service_path"` // The path to the service file
	StartScript string `json:"startScript"`  // The path to the start.sh script

	Hash common.Hash `json:"hash"` // The hash of the server
	ID   string      `json:"id"`   // The first 3 bytes of the hash (in hex) as a string
}

// NewServer constructs a new server struct.
func NewServer(version, name string, port, ram int) (*Server, error) {
	oldName := name

	// Replace spaces with '-' in server name
	name = strings.Replace(name, " ", "-", -1)

	// Determine the path for the server
	path, err := common.NewServerPath(name)
	if err != nil {
		return nil, err
	}

	// Create the new server
	newServer := &Server{
		Name:    name,
		Version: version,
		Port:    int(port),
		RAM:     int(ram),

		TimeCreated: time.Now().String(),
		Initialized: false,

		Properties: DefaultProperties(int(port), oldName+" on "+strconv.Itoa(port)),

		Path:        path,
		StartScript: "",
	}

	// Compute the hash and ID of the server
	newServer.Hash = common.Sha3(newServer.Bytes())
	newServer.ID = newServer.Hash.String()[0:common.ServerIDSize]

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

// GetCoreProperties returns a CoreProperties containing the core properties of the server.
func (server *Server) GetCoreProperties() CoreProperties {
	cp := CoreProperties{
		Name:    server.Name,
		Version: server.Version,
		Port:    server.Port,
		RAM:     server.RAM,
		ID:      server.ID,
	}
	return cp
}

/* -- END HELPER METHODS -- */
