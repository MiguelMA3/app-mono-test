package database

import (
	"log"

	"github.com/leandro-andrade-candido/api-go/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

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

	var userCount int64
	DB.Model(&models.User{}).Count(&userCount)

	if userCount == 0 {
		log.Println("Novo banco de dados detectado, criando usuario padrao...")
		adminUser := models.User{
			Username: "admin",
			Email:    "admin@example.com",
			Bio:      "Hello there",
		}
		if err := DB.Create(&adminUser).Error; err != nil {
			log.Fatalf("Falha ao criar usuario admin padrao: %v", err)
		}
		log.Println("Usuario 'admin' criado com sucesso.")
	}
}
