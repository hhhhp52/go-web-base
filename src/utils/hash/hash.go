package hash

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

//New takes a string and returns a hashed one
func New(data string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	return string(hashed)
}

// Verify decides whether hashed is generated from raw
func Verify(raw string, hashed string) bool {
	fmt.Println("hash :",[]byte(hashed))
	fmt.Println("raw  :",[]byte(raw))
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(raw))

	if err != nil {
		return false
	}
	return true
}
