package app

import (
	"github.com/LZBEddie/ranking_users-api/controllers/ping"
	"github.com/LZBEddie/ranking_users-api/controllers/users"
)

func mapUrls() {
	//GET methods
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", users.SearchUser)
	//POST methods
	router.POST("/users", users.CreateUser)

}
