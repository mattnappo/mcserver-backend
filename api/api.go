package api

import (
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// API represents the necessary data for the api (a REST server).
type API struct {
	Router *mux.Router    `json:"router"` // The API's mux router
	Log    *logrus.Logger `json:"logger"` // The API's logger
}

// NewAPI constructs a new API struct.
func NewAPI() *API {
	// Create and setup a new logger
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetOutput(os.Stdout)

	api := &API{
		Router: mux.NewRouter(), // Create a new mux router
		Log:    logger,          // The logger
	}

	api.SetupRoutes() // Setup the API's routes
	return api
}

// SetupRoutes initializes the necessary routes for the API's router.
func (api *API) SetupRoutes() {
	// Initialize the POST routes
	api.Router.HandleFunc("/api/createServer/", CreateServer).Methods("POST")
	api.Router.HandleFunc("/api/sendCommand/", SendCommand).Methods("POST")
	api.Router.HandleFunc("/api/changeProperty/", ChangeProperty).Methods("POST")

	// Initialize the GET routes
	api.Router.HandleFunc("/api/system/{method}/{hash}", SystemCommand).Methods("GET")
	api.Router.HandleFunc("/api/deleteServer/{hash}", DeleteServer).Methods("GET")
}
