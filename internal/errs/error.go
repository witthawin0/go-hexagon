package errs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// AppError represents a standardized error structure
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	return fmt.Sprintf("code: %d, message: %s, error: %v", e.Code, e.Message, e.Err)
}

// NewAppError creates a new AppError
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// WrapError wraps an existing error into an AppError
func WrapError(code int, message string, err error) *AppError {
	return NewAppError(code, message, err)
}

// ErrorResponse sends a JSON error response
func ErrorResponse(w http.ResponseWriter, err *AppError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Message})
}
