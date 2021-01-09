package users

import (
	"strings"

	"github.com/LZBEddie/ranking_users-api/utils/errors"
)

//User struct for user
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

//Validate validates user fields
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	return nil
}
