package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/leandro-andrade-candido/api-go/api"
	"github.com/leandro-andrade-candido/api-go/database"
)

func main() {
	database.ConnectDatabase()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello There!")
	})

	apiV1 := router.Group("/api/v1")

	apiV1.POST("/users", api.CreateUser)
	apiV1.GET("/users", api.GetUsers)
	apiV1.GET("/users/:id", api.GetUserByID)
	apiV1.PUT("/users/:id", api.UpdateUser)
	apiV1.DELETE("/users/:id", api.DeleteUser)

	apiV1.GET("/posts", api.GetPosts)
	apiV1.POST("/posts", api.CreatePost)
	apiV1.POST("/posts/:id/like", api.LikePost)

	fmt.Println("API rodando na porta 8080")
	router.Run(":8080")
}
