package api

// CreateServerRequest is the format for a POST request to CreateServer.
type CreateServerRequest struct {
	Version string `json:"version"`
	Name    string `json:"name"`
	Port    string `json:"port"`
	RAM     string `json:"ram"`
}
