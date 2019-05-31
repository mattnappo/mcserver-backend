package types

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// ServerDBName is the name of the server database.
const ServerDBName = "servers.json" // Make this an absolute path

// ServerDB contains the metadata for all of the servers.
type ServerDB struct {
	Servers []Server `json:"servers"`
}

// LoadDB returns the contents of a ServerDB file.
func LoadDB() (*ServerDB, error) {
	// Create the file (if it does not already exist)
	file, err := os.OpenFile(ServerDBName, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	file.Close()

	// Read the database file
	rawRead, err := ioutil.ReadFile(ServerDBName)
	if err != nil {
		return nil, err
	}

	// If the file is nil, load the empty array of servers
	if len(rawRead) == 0 {
		return &ServerDB{
			Servers: []Server{},
		}, nil
	}

	// Reconstruct the file
	buffer := &ServerDB{}
	err = json.Unmarshal(rawRead, buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil // Return
}

// AddServer adds a server to the database
func (db *ServerDB) AddServer(server *Server) error {
	for _, currentServer := range db.Servers {
		if currentServer.Name == server.Name {
			return errors.New("that server name is already in use")
		}
	}

	db.Servers = append((*db).Servers, *server)
	return nil
}

// UpdateServer updates a server in the database.
func (db *ServerDB) UpdateServer(oldServer, newServer *Server) error {
	// Check that the server exists
	exists := false
	for _, currentServer := range db.Servers {
		if currentServer.Hash.String() == oldServer.Hash.String() {
			exists = true
			break
		}
	}

	// If the server does not exist, throw an error
	if !exists {
		return errors.New("that server is not in the database")
	}

	// Remove the old server
	for i, server := range db.Servers {
		if server.Hash.String() == oldServer.Hash.String() {
			db.Servers[i] = Server{}
			break
		}
	}

	newServer.Recalculate() // Recalculate the hash

	// Add the new server
	db.AddServer(newServer)

	return nil
}

// GetServerFromHash returns the server belonging to the hash given.
func (db *ServerDB) GetServerFromHash(hash string) (*Server, error) {
	for _, currentServer := range db.Servers {
		// fmt.Printf("currentServer.Hash.String(): %s\nhash: %s\n", currentServer.Hash.String(), hash)
		if currentServer.Hash.String() == hash {
			return &currentServer, nil
		}
	}

	return nil, errors.New("a server with that hash does not exist")
}

// Close closes and writes changes to the database file
func (db *ServerDB) Close() error {
	// Marshall
	json, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return err
	}

	// Write to file
	err = ioutil.WriteFile(ServerDBName, json, 0644)
	if err != nil {
		return err
	}
	return nil
}

/* -- BEGIN HELPER METHODS -- */

// Bytes returns the raw bytes of the marshalled server database.
func (db *ServerDB) Bytes() []byte {
	json, _ := json.MarshalIndent(*db, "", "  ")
	return json
}

// String returns a string of the marshalled server database.
func (db *ServerDB) String() string {
	json, _ := json.MarshalIndent(*db, "", "  ")
	return string(json)
}

/* -- END HELPER METHODS -- */
