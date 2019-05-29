package models

import (
	"fmt"
)

// User exported
type User struct {
	ID     int64
	Name   string
	Emails []string
}

func (u User) String() string {
	return fmt.Sprintf("User: {ID: %d, Name: %s, Emails: %v}", u.ID, u.Name, u.Emails)
}
