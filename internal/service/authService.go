package service

import (
	"context"
	"errors"
	tokenModel "jorgerr9011/auth-golang/internal/model/refreshtoken"
	userModel "jorgerr9011/auth-golang/internal/model/user"
	"jorgerr9011/auth-golang/internal/model/user/dto"
	"jorgerr9011/auth-golang/internal/repository"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Authenticate(ctx context.Context, req *dto.LoginUserReq) (*userModel.User, error)
	Register(ctx context.Context, req *dto.CreateUserReq) (*userModel.User, error)
	SaveRefreshToken(ctx context.Context, userID uint, token string, expiresAt time.Time) error
}

type AuthService struct {
	repo             repository.IUserRepository
	refreshTokenRepo repository.IRefreshTokenRepository
}

func NewAuthService(
	repo repository.IUserRepository,
	refreshTokenRepo repository.IRefreshTokenRepository,
) *AuthService {
	return &AuthService{
		repo:             repo,
		refreshTokenRepo: refreshTokenRepo,
	}
}

func (s *AuthService) Authenticate(ctx context.Context, req *dto.LoginUserReq) (*userModel.User, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return nil, errors.New("credenciales inválidas")
	}

	// Verificar la contraseña
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	return user, nil
}

func (s *AuthService) Register(ctx context.Context, req *dto.CreateUserReq) (*userModel.User, error) {
	existingUser, _ := s.repo.GetByEmail(ctx, req.Email)
	if existingUser != nil {
		log.Println("Intento de registro con correo duplicado:", req.Email)
		return nil, errors.New("el correo electrónico ya está en uso")
	}

	// Cifrar la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error al cifrar la contraseña:", err)
		return nil, errors.New("error al procesar la contraseña")
	}

	user := &userModel.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Username: req.Username,
		Phone:    req.Phone,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		log.Println("Error al registrar el usuario:", err)
		return nil, errors.New("no se pudo registrar el usuario")
	}

	return user, nil
}

func (s *AuthService) SaveRefreshToken(ctx context.Context, userID uint, token string, expiresAt time.Time) error {
	rt := &tokenModel.RefreshToken{
		UserID:    userID,
		Token:     token,
		CreatedAt: time.Now(),
		ExpiresAt: expiresAt,
	}
	return s.refreshTokenRepo.Create(ctx, rt)
}
