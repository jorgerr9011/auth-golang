package model

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	Name string `gorm:"size:255"`
	//UserID uint64 `gorm:"index"`
	//User 	User
}
