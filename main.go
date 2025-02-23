package main

import "github.com/gin-gonic/gin"


func main() {
	r := gin.Default()
	// start building main file
	r.POST("/file-convert", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() 
}
