package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leandro-andrade-candido/api-go/database"
	"github.com/leandro-andrade-candido/api-go/database/models"
)

func CreatePost(c *gin.Context) {
	var newPost models.Post
	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := strconv.ParseUint(c.PostForm("user_id"), 10, 64)
	if err != nil {

		newPost.UserID = 1
	} else {
		newPost.UserID = uint(userID)
	}

	if err := models.CreatePost(database.DB, &newPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao criar post"})
		return
	}

	c.JSON(http.StatusCreated, newPost)
}

func GetPosts(c *gin.Context) {

	posts, err := models.GetAllPosts(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao buscar posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func LikePost(c *gin.Context) {
	id := c.Param("id")

	post, err := models.LikePost(database.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Post nao encontrado"})
		return
	}
	c.JSON(http.StatusOK, post)
}
