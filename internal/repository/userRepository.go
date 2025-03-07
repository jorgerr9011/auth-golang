package repository

import (
	"context"
	"log"
	"time"

	model "jorgerr9011/wiki-golang/internal/model/user"
	"jorgerr9011/wiki-golang/internal/model/user/dto"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetAll(ctx context.Context, req *dto.ListUserReq) ([]*model.User, error)
	GetById(ctx context.Context, id string) (*model.User, error)
	Delete(ctx context.Context, id string) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {

	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepository) GetAll(ctx context.Context, req *dto.ListUserReq) ([]*model.User, error) {

	var users []*model.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetById(ctx context.Context, id string) (*model.User, error) {

	var user model.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		log.Println("No se ha encontrado un usuario con el mismo correo:", err)
		return nil, err
	}
	return &user, nil
}
