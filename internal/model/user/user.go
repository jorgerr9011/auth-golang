package model

import (
	"time"
)

/*
	Con gorm.Model ya se crean los siguientes campos:
		- ID
		- CreatedAt
		- UpdatedAt
		- DeletedAt

	Seria interesante también manejar estos atributos de manera autónoma:

		type Product struct {
			ID          string     `json:"id" gorm:"unique;not null;index;primary_key"`
			CreatedAt   time.Time  `json:"created_at"`
			UpdatedAt   time.Time  `json:"updated_at"`
			DeletedAt   *time.Time `json:"deleted_at" gorm:"index"`
			Code        string     `json:"code" gorm:"uniqueIndex:idx_product_code,not null"`
			Name        string     `json:"name" gorm:"uniqueIndex:idx_product_name,not null"`
			Description string     `json:"description"`
			Price       float64    `json:"price"`
			Active      bool       `json:"active" gorm:"default:true"`
		}

		func (m *Product) BeforeCreate(tx *gorm.DB) error {
			m.ID = uuid.New().String()
			m.Code = utils.GenerateCode("P")
			m.Active = true
			return nil
		}
*/

type User struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"size:100;not null"`
	Email     string     `gorm:"size:100;unique;not null"`
	Password  string     `gorm:"size:255;not null"`
	Username  string     `gorm:"size:100;unique;not null"`
	Phone     string     `gorm:"size:9;not null"`
	LastLogin time.Time  `gorm:"default:null"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"index"`
}
