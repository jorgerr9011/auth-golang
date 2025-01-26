package model

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	ID   uint64 `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
	//UserID uint64 `gorm:"index"`
	//User 	User
}
