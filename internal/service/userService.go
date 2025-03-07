package service

import (
	"context"
	model "jorgerr9011/wiki-golang/internal/model/user"
	"jorgerr9011/wiki-golang/internal/model/user/dto"
	"jorgerr9011/wiki-golang/internal/repository"
	"log"
)

type IUserService interface {
	GetAllUsers(c context.Context, req *dto.ListUserReq) ([]*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, req *dto.CreateUserReq) (*model.User, error)
	UpdateUser(ctx context.Context, id string, req *dto.UpdateUserReq) (*model.User, error)
	DeleteUser(ctx context.Context, id uint) error
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(
	repo repository.IUserRepository,
) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context, req *dto.ListUserReq) ([]*model.User, error) {
	users, err := s.repo.GetAll(ctx, req)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *dto.CreateUserReq) (*model.User, error) {
	user := &model.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id string, req *dto.UpdateUserReq) (*model.User, error) {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		log.Printf("Error getting user by ID: %v", err)
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}

	if err := s.repo.Update(ctx, user); err != nil {
		log.Printf("Error updating user: %v", err)
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}
	return nil
}
