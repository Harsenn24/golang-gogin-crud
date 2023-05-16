package main

import (
	"go-api/config"
	"go-api/routes"


	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	config.ConnectDB()

	routes.UserRoute(router)
	routes.TestingRoute(router)

	router.Run(":5000")
}
