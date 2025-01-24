package handler

import (
	"jorgerr9011/wiki-golang/pkg/db"
	"jorgerr9011/wiki-golang/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Crear un nuevo usuario
func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Guardar el usuario en la base de datos
	result := db.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Obtener todos los usuarios
func GetUsers(c *gin.Context) {
	var users []model.User
	if err := db.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
