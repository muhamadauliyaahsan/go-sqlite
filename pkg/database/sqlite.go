package database

import (
	"log"
	"os"

	"github.com/ahsan/go-sqlite-crud/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "modernc.org/sqlite" // Pure-Go SQLite driver (no CGO required)
)

var DB *gorm.DB

func InitDB() {
	dbPath := "api.db"
	if os.Getenv("VERCEL") == "1" {
		dbPath = "/tmp/api.db"
	}

	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return
	}

	// Auto Migration
	if err = DB.AutoMigrate(&model.Product{}); err != nil {
		log.Printf("Failed to migrate database: %v", err)
		return
	}

	log.Println("Database connected and migrated.")
}
