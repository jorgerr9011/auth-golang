// pkg/db/db.go
package db

import (
	"fmt"
	"log"
	"time"

	"jorgerr9011/wiki-golang/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

const DatabaseTimeout = 5 * time.Second

type IDatabase interface {
	GetDB() *gorm.DB
}

type Database struct {
	database *gorm.DB
}

// Inicializa la conexión con la base de datos PostgreSQL
func NewDatabase(uri string) (*Database, error) {

	DB, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Warn),
	})

	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
		return nil, err
	}

	log.Println("Conexión a la base de datos establecida exitosamente.")

	return &Database{
		database: DB,
	}, nil
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

func (d *Database) GetDB() *gorm.DB {
	return d.database
}
