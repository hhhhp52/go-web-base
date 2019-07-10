package error

// ValidateError return a new ValidateError instance
func InvaildValueError(msg string) *Error {
	return &Error{
		code:    436,
		message: "Invalidate Error: " + msg,
	}
}
