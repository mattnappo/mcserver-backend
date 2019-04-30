package types

import (
	"encoding/json"
	"io/ioutil"
)

// ServerDBName is the name of the server database.
const ServerDBName = "servers.json"

// ServerDB contains the metadata for all of the servers.
type ServerDB struct {
	Servers []Server `json:"servers"`
}

// LoadDB returns the contents of a ServerDB file.
func LoadDB() (*ServerDB, error) {
	// Read the database file
	rawRead, err := ioutil.ReadFile(ServerDBName)
	if err != nil {
		return nil, err
	}
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
func (db *ServerDB) AddServer(server *Server) {
	(*db).Servers = append((*db).Servers, *server)
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
