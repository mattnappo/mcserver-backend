package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xoreo/mcserver-backend/commands"
	"github.com/xoreo/mcserver-backend/types"
)

/* ----- START POST ROUTES ----- */

// CreateServer is the api function to create a new server.
func CreateServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the proper header

	// Decode the post request
	var requestData CreateServerRequest
	json.NewDecoder(r.Body).Decode(&requestData)

	// Extract the data from the request
	port, err := strconv.Atoi(requestData.Port)
	if err != nil {
		log.Fatal(err.Error())
	}

	ram, err := strconv.Atoi(requestData.RAM)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create the new server
	server, err := types.NewServer(requestData.Version, requestData.Name, uint32(port), uint32(ram))
	if err != nil {
		log.Fatal(err.Error())
	}

	// Initialize the server
	err = commands.InitializeServer(server)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.NewEncoder(w).Encode(*server)
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

	// Extract the server hash from the request
	hashString := mux.Vars(r)["hash"]

	// Open the DB now
	serverDB, err := types.LoadDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Search for a server with the given hash
	server, err := serverDB.GetServerFromHash(hashString)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Execute the start command
	output, err := commands.Execute("start", *server)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Prepare the response
	res := NewGETResponse(output)

	// Write the response to the server
	json.NewEncoder(w).Encode(res)
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
