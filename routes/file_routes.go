package routes

import (
	"gin/controllers"

	"github.com/gin-gonic/gin"
)

func FileRoutes(r *gin.Engine) {
	files := r.Group("/files")
	{
		//files.GET("/", controllers.GetBooks)
		files.POST("/format", controllers.FormatFile)
		//files.DELETE("/:id", controllers.DeleteBook)
	}
}
