package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	JWTSecret  string
	JWTExp     string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Port:       Getenv("PORT"),
		DBName:     Getenv("DB_NAME"),
		DBUser:     Getenv("DB_USER"),
		DBPassword: Getenv("DB_PASSWORD"),
		DBHost:     Getenv("DB_HOST"),
		DBPort:     Getenv("DB_PORT"),
		JWTSecret:  Getenv("JWT_SECRET_KEY"),
		JWTExp:     Getenv("JWT_EXPIRE_IN"),
	}
}

func Getenv(key string) string {
	return os.Getenv(key)
}
