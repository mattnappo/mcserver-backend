package api

import (
	"fmt"
	"net/http"

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
	router := api.Router

	// Initialize the POST routes
	router.HandleFunc("/api/createServer/", CreateServer).Methods("POST")
	router.HandleFunc("/api/sendCommand/", SendCommand).Methods("POST")
	router.HandleFunc("/api/editProperties/", EditProperties).Methods("POST")

	// Initialize the GET routes
	router.HandleFunc("/api/startServer/{hash}", StartServer).Methods("GET")
	router.HandleFunc("/api/stopServer/{hash}", StopServer).Methods("GET")
	router.HandleFunc("/api/restartServer/{hash}", RestartServer).Methods("GET")
	router.HandleFunc("/api/serverStatus/{hash}", ServerStatus).Methods("GET")

	// A test route
	router.HandleFunc("/test/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "the test worked!")
		}).Methods("GET")
}
