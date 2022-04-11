package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	router := gin.Default()

	router.GET("/user/:id", getUser)
	router.POST("/user", createUser)
	router.PUT("/user/:id", updateUser)
	router.PATCH("/user/:id", updateUser)
	router.DELETE("/user/:id", deleteUser)

	router.Run()
}

func getUser(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func createUser(c *gin.Context) {
	var u User
	err := c.BindJSON(&u)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	u.ID = uuid.New().String()
	userList[u.ID] = &u
	c.JSON(http.StatusOK, u)
}

func updateUser(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func deleteUser(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
}

var userList = map[string]*User{}
