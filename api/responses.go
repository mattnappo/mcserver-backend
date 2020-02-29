package api

import "github.com/xoreo/mcserver-backend/types"

// GETResponse is the structure for the JSON responses for GET request responses.
type GETResponse struct {
	Output string `json:"output"`
}

// GETServerResponse is the structure for the JSON responses for the
// GET request responses containing a server.
type GETServerResponse struct {
	Timestamp string `json:"timestamp"`

	ID string `json:"id"`

	Properties     types.Properties     `json:"properties"`
	CoreProperties types.CoreProperties `json:"coreProperties"`
}

// GETServersResponse is the structure for the JSON responses for the
// GET request responses containing multiple servers.
type GETServersResponse struct {
	Server []types.Server `json:"servers"`
}

// ErrorResponse is the structure for when an error is thrown on the server side.
type ErrorResponse struct {
	Error string `json:"error"`
}

// DefaultResponse is the structure for an uncategorized response.
type DefaultResponse struct {
	Data string `json:"data"`
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

// NewDefaultResponse constructs a DefaultResponse.
func NewDefaultResponse(data string) DefaultResponse {
	// Return the new response
	return DefaultResponse{
		Data: data,
	}
}
