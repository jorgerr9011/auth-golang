package factories

import (
	"fmt"
	model "jorgerr9011/auth-golang/internal/model/user"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
)

// UserFactory genera un usuario falso
func UserFactory() model.User {
	return model.User{
		Name:      faker.Name(),
		Email:     fmt.Sprintf("%s@example.com", faker.Username()),
		Password:  fmt.Sprintf("password%d", rand.Intn(1000)), // Simulación de hash
		Username:  faker.Username(),
		Phone:     fmt.Sprintf("9%d", rand.Intn(100000000)), // Número de 9 dígitos
		LastLogin: time.Now().Add(-time.Duration(rand.Intn(1000)) * time.Hour),
	}
}
