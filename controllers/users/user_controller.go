package users

import (
	"net/http"
	"strconv"

	"github.com/LZBEddie/ranking_users-api/domain/users"
	"github.com/LZBEddie/ranking_users-api/services"
	"github.com/LZBEddie/ranking_users-api/utils/errors"

	"github.com/gin-gonic/gin"
)

//GetUser gets an user
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user_id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)

}

//CreateUser creates an user
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	//fmt.Println("user: ", user)

	c.JSON(http.StatusCreated, result)
}

//SearchUser searchs an user
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")

}
