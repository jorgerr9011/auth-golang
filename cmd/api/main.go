package main

import (
	"jorgerr9011/wiki-golang/internal/handler"
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
	db.NewDatabase(uri)

	router := gin.Default()

	router.GET("/hello", handler.Hello)
	router.GET("/users", handler.GetUsers)
	router.POST("/users", handler.CreateUser)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
