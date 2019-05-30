package api

import (
	"net/http"
	"strconv"
)

// StartAPIServer starts the API server.
func StartAPIServer(port int) {
	api := NewAPI() // Create a new API

	http.ListenAndServe(":"+strconv.Itoa(port), api.Router) // Start an HTTP server
}
