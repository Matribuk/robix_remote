package main

import (
	chatgpt "gin-backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:           []string{"*"},
		AllowHeaders:           []string{"*"},
		AllowAllOrigins:        true,
		AllowCredentials:       true,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
	}))
	router.GET("/api/:message", func(c *gin.Context) {
		message := c.Param("message")
		response := chatgpt.CallOpenAIChatAPI(message)
		c.JSON(200, gin.H{
			"message": response,
		})
	})

	router.Run(":8080")
}
