package common

// APIError generic struc for API error
type APIError struct {

	// error code
	ErrorCode int32 `json:"errorCode,omitempty"`

	// http code
	HTTPCode int32 `json:"httpCode,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}
