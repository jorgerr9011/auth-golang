package seeders

import (
	"fmt"
	"jorgerr9011/auth-golang/migrations/factories"
	"log"

	"gorm.io/gorm"
)

// UserSeeder inserta 10 usuarios en la BD
func UserSeeder(db *gorm.DB) {
	for i := 0; i < 100; i++ {
		user := factories.UserFactory()
		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("Error insertando usuario: %v", err)
		}
		fmt.Println("✔ Usuario creado:", user.Username, user.Email)
	}
}
