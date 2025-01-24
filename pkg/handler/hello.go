package handler

import "github.com/gin-gonic/gin"

// Hello responde a una solicitud GET en /hello
func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, world!",
	})
}
