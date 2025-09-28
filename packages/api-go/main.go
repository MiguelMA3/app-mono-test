// packages/api-go/main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/leandro-andrade-candido/api-go/api"
	"github.com/leandro-andrade-candido/api-go/database"

	_ "github.com/leandro-andrade-candido/api-go/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API App-Mono-Test
// @version 1.0
// @description API de Gerenciamento de Usuario e Timeline do desafio tecnico

// @contact.name MiguelMA3
// @contact.url https://github.com/MiguelMA3/
// @contact.email miguelm_avelar@hotmail.com

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Digite "Bearer " seguido do seu token JWT.

func main() {
	database.ConnectDatabase()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
			authorized.GET("/users/:id/posts", api.GetPostsByUserID)
		}
	}

	fmt.Println("API rodando na porta 8080")
	router.Run(":8080")
}
