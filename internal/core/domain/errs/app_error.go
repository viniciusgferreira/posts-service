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
	GetCode() string
	GetMessage() string
	GetType() Type
}

type appError struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
	Type    Type   `json:"-"`
}

var errorCodes = map[string]bool{}

// New creates a new appError instance with the provided code, message, and error type.
// Instantiate new sentinels errors from New function as it guarantees code uniqueness.
// Make sure code is unique throughout the application, otherwise it will panic.
//
// Note: The error Type defaults to InternalType if it's not one of the enums options.
func New(code, message string, errorType Type) AppErrorInterface {
	if errorCodes[code] {
		panic(fmt.Sprintf("App error with code %v already exists", code))
	}

	if errorType != ValidationType && errorType != InternalType && errorType != PermissionType {
		errorType = InternalType
	}

	err := appError{
		Message: message,
		Code:    code,
		Type:    errorType,
	}

	errorCodes[code] = true

	return err
}

func (e appError) Error() string {
	return fmt.Sprintf("%s - %v", e.Type, e.Message)
}

func (e appError) Is(target error) bool {
	if target == nil {
		return false
	}

	if targetErr, ok := target.(appError); ok {
		return e.Code == targetErr.Code
	}

	return false
}

func (e appError) GetCode() string {
	return e.Code
}

func (e appError) GetType() Type {
	return e.Type
}

func (e appError) GetMessage() string {
	return e.Message
}
