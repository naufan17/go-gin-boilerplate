package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/naufan17/go-gin-boilerplate/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = ConnectDB()
}

func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

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
