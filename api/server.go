package api

import (
	"fmt"
	"net/http"
	"strconv"
)

// StartAPIServer starts the API server.
func StartAPIServer(port int) {
	api := NewAPI(port) // Create a new API

	http.ListenAndServe(":"+strconv.Itoa(api.Port), api.Router) // Start an HTTP server
	fmt.Printf("== API SERVER LISTENING ON PORT %d ==\n", port)
}
