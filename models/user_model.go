package models

import (
	"fmt"
	"musicquiz-api/app/database"
	"musicquiz-api/utils/errors"
	"strings"

	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Email      string `gorm:"type:varchar(100);unique_index"`
	Password   string
	Username   string
}

// Validate validates a User before saving in the DB
func (u *User) Validate() *errors.RestErr {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	// @TODO Add proper passoword validation
	u.Password = strings.TrimSpace(u.Password)
	if u.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}
	return nil
}

// Get gets a User from the DB by its Id
func (u *User) Get(userID int64) *errors.RestErr {
	if err := database.ConnDB.First(&u, userID).Error; err != nil {
		return errors.NewNotFoundtError(fmt.Sprintf("User %d not found", userID))
	}
	return nil
}

// Save saves a new user in the DB
func (u *User) Save() *errors.RestErr {
	if err := database.ConnDB.Create(&u).Error; err != nil {
		return errors.NewBadRequestError(fmt.Sprintf("Email %s already registered", u.Email))
	}
	return nil
}
