package users

import (
	"fmt"

	"github.com/LZBEddie/ranking_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

//Get gets an user by ID
func (user *User) Get() *errors.RestErr {
	result := userDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.ID))
	}
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Nickname = result.Nickname
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

//Save saves an User on the DB
func (user *User) Save() *errors.RestErr {
	if userDB[user.ID] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.ID))
	}
	userDB[user.ID] = user
	return nil
}
