package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwtSecret []byte

func init() {
	// Carga el .env (puedes cargarlo también en main, pero por si acaso lo haces aquí)
	err := godotenv.Load()
	if err != nil {
		panic("Error cargando archivo .env")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET no está definido en .env")
	}

	jwtSecret = []byte(secret)
}

// GenerateToken genera un JWT con información del usuario
func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Expiración del token (24 horas)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken valida un token JWT y retorna los claims
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token con método de firma inválido")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("token inválido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("no se pudieron obtener los claims")
	}

	return claims, nil
}
