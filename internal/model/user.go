package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//ID    uint64 `gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
	//Documents 	[]Document
}
