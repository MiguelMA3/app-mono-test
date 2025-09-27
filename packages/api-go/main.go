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
	{
		apiV1.POST("/login", api.Login)

		authorized := apiV1.Group("/")
		authorized.Use(api.AuthMiddleware())
		{
			authorized.POST("/users", api.CreateUser)
			authorized.GET("/users", api.GetUsers)
			authorized.GET("/users/:id", api.GetUserByID)
			authorized.PUT("/users/:id", api.UpdateUser)
			authorized.DELETE("/users/:id", api.DeleteUser)

			authorized.GET("/posts", api.GetPosts)
			authorized.POST("/posts", api.CreatePost)
			authorized.POST("/posts/:id/like", api.LikePost)
		}
	}

	fmt.Println("API rodando na porta 8080")
	router.Run(":8080")
}
