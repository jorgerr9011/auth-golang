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

	"os"

	_ "github.com/lib/pq"
)

func main() {
	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		log.Fatal("No se pudo crear la carpeta logs:", err)
	}

	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("No se pudo abrir el archivo de log:", err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	defer logFile.Close()

	// Cargar las variables del .env
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	// Crear conexión a la BD y cerrarla al terminar
	uri := db.GenerateDSN(*cfg)
	database, err := db.NewDatabase(uri)
	if err != nil {
		log.Fatal("Error inicializando la base de datos: ", err)
	}

	defer func() {
		if err := database.Close(); err != nil {
			log.Println("Error cerrando la base de datos:", err)
		}
	}()

	userRepo := repository.NewUserRepository(database.GetDB())
	refreshTokenRepo := repository.NewRefreshTokenRepository(database.GetDB())
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo, refreshTokenRepo)
	userController := handler.NewUserController(userService)
	authController := handler.NewAuthController(authService)

	router := gin.Default()
	router.POST("/api/auth/register", authController.RegisterUser)
	router.POST("/api/auth/login", authController.LoginUser)
	router.POST("/api/auth/refresh", authController.RefreshToken)

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
