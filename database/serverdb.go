package database

import (
	"github.com/xoreo/mcserver-backend/types"
)

// ServerDBName is the name of the server database.
const ServerDBName = "serverdb"

// ServerDB contains the metadata for all of the servers.
type ServerDB struct {
	Servers []types.Server `json:"servers"`
}
