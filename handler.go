package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {
	var jsonUser User
	c.BindJSON(&jsonUser)

	var s userStorage
	s = newDataBaseStorage()

	insertedPrimaryKey, err := s.insert(jsonUser)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"pk": insertedPrimaryKey,
	})

}

func getUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)

	var s userStorage
	s = newDataBaseStorage()

	user, err := s.get(userId)

	if err != nil {
		c.JSON(404, gin.H{
			"MESSAGE": "RESOURCE NOT FOUND",
		})
	} else {
		c.Bind(*&user)
		c.JSON(200, gin.H{
			"id":        user.Id,
			"firstname": user.Firstname,
			"lastname":  user.Lastname,
			"name":      user.Name,
			"email":     user.Email,
		})
	}
}

//func updateUser(c *gin.Context) {
//	return nil
//}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)

	var s userStorage
	s = newDataBaseStorage()

	rowsAffected, err := s.delete(userId)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"rowsAffected": rowsAffected,
	})

}
