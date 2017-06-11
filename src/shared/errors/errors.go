package customErrors

// Error is a custom application error
type Error struct {
	ErrorCode string
	Message   string
}

// New creates a new custom error object
func New(code, message string) *Error {
	error := new(Error)
	error.ErrorCode = code
	error.Message = message
	return error
}

// Error returns the error's message
func (err *Error) Error() string {
	return err.Message
}

// Code returns the error's error code
func (err *Error) Code() string {
	return err.ErrorCode
}
