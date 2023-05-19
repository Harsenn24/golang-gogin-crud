package routes

import (
	"go-api/controller"
	"go-api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/register", controller.CreateUser) 
	router.POST("/login", controller.LoginUser) 
	
	router.Use(middleware.Authguard)
	
	router.GET("/user", controller.AggragateExample) 



}