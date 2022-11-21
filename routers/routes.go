package router

import (
	"net/http"

	"github.com/EmeraldLS/Gin_And_MongoDB/controllers"
	"github.com/gin-gonic/gin"
)

func Router(port string) {
	r := gin.New()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "You're on the homepage",
		})
	})
	r.GET("/api/students", controllers.GetAllStudent)
	r.GET("/api/student/:id", controllers.GetSingleStudent)
	r.POST("/api/students", controllers.AddOneStudent)
	r.DELETE("/api/student/:id", controllers.RemoveOneStudent)
	r.DELETE("/api/students", controllers.RemoveAllStudent)
	r.Run(port)
}
