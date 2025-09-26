package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
	Likes   int    `json:"likes"`
}

var posts = []Post{
	{ID: "1", UserID: "1", Content: "Hello there!", Likes: 0},
}

var nextPostID = 2

func CreatePost(c *gin.Context) {
	var newPost Post
	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPost.ID = strconv.Itoa(nextPostID)
	newPost.Likes = 0
	nextPostID++

	posts = append(posts, newPost)

	c.JSON(http.StatusCreated, newPost)
}

func GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, posts)
}

func LikePost(c *gin.Context) {
	id := c.Param("id")

	for i, p := range posts {
		if p.ID == id {
			posts[i].Likes++
			c.JSON(http.StatusOK, posts[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Postagem não encontrada"})
}
