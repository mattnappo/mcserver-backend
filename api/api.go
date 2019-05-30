package api

import (
	"github.com/gorilla/mux"
)

// API represents the necessary data for the api (a REST server).
type API struct {
	Router *mux.Router `json:"router"` // The API's mux router
}

// NewAPI constructs a new API struct.
func NewAPI() *API {
	api := &API{
		Router: mux.NewRouter(), // Create a new mux router
	}

	api.SetupRoutes() // Setup the API's routes
	return api
}

// SetupRoutes initializes the necessary routes for the API's router.
func (api *API) SetupRoutes() {
	// Initialize the POST routes
	api.Router.HandleFunc("/api/createServer/", CreateServer).Methods("POST")
	api.Router.HandleFunc("/api/sendCommand/", SendCommand).Methods("POST")
	api.Router.HandleFunc("/api/editProperties/", EditProperties).Methods("POST")

	// Initialize the GET routes
	api.Router.HandleFunc("/api/execute/{method}/{hash}", Execute).Methods("GET")
}
