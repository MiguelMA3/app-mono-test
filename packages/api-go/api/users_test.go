package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leandro-andrade-candido/api-go/auth"
	"github.com/leandro-andrade-candido/api-go/database"
	"github.com/leandro-andrade-candido/api-go/database/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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
	router.POST("/api/v1/users", CreateUser)

	newUser := `{"username": "testuser", "email": "test@example.com", "bio": "A test user"}`
	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBufferString(newUser))
	req.Header.Set("Content-Type", "application/json")

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
	authorized.Use(AuthMiddleware())
	{
		authorized.GET("/users", GetUsers)
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
	authorized.Use(AuthMiddleware())
	{
		authorized.GET("/users", GetUsers)
	}

	req, _ := http.NewRequest("GET", "/api/v1/users", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
