package main

import (
	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"privet": "poka",
		"123": 456,
	})
}

func main() {
	g := gin.Default()

	g.GET("/hi", getUser)
	g.Run(":8080")
}