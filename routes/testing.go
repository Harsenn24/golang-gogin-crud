package routes

import (
	"go-api/controller"

	"github.com/gin-gonic/gin"
)

func TestingRoute(router *gin.Engine) {
	v1 := router.Group("/v1")

	v1.GET("/hello", controller.HelloHandler)
	v1.GET("/hello/:id", controller.HelloParam)
	v1.POST("/hello/data", controller.NewData)
	v1.GET("/hello/query/:id", controller.HelloQuery)
	
}