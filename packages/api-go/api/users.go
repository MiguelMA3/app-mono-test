package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

var users = []User{
	{ID: "1", Username: "pedrosilva", Email: "pedrosilva@mail.com", Bio: "Bao?"},
}

var nextID = 2

func CreateUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser.ID = strconv.Itoa(nextID)
	nextID++

	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range users {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser User

	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, u := range users {
		if u.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser
			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Usuário não encontrado para atualização"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, u := range users {
		if u.ID == id {
			users[i] = users[len(users)-1]
			users = users[:len(users)-1]

			c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Usuário não encontrado para deleção"})
}
