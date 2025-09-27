package api

import (
	"net/http"
	"strconv"

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
