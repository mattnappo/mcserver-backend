package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Server initializes the necessary routes for the API's router.
func Server(port int) {
	router := mux.NewRouter()
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
	router.HandleFunc("/testGET/{data}",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json") // Set the proper header
			data := mux.Vars(r)["data"]

			formatted := fmt.Sprintf("[data] %s\n", data)
			// fmt.Fprintf(w, "somethin")
			json.NewEncoder(w).Encode(formatted)
			// fmt.Fprintf(w, )
		}).Methods("GET")

	http.ListenAndServe(":"+strconv.Itoa(port), router) // Start an HTTP server
}
