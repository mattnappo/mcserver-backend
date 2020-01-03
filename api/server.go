package api

import (
	"net/http"
	"strconv"

	"github.com/juju/loggo"
)

// StartAPIServer starts the API server.
func StartAPIServer(port int) {
	api := NewAPI(port) // Create a new API

	http.ListenAndServe(":"+strconv.Itoa(api.Port), api.Router) // Start an HTTP server
	api.Log.Logf(loggo.INFO, "API server listening on port %d", port)
}
