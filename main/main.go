package main

import (
	"gin/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil) // for development only, set to nil to disable trusted proxies
	// r.SetTrustedProxies([]string{"127.0.0.1"}) // for prod, might need to change ip


	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func main() {
	r := setupRouter()

	routes.FileRoutes(r)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	
}
