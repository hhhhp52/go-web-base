package error

// UnexpectedError return a new UnexpectedError instance
func UnexpectedError() *Error {
	return &Error{
		code:    999,
		message: "Unexpected Error",
	}
}
