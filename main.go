package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teplovyuriydeveloper/article/db"
)

func main() {
	db.Connect()

	r := gin.Default()

	r.GET("/article/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.Run(":8080")
}
