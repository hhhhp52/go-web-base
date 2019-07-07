package error

// ValidateError return a new ValidateError instance
func ValidateError(msg string) *Error {
	return &Error{
		code:    433,
		message: "Validate Error: " + msg,
	}
}
