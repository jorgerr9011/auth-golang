package main

import (
	"jorgerr9011/auth-golang/internal/handler"
	"jorgerr9011/auth-golang/internal/repository"
	"jorgerr9011/auth-golang/internal/service"
	"jorgerr9011/auth-golang/pkg/config"
	"jorgerr9011/auth-golang/pkg/db"
	"jorgerr9011/auth-golang/pkg/middleware"
	"log"

	"github.com/gin-gonic/gin"

	//"fmt"
	_ "github.com/lib/pq"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	uri := db.GenerateDSN(*cfg)
	database, err := db.NewDatabase(uri)

	if err != nil {
		log.Fatal("Error inicializando la base de datos: ", err)
	}

	// Ejecuta los seeders
	//seeders.RunSeeders(database.GetDB())

	userRepo := repository.NewUserRepository(database.GetDB())
	refreshTokenRepo := repository.NewRefreshTokenRepository(database.GetDB())
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo, refreshTokenRepo)
	userController := handler.NewUserController(userService)
	authController := handler.NewAuthController(authService)

	router := gin.Default()

	// Rutas públicas
	//auth := router.Group("/api/auth")
	//{
	router.POST("/api/auth/register", authController.RegisterUser)
	router.POST("/api/auth/login", authController.LoginUser)
	//}

	authorized := router.Group("/api/users")
	authorized.Use(middleware.JWTAuthMiddleware(*cfg))
	{
		authorized.POST("/", userController.CreateUser)
		authorized.GET("/", userController.GetUsers)
		authorized.GET("/:id", userController.GetUser)
		authorized.PUT("/:id", userController.UpdateUser)
		authorized.DELETE("/:id", userController.DeleteUser)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
