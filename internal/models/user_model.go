package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             uuid.UUID `json:"id" gorm:"type:char(36);not null"`
	Name           string    `json:"name" gorm:"type:varchar(128);not null"`
	Email          string    `json:"email" gorm:"type:varchar(128);not null;unique"`
	Address        string    `json:"address" gorm:"type:varchar(255);null"`
	PhoneNumber    string    `json:"phone_number" gorm:"type:varchar(20);null"`
	ProfilePicture string    `json:"profile_picture" gorm:"type:varchar(255);null"`
	IsVerified     bool      `json:"is_verified" gorm:"default:false"`
	Password       string    `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()

	return nil
}
