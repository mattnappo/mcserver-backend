package api

// GETResponse is the structure for the JSON responses for the GET requests.
type GETResponse struct {
	Output string `json:"output"`
}

// ErrorResponse is the structure for when an error is thrown on the server side.
type ErrorResponse struct {
	Error string `json:"error"`
}

// NewGETResponse constructs a new GETResponse struct.
func NewGETResponse(output string) GETResponse {
	// Return the new response
	return GETResponse{
		Output: output,
	}
}

// NewErrorResponse constructs an ErrorResponse.
func NewErrorResponse(err string) ErrorResponse {
	// Return the new response
	return ErrorResponse{
		Error: err,
	}
}
