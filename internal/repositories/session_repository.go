package repositories

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/naufan17/go-gin-boilerplate/config"
	"github.com/naufan17/go-gin-boilerplate/internal/models"
)

func CreateSession(userId uuid.UUID, ipAddress string, userAgent string) (models.Session, error) {
	jwtRefreshExpStr := config.LoadConfig().JWTRefreshExp
	jwtRefreshExp, _ := strconv.Atoi(jwtRefreshExpStr)

	session := models.Session{
		UserID:    userId,
		IPAddress: ipAddress,
		UserAgent: userAgent,
		ExpiresAt: time.Now().Add(time.Duration(jwtRefreshExp) * time.Millisecond),
	}

	if err := config.DB.Create(&session).Error; err != nil {
		return models.Session{}, err
	}

	return session, nil
}

func GetSessionByID(id uuid.UUID) (models.Session, error) {
	var session models.Session

	if err := config.DB.Where("id = ?", id).First(&session).Error; err != nil {
		return models.Session{}, err
	}

	return session, nil
}

func GetSessionByUserID(id uuid.UUID) ([]models.Session, error) {
	var session []models.Session

	if err := config.DB.Where("user_id = ?", id).Order("login_at desc").Find(&session).Error; err != nil {
		return []models.Session{}, err
	}

	return session, nil
}

func UpdateLastActive(id uuid.UUID) (models.Session, error) {
	var session models.Session

	if err := config.DB.Model(&session).Where("id = ?", id).Updates(models.Session{LastActiveAt: time.Now()}).Error; err != nil {
		return models.Session{}, err
	}

	return models.Session{}, nil
}

func UpdateExpiresAt(id uuid.UUID) (models.Session, error) {
	var session models.Session

	if err := config.DB.Model(&session).Where("user_id = ?", id).Updates(models.Session{ExpiresAt: time.Now()}).Error; err != nil {
		return models.Session{}, err
	}

	return models.Session{}, nil
}
