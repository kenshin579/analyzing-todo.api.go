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

// NotFoundErrorCode will be translated to a 404 HTTP response code
const NotFoundErrorCode string = "NOT_FOUND"

// InteralServerErrorCode willl be translated to a 500 HTTP response code
const InternalServerErrorCode string = "INTERNAL_SERVER_ERROR"

// ServiceUnavailableErrorCode will be translated to a 503 response code
const ServiceUnavailableErrorCode string = "SERVICE_UNAVAILABLE"
