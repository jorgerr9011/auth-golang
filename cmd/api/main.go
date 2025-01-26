package main

import (
	"jorgerr9011/wiki-golang/internal/handler"
	"jorgerr9011/wiki-golang/internal/repository"
	"jorgerr9011/wiki-golang/internal/service"
	"jorgerr9011/wiki-golang/pkg/config"
	"jorgerr9011/wiki-golang/pkg/db"
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

	repo := repository.NewUserRepository(database.GetDB())
	userService := service.NewUserService(repo)
	userController := handler.NewUserController(userService)

	router := gin.Default()

	api := router.Group("/api/users")
	{
		api.POST("/", userController.CreateUser)
		//api.GET("/", userController.GetUsers)
		api.GET("/:id", userController.GetUser)
		api.PUT("/:id", userController.UpdateUser)
		api.DELETE("/:id", userController.DeleteUser)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
