package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	router := gin.Default()

	router.GET("/users", searchUser)
	router.GET("/users/:id", getUserById)
	router.POST("/users", createUser)
	router.PUT("/users/:id", updateUser)
	router.PATCH("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	router.Run()
}

func searchUser(c *gin.Context) {
	c.JSON(http.StatusOK, userList)
}

func getUserById(c *gin.Context) {
	id := c.Param("id")
	u, ok := userList[id]
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, u)
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
	id := c.Param("id")
	u := userList[id]
	if u == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	payload := make(map[string]interface{})
	err := c.BindJSON(&payload)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if v, ok := payload["first_name"]; ok {
		u.FirstName = v.(string)
	}
	if v, ok := payload["last_name"]; ok {
		u.LastName = v.(string)
	}
	if v, ok := payload["email"]; ok {
		u.Email = v.(string)
	}

	c.JSON(http.StatusOK, u)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	_, ok := userList[id]
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	delete(userList, id)
	c.JSON(http.StatusAccepted, nil)
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var userList = map[string]*User{}
