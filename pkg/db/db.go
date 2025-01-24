// pkg/db/db.go
package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB es la variable global que mantiene la conexión con la base de datos
var DB *gorm.DB

// InitDB inicializa la conexión con la base de datos PostgreSQL
func InitDB() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),     // Nombre del host de la base de datos, que será 'db' en docker-compose
		os.Getenv("DB_USER"),     // Nombre de usuario, que es 'gorm' según docker-compose
		os.Getenv("DB_PASSWORD"), // Contraseña, que es 'gorm_password' según docker-compose
		os.Getenv("DB_NAME"),     // Nombre de la base de datos, que es 'gorm' según docker-compose
		os.Getenv("DB_PORT"),     // Puerto de la base de datos, que es '5432' según docker-compose
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	log.Println("Conexión a la base de datos establecida exitosamente.")
}
