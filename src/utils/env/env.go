package env

import (
	"os"
)

func Set() {
	os.Setenv("Gmail_Account", "xxxxx@gmail.com")
	os.Setenv("Gmail_Password", "xxxxx")
}
