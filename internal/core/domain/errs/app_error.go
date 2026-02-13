package errs

import "fmt"

type Type string

const (
	ValidationType Type = "ValidationError"
	InternalType   Type = "InternalError"
	PermissionType Type = "PermissionError"
)

type AppErrorInterface interface {
	error
	Is(error) bool
	GetCode() int
	GetMessage() string
	GetType() Type
}

type appError struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
	Type    Type   `json:"-"`
}

// New creates a new appError instance with the provided HTTP status code, message, and error type.
// Multiple errors can share the same HTTP status code as this is expected behavior.
//
// Note: The error Type defaults to InternalType if it's not one of the enums options.
func New(code int, message string, errorType Type) AppErrorInterface {
	if errorType != ValidationType && errorType != InternalType && errorType != PermissionType {
		errorType = InternalType
	}

	return appError{
		Message: message,
		Code:    code,
		Type:    errorType,
	}
}

func (e appError) Error() string {
	return fmt.Sprintf("%s - %v", e.Type, e.Message)
}

func (e appError) Is(target error) bool {
	if target == nil {
		return false
	}

	if targetErr, ok := target.(appError); ok {
		// Compare by both code and message to ensure proper error identification
		return e.Code == targetErr.Code && e.Message == targetErr.Message
	}

	return false
}

func (e appError) GetCode() int {
	return e.Code
}

func (e appError) GetType() Type {
	return e.Type
}

func (e appError) GetMessage() string {
	return e.Message
}
