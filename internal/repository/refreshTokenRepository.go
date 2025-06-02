package repository

import (
	"context"
	"time"

	model "jorgerr9011/auth-golang/internal/model/refreshtoken"

	"gorm.io/gorm"
)

type IRefreshTokenRepository interface {
	Create(ctx context.Context, token *model.RefreshToken) error
	FindByToken(ctx context.Context, tokenStr string) (*model.RefreshToken, error)
	DeleteByID(ctx context.Context, id uint) error
	DeleteExpired(ctx context.Context) error
}

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

// Guardar un refresh token
func (r *RefreshTokenRepository) Create(ctx context.Context, token *model.RefreshToken) error {
	return r.db.WithContext(ctx).Create(token).Error
}

// Buscar un refresh token por el token string
func (r *RefreshTokenRepository) FindByToken(ctx context.Context, tokenStr string) (*model.RefreshToken, error) {
	var token model.RefreshToken
	err := r.db.WithContext(ctx).Where("token = ?", tokenStr).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

// Eliminar un refresh token por ID
func (r *RefreshTokenRepository) DeleteByID(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.RefreshToken{}, id).Error
}

// Eliminar todos los refresh tokens expirados
func (r *RefreshTokenRepository) DeleteExpired(ctx context.Context) error {
	return r.db.WithContext(ctx).
		Where("expires_at <= ?", time.Now()).
		Delete(&model.RefreshToken{}).Error
}
