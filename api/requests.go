package api

// CreateServerRequest is the structure for a POST request to the CreateServer method.
type CreateServerRequest struct {
	Version string `json:"version"`
	Name    string `json:"name"`
	Port    string `json:"port"`
	RAM     string `json:"ram"`
}
