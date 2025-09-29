package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index" json:"deletedAt,omitempty"`
	Username  string     `gorm:"uniqueIndex;not null" json:"username"`
	Email     string     `gorm:"uniqueIndex;not null" json:"email"`
	Bio       string     `json:"bio"`
}

func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Find(&users)
	return users, result.Error
}

func GetUserByID(db *gorm.DB, id string) (User, error) {
	var user User
	result := db.First(&user, id)
	return user, result.Error
}

func CreateUser(db *gorm.DB, user *User) error {
	result := db.Create(user)
	return result.Error
}

func UpdateUser(db *gorm.DB, user *User) error {
	result := db.Save(user)
	return result.Error
}

func DeleteUser(db *gorm.DB, id string) error {
	result := db.Model(&User{}).Where("id = ?", id).Update("deleted_at", time.Now())

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
