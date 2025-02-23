package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GinMode          string
	Port             string
	DBName           string
	DBUser           string
	DBPassword       string
	DBHost           string
	DBPort           string
	DBSsl            string
	DBTimezone       string
	DBMaxIdle        string
	DBMaxOpen        string
	JWTAccessSecret  string
	JWTRefreshSecret string
	JWTAccessExp     string
	JWTRefreshExp    string
	AllowedOrigin    string
	MaxRequests      string
	WindowTime       string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		GinMode:          Getenv("GIN_MODE"),
		Port:             Getenv("PORT"),
		DBName:           Getenv("DB_NAME"),
		DBUser:           Getenv("DB_USER"),
		DBPassword:       Getenv("DB_PASSWORD"),
		DBHost:           Getenv("DB_HOST"),
		DBPort:           Getenv("DB_PORT"),
		DBSsl:            Getenv("DB_SSL"),
		DBTimezone:       Getenv("DB_TIMEZONE"),
		DBMaxIdle:        Getenv("DB_MAX_IDLE_CONNS"),
		DBMaxOpen:        Getenv("DB_MAX_OPEN_CONNS"),
		JWTAccessSecret:  Getenv("JWT_ACCESS_SECRET_KEY"),
		JWTRefreshSecret: Getenv("JWT_REFRESH_SECRET_KEY"),
		JWTAccessExp:     Getenv("JWT_ACCESS_EXPIRE_IN"),
		JWTRefreshExp:    Getenv("JWT_REFRESH_EXPIRE_IN"),
		AllowedOrigin:    Getenv("ALLOWED_ORIGIN"),
		MaxRequests:      Getenv("MAX_REQUESTS"),
		WindowTime:       Getenv("WINDOW_TIME"),
	}
}

func Getenv(key string) string {
	return os.Getenv(key)
}
