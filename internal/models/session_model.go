package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID           uuid.UUID `json:"id" gorm:"type:char(36);not null"`
	UserID       uuid.UUID `json:"user_id" gorm:"type:char(36);not null"`
	IPAddress    string    `json:"ip_address" gorm:"type:varchar(128);not null"`
	UserAgent    string    `json:"user_agent" gorm:"type:varchar(255);not null"`
	LoginAt      time.Time `json:"login_at" gorm:"autoCreateTime;not null"`
	LastActiveAt time.Time `json:"last_active_at" gorm:"autoUpdateTime;not null"`
	ExpiresAt    time.Time `json:"expires_at" gorm:"not null"`
}

func (session *Session) BeforeCreate(tx *gorm.DB) (err error) {
	session.ID = uuid.New()

	return nil
}
