package api

import "net/http"

// StartAPIServer starts the` API server.
func StartAPIServer() {
	api := NewAPI() // Create a new API

	http.ListenAndServe(":8000", api.Router) // Start an HTTP server
}
