package services

import (
	"github.com/LZBEddie/ranking_users-api/domain/users"
	"github.com/LZBEddie/ranking_users-api/utils/errors"
)

//GetUser manage to obtain the user
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

//CreateUser manage a user creation
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
