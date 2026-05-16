package database

import (
	"log"
	"os"

	"github.com/ahsan/go-sqlite-crud/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbPath := "api.db"
	if os.Getenv("VERCEL") == "1" {
		dbPath = "/tmp/api.db"
	}

	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migration
	err = DB.AutoMigrate(&model.Product{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connected and migrated.")
}
