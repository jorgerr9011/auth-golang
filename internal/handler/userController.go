package handler

import (
	"jorgerr9011/wiki-golang/internal/model/user/dto"
	"jorgerr9011/wiki-golang/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

// NewUserController crea una nueva instancia del controlador de usuario
func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var req dto.CreateUserReq

	// Bind the JSON body to the CreateUserReq struct
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := ctrl.userService.CreateUser(c, &req)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUser maneja la obtención de un usuario por ID
func (ctrl *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")

	// Llamar al servicio para obtener el usuario
	user, err := ctrl.userService.GetUserByID(c, id)
	if err != nil {
		log.Printf("Error getting user by ID: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Retornar el usuario encontrado
	c.JSON(http.StatusOK, user)
}

// UpdateUser maneja la actualización de un usuario
func (ctrl *UserController) UpdateUser(c *gin.Context) {

	id := c.Param("id")
	var req dto.UpdateUserReq

	// Bind the JSON body to the UpdateUserReq struct
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Llamar al servicio para actualizar el usuario
	user, err := ctrl.userService.UpdateUser(c, id, &req)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Retornar el usuario actualizado
	c.JSON(http.StatusOK, user)
}

/*
func (ctrl *UserController) GetUsers(c *gin.Context) {
	// Llamar al servicio para obtener la lista de usuarios
	users, err := ctrl.userService.GetAllUsers(c)
	if err != nil {
		log.Printf("Error getting users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	// Retornar la lista de usuarios
	c.JSON(http.StatusOK, users)
}
*/

// DeleteUser maneja la eliminación de un usuario por ID
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id") // Obtener el ID del parámetro en la URL

	// Llamar al servicio para eliminar el usuario
	err := ctrl.userService.DeleteUser(c, id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	// Retornar una respuesta exitosa
	c.JSON(http.StatusNoContent, gin.H{})
}
