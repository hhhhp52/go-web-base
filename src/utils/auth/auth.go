package auth

var (
	account string
	ip      string
)

// Set store token claim
func Set(params map[string]interface{}) {
	account = params["account"].(string)
	ip = params["ip"].(string)
}

// Account return the user name in token
func Account() string {
	return account
}

// IP return remote ip address
func IP() string {
	return ip
}
