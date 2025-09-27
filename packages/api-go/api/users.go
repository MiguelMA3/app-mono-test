package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leandro-andrade-candido/api-go/database"
	"github.com/leandro-andrade-candido/api-go/database/models"
)

func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateUser(database.DB, &newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao criar usuario"})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func GetUsers(c *gin.Context) {

	users, err := models.GetAllUsers(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao buscar usuarios"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := models.GetUserByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Usuario nao encontrado"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	user, err := models.GetUserByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Usuario nao encontrado"})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateUser(database.DB, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao atualizar usuario"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := models.DeleteUser(database.DB, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Usuario não encontrado para delecao"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Usuario deletado com sucesso"})
}
