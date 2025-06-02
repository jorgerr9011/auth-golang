package dto

import "time"

type RefreshTokenDTO struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
}
