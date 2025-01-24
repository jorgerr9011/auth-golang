package model

import "gorm.io/gorm"

// Modelo de Usuario
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}
