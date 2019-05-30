package api

// GETResponse is the structure for the JSON responses for the GET requests.
type GETResponse struct {
	Output string `json:"output"`
}

// NewGETResponse constructs a new GETResponse struct.
func NewGETResponse(data string) GETResponse {
	output := data

	// If the output is nil, tell the user
	if output == "" {
		output = "error getting output"
	}

	// Return the new response
	return GETResponse{
		Output: output,
	}
}
