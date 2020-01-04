package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/xoreo/mcserver-backend/commands"
	"github.com/xoreo/mcserver-backend/common"
	"github.com/xoreo/mcserver-backend/types"
)

/* ----- START POST ROUTES ----- */

// CreateServer is the api function to create a new server.
func CreateServer(w http.ResponseWriter, r *http.Request) {
	logger := common.NewLogger("api.CreateServer")
	w.Header().Set("Content-Type", "application/json") // Set the proper header

	// Decode the post request
	var requestData CreateServerRequest
	json.NewDecoder(r.Body).Decode(&requestData)

	// Extract the data from the request
	port, err := strconv.Atoi(requestData.Port)
	if err != nil {
		logger.Criticalf(err.Error())
	}

	ram, err := strconv.Atoi(requestData.RAM)
	if err != nil {
		logger.Criticalf(err.Error())
	}

	logger.Infof("request to create server with specs:\n%s", requestData.String())

	// Create the new server
	server, err := types.NewServer(requestData.Version, requestData.Name, port, ram)
	if err != nil {
		logger.Criticalf(err.Error())
	}

	logger.Debugf("created new server entry %s", server.Hash.String())

	// Initialize the server
	err = commands.InitializeServer(server)
	if err != nil {
		logger.Criticalf(err.Error())
	}

	logger.Debugf("server initialization complete")

	// Add the newly-created server to the database
	serverDB, err := types.LoadDB()
	if err != nil {
		logger.Criticalf(err.Error())
	}

	err = serverDB.AddServer(server)
	if err != nil {
		logger.Criticalf(err.Error())
	}
	serverDB.Close()

	logger.Debugf("added new server to the database")
	logger.Infof("created new server %s", server.Hash.String())

	json.NewEncoder(w).Encode(*server) // Write to the server
}

// SendCommand is the api function that sends a command to the server.
func SendCommand(w http.ResponseWriter, r *http.Request) {

}

// ChangeProperty is the api function that changes a property in a server
func ChangeProperty(w http.ResponseWriter, r *http.Request) {
	logger := common.NewLogger("api.ChangeProperty")
	w.Header().Set("Content-Type", "application/json") // Set the proper header

	// Decode the post request
	var requestData ChangePropertyRequest
	json.NewDecoder(r.Body).Decode(&requestData)

	hash := requestData.Hash // Extract the hash from the request
	if hash == "" {
		logger.Criticalf("hash cannot be nil")
	}

	logger.Debugf("request to change property:\n%s", requestData.String())

	// Add the newly-created server to the database
	serverDB, err := types.LoadDB()
	if err != nil {
		logger.Criticalf(err.Error())
	}

	// Search for a server with the given hash
	server, err := serverDB.GetServerFromHash(hash)
	if err != nil {
		logger.Criticalf(err.Error())
	}
	oldServer := server

	// Change the property in server to be the new server
	err = server.Properties.ChangeProperty(requestData.Property, requestData.NewValue)
	if err != nil {
		logger.Criticalf(err.Error())
	}

	logger.Debugf("changed property in new server")

	// Update the port field of the server itself, not just the property
	if requestData.Property == "ServerPort" {
		newPort, err := strconv.Atoi(requestData.NewValue)
		if err != nil {
			logger.Criticalf(err.Error())
		}

		server.Port = newPort
	}

	fmt.Println()

	// Update the server in the database
	err = serverDB.UpdateServer(oldServer, server, "")
	if err != nil {
		logger.Criticalf(err.Error())
	}

	serverDB.Close() // Save changes to the DB and close

	logger.Debugf("updated new server in database")

	// fmt.Println(server.Properties.GetFile())
	err = server.Properties.WriteToServer(server)
	if err != nil {
		logger.Criticalf(err.Error())
	}

	logger.Debugf("wrote new properties file to disk")
	logger.Infof("changed property %s in server %s", requestData.Property, server.Hash.String())

	// Send the response (the new server)
	res := NewDefaultResponse(server.String())
	json.NewEncoder(w).Encode(res)
}

/* ----- END POST ROUTES ----- */

/* ----- START GET ROUTES ----- */

// SystemCommand can start, stop, and restart a server as well as get its status.
func SystemCommand(w http.ResponseWriter, r *http.Request) {
	logger := common.NewLogger("api.SystemCommand")
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
		logger.Criticalf(errors.New("that is not a valid method").Error())
	}

	// Extract the server hash from the request
	hashString := mux.Vars(r)["hash"]

	logger.Infof("request to execute system command '%s' on server %s", method, hashString)

	// Open the DB
	serverDB, err := types.LoadDB()
	if err != nil {
		logger.Criticalf(err.Error())
	}

	// Search for a server with the given hash
	server, err := serverDB.GetServerFromHash(hashString)
	if err != nil {
		logger.Criticalf(err.Error())
	}

	logger.Debugf("found server %s in database", server.Hash.String())

	serverDB.Close() // Close the DB

	// Execute the command
	output, err := commands.Execute(method, *server)
	if err != nil {
		logger.Criticalf(err.Error())
	}

	logger.Infof("executed system command %s on server %s", method, server.Hash.String())

	// Prepare the response
	res := NewGETResponse(output)

	// Write the response to the server
	json.NewEncoder(w).Encode(res)
}

// GetServer returns a server given a hash.
func GetServer(w http.ResponseWriter, r *http.Request) {
	logger := common.NewLogger("api.GetServer")
	w.Header().Set("Content-Type", "application/json") // Set the proper header

	// Extract the server hash from the request
	hashString := mux.Vars(r)["hash"]

	logger.Infof("request to get %s", hashString)

	// Open the DB
	serverDB, err := types.LoadDB()
	if err != nil {
		logger.Criticalf(err.Error())
	}

	// Get the server from the DB
	server, err := serverDB.GetServerFromHash(hashString)
	if err != nil {
		logger.Criticalf(err.Error())
	}

	serverDB.Close() // Close the DB

	logger.Debugf("got server %s from database", server.Hash.String())

	// Prepare the response
	res := GETServerResponse{*server}

	// Write the response to the server
	json.NewEncoder(w).Encode(res)
}

// GetAllServers returns all servers in the database.
func GetAllServers(w http.ResponseWriter, r *http.Request) {
	logger := common.NewLogger("api.GetAllServers")
	w.Header().Set("Content-Type", "application/json") // Set the proper header

	logger.Infof("request to get all servers")

	// Open the DB
	serverDB, err := types.LoadDB()
	if err != nil {
		logger.Criticalf(err.Error())
	}

	// Get the server from the DB
	servers, err := serverDB.GetAllServers()
	if err != nil {
		logger.Criticalf(err.Error())
	}

	serverDB.Close() // Close the DB

	logger.Debugf("got all servers from database")

	// Prepare the response
	res := GETServersResponse{servers}

	// Write the response to the server
	json.NewEncoder(w).Encode(res)
}

// DeleteServer will delete a server given its hash.
func DeleteServer(w http.ResponseWriter, r *http.Request) {
	logger := common.NewLogger("api.DeleteServer")
	w.Header().Set("Content-Type", "application/json") // Set the proper header

	// Extract the server hash from the request
	hashString := mux.Vars(r)["hash"]

	logger.Infof("request to delete %s", hashString)

	// Open the DB
	serverDB, err := types.LoadDB()
	if err != nil {
		logger.Criticalf(err.Error())
	}

	// Delete the server (entry) from the database
	server, err := serverDB.DeleteServer(hashString)
	if err != nil {
		logger.Criticalf(err.Error())
	}

	logger.Debugf("deleted server %s from database", server.Hash.String())

	serverDB.Close() // Close the DB

	err = commands.Purge(server) // Purge the server files
	if err != nil {
		logger.Criticalf(err.Error())
	}

	logger.Debugf("purged server %s from system", server.Hash.String())
	logger.Infof("deleted %s", server.Hash.String())

	// Prepare the response
	res := NewGETResponse("success")

	// Write the response to the server
	json.NewEncoder(w).Encode(res)
}

/* ----- END GET ROUTES ----- */
