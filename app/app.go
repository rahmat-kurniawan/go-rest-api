package app

import "github.com/gin-gonic/gin"

func run() {
	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error":   false,
			"message": "Yayyyy I'am Gin Gonic",
		})
	})
	r.Run()
}
