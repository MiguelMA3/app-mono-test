package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leandro-andrade-candido/api-go/auth"
	"github.com/leandro-andrade-candido/api-go/database"
	"github.com/leandro-andrade-candido/api-go/database/models"
)

// CreateUser godoc
// @Summary      Cria um novo usuário
// @Description  Adiciona um novo usuário ao banco de dados com base no corpo da requisição.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "Dados do Usuário para Criar"
// @Success      201   {object}  models.User
// @Failure      400   {object}  object{error=string}
// @Failure      500   {object}  object{error=string}
// @Router       /users [post]
// @Security     BearerAuth
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

// GetUsers godoc
// @Summary      Lista todos os usuários
// @Description  Retorna um array com todos os usuários cadastrados no banco de dados.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.User
// @Failure      500  {object}  object{error=string}
// @Router       /users [get]
// @Security     BearerAuth
func GetUsers(c *gin.Context) {

	users, err := models.GetAllUsers(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao buscar usuarios"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID godoc
// @Summary      Busca um usuário pelo ID
// @Description  Retorna os detalhes de um usuário específico.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id    path      int  true  "ID do Usuário"
// @Success      200   {object}  models.User
// @Failure      404   {object}  object{message=string}
// @Router       /users/{id} [get]
// @Security     BearerAuth
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := models.GetUserByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Usuario nao encontrado"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary      Atualiza um usuário
// @Description  Atualiza os dados de um usuário existente com base no ID e no corpo da requisição.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id    path      int          true  "ID do Usuário"
// @Param        user  body      models.User  true  "Dados do Usuário para Atualizar"
// @Success      200   {object}  models.User
// @Failure      400   {object}  object{error=string}
// @Failure      404   {object}  object{message=string}
// @Failure      500   {object}  object{error=string}
// @Router       /users/{id} [put]
// @Security     BearerAuth
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

// DeleteUser godoc
// @Summary      Deleta um usuário
// @Description  Remove um usuário do banco de dados (soft delete).
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id    path      int  true  "ID do Usuário"
// @Success      200   {object}  object{message=string}
// @Failure      404   {object}  object{message=string}
// @Router       /users/{id} [delete]
// @Security     BearerAuth
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := models.DeleteUser(database.DB, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Usuario não encontrado para delecao"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Usuario deletado com sucesso"})
}

type LoginRequest struct {
	Username string `json:"username" biding:"required"`
	Password string `json:"password" biding:"required"`
}

// Login godoc
// @Summary      Autentica um usuário
// @Description  Verifica as credenciais e retorna um token JWT se forem válidas.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        credentials body      LoginRequest  true  "Credenciais de Login"
// @Success      200         {object}  object{token=string}
// @Failure      400         {object}  object{error=string}
// @Failure      401         {object}  object{error=string}
// @Failure      500         {object}  object{error=string}
// @Router       /login [post]
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "usuario e senha sao obrigatorios"})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "usuario ou senha incorretos"})
		return
	}

	if req.Password != "1234" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "usuario ou senha incorretos"})
		return
	}

	token, err := auth.GenerateToken(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":    token,
		"userId":   user.ID,
		"username": user.Username,
	})
}
