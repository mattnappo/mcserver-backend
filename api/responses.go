package api

// GETResponse is the structure for the JSON responses for the GET requests.
type GETResponse struct {
	Output string `json:"output"`
}

// NewGETResponse constructs a new GETResponse struct.
func NewGETResponse(data string) GETResponse {
	output := data
	if output == "" {
		output = "error getting output"
	}
	return GETResponse{
		Output: output,
	}
}
