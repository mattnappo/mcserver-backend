package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/* ----- START POST ROUTES ----- */

// CreateServer is the api function to create a new server.
func CreateServer(w http.ResponseWriter, r *http.Request) {

}

// SendCommand is the api function that sends a command to the server.
func SendCommand(w http.ResponseWriter, r *http.Request) {

}

// EditProperties is the api function that edits the server.properties file.
func EditProperties(w http.ResponseWriter, r *http.Request) {

}

/* ----- END POST ROUTES ----- */

/* ----- START GET ROUTES ----- */

// StartServer is the api function that starts a server.
func StartServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the proper header

	vars := mux.Vars(r)
	hash := vars["hash"]
	fmt.Println(hash)

}

// StopServer is the api function that stops a server.
func StopServer(w http.ResponseWriter, r *http.Request) {

}

// RestartServer is the api function that restarts a server.
func RestartServer(w http.ResponseWriter, r *http.Request) {

}

// ServerStatus is the api function that gets the status of a server.
func ServerStatus(w http.ResponseWriter, r *http.Request) {

}

/* ----- END GET ROUTES ----- */
