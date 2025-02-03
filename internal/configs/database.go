package configs

import (
	"github.com/naufan17/go-gin-boilerplate/database/seeders"
	"github.com/naufan17/go-gin-boilerplate/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
)

var DB *gorm.DB

func init() {
	DB = ConnectDB()
}

func ConnectDB() *gorm.DB {
	cfg := LoadConfig()

	dbHost := cfg.DBHost
	dbPort := cfg.DBPort
	dbUser := cfg.DBUser
	dbPassword := cfg.DBPassword
	dbName := cfg.DBName

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	} else {
		log.Println("Connected to database")
	}

	return db
}

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Failed to migrate database", err)
	} else {
		log.Println("Database migrated successfully")
	}
}

func SeedAll(db *gorm.DB) {
	seeders.SeedUsers(db)

	log.Println("Database seeded successfully")
}
