package model

import (
	"time"
)

type RefreshToken struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"index;not null"`
	Token     string `gorm:"not null"`
	CreatedAt time.Time
	ExpiresAt time.Time
}
