package api

import (
	"net/http"
	"strconv"
)

// StartAPIServer starts the API server.
func StartAPIServer(port int) error {
	api := NewAPI(port) // Create a new API

	err := http.ListenAndServe(":"+strconv.Itoa(api.Port), api.Router) // Start an HTTP server
	if err != nil {
		return err
	}

	api.Log.Infof("API server listening on port %d", port)
	return nil
}
