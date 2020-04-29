package services

import (
	"musicquiz-api/models"
	"musicquiz-api/utils/errors"
)

// GetUser returns a User
func GetUser(userID int64) (*models.User, *errors.RestErr) {
	var user = &models.User{}
	if err := user.Get(userID); err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser creates a new User
func CreateUser(user *models.User) (*models.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return user, nil
}
