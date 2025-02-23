package config

import (
	"strconv"

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
	dbUser := cfg.DBUser
	dbPassword := cfg.DBPassword
	dbName := cfg.DBName
	dbPort := cfg.DBPort
	dbSsl := cfg.DBSsl
	dbTimezone := cfg.DBTimezone
	dbMaxIdle, err := strconv.Atoi(cfg.DBMaxIdle)
	dbMaxOpen, err := strconv.Atoi(cfg.DBMaxOpen)

	if err != nil {
		log.Fatal("Invalid DBMaxIdle or DBMaxOpen value")
	}

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + dbSsl + " TimeZone=" + dbTimezone
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal("Failed to get SQL DB from GORM:", err)
	}

	sqlDB.SetMaxIdleConns(dbMaxIdle)
	sqlDB.SetMaxOpenConns(dbMaxOpen)

	log.Println("Connected to database")

	return db
}

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Session{})

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
