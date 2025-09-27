package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index" json:"deletedAt,omitempty"` // <-- MUDANÇA AQUI
	UserID    uint       `json:"userId"`
	Content   string     `gorm:"not null" json:"content"`
	Likes     int        `gorm:"default:0" json:"likes"`
}

func GetAllPosts(db *gorm.DB) ([]Post, error) {
	var posts []Post
	result := db.Order("created_at desc").Find(&posts)
	return posts, result.Error
}

func CreatePost(db *gorm.DB, post *Post) error {
	result := db.Create(post)
	return result.Error
}

func LikePost(db *gorm.DB, id string) (Post, error) {
	var post Post
	result := db.First(&post, id)
	if result.Error != nil {
		return post, result.Error
	}

	post.Likes++
	db.Save(&post)

	return post, nil
}
