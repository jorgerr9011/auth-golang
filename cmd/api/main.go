package main

import (
	"jorgerr9011/wiki-golang/internal/handler"
	"jorgerr9011/wiki-golang/internal/repository"
	"jorgerr9011/wiki-golang/internal/service"
	"jorgerr9011/wiki-golang/pkg/config"
	"jorgerr9011/wiki-golang/pkg/db"
	"jorgerr9011/wiki-golang/pkg/middleware"
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

	repo := repository.NewUserRepository(database.GetDB())
	userService := service.NewUserService(repo)
	authService := service.NewAuthService(repo)
	userController := handler.NewUserController(userService)
	authController := handler.NewAuthController(authService)

	router := gin.Default()

	// Este uso de variable de entorno es temporal
	//router.Use(middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET_KEY")))

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
