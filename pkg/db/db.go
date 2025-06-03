// pkg/db/db.go
package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"jorgerr9011/auth-golang/pkg/config"

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
	sqlDB    *sql.DB
}

func NewDatabase(uri string) (*Database, error) {
	var DB *gorm.DB
	var err error

	const maxAttempts = 10
	const waitBetween = 3 * time.Second

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		DB, err = gorm.Open(postgres.Open(uri), &gorm.Config{
			Logger: gormLogger.Default.LogMode(gormLogger.Warn),
		})
		if err == nil {
			sqlDB, err := DB.DB()
			if err != nil {
				return nil, err
			}

			// Verificar conexión ping
			err = sqlDB.Ping()
			if err == nil {
				log.Println("Conexión a la base de datos establecida exitosamente.")
				return &Database{
					database: DB,
					sqlDB:    sqlDB,
				}, nil
			}
		}

		log.Printf("Intento %d/%d: error al conectar a la base de datos: %v. Reintentando en %s...\n",
			attempts, maxAttempts, err, waitBetween)
		time.Sleep(waitBetween)
	}

	return nil, fmt.Errorf("no se pudo conectar a la base de datos después de %d intentos: %w", maxAttempts, err)
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

func (d *Database) Close() error {
	return d.sqlDB.Close()
}
