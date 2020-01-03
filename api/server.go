package api

import (
	"net/http"
	"strconv"
	"fmt"
)

// StartAPIServer starts the API server.
func StartAPIServer(port int) {
	api := NewAPI() // Create a new API

	http.ListenAndServe(":"+strconv.Itoa(port), api.Router) // Start an HTTP server
	fmt.Printf("== API SERVER LISTENING ON PORT %d ==\n", port)
}
