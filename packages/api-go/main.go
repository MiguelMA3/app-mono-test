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

	// Configuração do CORS para permitir que o frontend (rodando em localhost:3000)
	// se comunique com a API. Sem isso, o navegador bloquearia as requisições
	// por questões de segurança (Same-Origin Policy).
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Rota para a documentação da API gerada pelo Swagger.
	// O `ginSwagger.WrapHandler` serve a interface do Swagger UI.
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Rota de health-check para verificar se a API está no ar.
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello There!")
	})

	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/login", api.Login)

		// Grupo de rotas que exigem autenticação.
		// O `api.AuthMiddleware()` é aplicado a todas as rotas dentro deste grupo,
		// garantindo que apenas usuários autenticados possam acessá-las.
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
