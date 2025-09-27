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
}
