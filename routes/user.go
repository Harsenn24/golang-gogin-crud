package routes

import (
	"go-api/controller/auth"
	"go-api/middleware"
	"go-api/controller/product"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	
	router.POST("/register", auth.CreateUser) 

	router.POST("/login", auth.LoginUser) 

	router.PUT("/active-user", auth.ActiveUser) 

	router.GET("/product-list", product.ListProduct)

	router.Use(middleware.Authguard)
	
	router.POST("/product", product.CreateProduct) 

	router.PUT("/product/:id", product.UpdatePRoduct) 

	router.PUT("/product-image/:id", product.UpdateImageProduct)

	router.DELETE("/product/:id", product.DeleteProduct) 

}