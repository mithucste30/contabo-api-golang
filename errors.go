package contabo

import (
	"errors"
	"fmt"
)

// Common errors
var (
	ErrMissingClientID     = errors.New("client ID is required")
	ErrMissingClientSecret = errors.New("client secret is required")
	ErrMissingUsername     = errors.New("username (API user email) is required")
	ErrMissingPassword     = errors.New("password (API password) is required")
	ErrAuthenticationFailed = errors.New("authentication failed")
	ErrInvalidToken        = errors.New("invalid or expired token")
)

// APIError represents an error returned by the Contabo API
type APIError struct {
	StatusCode int
	Message    string
	TraceID    string
	RequestID  string
}

// Error implements the error interface
func (e *APIError) Error() string {
	return fmt.Sprintf("API error (status %d): %s", e.StatusCode, e.Message)
}

// NewAPIError creates a new APIError
func NewAPIError(statusCode int, message, requestID, traceID string) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Message:    message,
		RequestID:  requestID,
		TraceID:    traceID,
	}
}
