package routes

import (
	"go-api/controller/testing"

	"github.com/gin-gonic/gin"
)

func TestingRoute(router *gin.Engine) {
	v1 := router.Group("/v1")

	v1.GET("/hello", testing.HelloHandler)
	v1.GET("/hello/:id", testing.HelloParam)
	v1.POST("/hello/data", testing.NewData)
	v1.GET("/hello/query/:id", testing.HelloQuery)
	
}