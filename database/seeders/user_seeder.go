package seeders

import (
	"log"

	"github.com/naufan17/go-gin-boilerplate/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	users := []models.User{
		{Name: "John Doe", Email: "jhon@example.com", Password: string(hashedPassword)},
		{Name: "Jane Doe", Email: "jane@example.com", Password: string(hashedPassword)},
		{Name: "Mark Doe", Email: "mark@example.com", Password: string(hashedPassword)},
		{Name: "Alice Doe", Email: "alice@example.com", Password: string(hashedPassword)},
	}

	if err := db.Create(&users).Error; err != nil {
		log.Fatalf("Failed to seed users: %v", err)
	}
}
