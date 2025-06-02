package handler

import (
	"jorgerr9011/auth-golang/internal/model/user/dto"
	"jorgerr9011/auth-golang/internal/service"
	"jorgerr9011/auth-golang/pkg/auth"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ctrl *AuthController) RegisterUser(c *gin.Context) {
	var req dto.CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.authService.Register(c.Request.Context(), &req)
	if err != nil {
		log.Printf("Error registrando usuario: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	// Generar JWT para el usuario registrado
	token, err := auth.GenerateAccessToken(user.ID)
	refreshToken, errRefresh := auth.GenerateRefreshToken(user.ID)
	if err != nil || errRefresh != nil {
		log.Println("Error generating tokens: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating tokens"})
		return
	}

	expiresAt := time.Now().Add(7 * 24 * time.Hour) // 7 días
	if err := ctrl.authService.SaveRefreshToken(c.Request.Context(), user.ID, refreshToken, expiresAt); err != nil {
		log.Println("Error guardando refresh token:", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user":          user,
		"access_token":  token,
		"refresh_token": refreshToken,
	})
}

func (ctrl *AuthController) LoginUser(c *gin.Context) {
	var req dto.LoginUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	user, err := ctrl.authService.Authenticate(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	// Generar token JWT
	token, err := auth.GenerateAccessToken(user.ID)
	refreshToken, errRefresh := auth.GenerateRefreshToken(user.ID)
	if err != nil || errRefresh != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
		return
	}

	expiresAt := time.Now().Add(7 * 24 * time.Hour) // 7 días
	if err := ctrl.authService.SaveRefreshToken(c.Request.Context(), user.ID, refreshToken, expiresAt); err != nil {
		log.Println("Error guardando refresh token:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":          user,
		"access_token":  token,
		"refresh_token": refreshToken,
	})
}
