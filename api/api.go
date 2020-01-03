package api

import (
	"os"
	"path"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/xoreo/mcserver-backend/common"
)

// API represents the necessary data for the api (a REST server).
type API struct {
	Router *mux.Router    `json:"router"` // The API's mux router
	Log    *logrus.Logger `json:"logger"` // The API's logger

	Root string `json:"root"` // The root route
	Port int    `json:"port"` // The port to listen on
}

// NewAPI constructs a new API struct.
func NewAPI(port int) *API {
	// Create and setup a new logger
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetOutput(os.Stdout)

	api := &API{
		Router: mux.NewRouter(), // Create a new mux router
		Log:    logger,          // The logger

		Root: common.APIServerRoot, // The default root
		Port: port,                 // The given port
	}

	api.SetupRoutes() // Setup the API's routes
	return api
}

// SetupRoutes initializes the necessary routes for the API's router.
func (api *API) SetupRoutes() {
	// Initialize the POST routes
	api.Router.HandleFunc(path.Join(api.Root, "createServer"), CreateServer).Methods("POST")
	api.Router.HandleFunc(path.Join(api.Root, "sendCommand"), SendCommand).Methods("POST")
	api.Router.HandleFunc(path.Join(api.Root, "changeProperty"), ChangeProperty).Methods("POST")

	// Initialize the GET routes
	api.Router.HandleFunc(path.Join(api.Root, "system/{method}/{hash}"), SystemCommand).Methods("GET")
	api.Router.HandleFunc(path.Join(api.Root, "deleteServer/{hash}"), DeleteServer).Methods("GET")
}
