package main

import "github.com/gin-gonic/gin"

func HandleGetUsers(c *gin.Context) {
	c.JSON(200,
		"user dami",
	)
}
