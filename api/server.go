package api

import (
	"net/http"
	"strconv"
)

// StartAPIServer starts the API server.
func StartAPIServer(port int) error {
	api := NewAPI(port) // Create a new API

	api.Log.Infof("API server to listen on port %d", port)
	err := http.ListenAndServe(":"+strconv.Itoa(api.Port), api.Router) // Start an HTTP server
	if err != nil {
		return err
	}

	return nil
}
