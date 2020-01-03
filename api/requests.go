package api

import "encoding/json"

// CreateServerRequest is the structure for a POST request to the CreateServer method.
type CreateServerRequest struct {
	Version string `json:"version"`
	Name    string `json:"name"`
	Port    string `json:"port"`
	RAM     string `json:"ram"`
}

// ChangePropertyRequest is the structure for a POST request to the ChangeProperty method.
type ChangePropertyRequest struct {
	Hash     string `json:"hash"`
	Property string `json:"property"`
	NewValue string `json:"newValue"`
}

func (csr *CreateServerRequest) String() string {
	json, _ := json.MarshalIndent(*csr, "", "  ")
	return string(json)
}

func (cpr *ChangePropertyRequest) String() string {
	json, _ := json.MarshalIndent(*cpr, "", "  ")
	return string(json)
}
