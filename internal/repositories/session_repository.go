package repositories

import (
	"github.com/google/uuid"
	"github.com/naufan17/go-gin-boilerplate/config"
	"github.com/naufan17/go-gin-boilerplate/internal/models"
)

func CreateSession(session *models.Session) (uuid.UUID, error) {
	if err := config.DB.Create(&session).Error; err != nil {
		return uuid.Nil, err
	}

	return session.ID, nil
}

func GetSessionByUserID(id uuid.UUID) (models.Session, error) {
	var session models.Session

	if err := config.DB.Where("user_id = ?", id).First(&session).Error; err != nil {
		return models.Session{}, err
	}

	return session, nil
}

func UpdateLastActive(id uuid.UUID) (models.Session, error) {
	var session models.Session

	if err := config.DB.Model(&session).Where("id = ?", id).Update("last_active_at", session.LastActiveAt).Error; err != nil {
		return models.Session{}, err
	}

	return models.Session{}, nil
}

func UpdateExpiresAt(id uuid.UUID) (models.Session, error) {
	var session models.Session

	if err := config.DB.Model(&session).Where("id = ?", id).Update("expires_at", session.ExpiresAt).Error; err != nil {
		return models.Session{}, err
	}

	return models.Session{}, nil
}
