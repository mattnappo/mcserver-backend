package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	// "github.com/xoreo/mcserver-backend/commands"

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

	fmt.Printf("\n\n%s\n%s\n%s\n%s\n\n\n", requestData.Version, requestData.Name, requestData.Port, requestData.RAM)

	// Extract the data from the request
	// port, err := strconv.Atoi(requestData.Port)
	port, err := strconv.Atoi("23")
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
	fmt.Println("this worked!!!")
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
	w.Header().Set("Content-Type", "application/text") // Set the proper header

	// Extract the server hash from the request
	hashString := mux.Vars(r)["hash"]

	// Open the DB now
	serverDB, _ := types.LoadDB() // THE BUG IS ON THIS LINE
	// log.Fatal(err.Error())

	// Search for a server with the given hash
	server, _ := serverDB.GetServerFromHash(hashString)
	fmt.Printf("server: %s\n\n", server.String())
	// log.Fatal(err.Error())

	// Execute the start command
	output, err := commands.Execute("start", *server)
	fmt.Printf("output: %s\n\n", output)
	log.Fatal(err.Error())

	// Write to the api server
	fmt.Fprintf(w, output)
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
