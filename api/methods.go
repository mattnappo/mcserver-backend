package api

import (
	"encoding/json"
	"errors"
	"fmt"
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
	server, err := types.NewServer(requestData.Version, requestData.Name, port, ram)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Initialize the server
	err = commands.InitializeServer(server)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Add the newly-created server to the database
	serverDB, err := types.LoadDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = serverDB.AddServer(server)
	if err != nil {
		log.Fatal(err.Error())
	}
	serverDB.Close()

	fmt.Printf("-- GENERATED NEW SERVER --\n[hash] %s\n\n", server.Hash.String()) // Log

	json.NewEncoder(w).Encode(*server) // Write to the server
}

// SendCommand is the api function that sends a command to the server.
func SendCommand(w http.ResponseWriter, r *http.Request) {

}

// ChangeProperty is the api function that changes a property in a server
func ChangeProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the proper header

	// Decode the post request
	var requestData ChangePropertyRequest
	json.NewDecoder(r.Body).Decode(&requestData)

	hash := requestData.Hash // Extract the hash from the request
	if hash == "" {
		log.Fatal("hash cannot be nil")
	}

	// Add the newly-created server to the database
	serverDB, err := types.LoadDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Search for a server with the given hash
	server, err := serverDB.GetServerFromHash(hash)
	if err != nil {
		log.Fatal(err.Error())
	}
	oldServer := server

	err = server.Properties.ChangeProperty(requestData.Property, requestData.NewValue)
	if err != nil {
		log.Fatal(err.Error())
	}

	serverDB.UpdateServer(oldServer, server)

	fmt.Println(server.Properties.GetFile())

}

/* ----- END POST ROUTES ----- */

/* ----- START GET ROUTES ----- */

// Execute can start, stop, and restart a server as well as get its status.
func Execute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the proper header

	// Determine which method to call
	method := mux.Vars(r)["method"]
	switch method {
	case "start":
		break
	case "stop":
		break
	case "status":
		break
	case "restart":
		break
	default:
		log.Fatal(errors.New("that is not a valid method"))
	}

	// Extract the server hash from the request
	hashString := mux.Vars(r)["hash"]

	// Open the DB
	serverDB, err := types.LoadDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Search for a server with the given hash
	server, err := serverDB.GetServerFromHash(hashString)
	if err != nil {
		log.Fatal(err.Error())
	}

	serverDB.Close() // Close the DB

	// Execute the command
	output, err := commands.Execute(method, *server)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("[%s output] %s\n", method, output)

	// Prepare the response
	res := NewGETResponse(output)

	// Write the response to the server
	json.NewEncoder(w).Encode(res)

}

/* ----- END GET ROUTES ----- */
