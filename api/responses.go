package api

// GETResponse is the structure for the JSON responses for the GET requests.
type GETResponse struct {
	Output string `json:"output"`
}

// NewGETResponse constructs a new GETResponse struct.
func NewGETResponse(output string) GETResponse {
	// Return the new response
	return GETResponse{
		Output: output,
	}
}
