package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "github.com/xoreo/mcserver-backend/commands"

	"github.com/gorilla/mux"
	"github.com/xoreo/mcserver-backend/commands"
	"github.com/xoreo/mcserver-backend/types"
)

/* ----- START POST ROUTES ----- */

// CreateServer is the api function to create a new server.
func CreateServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text") // Set the proper header
	fmt.Printf("formvalue: %s\n\n\n", r.FormValue("a"))

	server, err := types.NewServer("1.7.2", "server test", 25565, 1024)
	log.Fatal(err.Error())

	err = commands.InitializeServer(server)
	log.Fatal(err.Error())
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

// POSTRequest is a temp post request JSON format.
type POSTRequest struct {
	Val string `json:"val"`
}

// TestPOST is a a random testing method.
func TestPOST(w http.ResponseWriter, r *http.Request) {
	fmt.Print("test\n\n\n\n")
	var res POSTRequest
	json.NewDecoder(r.Body).Decode(&res)
	_ = r.FormValue("val")
	json.NewEncoder(w).Encode("test")
	// fmt.Fprintf(w, "sometig")
	// fmt.Fprintf(w, fmt.Sprintf("[val] %s\n", val))
}
