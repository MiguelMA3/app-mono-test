package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leandro-andrade-candido/api-go/api"
	"github.com/leandro-andrade-candido/api-go/auth"
	"github.com/leandro-andrade-candido/api-go/database"
	"github.com/leandro-andrade-candido/api-go/database/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createUserAndToken(t *testing.T, db *gorm.DB, username string) (*models.User, string) {
	user := &models.User{Username: username, Email: username + "@test.com"}
	err := db.Create(user).Error
	assert.NoError(t, err)

	token, err := auth.GenerateToken(user.Username)
	assert.NoError(t, err)

	return user, token
}

func setupTestDBAndRouter(t *testing.T) *gin.Engine {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.Migrator().DropTable(&models.User{}, &models.Post{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&models.User{}, &models.Post{})
	assert.NoError(t, err)

	database.DB = db

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func TestCreateUser(t *testing.T) {
	router := setupTestDBAndRouter(t)
	_, token := createUserAndToken(t, database.DB, "admin_user")

	authorized := router.Group("/api/v1")
	authorized.Use(api.AuthMiddleware())
	{
		authorized.POST("/users", api.CreateUser)
	}

	newUser := `{"username": "testuser", "email": "test@example.com", "bio": "A test user"}`
	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBufferString(newUser))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var user models.User
	err := json.Unmarshal(w.Body.Bytes(), &user)
	assert.NoError(t, err)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "test@example.com", user.Email)
}

func TestGetUsers_Authorized(t *testing.T) {
	router := setupTestDBAndRouter(t)

	testUser := models.User{Username: "authtest", Email: "auth@test.com"}
	database.DB.Create(&testUser)

	authorized := router.Group("/api/v1")
	authorized.Use(api.AuthMiddleware())
	{
		authorized.GET("/users", api.GetUsers)
	}

	token, err := auth.GenerateToken(testUser.Username)
	assert.NoError(t, err)

	req, _ := http.NewRequest("GET", "/api/v1/users", nil)

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var users []models.User
	err = json.Unmarshal(w.Body.Bytes(), &users)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "authtest", users[0].Username)
}

func TestGetUsers_Unauthorized(t *testing.T) {
	router := setupTestDBAndRouter(t)

	authorized := router.Group("/api/v1")
	authorized.Use(api.AuthMiddleware())
	{
		authorized.GET("/users", api.GetUsers)
	}

	req, _ := http.NewRequest("GET", "/api/v1/users", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestGetUserByID(t *testing.T) {

	router := setupTestDBAndRouter(t)
	createdUser, token := createUserAndToken(t, database.DB, "getbyid_user")

	authorized := router.Group("/api/v1")
	authorized.Use(api.AuthMiddleware())
	{
		authorized.GET("/users/:id", api.GetUserByID)
	}

	url := fmt.Sprintf("/api/v1/users/%d", createdUser.ID)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var foundUser models.User
	err := json.Unmarshal(w.Body.Bytes(), &foundUser)
	assert.NoError(t, err)
	assert.Equal(t, createdUser.Username, foundUser.Username)
	assert.Equal(t, createdUser.ID, foundUser.ID)
}

func TestUpdateUser(t *testing.T) {

	router := setupTestDBAndRouter(t)
	userToUpdate, token := createUserAndToken(t, database.DB, "update_user")

	authorized := router.Group("/api/v1")
	authorized.Use(api.AuthMiddleware())
	{
		authorized.PUT("/users/:id", api.UpdateUser)
	}

	updatePayload := `{"username": "updated_user", "email": "update@test.com", "bio": "Bio atualizada"}`
	url := fmt.Sprintf("/api/v1/users/%d", userToUpdate.ID)
	req, _ := http.NewRequest("PUT", url, bytes.NewBufferString(updatePayload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var updatedUser models.User
	err := json.Unmarshal(w.Body.Bytes(), &updatedUser)
	assert.NoError(t, err)
	assert.Equal(t, "updated_user", updatedUser.Username)
	assert.Equal(t, "Bio atualizada", updatedUser.Bio)
}

func TestDeleteUser(t *testing.T) {

	router := setupTestDBAndRouter(t)
	userToDelete, token := createUserAndToken(t, database.DB, "delete_user")

	authorized := router.Group("/api/v1")
	authorized.Use(api.AuthMiddleware())
	{
		authorized.DELETE("/users/:id", api.DeleteUser)
	}

	url := fmt.Sprintf("/api/v1/users/%d", userToDelete.ID)
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var deletedUser models.User
	err := database.DB.Unscoped().First(&deletedUser, userToDelete.ID).Error
	assert.NoError(t, err)
	assert.NotNil(t, deletedUser.DeletedAt)
}
