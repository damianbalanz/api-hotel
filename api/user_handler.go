package api

import (
	"github.com/damianbalanz/api-hotel/types"
	"github.com/gin-gonic/gin"
)

func HandleGetUsers(c *gin.Context) {
	u := types.User{
		Firstname: "Dami",
		LastName:  "Abalos",
	}
	c.JSON(200, u)
}

func HandleGetUsersById(c *gin.Context) {
	c.JSON(200,
		"user dami",
	)
}
