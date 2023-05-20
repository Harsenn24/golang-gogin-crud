package routes

import (
	"go-api/controller/auth"
	"go-api/controller/product"
	"go-api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/register", auth.CreateUser) 
	router.POST("/login", auth.LoginUser) 
	router.PUT("/active-user", auth.ActiveUser) 

	
	router.Use(middleware.Authguard)
	
	router.POST("/product", product.CreateProduct) 



}