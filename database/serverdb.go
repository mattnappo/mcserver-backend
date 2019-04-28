package database

import (
	"encoding/json"
	"io/ioutil"

	"github.com/xoreo/mcserver-backend/types"
)

// ServerDBName is the name of the server database.
const ServerDBName = "serverdb"

// ServerDB contains the metadata for all of the servers.
type ServerDB struct {
	Servers []types.Server `json:"servers"`
}

// LoadDB returns the contents of a ServerDB file.
func LoadDB() (*ServerDB, error) {
	// Read the database file
	rawRead, err := ioutil.ReadFile(ServerDBName + ".json")
	if err != nil {
		return nil, err
	}

	// Reconstruct the file
	buffer := &ServerDB{}
	err = json.Unmarshal(rawRead, buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil // Return
}
