package main

import (
	"go-api/config"
	"go-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	config.ConnectDB()

	routes.TestingRoute(router)
	routes.UserRoute(router)

	router.Run(":5000")
}
