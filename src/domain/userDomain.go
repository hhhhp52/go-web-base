package domain

import (
	"time"
)

// User is the structure for accessing table user
type User struct {
	ID        int64
	Account   string
	Password  string
	Name      string
	Nickname  string
	Email     string
	Status    string
	CreatedAt time.Time
	UpdatedAt *time.Time
	CreatedBy string
	UpdatedBy string
}
