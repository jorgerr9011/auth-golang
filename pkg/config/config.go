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
	// Cargar variables del archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env, asegurando uso de variables de entorno existentes")
		return nil, err
	}

	// Leer las variables de entorno
	return &Config{
		Db_user:     getEnv("db_user", "default_user"),
		Db_password: getEnv("db_password", "default_password"),
		Db_name:     getEnv("db_name", "default_db"),
		Db_host:     getEnv("db_host", "localhost"),
		Db_port:     getEnv("db_port", "5432"),
		Jwt_secret:  getEnv("JWT_SECRET", "123456789"),
	}, nil
}

// getEnv obtiene una variable de entorno o devuelve un valor predeterminado
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
