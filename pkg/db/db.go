// pkg/db/db.go
package db

import (
	"fmt"
	"log"

	"jorgerr9011/wiki-golang/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

// DB es la variable global que mantiene la conexión con la base de datos
var DB *gorm.DB

// InitDB inicializa la conexión con la base de datos PostgreSQL
func NewDatabase(uri string) {
	var err error
	DB, err = gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	log.Println("Conexión a la base de datos establecida exitosamente.")
}

// GenerateDSN genera un Data Source Name (DSN) para la conexión a la base de datos.
func GenerateDSN(cfg config.Config) string {

	if cfg.Db_host == "" || cfg.Db_user == "" || cfg.Db_password == "" || cfg.Db_name == "" || cfg.Db_port == "" {
		log.Fatal("Error: faltan algunas variables de entorno para la conexión a la base de datos.")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Db_host, cfg.Db_user, cfg.Db_password, cfg.Db_name, cfg.Db_port)

	return dsn
}
