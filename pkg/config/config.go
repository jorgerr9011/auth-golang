package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db_user     string
	Db_password string
	Db_name     string
	Db_host     string
	Db_port     string
	Jwt_secret  string
}

func LoadConfig() (*Config, error) {
	// Carga condicional del .env (solo en desarrollo)
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("No se pudo cargar el archivo .env, usando variables del entorno")
		}
	}

	// Leer las variables de entorno
	return &Config{
		Db_user:     getEnv("DB_USER", "default_user"),
		Db_password: getEnv("DB_PASSWORD", "default_password"),
		Db_name:     getEnv("DB_NAME", "default_db"),
		Db_host:     getEnv("DB_HOST", "localhost"),
		Db_port:     getEnv("DB_PORT", "5432"),
		Jwt_secret:  getEnv("JWT_SECRET", "123456789"),
	}, nil
}

// getEnv obtiene una variable de entorno o devuelve un valor predeterminado
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
