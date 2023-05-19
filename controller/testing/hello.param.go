package testing

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func HelloParam(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"name":   "harsenn",
		"status": "single",
		"id":     id,
	})
}