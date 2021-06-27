package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/user", createUser)
	router.GET("/user/:id", getUser)
	router.DELETE("/user/:id", deleteUser)

	router.Run()
}
