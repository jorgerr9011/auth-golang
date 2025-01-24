package main

import (
	"jorgerr9011/wiki-golang/pkg/db"
	"jorgerr9011/wiki-golang/pkg/handler" // Importa el paquete de los controladores
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la conexión a la base de datos
	db.InitDB()

	router := gin.Default()

	router.GET("/hello", handler.Hello)
	router.GET("/users", handler.GetUsers)
	router.POST("/users", handler.CreateUser)
	/*
		router.GET("/albums", getAlbums)
		router.GET("/albums/:id", getAlbumByID)
		router.POST("/albums", postAlbums)
	*/
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
