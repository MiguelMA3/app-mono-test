package database

import (
	"log"

	"github.com/leandro-andrade-candido/api-go/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func seedDB() {
	var userCount int64
	DB.Model(&models.User{}).Count(&userCount)

	if userCount > 0 {
		return
	}

	log.Println("Nenhum usuario encontrado, populando database...")

	users := []models.User{
		{Username: "admin", Email: "admin@example.com", Bio: "Administrador do sistema."},
		{Username: "jedi_master", Email: "jedimaster@example.com", Bio: "May the 4th be with you"},
		{Username: "droid_general", Email: "droidgeneral@example.com", Bio: "Hello there!"},
	}

	if err := DB.Create(&users).Error; err != nil {
		log.Fatalf("Falha ao criar usuarios padrao: %v", err)
	}
	log.Println("Usuarios padrao criados com sucesso.")

	posts := []models.Post{
		{UserID: users[1].ID, Content: "Hello there!", Likes: 15},
		{UserID: users[2].ID, Content: "General Kenobi!", Likes: 22},
		{UserID: users[1].ID, Content: "I have the high ground", Likes: 8},
		{UserID: users[0].ID, Content: "A long time ago in a galaxy far, far away", Likes: 50},
	}

	if err := DB.Create(&posts).Error; err != nil {
		log.Fatalf("Falha ao criar posts padrao: %v", err)
	}
	log.Println("Posts padrao criados com sucesso.")
}

func ConnectDatabase() {
	var err error
	dbPath := "database.sqlite"

	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha na conexao com Database: %v", err)
	}
	log.Println("Banco de Dados Conectado! BOYAH!")

	err = DB.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		log.Fatalf("Falha na migracao do Database: %v", err)
	}
	log.Println("Migracao do Database OK!")

	seedDB()
}
