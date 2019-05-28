package api

import (
	"github.com/gorilla/mux"
)

// API represents the necessary data for the api (a REST server).
type API struct {
	Router mux.Router `json:"router"`
}

// InitializeRoutes initializes the routes for the API.
func InitializeRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/api/createServer", CreateServer).Methods("POST")
	router.HandleFunc("/api/serverStatus/{hash}", ServerStatus).Methods("GET")
}
