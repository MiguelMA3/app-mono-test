package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID  uint
	Content string `gorm:"not null"`
	Likes   int    `gorm:"default:0"`
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
