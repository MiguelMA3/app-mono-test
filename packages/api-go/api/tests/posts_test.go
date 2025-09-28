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

func setupTestDBAndRouterForPosts(t *testing.T) (*gin.Engine, *models.User) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.Migrator().DropTable(&models.User{}, &models.Post{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&models.User{}, &models.Post{})
	assert.NoError(t, err)

	database.DB = db
	testUser := &models.User{Username: "postuser", Email: "post@test.com"}
	database.DB.Create(testUser)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router, testUser
}

func TestCreatePost(t *testing.T) {
	router, testUser := setupTestDBAndRouterForPosts(t)

	authorized := router.Group("/api/v1")
	authorized.Use(api.AuthMiddleware())
	{
		authorized.POST("/posts", api.CreatePost)
	}

	token, err := auth.GenerateToken(testUser.Username)
	assert.NoError(t, err)

	postContent := `{"content": "Este é o meu primeiro post!"}`
	req, _ := http.NewRequest("POST", "/api/v1/posts", bytes.NewBufferString(postContent))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var post models.Post
	err = json.Unmarshal(w.Body.Bytes(), &post)
	assert.NoError(t, err)
	assert.Equal(t, "Este é o meu primeiro post!", post.Content)
	assert.Equal(t, 0, post.Likes)
}

func TestLikePost(t *testing.T) {
	router, testUser := setupTestDBAndRouterForPosts(t)

	initialPost := models.Post{Content: "Post para curtir", UserID: testUser.ID, Likes: 5}
	database.DB.Create(&initialPost)

	authorized := router.Group("/api/v1")
	authorized.Use(api.AuthMiddleware())
	{
		authorized.POST("/posts/:id/like", api.LikePost)
	}

	token, err := auth.GenerateToken(testUser.Username)
	assert.NoError(t, err)

	url := fmt.Sprintf("/api/v1/posts/%d/like", initialPost.ID)
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var likedPost models.Post
	err = json.Unmarshal(w.Body.Bytes(), &likedPost)
	assert.NoError(t, err)
	assert.Equal(t, initialPost.Likes+1, likedPost.Likes)
}
