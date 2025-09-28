package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leandro-andrade-candido/api-go/database"
	"github.com/leandro-andrade-candido/api-go/database/models"
)

// CreatePost godoc
// @Summary      Cria uma nova postagem
// @Description  Adiciona uma nova postagem à timeline. O ID do usuário é extraído do token.
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        post  body      models.Post  true  "Conteúdo da Postagem"
// @Success      201   {object}  models.Post
// @Failure      400   {object}  object{error=string}
// @Failure      500   {object}  object{error=string}
// @Router       /posts [post]
// @Security     BearerAuth
func CreatePost(c *gin.Context) {
	var newPost models.Post
	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// --- LÓGICA CORRIGIDA E SEGURA ---
	// 1. Obtém o username do contexto, que foi adicionado pelo middleware de autenticação.
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	// 2. Busca o usuário no banco de dados para obter seu ID.
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao encontrar usuário autenticado"})
		return
	}

	// 3. Associa o ID do usuário autenticado ao novo post, garantindo que ninguém poste em nome de outrem.
	newPost.UserID = user.ID

	if err := models.CreatePost(database.DB, &newPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao criar post"})
		return
	}

	// Recarrega o post com os dados do usuário para que a resposta ao frontend seja completa.
	database.DB.Preload("User").First(&newPost, newPost.ID)

	c.JSON(http.StatusCreated, newPost)
}

// GetPosts godoc
// @Summary      Lista todas as postagens
// @Description  Retorna um array com todas as postagens na timeline, ordenadas da mais recente para a mais antiga.
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Post
// @Failure      500  {object}  object{error=string}
// @Router       /posts [get]
// @Security     BearerAuth
func GetPosts(c *gin.Context) {
	posts, err := models.GetAllPosts(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao buscar posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// LikePost godoc
// @Summary      Curte uma postagem
// @Description  Incrementa o contador de curtidas de uma postagem específica.
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        id    path      int  true  "ID da Postagem"
// @Success      200   {object}  models.Post
// @Failure      404   {object}  object{message=string}
// @Router       /posts/{id}/like [post]
// @Security     BearerAuth
func LikePost(c *gin.Context) {
	id := c.Param("id")

	post, err := models.LikePost(database.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Post nao encontrado"})
		return
	}
	c.JSON(http.StatusOK, post)
}

// GetPostsByUserID godoc
// @Summary      Lista os posts de um usuário específico
// @Description  Retorna um array com todos os posts de um usuário.
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID do Usuário"
// @Success      200  {array}   models.Post
// @Failure      500  {object}  object{error=string}
// @Router       /users/{id}/posts [get]
// @Security     BearerAuth
func GetPostsByUserID(c *gin.Context) {
	userID := c.Param("id")

	posts, err := models.GetPostsByUserID(database.DB, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao buscar posts do usuário"})
		return
	}
	c.JSON(http.StatusOK, posts)
}
