package error

// LoginError return a new LoginError instance
func LoginError() *Error {
	return &Error{
		code:    432,
		message: "Login Error, invalid username or password",
	}
}
