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
	// Initialize the POST routes
	api.Router.HandleFunc("/api/createServer/", CreateServer).Methods("POST")
	api.Router.HandleFunc("/api/sendCommand/", SendCommand).Methods("POST")
	api.Router.HandleFunc("/api/editProperties/", EditProperties).Methods("POST")

	// Initialize the GET routes
	api.Router.HandleFunc("/api/startServer/{hash}", StartServer).Methods("GET")
	api.Router.HandleFunc("/api/stopServer/{hash}", StopServer).Methods("GET")
	api.Router.HandleFunc("/api/restartServer/{hash}", RestartServer).Methods("GET")
	api.Router.HandleFunc("/api/serverStatus/{hash}", ServerStatus).Methods("GET")

	// A test route
	api.Router.HandleFunc("/testGET/{data}",
		func(w http.ResponseWriter, r *http.Request) {
			// data := mux.Vars(r)["data"]
			fmt.Fprintf(w, "somethin")
			// fmt.Fprintf(w, fmt.Sprintf("[data] %s\n", data))
		}).Methods("GET")

	api.Router.HandleFunc("/testPOST/", TestPOST).Methods("POST")
}
