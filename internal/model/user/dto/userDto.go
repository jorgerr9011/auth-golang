package dto

import "time"

// UserDto representa el modelo de usuario adaptado al modelo GORM
type UserDto struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Username  string     `json:"username"`
	Phone     string     `json:"phone"`
	LastLogin *time.Time `json:"last_login,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// ListUserReq es la solicitud para listar usuarios con filtros opcionales
type ListUserReq struct {
	Name  string `json:"name,omitempty" form:"name"`
	Email string `json:"email,omitempty" form:"email"`
}

// ListUserRes es la respuesta que contiene la lista de usuarios
type ListUserRes struct {
	Users []*UserDto `json:"users"`
}

// CreateUserReq es la solicitud para crear un nuevo usuario
type CreateUserReq struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Username string `json:"username" validate:"required,alphanum"`
	Phone    string `json:"phone" validate:"required,len=9,numeric"`
}

// UpdateUserReq es la solicitud para actualizar un usuario
type UpdateUserReq struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

// LoginUserReq es la solicitud para iniciar sesión
type LoginUserReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
