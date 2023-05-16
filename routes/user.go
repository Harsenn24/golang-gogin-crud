package routes

import (
	"go-api/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controller.CreateUser) 
	router.GET("/user", controller.AggragateExample) 

}