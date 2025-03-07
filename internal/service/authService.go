package service

import (
	"context"
	"errors"
	model "jorgerr9011/wiki-golang/internal/model/user"
	"jorgerr9011/wiki-golang/internal/model/user/dto"
	"jorgerr9011/wiki-golang/internal/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Authenticate(ctx context.Context, req *dto.LoginUserReq) (*model.User, error)
	Register(ctx context.Context, req *dto.CreateUserReq) (*model.User, error)
}

type AuthService struct {
	repo repository.IUserRepository
}

func NewAuthService(
	repo repository.IUserRepository,
) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Authenticate(ctx context.Context, req *dto.LoginUserReq) (*model.User, error) {
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

func (s *AuthService) Register(ctx context.Context, req *dto.CreateUserReq) (*model.User, error) {
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

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		log.Println("Error al registrar el usuario:", err)
		return nil, errors.New("no se pudo registrar el usuario")
	}

	return user, nil
}
