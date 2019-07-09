package error

// DataNotFoundError return a new DataNotFoundError instance
func DataNotFoundError(msg string) *Error {
	return &Error{
		code:    437,
		message: "Data Not Found: " + msg,
	}
}
