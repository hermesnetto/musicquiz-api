package models

import (
	"fmt"
	"musicquiz-api/utils/date"
	"musicquiz-api/utils/errors"
	"strings"
)

// User model
type User struct {
	ID         int64  `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	CreatedAt  string `json:"created_at"`
}

var usersDB = map[int64]*User{
	1: &User{ID: 1, Email: "john@mail.com", Username: "john_doe", CreatedAt: "hoje"},
	2: &User{ID: 2, Email: "maria@mail.com", Username: "maria10", CreatedAt: "hoje"},
	3: &User{ID: 3, Email: "lucas@mail.com", Username: "luquinhas", CreatedAt: "hoje"},
	4: &User{ID: 4, Email: "homer@mail.com", Username: "homer", CreatedAt: "hoje"},
}

// Validate validates a User before saving in the DB
func (u *User) Validate() *errors.RestErr {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	return nil
}

// Get gets a User from the DB by its Id
func (u *User) Get() *errors.RestErr {
	user := usersDB[u.ID]
	if user == nil {
		return errors.NewNotFoundtError(fmt.Sprintf("User %d not found", u.ID))
	}
	u.ID = user.ID
	u.Email = user.Email
	u.Username = user.Username
	u.CreatedAt = user.CreatedAt
	return nil
}

// Save saves a new user in the DB
func (u *User) Save() *errors.RestErr {
	current := usersDB[u.ID]
	if current != nil {
		if current.Email == u.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %d already registered", u.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already registered", u.ID))
	}
	u.CreatedAt = date.GetNowFormatted()
	usersDB[u.ID] = u
	return nil
}
