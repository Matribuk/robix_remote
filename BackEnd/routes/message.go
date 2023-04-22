package routes

import (
	chatgpt "robix-backend/controllers"

	"github.com/gin-gonic/gin"
)

func message(c *gin.Context) {
	message := c.Param("message")
	response := chatgpt.CallOpenAIChatAPI(message)
	c.JSON(200, gin.H{
		"message": response,
	})
}
